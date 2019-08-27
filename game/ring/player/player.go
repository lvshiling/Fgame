package player

import (
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/event"
	"fgame/fgame/game/global"
	inventorytypes "fgame/fgame/game/inventory/types"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	ringcommon "fgame/fgame/game/ring/common"
	"fgame/fgame/game/ring/dao"
	ringeventtypes "fgame/fgame/game/ring/event/types"
	ringtemplate "fgame/fgame/game/ring/template"
	ringtypes "fgame/fgame/game/ring/types"
	viplogic "fgame/fgame/game/vip/logic"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
)

// 凌晨5点刷新
const (
	refreshHour int = 5
)

type PlayerRingDataManager struct {
	pl                    player.Player
	playerRingObjectMap   map[ringtypes.RingType]*PlayerRingObject
	playerRingBaoKuObject map[ringtypes.BaoKuType]*PlayerRingBaoKuObject
}

// 加载
func (m *PlayerRingDataManager) Load() (err error) {
	m.playerRingObjectMap = make(map[ringtypes.RingType]*PlayerRingObject)
	m.playerRingBaoKuObject = make(map[ringtypes.BaoKuType]*PlayerRingBaoKuObject)

	// 加载特戒数据
	ringEntityList, err := dao.GetRingDao().GetPlayerRingEntity(m.pl.GetId())
	if err != nil {
		return
	}
	for _, ringEntity := range ringEntityList {
		ringObj := NewRingObject(m.pl)
		ringObj.FromEntity(ringEntity)
		m.playerRingObjectMap[ringObj.typ] = ringObj
	}

	// 加载特权寻宝数据
	xunBaoEntityList, err := dao.GetRingDao().GetPlayerRingBaoKuEntity(m.pl.GetId())
	if err != nil {
		return
	}
	for _, ringEntity := range xunBaoEntityList {
		ringObj := NewRingBaoKuObject(m.pl)
		ringObj.FromEntity(ringEntity)
		m.playerRingBaoKuObject[ringObj.typ] = ringObj
	}

	return nil
}

func (m *PlayerRingDataManager) initRingBaoKuObject(typ ringtypes.BaoKuType) *PlayerRingBaoKuObject {
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj := NewRingBaoKuObject(m.pl)
	obj.id = id
	obj.typ = typ
	obj.lastSystemRefreshTime = now
	obj.createTime = now
	obj.SetModified()
	return obj
}

func (m *PlayerRingDataManager) Player() player.Player {
	return m.pl
}

//加载后
func (m *PlayerRingDataManager) AfterLoad() (err error) {
	m.RefreshLuckyPoints()
	return
}

//心跳
func (m *PlayerRingDataManager) Heartbeat() {
}

// 获取特戒信息
func (m *PlayerRingDataManager) GetPlayerRingObjectMap() map[ringtypes.RingType]*PlayerRingObject {
	return m.playerRingObjectMap
}

// 获取宝库信息
func (m *PlayerRingDataManager) GetPlayerBaoKuObject(typ ringtypes.BaoKuType) *PlayerRingBaoKuObject {
	m.RefreshLuckyPoints()
	obj, ok := m.playerRingBaoKuObject[typ]
	if !ok {
		obj = m.initRingBaoKuObject(typ)
		m.playerRingBaoKuObject[typ] = obj
	}
	return obj
}

// 根据类型获取特戒信息
func (m *PlayerRingDataManager) GetPlayerRingObject(typ ringtypes.RingType) *PlayerRingObject {
	return m.getPlayerRingObject(typ)
}

