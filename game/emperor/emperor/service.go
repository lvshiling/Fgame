package emperor

import (
	"fgame/fgame/game/common/common"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/emperor/dao"
	emperortypes "fgame/fgame/game/emperor/emperor/types"
	emperoreventtypes "fgame/fgame/game/emperor/event/types"
	emperortemplate "fgame/fgame/game/emperor/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	"fgame/fgame/pkg/idutil"
	"sort"
	"sync"
)

//抢龙椅接口处理
type EmperorService interface {
	Heartbeat()
	//获取龙椅信息
	GetEmperorInfo() *EmperorObject
	//膜拜增加帝王金库
	EmperorStorageUpdate()
	//获取当前帝王国库
	GetEmperorStorage() (storage int64)
	//获取帝王id和抢夺时间
	GetEmperorIdAndRobTime() (emperorId int64, robTime int64)
	//获取帝王抢夺次数
	GetEmperorIdAndRobNum() (emperorId int64, robNum int64)
	//获取抢夺记录
	GetEmperorRobListByLogTime(logTime int64) []*EmperorRecordsObject
	//领取帝王金库
	EmperorStorageGet(emperorId int64, now int64) (storage int64, successful bool)
	//帝王被抢
	EmperorRobbed(pl player.Player, spouseName string, robNum int64) (dropItemList []*droptemplate.DropItemData, err error)
	//设置配偶名字
	SetEmperorSpouseName(proposalId int64, name string, dealId int64, peerName string)
	//重置配偶名字
	ResetEmperorSpouseName(playerId int64, spouseId int64)
	//获取帝王宝箱次数
	GetEmperorIdAndBoxNum() (emperorId int64, boxNum int64)
	//帝王开宝箱
	OpenBox(pl player.Player) (emperorInfo *EmperorObject, dropItemList []*droptemplate.DropItemData, err error)
	//玩家名字变化
	PlayerNameChanged(pl player.Player)
	//玩家性别变化
	PlayerSexChanged(pl player.Player)
	//获取上一个帝王的消耗元宝
	GetLastEmperorCostGold() int32
	//仅gm使用
	GMClearEmperor()
	GmSetEmperorSilver(silver int64)
	GmSetEmperorBox(boxNum int64)
	GMResetLastTime()
}

const (
	//抢龙椅日志记录大小
	RobEmperorLogSize = int(50)
	specialBoxLeftNum = int32(2)
)

type emperorConfig struct {
	//国库存银最大
	chestMax int64
	//膜拜给金库增加银两数量
	chestSilver int64
}

//读写锁
type emperorService struct {
	rwm sync.RWMutex
	//龙椅
	emperorObj *EmperorObject
	//龙椅抢夺记录
	emperorRecordsList []*EmperorRecordsObject
	//抢龙椅常量配置
	emperorConst emperorConfig
}

//初始化
func (es *emperorService) init() (err error) {
	serverId := global.GetGame().GetServerIndex()
	//抢龙椅记录
	emperorRecordsList, err := dao.GetEmperorDao().GetEmperorRecordsList(serverId)
	if err != nil {
		return
	}
	for _, emperorRecord := range emperorRecordsList {
		ero := NewEmperorRecordsObject()
		ero.FromEntity(emperorRecord)
		es.emperorRecordsList = append(es.emperorRecordsList, ero)
	}

	//合服
	isMerge := merge.GetMergeService().IsMerge()
	if isMerge {
		err = es.mergeServer(serverId)
		if err != nil {
			return
		}
	} else {
		err = es.initEmperor(serverId)
		if err != nil {
			return
		}
	}

	es.emperorConst.chestMax = emperortemplate.GetEmperorTemplateService().GetEmperorChestMax()
	es.emperorConst.chestSilver = emperortemplate.GetEmperorTemplateService().GetEmperorWorshipChestSilver()
	return
}

