package treasurebox

import (
	"fgame/fgame/cross/treasurebox/dao"
	treasureboxpb "fgame/fgame/cross/treasurebox/pb"
	arenatemplate "fgame/fgame/game/arena/template"
	droptemplate "fgame/fgame/game/drop/template"
	dummytemplate "fgame/fgame/game/dummy/template"
	"fgame/fgame/game/global"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	playertypes "fgame/fgame/game/player/types"
	boxlogic "fgame/fgame/game/treasurebox/logic"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/mathutils"
	"sort"
	"sync"
)

var (
	minServerId = int(1)
	maxServerId = int(10)
)

//跨服宝箱接口处理
type TreasureBoxService interface {
	Heartbeat()
	//获取跨服宝箱日志列表
	GetLogList() []*TreasureBoxLogObject
	//开跨服宝箱
	OpenTreasureBox(serverId int32, playerName string, itemList []*treasureboxpb.ItemInfo)
}

type treasureBoxService struct {
	logBoxList []*TreasureBoxLogObject
	//读写锁
	rwm sync.RWMutex
	//虚假日志记录时间
	lastAddDummyLogLime int64
}

//初始化
func (rs *treasureBoxService) init() (err error) {
	rs.logBoxList = make([]*TreasureBoxLogObject, 0, 8)
	platform := global.GetGame().GetPlatform()
	riZhiMax := int32(arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().RiZhiMax)
	boxLogList, err := dao.GetTreasureBoXDao().GetTreasureBoxLogList(platform, riZhiMax)
	if err != nil {
		return
	}
	for _, boxLog := range boxLogList {
		tbo := NewTreasureBoxLogObject()
		tbo.FromEntity(boxLog)
		rs.logBoxList = append(rs.logBoxList, tbo)
	}

	return
}

//获取跨服宝箱日志
func (rs *treasureBoxService) GetLogList() []*TreasureBoxLogObject {
	return rs.logBoxList
}

func (rs *treasureBoxService) treasureboxLog(serverId int32, playerName string, lastTime int64) *TreasureBoxLogObject {
	tbo := NewTreasureBoxLogObject()
	id, _ := idutil.GetId()
	tbo.Id = id
	tbo.AreaId = global.GetGame().GetPlatform()
	tbo.ServerId = serverId
	tbo.PlayerName = playerName
	tbo.ItemMap = make(map[int32]int32)
	tbo.LastTime = lastTime
	tbo.CreateTime = lastTime
	return tbo
}

func (rs *treasureBoxService) OpenTreasureBox(serverId int32, playerName string, itemList []*treasureboxpb.ItemInfo) {
	rs.rwm.Lock()
	defer rs.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	logLen := len(rs.logBoxList)
	var tbo *TreasureBoxLogObject
	riZhiMax := int(arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().RiZhiMax)
	if logLen >= riZhiMax {
		tbo = rs.logBoxList[0]
		tbo.ServerId = serverId
		tbo.PlayerName = playerName
		tbo.LastTime = now
		tbo.UpdateTime = now
		tbo.ItemMap = make(map[int32]int32)
		for _, itemInfo := range itemList {
			itemId := itemInfo.GetItemId()
			num := itemInfo.GetNum()
			tbo.ItemMap[itemId] = num
		}
		tbo.SetModified()
	} else {
		tbo = rs.treasureboxLog(serverId, playerName, now)
		for _, itemInfo := range itemList {
			itemId := itemInfo.GetItemId()
			num := itemInfo.GetNum()
			tbo.ItemMap[itemId] = num
		}
		tbo.SetModified()
		rs.logBoxList = append(rs.logBoxList, tbo)
	}
	sort.Sort(TreasureBoxLogObjectList(rs.logBoxList))
}

//心跳
func (rs *treasureBoxService) Heartbeat() {
	//TODO 假日志
	now := global.GetGame().GetTimeService().Now()
	lastTime := rs.lastAddDummyLogLime
	diffTime := now - lastTime
	randTime := rs.getRandomLogTime()
	if diffTime < randTime {
		return
	}
	rs.lastAddDummyLogLime = now

	serverId := rs.getRandomLogServerId()
	name := dummytemplate.GetDummyTemplateService().GetRandomDummyName()

	itemTemplate := item.GetItemService().GetItemTemplate(itemtypes.ItemTypeGiftBag, itemtypes.ItemGiftBagSubTypeCorss)
	if itemTemplate == nil {
		return
	}

	starBoxTemplate := itemTemplate.GetBoxTemplate()
	boxTemplate := boxlogic.GetRandomBoxTemplate(starBoxTemplate)

	sexType := playertypes.RandomSex()
	roleType := playertypes.RandomRole()
	dropIdList := boxTemplate.GetDropIdList(roleType, sexType)
	dropList := droptemplate.GetDropTemplateService().GetDropListItemLevelList(dropIdList)
	if len(dropList) == 0 {
		return
	}

	rs.rwm.Lock()
	defer rs.rwm.Unlock()
	var tbo *TreasureBoxLogObject
	logLen := len(rs.logBoxList)

	riZhiMax := int(arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().RiZhiMax)
	if logLen >= riZhiMax {
		tbo = rs.logBoxList[0]
		tbo.ServerId = serverId
		tbo.PlayerName = name
		tbo.LastTime = now
		tbo.UpdateTime = now
		tbo.ItemMap = make(map[int32]int32)
		for _, itemInfo := range dropList {
			itemId := itemInfo.GetItemId()
			num := itemInfo.GetNum()
			tbo.ItemMap[itemId] = num
		}
		tbo.SetModified()
	} else {
		tbo = rs.treasureboxLog(serverId, name, now)
		for _, itemInfo := range dropList {
			itemId := itemInfo.GetItemId()
			num := itemInfo.GetNum()
			tbo.ItemMap[itemId] = num
		}
		tbo.SetModified()
		rs.logBoxList = append(rs.logBoxList, tbo)
	}
	sort.Sort(TreasureBoxLogObjectList(rs.logBoxList))
}

func (rs *treasureBoxService) getRandomLogServerId() int32 {
	randServerId := int32(mathutils.RandomRange(minServerId, maxServerId))
	return randServerId
}

//系统假日志生成间隔
func (rs *treasureBoxService) getRandomLogTime() int64 {
	minTime := int(arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().RiZhiTimeMin)
	maxTime := int(arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().RiZhiTimeMax)
	randTime := int64(mathutils.RandomRange(minTime, maxTime))
	return randTime
}

var (
	once sync.Once
	cs   *treasureBoxService
)

func Init() (err error) {
	once.Do(func() {
		cs = &treasureBoxService{}
		err = cs.init()
	})
	return err
}

func GetTreasureBoxService() TreasureBoxService {
	return cs
}