// 获取宝库掉落包
func (m *PlayerRingDataManager) GetRingBaoKuDrop(times int32, typ ringtypes.BaoKuType) (rewList []*droptemplate.DropItemData) {
	baoKuTemp := ringtemplate.GetRingTemplateService().GetRingBaoKuTemplate(typ)
	if baoKuTemp == nil {
		return
	}
	obj := m.GetPlayerBaoKuObject(typ)
	curTotal := obj.totalAttendTimes
	for index := int32(0); index < times; index++ {
		curTotal += 1
		dropId := baoKuTemp.DropId

		rewDropByTimesMap := baoKuTemp.GetRewDropMap()
		timesDescList := baoKuTemp.GetDropTimesDescList()

		vipType, flag := ringtypes.RingTypeToCostLevelRuleType(typ)
		if !flag {
			return
		}

		ruleTimesMap := viplogic.CountDropTimesWithCostLevel(m.pl, vipType, timesDescList)
		for _, times := range timesDescList {
			ruleTimes := ruleTimesMap[times]
			ret := curTotal % ruleTimes
			if ret == 0 {
				dropId = rewDropByTimesMap[int32(times)]
				break
			}
		}

		dropData := droptemplate.GetDropTemplateService().GetDropBaoKuItemLevel(dropId)
		if dropData != nil {
			rewList = append(rewList, dropData)
		}
	}

	return
}

// 寻宝完成
func (m *PlayerRingDataManager) AttendRingBaoKu(luckyVal, attendVal, attendCount int32, typ ringtypes.BaoKuType) (flag bool) {
	m.RefreshLuckyPoints()
	obj := m.GetPlayerBaoKuObject(typ)
	if luckyVal > 0 {
		obj.luckyPoints += luckyVal

		// 判断幸运值是否已满
		baoKuTemp := ringtemplate.GetRingTemplateService().GetRingBaoKuTemplate(typ)
		if baoKuTemp == nil {
			return
		}
		if obj.luckyPoints > baoKuTemp.NeedXingYunZhi {
			obj.luckyPoints -= baoKuTemp.NeedXingYunZhi
		}
	}

	if attendVal > 0 {
		obj.attendPoints += attendVal
	}

	if attendCount > 0 {
		obj.totalAttendTimes += attendCount
	}

	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
	flag = true
	return
}

//刷新装备宝库幸运值
func (m *PlayerRingDataManager) RefreshLuckyPoints() (err error) {
	now := global.GetGame().GetTimeService().Now()
	for _, obj := range m.playerRingBaoKuObject {
		isSame, err := timeutils.IsSameDay(obj.lastSystemRefreshTime, now)
		if err != nil {
			continue
		}
		hour := timeutils.MillisecondToTime(now).Hour()
		if !isSame && hour >= refreshHour && obj.luckyPoints != 0 {
			obj.luckyPoints = 0
			obj.lastSystemRefreshTime = now
			obj.updateTime = now
			obj.SetModified()

			event.Emit(ringeventtypes.EventTypeRingLuckyPointsChange, m.pl, obj.GetType())
		}
	}
	return
}

