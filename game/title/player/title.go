package player

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fmt"

	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/title/dao"
	titleentity "fgame/fgame/game/title/entity"
	titleeventtypes "fgame/fgame/game/title/event/types"
	"fgame/fgame/game/title/title"
	titletypes "fgame/fgame/game/title/types"
)

//称号对象
type PlayerTitleObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	TitleId    int32
	ActiveFlag int32
	ActiveTime int64
	ValidTime  int64
	StarLev    int32
	StarNum    int32
	StarBless  int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerTitleObject(pl player.Player) *PlayerTitleObject {
	pto := &PlayerTitleObject{
		player: pl,
	}
	return pto
}

func convertNewPlayerTitleObjectToEntity(pto *PlayerTitleObject) (*titleentity.PlayerTitleEntity, error) {
	e := &titleentity.PlayerTitleEntity{
		Id:         pto.Id,
		PlayerId:   pto.PlayerId,
		TitleId:    pto.TitleId,
		ActiveFlag: pto.ActiveFlag,
		ActiveTime: pto.ActiveTime,
		ValidTime:  pto.ValidTime,
		StarLev:    pto.StarLev,
		StarNum:    pto.StarNum,
		StarBless:  pto.StarBless,
		UpdateTime: pto.UpdateTime,
		CreateTime: pto.CreateTime,
		DeleteTime: pto.DeleteTime,
	}
	return e, nil
}

func (pto *PlayerTitleObject) GetPlayerId() int64 {
	return pto.PlayerId
}

func (pto *PlayerTitleObject) GetDBId() int64 {
	return pto.Id
}

func (pto *PlayerTitleObject) GetStarLev() int32 {
	return pto.StarLev
}

func (pto *PlayerTitleObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerTitleObjectToEntity(pto)
	return e, err
}

func (pto *PlayerTitleObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*titleentity.PlayerTitleEntity)

	pto.Id = pse.Id
	pto.PlayerId = pse.PlayerId
	pto.TitleId = pse.TitleId
	pto.ActiveFlag = pse.ActiveFlag
	pto.ActiveTime = pse.ActiveTime
	pto.ValidTime = pse.ValidTime
	pto.StarLev = pse.StarLev
	pto.StarNum = pse.StarNum
	pto.StarBless = pse.StarBless
	pto.UpdateTime = pse.UpdateTime
	pto.CreateTime = pse.CreateTime
	pto.DeleteTime = pse.DeleteTime
	return nil
}

