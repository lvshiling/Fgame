package player

import (
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	chargetemplate "fgame/fgame/game/charge/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"

	"fgame/fgame/game/charge/dao"
)

//玩家充值管理器
type PlayerChargeDataManager struct {
	p player.Player
	//充值记录
	playerChargeList []*PlayerChargeObject
	//扶持充值记录
	privilegePlayerChargeList []*PlayerPrivilegeChargeObject
	//档次首充记录
	playerFirstChargeRecordMap map[int32]*PlayerFirstChargeRecordObject
	//每日充值记录
	playerCycleChargeRecord *PlayerCycleChargeRecordObject
	//新档次首充记录
	playerNewFirstChargeRecord *PlayerNewFirstChargeRecordObject
}

func (m *PlayerChargeDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerChargeDataManager) Load() (err error) {
	//加载玩家充值信息
	chargeEntityList, err := dao.GetChargeDao().GetPlayerChargeEntityList(m.p.GetId())
	if err != nil {
		return
	}

	for _, chargeEntity := range chargeEntityList {
		obj := NewPlayerChargeObject(m.p)
		obj.FromEntity(chargeEntity)
		m.playerChargeList = append(m.playerChargeList, obj)
	}

	//加载玩家充值信息
	chargePrivilegeEntityList, err := dao.GetChargeDao().GetPlayerPrivilegeChargeEntityList(m.p.GetId())
	if err != nil {
		return
	}

	for _, chargePrivilegeEntity := range chargePrivilegeEntityList {
		obj := NewPlayerPrivilegeChargeObject(m.p)
		obj.FromEntity(chargePrivilegeEntity)
		m.privilegePlayerChargeList = append(m.privilegePlayerChargeList, obj)
	}

	//加载玩家档次首充信息
	firstRecordEntityList, err := dao.GetChargeDao().GetPlayerFirstChargeRecordEntityList(m.p.GetId())
	if err != nil {
		return
	}

	for _, entity := range firstRecordEntityList {
		obj := NewPlayerFirstChargeRecordObject(m.p)
		obj.FromEntity(entity)

		m.playerFirstChargeRecordMap[obj.chargeId] = obj
	}
	//加载玩家档次首充信息
	cycleRecordEntity, err := dao.GetChargeDao().GetPlayerCycleChargeRecordEntity(m.p.GetId())
	if err != nil {
		return
	}

	if cycleRecordEntity != nil {
		obj := NewPlayerCycleChargeRecordObject(m.p)
		obj.FromEntity(cycleRecordEntity)
		m.playerCycleChargeRecord = obj
	} else {
		m.initCycleChargeObject()
	}
	//加载玩家新档次首充信息
	newFirstRecordEntity, err := dao.GetChargeDao().GetPlayerNewFirstChargeRecordEntity(m.p.GetId())
	if err != nil {
		return
	}

	if newFirstRecordEntity != nil {
		obj := NewPlayerNewFirstChargeRecordObject(m.p)
		obj.FromEntity(newFirstRecordEntity)
		m.playerNewFirstChargeRecord = obj
	} else {
		m.initNewFisrtChargeObject()
	}

	return nil
}

//加载后
func (m *PlayerChargeDataManager) AfterLoad() (err error) {
	m.refreshCycleCharge()
	return
}

//心跳
func (m *PlayerChargeDataManager) Heartbeat() {

}

//充值
func (m *PlayerChargeDataManager) AddCharge(orderId string, chargeId int32) bool {
	chargeTmep := chargetemplate.GetChargeTemplateService().GetChargeTemplate(chargeId)
	if chargeTmep == nil {
		return false
	}
	orderObj := m.getChargeByOrderId(orderId)
	if orderObj != nil {
		return true
	}
	chargeGold := chargeTmep.Gold
	now := global.GetGame().GetTimeService().Now()
	orderObj = NewPlayerChargeObject(m.p)
	id, _ := idutil.GetId()
	orderObj.id = id
	orderObj.orderId = orderId
	orderObj.chargeType = chargeTmep.Type
	orderObj.chargeId = chargeId
	orderObj.chargeNum = chargeGold
	orderObj.createTime = now
	orderObj.SetModified()
	m.playerChargeList = append(m.playerChargeList, orderObj)
	m.addFirstChargeRecord(chargeId)

	//每日充值
	m.updateCycleChargeInfo(int64(chargeGold))

	//充值事件
	gameevent.Emit(chargeeventtypes.ChargeEventTypeChargeGold, m.p, chargeGold)
	eventData := chargeeventtypes.CreatePlayerChargeSuccessEventData(chargeId, chargeGold)
	gameevent.Emit(chargeeventtypes.ChargeEventTypeChargeSuccess, m.p, eventData)
	return true
}

