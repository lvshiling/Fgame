package battle

import (
	"fgame/fgame/core/heartbeat"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
)

//玩家血池管理器
type PlayerXueChiManager struct {
	p         scene.Player
	bloodLine int32
	blood     int64
	lastTime  int64
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
}

//心跳
func (m *PlayerXueChiManager) Heartbeat() {
	m.heartbeatRunner.Heartbeat()
}

func (m *PlayerXueChiManager) SetBloodLine(bloodLine int32) {
	m.bloodLine = bloodLine

	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerXueChiBloodLineChanged, m.p, nil)
}

func (m *PlayerXueChiManager) GetBloodLine() int32 {
	return m.bloodLine
}

func (m *PlayerXueChiManager) AddBlood(addBlood int64) {
	m.blood += addBlood
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerXueChiBloodAdd, m.p, addBlood)
}

func (m *PlayerXueChiManager) SyncBlood(blood int64, bloodLine int32) {
	m.blood = blood
	m.bloodLine = bloodLine
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerXueChiBloodSync, m.p, nil)
}

func (m *PlayerXueChiManager) GetBlood() int64 {
	return m.blood
}

func (m *PlayerXueChiManager) GetLastBloodTime() int64 {
	return m.lastTime
}

func (m *PlayerXueChiManager) RecoverHp(recover int64) {
	realRecover := recover
	now := global.GetGame().GetTimeService().Now()
	if m.blood < recover {
		realRecover = m.blood
	}
	m.blood -= realRecover
	m.lastTime = now
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerXueChiRecover, m.p, realRecover)
}

func CreatePlayerXueChiManager(p scene.Player, blood int64, bloodLine int32) *PlayerXueChiManager {
	m := &PlayerXueChiManager{}
	m.p = p
	m.blood = blood
	m.bloodLine = bloodLine
	m.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	xueChiTask := CreateXueChiTask(m.p)
	m.heartbeatRunner.AddTask(xueChiTask)
	return m
}
