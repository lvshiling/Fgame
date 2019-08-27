package player

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/game/constant/constant"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/fashion/dao"
	fashioneventtypes "fgame/fgame/game/fashion/event/types"
	"fgame/fgame/game/fashion/fashion"
	fashiontypes "fgame/fgame/game/fashion/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fmt"
)

//玩家时装管理器
type PlayerFashionDataManager struct {
	p player.Player
	//玩家时装map
	fashionMap map[fashiontypes.FashionType]map[int32]*PlayerFashionObject
	//未激活时效时装
	noExpireFahionMap map[int32]*PlayerFashionObject
	//玩家穿戴时装
	PlayerFashionWearObject *PlayerFashionWearObject
	//玩家时装试用
	playerFashionTrialMap map[int32]*PlayerFashionTrialObject
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
}

func (m *PlayerFashionDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerFashionDataManager) Load() (err error) {
	err = m.loadFashionObj()
	if err != nil {
		return
	}

	err = m.loadFashionWearObj()
	if err != nil {
		return
	}

	err = m.loadFashionTrialObj()
	if err != nil {
		return
	}

	return
}

//加载玩家时装信息
func (m *PlayerFashionDataManager) loadFashionObj() (err error) {
	now := global.GetGame().GetTimeService().Now()

	fashionList, err := dao.GetFashionDao().GetFashionList(m.p.GetId())
	if err != nil {
		return
	}
	for _, item := range fashionList {
		pfo := NewPlayerFashionObject(m.p)
		pfo.FromEntity(item)

		fashionType, activeFlag := m.fashionRefreshCheck(pfo, now)
		if activeFlag {
			//添加时装信息
			fashionTypeMap, exist := m.fashionMap[fashionType]
			if !exist {
				fashionTypeMap = make(map[int32]*PlayerFashionObject)
				m.fashionMap[fashionType] = fashionTypeMap
			}
			fashionTypeMap[pfo.FashionId] = pfo
		} else {
			m.noExpireFahionMap[pfo.FashionId] = pfo
		}
	}

	return
}

//加载玩家穿戴时装信息
func (m *PlayerFashionDataManager) loadFashionWearObj() (err error) {
	now := global.GetGame().GetTimeService().Now()
	fashionWearEntity, err := dao.GetFashionDao().GetFashionWearEntity(m.p.GetId())
	if err != nil {
		return
	}
	if fashionWearEntity == nil {
		m.initPlayerFashionWearObject()
	} else {
		m.PlayerFashionWearObject = NewPlayerFashionWearObject(m.p)
		m.PlayerFashionWearObject.FromEntity(fashionWearEntity)
		m.fashionWearRefreshCheck(m.PlayerFashionWearObject, now)
	}

	return
}

//加载玩家时装试用信息
func (m *PlayerFashionDataManager) loadFashionTrialObj() (err error) {
	fashionTrialEntityList, err := dao.GetFashionDao().GetFashionTrialList(m.p.GetId())
	if err != nil {
		return
	}

	for _, entity := range fashionTrialEntityList {
		obj := NewPlayerFashionTrialObject(m.p)
		obj.FromEntity(entity)
		m.playerFashionTrialMap[obj.trialFashionId] = obj
	}

	return
}

func (m *PlayerFashionDataManager) fashionWearRefreshCheck(pfwo *PlayerFashionWearObject, now int64) {
	fashionTemplate := fashion.GetFashionService().GetFashionTemplate(int(pfwo.FashionWear))
	if fashionTemplate == nil {
		return
	}
	playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(m.p.GetRole(), m.p.GetSex())
	bornFashionId := playerCreateTemplate.FashionId

	//时装试用判断
	trialObj := m.getTrialFashionObj(m.PlayerFashionWearObject.FashionWear)
	if trialObj == nil {
		return
	}
	if now > trialObj.expireTime {
		m.PlayerFashionWearObject.FashionWear = bornFashionId
		m.PlayerFashionWearObject.UpdateTime = now
		m.PlayerFashionWearObject.SetModified()

		m.TrialFashionOverdue(trialObj.trialFashionId, fashiontypes.FashionTrialOverdueTypeExpire)
	}

	//时装判断
	fashionType := fashionTemplate.GetFashionType()
	if fashionType != fashiontypes.FashionTypeEffective {
		return
	}
	fashionTypeMap, exist := m.fashionMap[fashiontypes.FashionTypeEffective]
	if !exist {
		m.PlayerFashionWearObject.FashionWear = bornFashionId
		m.PlayerFashionWearObject.UpdateTime = now
		m.PlayerFashionWearObject.SetModified()
		return
	}
	fashionObj, exist := fashionTypeMap[pfwo.FashionWear]
	if !exist {
		m.PlayerFashionWearObject.FashionWear = bornFashionId
		m.PlayerFashionWearObject.UpdateTime = now
		m.PlayerFashionWearObject.SetModified()
		return
	}

	if fashionObj.IsExpire == 1 {
		m.PlayerFashionWearObject.FashionWear = bornFashionId
		m.PlayerFashionWearObject.UpdateTime = now
		m.PlayerFashionWearObject.SetModified()
	}
	return

}