//后台充值(废弃)
func (m *PlayerChargeDataManager) AddPrivilegeCharge(chargeGold int32) {
	m.updateCycleChargeInfo(int64(chargeGold))

	//充值事件
	gameevent.Emit(chargeeventtypes.ChargeEventTypeChargeGold, m.p, chargeGold)
	return
}

//后台充值
func (m *PlayerChargeDataManager) AddPrivilegeChargeId(chargeGold int32, chargeId int32) {
	chargeTmep := chargetemplate.GetChargeTemplateService().GetChargeTemplate(chargeId)
	if chargeTmep == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	orderObj := NewPlayerPrivilegeChargeObject(m.p)
	id, _ := idutil.GetId()
	orderObj.id = id
	orderObj.chargeType = chargeTmep.Type
	orderObj.chargeId = chargeId
	orderObj.chargeNum = chargeGold
	orderObj.createTime = now
	orderObj.SetModified()
	m.privilegePlayerChargeList = append(m.privilegePlayerChargeList, orderObj)

	m.addFirstChargeRecord(chargeId)
	//每日充值
	m.updateCycleChargeInfo(int64(chargeGold))
	//充值事件
	gameevent.Emit(chargeeventtypes.ChargeEventTypeChargeGold, m.p, chargeGold)
	eventData := chargeeventtypes.CreatePlayerChargeSuccessEventData(chargeId, chargeGold)
	gameevent.Emit(chargeeventtypes.ChargeEventTypeChargeSuccess, m.p, eventData)
	return
}

func (m *PlayerChargeDataManager) getChargeByOrderId(orderId string) *PlayerChargeObject {
	for _, obj := range m.playerChargeList {
		if obj.orderId == orderId {
			return obj
		}
	}
	return nil
}

//是否档次充值记录
func (m *PlayerChargeDataManager) IsHadFirstChargeRecord(chargeId int32) bool {
	_, ok := m.playerFirstChargeRecordMap[chargeId]
	if !ok {
		return false
	}
	return true
}

//添加档次充值记录
func (m *PlayerChargeDataManager) addFirstChargeRecord(chargeId int32) {
	chargeTmep := chargetemplate.GetChargeTemplateService().GetChargeTemplate(chargeId)
	_, ok := m.playerFirstChargeRecordMap[chargeId]
	if !ok {
		now := global.GetGame().GetTimeService().Now()
		newObj := NewPlayerFirstChargeRecordObject(m.p)
		id, _ := idutil.GetId()
		newObj.id = id
		newObj.chargeType = chargeTmep.Type
		newObj.chargeId = chargeId
		newObj.createTime = now
		newObj.SetModified()
		m.playerFirstChargeRecordMap[chargeId] = newObj
	}
}

//充值档次首充
func (m *PlayerChargeDataManager) ResetFirstChargeRecord(newChargeTime int64) {
	now := global.GetGame().GetTimeService().Now()
	for chargeId, chargeRecord := range m.playerFirstChargeRecordMap {
		if chargeRecord.createTime < newChargeTime {
			chargeRecord.deleteTime = now
			chargeRecord.SetModified()
			delete(m.playerFirstChargeRecordMap, chargeId)
		}
	}
}

//获取档次充值记录
func (m *PlayerChargeDataManager) GetFirstChargeRecord() map[int32]*PlayerFirstChargeRecordObject {
	return m.playerFirstChargeRecordMap
}