func (pto *PlayerTitleObject) SetModified() {
	e, err := pto.ToEntity()
	if err != nil {
		panic(fmt.Errorf("title: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pto.player.AddChangedObject(obj)
	return
}

//穿戴称号对象
type PlayerTitleWearObject struct {
	player     player.Player
	Id         int64
	PlayerId   int64
	TitleWear  int32
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerTitleWearObject(pl player.Player) *PlayerTitleWearObject {
	pwto := &PlayerTitleWearObject{
		player: pl,
	}
	return pwto
}

func convertNewPlayerTitleWearObjectToEntity(pwto *PlayerTitleWearObject) (*titleentity.PlayerWearTitleEntity, error) {
	e := &titleentity.PlayerWearTitleEntity{
		Id:         pwto.Id,
		PlayerId:   pwto.PlayerId,
		TitleWear:  pwto.TitleWear,
		UpdateTime: pwto.UpdateTime,
		CreateTime: pwto.CreateTime,
		DeleteTime: pwto.DeleteTime,
	}
	return e, nil
}

func (pwto *PlayerTitleWearObject) GetPlayerId() int64 {
	return pwto.PlayerId
}

func (pwto *PlayerTitleWearObject) GetDBId() int64 {
	return pwto.Id
}

func (pwto *PlayerTitleWearObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerTitleWearObjectToEntity(pwto)
	return e, err
}

func (pwto *PlayerTitleWearObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*titleentity.PlayerWearTitleEntity)

	pwto.Id = pse.Id
	pwto.PlayerId = pse.PlayerId
	pwto.TitleWear = pse.TitleWear
	pwto.UpdateTime = pse.UpdateTime
	pwto.CreateTime = pse.CreateTime
	pwto.DeleteTime = pse.DeleteTime
	return nil
}

func (pwto *PlayerTitleWearObject) SetModified() {
	e, err := pwto.ToEntity()
	if err != nil {
		panic(fmt.Errorf("title: err[%s]", err.Error()))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pwto.player.AddChangedObject(obj)
	return
}

//玩家称号管理器
type PlayerTitleDataManager struct {
	p player.Player
	//玩家称号列表
	titleMap map[titletypes.TitleType]map[int32]*PlayerTitleObject
	//玩家所有激活称号id
	titleIdMap map[int32]*TitleData
	//玩家穿戴称号
	PlayerTitleWearObject *PlayerTitleWearObject
	//未激活活动称号
	noActivityTitleMap map[int32]*PlayerTitleObject
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
}

type TitleData struct {
	activeTime int64 //激活时间
	validTime  int64 //持续时间
}

func (d *TitleData) GetActiveTime() int64 {
	return d.activeTime
}

func (d *TitleData) GetValidTime() int64 {
	return d.validTime
}

func (ptdm *PlayerTitleDataManager) Player() player.Player {
	return ptdm.p
}

//加载
func (ptdm *PlayerTitleDataManager) Load() (err error) {
	ptdm.titleMap = make(map[titletypes.TitleType]map[int32]*PlayerTitleObject)
	ptdm.titleIdMap = make(map[int32]*TitleData)
	ptdm.noActivityTitleMap = make(map[int32]*PlayerTitleObject)
	now := global.GetGame().GetTimeService().Now()
	//加载玩家称号信息
	titleItems, err := dao.GetTitleDao().GetTitleList(ptdm.p.GetId())
	if err != nil {
		return
	}
	//称号信息
	for _, item := range titleItems {
		pto := NewPlayerTitleObject(ptdm.p)
		pto.FromEntity(item)

		titleType, activeFlag := ptdm.titleRefreshCheck(pto, now)
		if activeFlag {
			//添加称号信息
			titleTypeMap, exist := ptdm.titleMap[titleType]
			if !exist {
				titleTypeMap = make(map[int32]*PlayerTitleObject)
				ptdm.titleMap[titleType] = titleTypeMap
			}
			titleTypeMap[pto.TitleId] = pto
			ptdm.addTitleId(pto.TitleId, pto.ActiveTime, pto.ValidTime)
		} else {
			ptdm.noActivityTitleMap[pto.TitleId] = pto
		}
	}

	//加载玩家穿戴称号信息
	titleWearEntity, err := dao.GetTitleDao().GetTitleWearEntity(ptdm.p.GetId())
	if err != nil {
		return
	}
	if titleWearEntity == nil {
		ptdm.initPlayerTitleWearObject()
	} else {
		ptdm.PlayerTitleWearObject = NewPlayerTitleWearObject(ptdm.p)
		ptdm.PlayerTitleWearObject.FromEntity(titleWearEntity)
		ptdm.titleWearRefreshCheck(ptdm.PlayerTitleWearObject, now)
	}

	return nil
}

func (ptdm *PlayerTitleDataManager) titleWearRefreshCheck(ptwo *PlayerTitleWearObject, now int64) {
	titleTemplate := title.GetTitleService().GetTitleTemplate(int(ptwo.TitleWear))
	if titleTemplate == nil {
		return
	}
	//活动称号判断
	titleType := titleTemplate.GetTitleType()
	if titleType != titletypes.TitleTypeActivity {
		return
	}

	_, exist := ptdm.titleIdMap[ptwo.TitleWear]
	if exist {
		return
	}

	ptdm.PlayerTitleWearObject.TitleWear = 0
	ptdm.PlayerTitleWearObject.UpdateTime = now
	ptdm.PlayerTitleWearObject.SetModified()

}

func (ptdm *PlayerTitleDataManager) titleRefreshCheck(pto *PlayerTitleObject, now int64) (titleType titletypes.TitleType, activeFlag bool) {
	activeFlag = true
	titleTemplate := title.GetTitleService().GetTitleTemplate(int(pto.TitleId))
	//活动称号判断
	titleType = titleTemplate.GetTitleType()
	if titleType == titletypes.TitleTypeNormal {
		return
	}

	if pto.ActiveFlag == 0 {
		activeFlag = false
		return
	}

	// TODO xzk: 新增字段，是否要做修正，已激活的时效性称号 赋值validTime
	if pto.ValidTime != 0 {
		existTime := pto.ActiveTime + pto.ValidTime
		if now >= existTime {
			pto.ActiveFlag = 0
			pto.ActiveTime = 0
			pto.ValidTime = 0
			pto.SetModified()
			activeFlag = false

			// 发送称号过期事件
			data := titleeventtypes.CreatePlayerTitleTimeExpireEventData(pto.TitleId, existTime)
			gameevent.Emit(titleeventtypes.EventTypeTitleTimeExpire, ptdm.p, data)
		}
	}

	return
}

//第一次初始化
func (ptdm *PlayerTitleDataManager) initPlayerTitleWearObject() {
	pwto := NewPlayerTitleWearObject(ptdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pwto.Id = id
	//生成id
	pwto.PlayerId = ptdm.p.GetId()
	pwto.TitleWear = int32(0)
	pwto.CreateTime = now
	ptdm.PlayerTitleWearObject = pwto
	pwto.SetModified()
}

//加载后
func (ptdm *PlayerTitleDataManager) AfterLoad() (err error) {
	ptdm.heartbeatRunner.AddTask(CreateActivityTitleTask(ptdm.p))
	return nil
}

//心跳
func (ptdm *PlayerTitleDataManager) Heartbeat() {
	ptdm.heartbeatRunner.Heartbeat()
}

//添加称号
func (ptdm *PlayerTitleDataManager) addTitleId(titleId int32, activeTime, validTime int64) {
	data := &TitleData{
		activeTime: activeTime,
		validTime:  validTime,
	}
	ptdm.titleIdMap[titleId] = data
}

//移除临时称号
func (ptdm *PlayerTitleDataManager) removeTempTitleId(tempTitleId int32) {
	delete(ptdm.titleIdMap, tempTitleId)
}

//移除活动称号
func (ptdm *PlayerTitleDataManager) RemoveActivity(titleId int32) {
	titleTypeMap, exist := ptdm.titleMap[titletypes.TitleTypeActivity]
	if !exist {
		return
	}
	titleObj, exist := titleTypeMap[titleId]
	if !exist {
		return
	}
	titleObj.ActiveFlag = 0
	titleObj.ActiveTime = 0
	titleObj.SetModified()

	ptdm.noActivityTitleMap[titleId] = titleObj
	delete(titleTypeMap, titleId)
	ptdm.removeTempTitleId(titleId)

	titleWear := ptdm.GetTitleId()
	if titleWear == titleId {
		ptdm.TitleNoWear()
	}
	gameevent.Emit(titleeventtypes.EventTypeTitleActivityOverdue, ptdm.p, titleId)
}

//获取活动称号
func (ptdm *PlayerTitleDataManager) GetActivityMap() map[int32]*PlayerTitleObject {
	titleActivityMap, exist := ptdm.titleMap[titletypes.TitleTypeActivity]
	if !exist {
		return nil
	}
	return titleActivityMap
}

//获取活动称号
func (ptdm *PlayerTitleDataManager) GetTitleInfo(titleType titletypes.TitleType, titleId int32) *PlayerTitleObject {
	subMap, exist := ptdm.titleMap[titleType]
	if !exist {
		return nil
	}
	obj, ok := subMap[titleId]
	if !ok {
		return nil
	}
	return obj
}

//获取称号map
func (ptdm *PlayerTitleDataManager) GetTitleIdMap() map[int32]*TitleData {
	return ptdm.titleIdMap
}

//称号穿戴信息
func (ptdm *PlayerTitleDataManager) GetTitleWear() *PlayerTitleWearObject {
	return ptdm.PlayerTitleWearObject
}

//称号穿戴信息
func (ptdm *PlayerTitleDataManager) GetTitleId() int32 {
	return ptdm.PlayerTitleWearObject.TitleWear
}

//校验titleId
func (ptdm *PlayerTitleDataManager) IsValid(titleId int32) bool {
	if titleId <= 0 {
		return false
	}
	titleTemplate := title.GetTitleService().GetTitleTemplate(int(titleId))
	if titleTemplate == nil {
		return false
	}

	titleType := titleTemplate.GetTitleType()
	switch titleType {
	case titletypes.TitleTypeNormal,
		titletypes.TitleTypeActivity:
		{
			return true
		}
	}
	return false
}

func (ptdm *PlayerTitleDataManager) IsWearValid(titleId int32) bool {
	if titleId <= 0 {
		return false
	}
	titleTemplate := title.GetTitleService().GetTitleTemplate(int(titleId))
	if titleTemplate == nil {
		return false
	}

	return true
}

//是否已拥有该称号
func (ptdm *PlayerTitleDataManager) IfTitleExist(titleId int32) bool {
	_, exist := ptdm.titleIdMap[titleId]
	if exist {
		return true
	}
	return false
}

//是否已穿戴
func (ptdm *PlayerTitleDataManager) HasedWeared(titleId int32) bool {
	return ptdm.PlayerTitleWearObject.TitleWear == titleId
}

//称号激活
func (ptdm *PlayerTitleDataManager) TitleActive(titleId int32) (*PlayerTitleObject, bool) {
	flag := ptdm.IsValid(titleId)
	if !flag {
		return nil, false
	}
	flag = ptdm.IfTitleExist(titleId)
	if flag {
		return nil, false
	}

	titleTemplate := title.GetTitleService().GetTitleTemplate(int(titleId))
	titleType := titleTemplate.GetTitleType()

	now := global.GetGame().GetTimeService().Now()
	titleTypeMap, exist := ptdm.titleMap[titleType]
	if !exist {
		titleTypeMap = make(map[int32]*PlayerTitleObject)
		ptdm.titleMap[titleType] = titleTypeMap
	}
	pto, exist := ptdm.noActivityTitleMap[titleId]
	if exist {
		pto.ActiveFlag = 1
		pto.ActiveTime = now
		pto.ValidTime = titleTemplate.Time
		pto.SetModified()
		titleTypeMap[titleId] = pto
		ptdm.addTitleId(titleId, pto.ActiveTime, pto.ValidTime)
		delete(ptdm.noActivityTitleMap, titleId)
	} else {
		id, err := idutil.GetId()
		if err != nil {
			return nil, false
		}

		pto = NewPlayerTitleObject(ptdm.p)
		pto.Id = id
		pto.PlayerId = ptdm.p.GetId()
		pto.TitleId = titleId
		pto.ActiveFlag = 1
		pto.ActiveTime = now
		pto.ValidTime = titleTemplate.Time
		pto.CreateTime = now
		pto.SetModified()
		titleTypeMap[titleId] = pto
		ptdm.addTitleId(titleId, pto.ActiveTime, pto.ValidTime)
	}

	gameevent.Emit(titleeventtypes.EventTypeTitleActivate, ptdm.p, titleId)
	return pto, true
}

//称号叠加时效性
func (ptdm *PlayerTitleDataManager) TitleAddValid(titleId int32) bool {
	flag := ptdm.IsValid(titleId)
	if !flag {
		return false
	}
	titleTemplate := title.GetTitleService().GetTitleTemplate(int(titleId))
	obj := ptdm.GetTitleInfo(titleTemplate.GetTitleType(), titleId)
	if obj == nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	obj.ValidTime += titleTemplate.Time
	obj.UpdateTime = now
	obj.SetModified()

	return true
}

//称号穿戴
func (ptdm *PlayerTitleDataManager) TitleWear(titleId int32) bool {
	flag := ptdm.IfTitleExist(titleId)
	if !flag {
		return false
	}

	flag = ptdm.HasedWeared(titleId)
	if flag {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	ptdm.PlayerTitleWearObject.TitleWear = titleId
	ptdm.PlayerTitleWearObject.UpdateTime = now
	ptdm.PlayerTitleWearObject.SetModified()
	//发送事件
	gameevent.Emit(titleeventtypes.EventTypeTitleChanged, ptdm.p, nil)
	return true
}

// 通过称号id获取称号对象
func (ptdm *PlayerTitleDataManager) GetTitleObjectById(titleId int32) *PlayerTitleObject {
	titleTemplate := title.GetTitleService().GetTitleTemplate(int(titleId))
	if titleTemplate == nil {
		return nil
	}
	obj := ptdm.GetTitleInfo(titleTemplate.GetTitleType(), titleId)
	return obj
}

// 称号升星
func (ptdm *PlayerTitleDataManager) Upstar(titleId int32, bless int32, sucess bool) bool {
	obj := ptdm.GetTitleObjectById(titleId)
	if obj == nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	if sucess {
		obj.StarLev += 1
		obj.StarNum = 0
		obj.StarBless = 0
	} else {
		obj.StarNum += 1
		obj.StarBless += bless
	}
	obj.UpdateTime = now
	obj.SetModified()
	return true
}

//称号卸下
func (ptdm *PlayerTitleDataManager) TitleNoWear() {
	titleWear := ptdm.GetTitleId()
	if titleWear == 0 {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	ptdm.PlayerTitleWearObject.TitleWear = 0
	ptdm.PlayerTitleWearObject.UpdateTime = now
	ptdm.PlayerTitleWearObject.SetModified()

	//发送事件
	gameevent.Emit(titleeventtypes.EventTypeTitleChanged, ptdm.p, nil)
	return
}

//临时称号增加
func (ptdm *PlayerTitleDataManager) TempTitleAdd(titleId int32) {
	if !ptdm.TitleWear(titleId) {
		return
	}

	ptdm.addTitleId(titleId, 0, 0)
	return
}

//临时称号移除
func (ptdm *PlayerTitleDataManager) TempTitleRemove(titleId int32) {
	titleTemplate := title.GetTitleService().GetTitleTemplate(int(titleId))
	if titleTemplate == nil {
		return
	}
	ptdm.removeTempTitleId(titleId)
	titleWear := ptdm.GetTitleId()
	if titleWear == titleId {
		ptdm.TitleNoWear()
	}
	return
}

//增加临时称号
func (ptdm *PlayerTitleDataManager) AddTempTitleIdList(titleIdList []int32) {
	for _, titleId := range titleIdList {
		ptdm.addTitleId(titleId, 0, 0)
		// ptdm.TitleWear(titleId)
	}
}

func (ptdm *PlayerTitleDataManager) HasTitle() bool {
	if len(ptdm.noActivityTitleMap) > 0 {
		return true
	}
	for _, titleMap := range ptdm.titleMap {
		if len(titleMap) > 0 {
			return true
		}
	}
	return false
}

func CreatePlayerTitleDataManager(p player.Player) player.PlayerDataManager {
	ptdm := &PlayerTitleDataManager{}
	ptdm.p = p
	ptdm.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	return ptdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerTitleDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerTitleDataManager))
}