func (m *PlayerFashionDataManager) fashionRefreshCheck(pfo *PlayerFashionObject, now int64) (fashionType fashiontypes.FashionType, activeFlag bool) {
	activeFlag = true
	fashionTemplate := fashion.GetFashionService().GetFashionTemplate(int(pfo.FashionId))
	//时装判断
	fashionType = fashionTemplate.GetFashionType()
	if fashionType == fashiontypes.FashionTypeNormal {
		return
	}

	if pfo.IsExpire == 1 {
		activeFlag = false
		return
	}

	existTime := fashionTemplate.Time
	diffTime := now - pfo.ActiveTime
	if diffTime >= existTime {
		pfo.IsExpire = 1
		pfo.ActiveTime = 0
		pfo.SetModified()
		activeFlag = false
	}
	return
}

//第一次初始化
func (m *PlayerFashionDataManager) initPlayerFashionWearObject() {
	playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(m.p.GetRole(), m.p.GetSex())
	//初始化第一套时装
	if !m.IfFashionExist(playerCreateTemplate.FashionId) {
		_, flag := m.FashionActive(playerCreateTemplate.FashionId, false)
		if !flag {
			panic(fmt.Errorf("fashion:初始化第一套时装应该成功"))
		}
	}

	pfwo := NewPlayerFashionWearObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pfwo.Id = id
	//生成id
	pfwo.PlayerId = m.p.GetId()
	pfwo.FashionWear = int32(playerCreateTemplate.FashionId)
	pfwo.CreateTime = now
	m.PlayerFashionWearObject = pfwo
	pfwo.SetModified()
}

//加载后
func (m *PlayerFashionDataManager) AfterLoad() (err error) {
	m.heartbeatRunner.AddTask(CreateFashionTask(m.p))
	m.heartbeatRunner.AddTask(CreateFashionTrialTask(m.p))
	return nil
}

//心跳
func (m *PlayerFashionDataManager) Heartbeat() {
	m.heartbeatRunner.Heartbeat()
}

//时装map信息
func (m *PlayerFashionDataManager) GetFashionMap() map[fashiontypes.FashionType]map[int32]*PlayerFashionObject {
	return m.fashionMap
}

//时装穿戴信息
func (m *PlayerFashionDataManager) GetFashionWear() *PlayerFashionWearObject {
	return m.PlayerFashionWearObject
}

//获取当前时装
func (m *PlayerFashionDataManager) GetFashionId() int32 {
	if m.PlayerFashionWearObject.FashionWear == 0 {
		playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(m.p.GetRole(), m.p.GetSex())
		return playerCreateTemplate.FashionId
	}
	return m.PlayerFashionWearObject.FashionWear
}

//根据时装id获取时装信息
func (m *PlayerFashionDataManager) getFashionById(fashionType fashiontypes.FashionType, fashionId int32) (pfo *PlayerFashionObject) {
	fashionTypeMap, exist := m.fashionMap[fashionType]
	if !exist {
		return
	}
	fashionObj, exist := fashionTypeMap[fashionId]
	if !exist {
		return
	}
	pfo = fashionObj
	return

}

//校验fashionId
func (m *PlayerFashionDataManager) IsValid(fashionId int32) bool {
	if fashionId <= 0 {
		return false
	}
	fashionTemplate := fashion.GetFashionService().GetFashionTemplate(int(fashionId))
	if fashionTemplate == nil {
		return false
	}
	return true
}

