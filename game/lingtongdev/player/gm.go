package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	lingtongdeveventtypes "fgame/fgame/game/lingtongdev/event/types"
	"fgame/fgame/game/lingtongdev/types"
	"sort"
)

//仅gm使用 灵童养成进阶
func (m *PlayerLingTongDevDataManager) GmSetLingTongDevAdvanced(classType types.LingTongDevSysType, advancedId int) {
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		obj = m.initPlayerLingTongDevObj(classType)
	}
	obj.advanceId = advancedId
	obj.timesNum = 0
	obj.bless = 0
	obj.blessTime = 0
	obj.seqId = 0
	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
	gameevent.Emit(lingtongdeveventtypes.EventTypeLingTongDevChanged, m.p, classType)
	gameevent.Emit(lingtongdeveventtypes.EventTypeLingTongDevAdvanced, m.p, obj)
	return
}

//仅gm使用 灵童养成幻化
func (m *PlayerLingTongDevDataManager) GmSetLingTongDevUnreal(classType types.LingTongDevSysType, seqId int) {
	now := global.GetGame().GetTimeService().Now()
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		obj = m.initPlayerLingTongDevObj(classType)
	}
	if !m.IsUnrealed(classType, seqId) {
		obj.unrealList = append(obj.unrealList, seqId)
		sort.Ints(obj.unrealList)
		obj.updateTime = now
		obj.SetModified()
	}

	obj.seqId = int32(seqId)
	obj.updateTime = now
	obj.SetModified()
	gameevent.Emit(lingtongdeveventtypes.EventTypeLingTongDevChanged, m.p, classType)
	gameevent.Emit(lingtongdeveventtypes.EventTypeLingTongDevUnrealActivate, m.p, obj)
}

//仅gm使用 灵童养成食幻化丹
func (m *PlayerLingTongDevDataManager) GmSetLingTongDevUnrealDanLevel(classType types.LingTongDevSysType, level int32) {
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		obj = m.initPlayerLingTongDevObj(classType)
	}

	obj.unrealLevel = level
	obj.unrealNum = 0
	obj.unrealPro = 0
	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
}

//仅gm使用 灵童养成培养
func (m *PlayerLingTongDevDataManager) GmSetLingTongDevCulLevel(classType types.LingTongDevSysType, level int32) {
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		obj = m.initPlayerLingTongDevObj(classType)
	}
	obj.culLevel = level
	obj.culNum = 0
	obj.culPro = 0

	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
}

//仅gm使用 灵童养成通灵
func (m *PlayerLingTongDevDataManager) GmSetLingTongDevTongLing(classType types.LingTongDevSysType, level int32) {
	obj := m.getLingTongInfo(classType)
	if obj == nil {
		obj = m.initPlayerLingTongDevObj(classType)
	}
	obj.tongLingLevel = level
	obj.tongLingNum = 0
	obj.tongLingPro = 0

	now := global.GetGame().GetTimeService().Now()
	obj.updateTime = now
	obj.SetModified()
}