func (es *emperorService) mergeServer(serverId int32) (err error) {
	emperorEntityList, err := dao.GetEmperorDao().GetEmperorList(serverId)
	if err != nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	for _, emperorEmtity := range emperorEntityList {
		emperorObj := NewEmperorObject()
		emperorObj.FromEntity(emperorEmtity)
		emperorObj.DeleteTime = now
		emperorObj.SetModified()
		if emperorObj.EmperorId == 0 {
			continue
		}
		itemMap := make(map[int32]int32)
		needGold := int32(emperortemplate.GetEmperorTemplateService().GetEmperorRobNeedGold(emperorObj.RobNum))
		itemMap[constanttypes.GoldItem] = needGold
		gameevent.Emit(emperoreventtypes.EmperorMergeServer, emperorObj.EmperorId, itemMap)
	}
	es.initEmperorObject()
	return
}

func (es *emperorService) initEmperor(serverId int32) (err error) {
	//龙椅数据
	emperorEntity, err := dao.GetEmperorDao().GetEmperorEntity(serverId)
	if err != nil {
		return err
	}
	if emperorEntity == nil {
		es.initEmperorObject()
	} else {
		es.emperorObj = NewEmperorObject()
		err = es.emperorObj.FromEntity(emperorEntity)
		if err != nil {
			return err
		}
	}
	return
}

//第一次初始化
func (es *emperorService) initEmperorObject() {
	peo := NewEmperorObject()
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	peo.Id = id
	peo.ServerId = global.GetGame().GetServerIndex()
	peo.EmperorId = int64(0)
	peo.Name = ""
	peo.SpouseName = ""
	peo.Storage = int64(0)
	peo.RobNum = int64(0)
	peo.RobTime = int64(0)
	peo.CreateTime = now
	es.emperorObj = peo
	peo.SetModified()
}

//心跳
func (es *emperorService) Heartbeat() {
	es.rwm.Lock()
	defer es.rwm.Unlock()
	es.addStorageTimePhase()
	es.addBoxTimePhase()
}

func (es *emperorService) getAddSilver(now int64) (autoSilver int64, offTime int64) {
	autoSilver = 0
	robTime := es.emperorObj.RobTime
	robElapseTime := now - robTime
	elapseTime := now - es.emperorObj.LastTime
	autoSilver, offTime = emperortemplate.GetEmperorTemplateService().GetEmperorChanChuSilver(robElapseTime, elapseTime)
	return
}

//定时往国库加银两
func (es *emperorService) addStorageTimePhase() {
	if es.emperorObj.EmperorId == 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	elapseTime := now - es.emperorObj.LastTime
	if elapseTime < int64(common.MINUTE) {
		return
	}

	autoSilver, offTime := es.getAddSilver(now)
	storage := es.emperorObj.Storage + autoSilver
	if storage >= es.emperorConst.chestMax {
		storage = es.emperorConst.chestMax
	}

	es.emperorObj.Storage = storage
	es.emperorObj.LastTime += offTime
	es.emperorObj.UpdateTime = now
	es.emperorObj.SetModified()
	return
}

//定时增加帝王宝箱
func (es *emperorService) addBoxTimePhase() {
	if es.emperorObj.EmperorId == 0 {
		return
	}
	isAdd := false
	addNum := int32(0)
	offTime := int64(0)
	now := global.GetGame().GetTimeService().Now()
	robElapseTime := now - es.emperorObj.RobTime
	elapseTime := now - es.emperorObj.BoxLastTime
	if elapseTime <= 0 {
		return
	}
	boxOutNum := es.emperorObj.BoxOutNum
	addNum, offTime, isAdd = emperortemplate.GetEmperorTemplateService().IfAddEmperorBox(elapseTime, robElapseTime, boxOutNum)
	if !isAdd {
		return
	}
	es.emperorObj.BoxNum += int64(addNum)
	es.emperorObj.BoxOutNum += int64(addNum)
	es.emperorObj.BoxLastTime += offTime
	es.emperorObj.UpdateTime = now
	es.emperorObj.SetModified()
}