//消耗积分
func (m *PlayerRingDataManager) UseJiFen(typ ringtypes.BaoKuType, num int32) bool {
	m.RefreshLuckyPoints()
	obj := m.GetPlayerBaoKuObject(typ)
	if !obj.IfEnoughJiFen(num) {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	obj.attendPoints -= num
	obj.updateTime = now
	obj.SetModified()

	event.Emit(ringeventtypes.EventTypeRingAttendPointsChange, m.pl, typ)
	return true
}

// 特戒进阶成功
func (m *PlayerRingDataManager) RingAdvanceSuccess(typ ringtypes.RingType, success bool, pro int32) bool {
	obj := m.getPlayerRingObject(typ)
	if obj == nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	data := obj.propertyData.(*ringtypes.RingPropertyData)
	if success {
		data.Advance++

		// 判断祝福进度值是否清空
		advanceTemp := ringtemplate.GetRingTemplateService().GetRingAdvanceTemplate(obj.itemId, data.Advance)
		if advanceTemp == nil {
			return false
		}
		if advanceTemp.IsClearBless() {
			data.AdvancePro = 0
		}

		data.AdvanceNum = 0
	} else {
		data.AdvanceNum++
		data.AdvancePro += pro
	}

	obj.updateTime = now
	obj.SetModified()
	return true
}

// 特戒强化成功
func (m *PlayerRingDataManager) RingStrengthenSuccess(typ ringtypes.RingType, success bool) bool {
	obj := m.getPlayerRingObject(typ)
	if obj == nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	data := obj.propertyData.(*ringtypes.RingPropertyData)
	if success {
		data.StrengthLevel++
		data.StrengthNum = 0
	} else {
		data.StrengthNum++
	}

	obj.updateTime = now
	obj.SetModified()
	return true
}

// 特戒净灵成功
func (m *PlayerRingDataManager) RingJingLingSuccess(typ ringtypes.RingType, success bool) bool {
	obj := m.getPlayerRingObject(typ)
	if obj == nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	data := obj.propertyData.(*ringtypes.RingPropertyData)
	if success {
		data.JingLingLevel++
		data.JingLingNum = 0
	} else {
		data.JingLingNum++
	}

	obj.updateTime = now
	obj.SetModified()
	return true
}

// 特戒融合成功
func (m *PlayerRingDataManager) RingFuseSuccess(typ ringtypes.RingType, itemId int32) bool {
	obj := m.getPlayerRingObject(typ)
	if obj == nil {
		return false
	}
	lastItemId := obj.itemId
	now := global.GetGame().GetTimeService().Now()
	obj.itemId = itemId
	obj.updateTime = now
	obj.SetModified()

	data := ringeventtypes.CreatePlayerRingFuseChangeEventData(lastItemId, itemId)
	event.Emit(ringeventtypes.EventTypeRingFuseChange, m.pl, data)
	return true
}

// 特戒穿戴成功
func (m *PlayerRingDataManager) RingEquipSuccess(typ ringtypes.RingType, itemId int32, propertyData inventorytypes.ItemPropertyData, bindType itemtypes.ItemBindType) bool {
	obj := m.getPlayerRingObject(typ)
	if obj != nil {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj = NewRingObject(m.pl)
	obj.id = id
	obj.typ = typ
	obj.itemId = itemId
	obj.bindType = bindType
	obj.propertyData = propertyData
	obj.createTime = now
	obj.SetModified()
	m.playerRingObjectMap[typ] = obj

	data := ringeventtypes.CreatePlayerRingFuseChangeEventData(0, itemId)
	event.Emit(ringeventtypes.EventTypeRingFuseChange, m.pl, data)
	return true
}

// 卸下特戒成功
func (m *PlayerRingDataManager) RingUnloadSuccess(typ ringtypes.RingType) bool {
	obj := m.getPlayerRingObject(typ)
	if obj == nil {
		return false
	}
	lastItemId := obj.itemId

	now := global.GetGame().GetTimeService().Now()
	obj.deleteTime = now
	obj.SetModified()

	delete(m.playerRingObjectMap, typ)

	data := ringeventtypes.CreatePlayerRingFuseChangeEventData(lastItemId, 0)
	event.Emit(ringeventtypes.EventTypeRingFuseChange, m.pl, data)
	return true
}

func (m *PlayerRingDataManager) GetRingInfoList() []*ringcommon.RingInfo {
	ringInfoList := make([]*ringcommon.RingInfo, 0, 8)
	for _, obj := range m.playerRingObjectMap {
		info := &ringcommon.RingInfo{}
		info.Typ = obj.typ
		info.ItemId = obj.itemId
		data, _ := obj.propertyData.(*ringtypes.RingPropertyData)
		info.PropertyData = data
		ringInfoList = append(ringInfoList, info)
	}
	return ringInfoList
}

func (m *PlayerRingDataManager) getPlayerRingObject(typ ringtypes.RingType) *PlayerRingObject {
	obj, ok := m.playerRingObjectMap[typ]
	if !ok {
		return nil
	}
	return obj
}

func CreatePlayerRingDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerRingDataManager{}
	m.pl = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerRingDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerRingDataManager))
}
