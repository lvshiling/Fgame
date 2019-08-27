package battle

import (
	battleeventtypes "fgame/fgame/game/battle/event/types"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

type PlayerZhenYingManager struct {
	p           scene.Player
	campType    chuangshitypes.ChuangShiCampType
	guanZhi     chuangshitypes.ChuangShiGuanZhi
	reliveTimes int32 //攻城原地复活次数
}

func (m *PlayerZhenYingManager) GetCamp() chuangshitypes.ChuangShiCampType {
	return m.campType
}

func (m *PlayerZhenYingManager) SetCamp(camp chuangshitypes.ChuangShiCampType) {
	m.campType = camp
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerCampChanged, m.p, nil)
}

func (m *PlayerZhenYingManager) GetGuanZhi() chuangshitypes.ChuangShiGuanZhi {
	return m.guanZhi
}

func (m *PlayerZhenYingManager) SetGuanZhi(guanZhi chuangshitypes.ChuangShiGuanZhi) {
	m.guanZhi = guanZhi
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerGuanZhiChanged, m.p, nil)
}

func (m *PlayerZhenYingManager) SetChuangShiReliveTimes(reliveTimes int32) {
	m.reliveTimes = reliveTimes
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerChuangShiReliveTimesChanged, m.p, nil)
}

func (m *PlayerZhenYingManager) GetChuangShiReliveTimes() int32 {
	return m.reliveTimes
}

func CreatePlayerZhenYingManager(p scene.Player, camp chuangshitypes.ChuangShiCampType, guanZhi chuangshitypes.ChuangShiGuanZhi) *PlayerZhenYingManager {
	m := &PlayerZhenYingManager{
		p:        p,
		campType: camp,
		guanZhi:  guanZhi,
	}
	return m
}