//更新每日充值信息
func (m *PlayerChargeDataManager) updateCycleChargeInfo(chargeNum int64) {
	m.refreshCycleCharge()

	if m.playerCycleChargeRecord.chargeNum == 0 {
		gameevent.Emit(chargeeventtypes.ChargeEventTypeFirstCycleCharge, m.p, chargeNum)
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerCycleChargeRecord.chargeNum += chargeNum
	m.playerCycleChargeRecord.updateTime = now
	m.playerCycleChargeRecord.SetModified()
}

// 判断是否有新首充领取记录
func (m *PlayerChargeDataManager) IsHasNewFirstChargeRecord(chargeId int32, startTime int64) bool {
	// 活动时间不一致，清空记录并更新信息
	if m.playerNewFirstChargeRecord.startTime != startTime {
		m.clearNewFirstChargeRecrodAndUpdate(startTime)
		return false
	}

	for _, record := range m.playerNewFirstChargeRecord.record {
		if record == chargeId {
			return true
		}
	}
	return false
}

// 清空新首充领取信息,更新开始时间
func (m *PlayerChargeDataManager) clearNewFirstChargeRecrodAndUpdate(startTime int64) {
	now := global.GetGame().GetTimeService().Now()
	if len(m.playerNewFirstChargeRecord.record) != 0 {
		m.playerNewFirstChargeRecord.record = m.playerNewFirstChargeRecord.record[:0]
	}
	m.playerNewFirstChargeRecord.startTime = startTime
	m.playerNewFirstChargeRecord.updateTime = now
	m.playerNewFirstChargeRecord.SetModified()
}

// 更新新首充领取信息
func (m *PlayerChargeDataManager) UpdateNewFirstChargeRecord(chargeId int32, startTime int64) bool {
	now := global.GetGame().GetTimeService().Now()

	// if m.playerNewFirstChargeRecord.startTime != startTime {
	// 	return false
	// }

	// // 判断是否领取过
	// for _, record := range m.playerNewFirstChargeRecord.record {
	// 	if record == chargeId {
	// 		return false
	// 	}
	// }

	if m.IsHasNewFirstChargeRecord(chargeId, startTime) {
		return false
	}

	m.playerNewFirstChargeRecord.record = append(m.playerNewFirstChargeRecord.record, chargeId)
	m.playerNewFirstChargeRecord.updateTime = now
	m.playerNewFirstChargeRecord.SetModified()
	return true
}

//刷新每日充值信息
func (m *PlayerChargeDataManager) refreshCycleCharge() {
	now := global.GetGame().GetTimeService().Now()

	diff, _ := timeutils.DiffDay(now, m.playerCycleChargeRecord.updateTime)
	if diff != 0 {
		if diff == 1 {
			m.playerCycleChargeRecord.preDayChargeNum = m.playerCycleChargeRecord.chargeNum
		} else {
			m.playerCycleChargeRecord.preDayChargeNum = 0
		}
		m.playerCycleChargeRecord.chargeNum = 0
		m.playerCycleChargeRecord.updateTime = now
		m.playerCycleChargeRecord.SetModified()
	}
}

//初始化每日首充对象
func (m *PlayerChargeDataManager) initCycleChargeObject() {
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()

	obj := NewPlayerCycleChargeRecordObject(m.p)
	obj.id = id
	obj.chargeNum = 0
	obj.preDayChargeNum = 0
	obj.createTime = now
	obj.updateTime = now
	obj.SetModified()
	m.playerCycleChargeRecord = obj
}

func (m *PlayerChargeDataManager) initNewFisrtChargeObject() {
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj := NewPlayerNewFirstChargeRecordObject(m.p)
	obj.id = id
	obj.record = make([]int32, 0, 8)
	obj.createTime = now
	obj.updateTime = now
	obj.SetModified()
	m.playerNewFirstChargeRecord = obj
}

//获取今日充值数
func (m *PlayerChargeDataManager) GetTodayChargeNum() int64 {
	m.refreshCycleCharge()
	return m.playerCycleChargeRecord.chargeNum
}

//获取今日充值数
func (m *PlayerChargeDataManager) GetPreDayChargeNum() int64 {
	m.refreshCycleCharge()
	return m.playerCycleChargeRecord.preDayChargeNum
}

//获取新首充记录信息
func (m *PlayerChargeDataManager) GetNewFirstChargeRecordInfo(startTime int64) *PlayerNewFirstChargeRecordObject {
	if m.playerNewFirstChargeRecord.startTime != startTime {
		m.clearNewFirstChargeRecrodAndUpdate(startTime)
	}
	return m.playerNewFirstChargeRecord
}

//获取今日充值数
func (m *PlayerChargeDataManager) GetTodayChargeList() (chargeMap map[int32]int32) {
	chargeMap = make(map[int32]int32)
	now := global.GetGame().GetTimeService().Now()
	for _, chargeObj := range m.playerChargeList {
		isSame, _ := timeutils.IsSameDay(now, chargeObj.createTime)
		if !isSame {
			continue
		}
		chargeMap[chargeObj.chargeNum] += 1
	}
	for _, chargeObj := range m.privilegePlayerChargeList {
		isSame, _ := timeutils.IsSameDay(now, chargeObj.createTime)
		if !isSame {
			continue
		}
		chargeMap[chargeObj.chargeNum] += 1
	}
	return
}

// 获取今日最大单笔充值
func (m *PlayerChargeDataManager) GetTodayMaxSingleCharge() (maxSingle int32) {
	now := global.GetGame().GetTimeService().Now()
	for _, chargeObj := range m.playerChargeList {
		isSame, _ := timeutils.IsSameDay(now, chargeObj.createTime)
		if !isSame {
			continue
		}
		if chargeObj.chargeNum > maxSingle {
			maxSingle = chargeObj.chargeNum
		}
	}
	for _, chargeObj := range m.privilegePlayerChargeList {
		isSame, _ := timeutils.IsSameDay(now, chargeObj.createTime)
		if !isSame {
			continue
		}
		if chargeObj.chargeNum > maxSingle {
			maxSingle = chargeObj.chargeNum
		}
	}
	return
}

func CreatePlayerChargeDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerChargeDataManager{}
	m.p = p
	m.playerFirstChargeRecordMap = make(map[int32]*PlayerFirstChargeRecordObject)
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerChargeDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerChargeDataManager))
}
