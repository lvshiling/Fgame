package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	xiantaoeventtypes "fgame/fgame/game/xiantao/event/types"
)

//玩家设置百年仙桃数量
func (pddm *PlayerXianTaoDataManager) GmSetJuniorPeachCount(count int32) {
	now := global.GetGame().GetTimeService().Now()
	pddm.xianTaoObject.JuniorPeachCount = count
	pddm.xianTaoObject.UpdateTime = now
	pddm.xianTaoObject.SetModified()
	gameevent.Emit(xiantaoeventtypes.EventTypeBaiNianXianTaoChange, pddm.p, nil)
	return
}

//玩家设置千年仙桃数量
func (pddm *PlayerXianTaoDataManager) GmSetHighPeachCount(count int32) {
	now := global.GetGame().GetTimeService().Now()
	pddm.xianTaoObject.HighPeachCount = count
	pddm.xianTaoObject.UpdateTime = now
	pddm.xianTaoObject.SetModified()
	gameevent.Emit(xiantaoeventtypes.EventTypeQianNianXianTaoChange, pddm.p, nil)
	return
}

//玩家设置劫取次数
func (pddm *PlayerXianTaoDataManager) GmSetRobCount(count int32) {
	now := global.GetGame().GetTimeService().Now()
	pddm.xianTaoObject.RobCount = count
	pddm.xianTaoObject.UpdateTime = now
	pddm.xianTaoObject.SetModified()
	return
}

//玩家设置被劫取次数
func (pddm *PlayerXianTaoDataManager) GmSetBeRobCount(count int32) {
	now := global.GetGame().GetTimeService().Now()
	pddm.xianTaoObject.BeRobCount = count
	pddm.xianTaoObject.UpdateTime = now
	pddm.xianTaoObject.SetModified()
	return
}