func (es *emperorService) newEmperorRecordsObject(operType int32, emperorName string, robberdName string, robTime int64, dropItemList []*droptemplate.DropItemData) *EmperorRecordsObject {
	peo := NewEmperorRecordsObject()
	id, _ := idutil.GetId()
	peo.Id = id
	peo.Type = operType
	peo.ItemMap = make(map[int32]int32)
	for _, dropItem := range dropItemList {
		itemId := dropItem.ItemId
		num := dropItem.Num
		peo.ItemMap[itemId] += num
	}
	peo.ServerId = global.GetGame().GetServerIndex()
	peo.EmperorName = emperorName
	peo.RobbedName = robberdName
	peo.RobTime = robTime
	peo.CreateTime = robTime
	peo.SetModified()
	return peo
}

//获取龙椅信息
func (es *emperorService) GetEmperorInfo() *EmperorObject {
	return es.emperorObj
}

//获取帝王id和抢夺时间
func (es *emperorService) GetEmperorIdAndRobTime() (emperorId int64, robTime int64) {
	return es.emperorObj.EmperorId, es.emperorObj.RobTime
}

//获取帝王抢夺次数
func (es *emperorService) GetEmperorIdAndRobNum() (emperorId int64, robNum int64) {
	return es.emperorObj.EmperorId, es.emperorObj.RobNum
}

//获取帝王id和宝箱次数
func (es *emperorService) GetEmperorIdAndBoxNum() (emperorId int64, boxNum int64) {
	return es.emperorObj.EmperorId, es.emperorObj.BoxNum
}

//领取帝王金库
func (es *emperorService) EmperorStorageGet(emperorId int64, now int64) (storage int64, successful bool) {
	es.rwm.Lock()
	defer es.rwm.Unlock()

	if es.emperorObj.EmperorId != emperorId {
		return 0, false
	}
	storage = es.emperorObj.Storage
	es.emperorObj.RobTime = now
	es.emperorObj.UpdateTime = now
	es.emperorObj.Storage = 0
	es.emperorObj.SetModified()
	successful = true
	return
}

//膜拜更新帝王金库
func (es *emperorService) EmperorStorageUpdate() {
	es.rwm.Lock()
	defer es.rwm.Unlock()

	storage := es.emperorObj.Storage + es.emperorConst.chestSilver
	if storage >= es.emperorConst.chestMax {
		storage = es.emperorConst.chestMax
	}
	now := global.GetGame().GetTimeService().Now()
	es.emperorObj.Storage = storage
	es.emperorObj.UpdateTime = now
	es.emperorObj.SetModified()
	return
}

//帝王被抢
func (es *emperorService) EmperorRobbed(pl player.Player, spouseName string, robNum int64) (dropItemList []*droptemplate.DropItemData, err error) {
	es.rwm.Lock()
	defer es.rwm.Unlock()
	now := global.GetGame().GetTimeService().Now()
	curNum := es.emperorObj.RobNum
	diffNum := robNum - curNum
	if diffNum != 1 {
		err = errorEmperorRobbedByOther
		return
	}
	oldPlayerId := es.emperorObj.EmperorId
	robberName := es.emperorObj.Name
	es.emperorObj.RobNum++
	es.emperorObj.EmperorId = pl.GetId()
	es.emperorObj.UpdateTime = now
	es.emperorObj.RobTime = now
	es.emperorObj.Name = pl.GetName()
	es.emperorObj.Sex = pl.GetSex()
	es.emperorObj.SpouseName = spouseName
	es.emperorObj.SpecialBoxLeftNum = specialBoxLeftNum
	es.emperorObj.BoxLastTime = now
	es.emperorObj.LastTime = now
	es.emperorObj.BoxOutNum = 0
	es.emperorObj.SetModified()
	dropItemList = es.getBoxItemMap(true)
	es.emperorRecords(emperortypes.EmperorRecordLogTypeRob, pl.GetName(), robberName, now, nil)
	es.emperorRecords(emperortypes.EmperorRecordLogTypeOpen, pl.GetName(), "", now, dropItemList)
	gameevent.Emit(emperoreventtypes.EmperorEventTypeRobed, pl, oldPlayerId)
	return
}