//是否已拥有该时装
func (m *PlayerFashionDataManager) IfFashionExist(fashionId int32) bool {
	fashionTemplate := fashion.GetFashionService().GetFashionTemplate(int(fashionId))
	if fashionTemplate == nil {
		return false
	}
	fashionType := fashionTemplate.GetFashionType()
	//该时装是否已获得
	fashion := m.getFashionById(fashionType, fashionId)
	if fashion == nil {
		return false
	}
	return true
}

//是否可以穿戴
func (m *PlayerFashionDataManager) IfFashionWear(fashionId int32) bool {
	isExist := m.IfFashionExist(fashionId)
	if isExist {
		return true
	}

	if m.IsFashionTrial(fashionId) {
		return true
	}

	return false
}

//是否已穿戴
func (m *PlayerFashionDataManager) HasedWeared(fashionId int32) bool {
	return m.PlayerFashionWearObject.FashionWear == fashionId
}

//时装激活
func (m *PlayerFashionDataManager) FashionActive(fashionId int32, sendEvent bool) (acitveTime int64, flag bool) {
	flag = m.IsValid(fashionId)
	if !flag {
		return 0, false
	}
	fashionTemplate := fashion.GetFashionService().GetFashionTemplate(int(fashionId))
	fashionType := fashionTemplate.GetFashionType()
	flag = m.IfFashionExist(fashionId)
	if flag {
		return 0, false
	}

	now := global.GetGame().GetTimeService().Now()
	fashionTypeMap, exist := m.fashionMap[fashionType]
	if !exist {
		fashionTypeMap = make(map[int32]*PlayerFashionObject)
		m.fashionMap[fashionType] = fashionTypeMap
	}
	fashionObj, exist := m.noExpireFahionMap[fashionId]
	if exist {
		fashionObj.IsExpire = 0
		fashionObj.ActiveTime = now
		fashionObj.Star = 0
		fashionObj.UpStarNum = 0
		fashionObj.UpStarPro = 0
		fashionObj.SetModified()
		fashionTypeMap[fashionId] = fashionObj
		delete(m.noExpireFahionMap, fashionId)
	} else {
		id, err := idutil.GetId()
		if err != nil {
			return 0, false
		}

		pfo := NewPlayerFashionObject(m.p)
		pfo.Id = id
		pfo.PlayerId = m.p.GetId()
		pfo.FashionId = fashionId
		pfo.IsExpire = 0
		pfo.Star = 0
		pfo.UpStarNum = 0
		pfo.UpStarPro = 0
		pfo.ActiveTime = now
		pfo.CreateTime = now
		pfo.SetModified()
		fashionTypeMap[fashionId] = pfo
	}

	m.TrialFashionOverdue(fashionId, fashiontypes.FashionTrialOverdueTypeActivate)
	if sendEvent {
		gameevent.Emit(fashioneventtypes.EventTypeFashionActivate, m.p, fashionId)
	}
	return now, true
}