func (es *emperorService) emperorRecords(typ emperortypes.EmperorRecordLogType, emperorName string, robberName string, robTime int64, dropItemList []*droptemplate.DropItemData) {
	maxLen := RobEmperorLogSize
	operType := int32(typ)
	if typ == emperortypes.EmperorRecordLogTypeOpen && len(dropItemList) == 0 {
		return
	}
	ero := es.newEmperorRecordsObject(operType, emperorName, robberName, robTime, dropItemList)
	es.emperorRecordsList = append(es.emperorRecordsList, ero)
	curLen := len(es.emperorRecordsList)
	sort.Sort(sort.Reverse(EmperorRecordsObjectList(es.emperorRecordsList)))
	if curLen > maxLen {
		for index := maxLen; index < curLen; index++ {
			es.emperorRecordsList[index].DeleteTime = robTime
			es.emperorRecordsList[index].UpdateTime = robTime
			es.emperorRecordsList[index].SetModified()
		}
		es.emperorRecordsList = es.emperorRecordsList[:maxLen]
	}
	return
}

//获取当前帝王国库
func (es *emperorService) GetEmperorStorage() (storage int64) {
	return es.emperorObj.Storage
}

//获取抢夺记录
func (es *emperorService) GetEmperorRobListByLogTime(logTime int64) []*EmperorRecordsObject {
	starIndex := int(-1)
	for index, emperorRecord := range es.emperorRecordsList {
		if emperorRecord.RobTime <= logTime {
			break
		}
		starIndex = index
	}
	if starIndex >= 0 {
		return es.emperorRecordsList[0 : starIndex+1]
	}
	return nil
}