//时装穿戴
func (m *PlayerFashionDataManager) FashionWear(fashionId int32) bool {
	flag := m.HasedWeared(fashionId)
	if flag {
		return false
	}

	flag = m.IfFashionWear(fashionId)
	if !flag {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	m.PlayerFashionWearObject.FashionWear = fashionId
	m.PlayerFashionWearObject.UpdateTime = now
	m.PlayerFashionWearObject.SetModified()

	gameevent.Emit(fashioneventtypes.EventTypeFashionChanged, m.p, nil)
	return true
}

//获取出生时装
func (m *PlayerFashionDataManager) GetBornFashion() int32 {
	playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(m.p.GetRole(), m.p.GetSex())
	return playerCreateTemplate.FashionId
}

//时装卸下
func (m *PlayerFashionDataManager) Unload() {
	bornfashion := m.GetBornFashion()
	if m.PlayerFashionWearObject.FashionWear == bornfashion {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.PlayerFashionWearObject.FashionWear = bornfashion
	m.PlayerFashionWearObject.UpdateTime = now
	m.PlayerFashionWearObject.SetModified()

	gameevent.Emit(fashioneventtypes.EventTypeFashionChanged, m.p, nil)
	return
}

//是否能升星
func (m *PlayerFashionDataManager) IfCanUpStar(fashionId int32) bool {
	fashionTemplate := fashion.GetFashionService().GetFashionTemplate(int(fashionId))
	if fashionTemplate == nil {
		return false
	}
	fashionType := fashionTemplate.GetFashionType()
	flag := m.IfFashionExist(fashionId)
	if !flag {
		return false
	}
	if fashionTemplate.FashionUpgradeBeginId == 0 {
		return false
	}

	fashionObj := m.getFashionById(fashionType, fashionId)
	if fashionObj == nil {
		return false
	}
	star := fashionObj.Star
	if star <= 0 {
		return true
	}
	nextTo := fashionTemplate.GetFashionUpstarByLevel(star)
	if nextTo.NextId != 0 {
		return true
	}
	return false
}

//获取时装
func (m *PlayerFashionDataManager) GetFashion(fashionId int32) *PlayerFashionObject {
	fashionTemplate := fashion.GetFashionService().GetFashionTemplate(int(fashionId))
	if fashionTemplate == nil {
		return nil
	}
	fashionType := fashionTemplate.GetFashionType()
	return m.getFashionById(fashionType, fashionId)
}

//时装升星
func (m *PlayerFashionDataManager) Upstar(fashionId int32, pro int32, sucess bool) bool {
	flag := m.IfCanUpStar(fashionId)
	if !flag {
		return false
	}
	obj := m.GetFashion(fashionId)
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		obj.Star += 1
		obj.UpStarNum = 0
		obj.UpStarPro = pro
	} else {
		obj.UpStarNum += 1
		obj.UpStarPro += pro
	}
	obj.UpdateTime = now
	obj.SetModified()
	return true
}

//获取所有时效时装
func (m *PlayerFashionDataManager) GetExpireMap() map[int32]*PlayerFashionObject {
	fashionExpireMap, exist := m.fashionMap[fashiontypes.FashionTypeEffective]
	if !exist {
		return nil
	}
	return fashionExpireMap
}

//移除时效时装
func (m *PlayerFashionDataManager) RemoveExpireFashion(fashionId int32) {
	fashionTypeMap, exist := m.fashionMap[fashiontypes.FashionTypeEffective]
	if !exist {
		return
	}
	fashionObj, exist := fashionTypeMap[fashionId]
	if !exist {
		return
	}
	fashionObj.IsExpire = 1
	fashionObj.ActiveTime = 0
	fashionObj.SetModified()

	m.noExpireFahionMap[fashionId] = fashionObj
	delete(fashionTypeMap, fashionId)

	fashionWear := m.GetFashionId()
	if fashionWear == fashionId {
		m.Unload()
	}
	gameevent.Emit(fashioneventtypes.EventTypeFashionOverdue, m.p, fashionId)
}

func (m *PlayerFashionDataManager) removeCampFashion() {
	fashionTypeMap := m.fashionMap[fashiontypes.FashionTypeCamp]
	for _, fashionObj := range fashionTypeMap {
		fashionObj.IsExpire = 1
		fashionObj.ActiveTime = 0
		fashionObj.SetModified()

		m.noExpireFahionMap[fashionObj.FashionId] = fashionObj
		delete(fashionTypeMap, fashionObj.FashionId)
	}

	return
}

// 是否时装试用
func (m *PlayerFashionDataManager) IsFashionTrial(fashionId int32) bool {
	trialObj, isExist := m.playerFashionTrialMap[fashionId]
	if isExist {
		now := global.GetGame().GetTimeService().Now()
		if now < trialObj.expireTime {
			return true
		}

		return false
	}

	return false
}

// 时装试用
func (m *PlayerFashionDataManager) UseFashionTrialCard(trialCardItemId int32) {
	itemTemplate := item.GetItemService().GetItem(int(trialCardItemId))
	if itemTemplate == nil {
		return
	}

	trialFashionId := itemTemplate.TypeFlag1
	expireTime := int64(itemTemplate.TypeFlag2)
	now := global.GetGame().GetTimeService().Now()

	trialObj := m.getTrialFashionObj(trialFashionId)
	if trialObj == nil {
		trialObj = NewPlayerFashionTrialObject(m.p)
		id, _ := idutil.GetId()
		trialObj.id = id
		trialObj.trialFashionId = trialFashionId
		trialObj.createTime = now

		m.playerFashionTrialMap[trialFashionId] = trialObj
	} else {
		trialObj.updateTime = now
	}

	trialObj.expireTime = now + expireTime
	trialObj.SetModified()
}

func (m *PlayerFashionDataManager) getTrialFashionObj(trialFashionId int32) *PlayerFashionTrialObject {
	trialObj, ok := m.playerFashionTrialMap[trialFashionId]
	if ok {
		return trialObj
	}
	return nil
}

// 时装试用过期
func (m *PlayerFashionDataManager) TrialFashionOverdue(trialId int32, overdueType fashiontypes.FashionTrialOverdueType) {
	trialObj := m.getTrialFashionObj(trialId)
	if trialObj == nil {
		return
	}

	if trialObj.expireTime == 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	trialObj.expireTime = 0
	trialObj.updateTime = now
	trialObj.SetModified()

	curWear := m.GetFashionId()
	if curWear == trialId {
		m.Unload()
	}

	eventData := fashioneventtypes.CreateFashionTrialOverdueEventData(trialId, overdueType)
	gameevent.Emit(fashioneventtypes.EventTypeFashionTrialOverdue, m.p, eventData)
}

//时装试用map信息
func (m *PlayerFashionDataManager) GetTrialFashionMap() map[int32]*PlayerFashionTrialObject {
	newMap := make(map[int32]*PlayerFashionTrialObject)
	for fashionId, trialObj := range m.playerFashionTrialMap {
		if trialObj.expireTime == 0 {
			continue
		}

		newMap[fashionId] = trialObj
	}
	return newMap
}

//阵营时装激活
func (m *PlayerFashionDataManager) CampFashionActivate(fashionId int32) (flag bool) {
	m.removeCampFashion()

	_, flag = m.FashionActive(fashionId, false)
	if !flag {
		return
	}

	flag = m.FashionWear(fashionId)
	return
}

//仅gm 使用 激活时装
func (m *PlayerFashionDataManager) GmFashionActive(fashionId int32) (activeTime int64, flag bool) {
	flag = m.IsValid(fashionId)
	if !flag {
		return 0, false
	}
	fashionTemplate := fashion.GetFashionService().GetFashionTemplate(int(fashionId))
	fashionType := fashionTemplate.GetFashionType()
	flag = m.IfFashionExist(fashionId)
	if flag {
		return 0, false
	}

	now := global.GetGame().GetTimeService().Now()
	fashionTypeMap, exist := m.fashionMap[fashionType]
	if !exist {
		fashionTypeMap = make(map[int32]*PlayerFashionObject)
		m.fashionMap[fashionType] = fashionTypeMap
	}
	fashionObj, exist := m.noExpireFahionMap[fashionId]
	if exist {
		fashionObj.IsExpire = 0
		fashionObj.ActiveTime = now
		fashionObj.Star = 0
		fashionObj.UpStarNum = 0
		fashionObj.UpStarPro = 0
		fashionObj.SetModified()
		fashionTypeMap[fashionId] = fashionObj
		delete(m.noExpireFahionMap, fashionId)
	} else {
		id, err := idutil.GetId()
		if err != nil {
			return 0, false
		}

		pfo := NewPlayerFashionObject(m.p)
		pfo.Id = id
		pfo.PlayerId = m.p.GetId()
		pfo.FashionId = fashionId
		pfo.IsExpire = 0
		pfo.Star = 0
		pfo.UpStarNum = 0
		pfo.UpStarPro = 0
		pfo.ActiveTime = now
		pfo.CreateTime = now
		pfo.SetModified()
		fashionTypeMap[fashionId] = pfo
	}

	m.TrialFashionOverdue(fashionId, fashiontypes.FashionTrialOverdueTypeActivate)
	gameevent.Emit(fashioneventtypes.EventTypeFashionActivate, m.p, fashionId)
	return now, true
}

func CreatePlayerFashionDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerFashionDataManager{}
	m.p = p
	m.fashionMap = make(map[fashiontypes.FashionType]map[int32]*PlayerFashionObject)
	m.playerFashionTrialMap = make(map[int32]*PlayerFashionTrialObject)
	m.noExpireFahionMap = make(map[int32]*PlayerFashionObject)
	m.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerFashionDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerFashionDataManager))
}