//设置配偶名字
func (es *emperorService) SetEmperorSpouseName(proposalId int64, name string, dealId int64, peerName string) {
	es.rwm.Lock()
	defer es.rwm.Unlock()

	if es.emperorObj.EmperorId != proposalId && es.emperorObj.EmperorId != dealId {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if es.emperorObj.EmperorId == proposalId {
		es.emperorObj.SpouseName = name
	} else {
		es.emperorObj.SpouseName = peerName
	}
	es.emperorObj.UpdateTime = now
	es.emperorObj.SetModified()
}

//重置配偶名字
func (es *emperorService) ResetEmperorSpouseName(playerId int64, spouseId int64) {
	es.rwm.Lock()
	defer es.rwm.Unlock()

	if es.emperorObj.EmperorId != playerId && es.emperorObj.EmperorId != spouseId {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	es.emperorObj.SpouseName = ""
	es.emperorObj.UpdateTime = now
	es.emperorObj.SetModified()
}

//玩家名字变化
func (es *emperorService) PlayerNameChanged(pl player.Player) {
	es.rwm.Lock()
	defer es.rwm.Unlock()
	playerId := pl.GetId()
	spouseId := pl.GetSpouseId()
	now := global.GetGame().GetTimeService().Now()
	if es.emperorObj.EmperorId == playerId {
		es.emperorObj.Name = pl.GetName()
		es.emperorObj.UpdateTime = now
		es.emperorObj.SetModified()
		return
	}
	if es.emperorObj.EmperorId == spouseId {
		es.emperorObj.SpouseName = pl.GetName()
		es.emperorObj.UpdateTime = now
		es.emperorObj.SetModified()
		return
	}

}

func (es *emperorService) PlayerSexChanged(pl player.Player) {
	es.rwm.Lock()
	defer es.rwm.Unlock()
	playerId := pl.GetId()
	now := global.GetGame().GetTimeService().Now()
	if es.emperorObj.EmperorId == playerId {
		es.emperorObj.Sex = pl.GetSex()
		es.emperorObj.UpdateTime = now
		es.emperorObj.SetModified()
		return
	}

}

func (es *emperorService) getBoxItemMap(isRob bool) (dropItemList []*droptemplate.DropItemData) {
	emperorTempate := emperortemplate.GetEmperorTemplateService().GetEmperorTemplate()
	if emperorTempate == nil {
		return
	}
	var dropIdList []int32
	if isRob {
		dropIdList = emperorTempate.GetSpecialDropList()
	} else {
		if es.emperorObj.SpecialBoxLeftNum > 0 {
			es.emperorObj.SpecialBoxLeftNum--
			dropIdList = emperorTempate.GetSpecialDropList()
		} else {
			dropIdList = emperorTempate.GetCommonDropList()
		}
	}

	if len(dropIdList) != 0 {
		dropItemList = droptemplate.GetDropTemplateService().GetDropListItemLevelList(dropIdList)
	}
	return
}

func (es *emperorService) OpenBox(pl player.Player) (emperorInfo *EmperorObject, dropItemList []*droptemplate.DropItemData, err error) {
	es.rwm.Lock()
	defer es.rwm.Unlock()

	if es.emperorObj.EmperorId != pl.GetId() {
		err = errorEmperorOpenBoxRobbed
		return
	}
	if es.emperorObj.BoxNum <= 0 {
		err = errorEmperorOpenBoxNoStorage
		return
	}
	now := global.GetGame().GetTimeService().Now()
	dropItemList = es.getBoxItemMap(false)
	es.emperorObj.BoxNum -= 1
	es.emperorObj.UpdateTime = now
	es.emperorObj.SetModified()
	es.emperorRecords(emperortypes.EmperorRecordLogTypeOpen, pl.GetName(), "", now, dropItemList)
	emperorInfo = es.emperorObj
	return
}

//获取上一个帝王的消耗元宝
func (es *emperorService) GetLastEmperorCostGold() int32 {
	_, robNum := es.GetEmperorIdAndRobNum()
	if robNum == 1 {
		return 0
	}
	return int32(emperortemplate.GetEmperorTemplateService().GetEmperorRobNeedGold(robNum - 1))
}

//仅gm使用
func (es *emperorService) GMClearEmperor() {
	es.rwm.Lock()
	defer es.rwm.Unlock()
	now := global.GetGame().GetTimeService().Now()
	es.emperorObj.RobNum = 0
	es.emperorObj.EmperorId = 0
	es.emperorObj.UpdateTime = now
	es.emperorObj.RobTime = 0
	es.emperorObj.Name = ""
	es.emperorObj.SpouseName = ""
	es.emperorObj.Storage = 0
	es.emperorObj.BoxNum = 0
	es.emperorObj.BoxOutNum = 0
	es.emperorObj.SpecialBoxLeftNum = 0
	es.emperorObj.LastTime = 0
	es.emperorObj.BoxLastTime = 0
	es.emperorObj.SetModified()

	for _, recordObj := range es.emperorRecordsList {
		recordObj.DeleteTime = now
		recordObj.SetModified()
	}
	es.emperorRecordsList = make([]*EmperorRecordsObject, 0, 8)
	return
}

func (es *emperorService) GmSetEmperorSilver(silver int64) {
	es.rwm.Lock()
	defer es.rwm.Unlock()
	now := global.GetGame().GetTimeService().Now()

	if silver < 0 || silver > int64(es.emperorConst.chestMax) {
		return
	}
	es.emperorObj.Storage = silver
	es.emperorObj.UpdateTime = now
	es.emperorObj.SetModified()
}

func (es *emperorService) GmSetEmperorBox(boxNum int64) {
	es.rwm.Lock()
	defer es.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	if boxNum < 0 {
		return
	}
	es.emperorObj.BoxNum = boxNum
	es.emperorObj.UpdateTime = now
	es.emperorObj.SetModified()
}

func (es *emperorService) GMResetLastTime() {
	es.rwm.Lock()
	defer es.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	es.emperorObj.LastTime = now
	es.emperorObj.BoxLastTime = now
	es.emperorObj.UpdateTime = now
	es.emperorObj.SetModified()
}

var (
	once sync.Once
	cs   *emperorService
)

func Init() (err error) {
	once.Do(func() {
		cs = &emperorService{}
		err = cs.init()
	})
	return err
}

func GetEmperorService() EmperorService {
	return cs
}
