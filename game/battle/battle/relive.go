package battle

import (
	"fgame/fgame/core/heartbeat"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	relivecommon "fgame/fgame/game/relive/common"
	"fgame/fgame/game/scene/scene"
)

type PlayerReliveManager struct {
	p              scene.Player
	culTime        int32
	lastReliveTime int64
	runner         heartbeat.HeartbeatTaskRunner
}

func (m *PlayerReliveManager) RefreshReliveTime() bool {
	now := global.GetGame().GetTimeService().Now()
	elapse := now - m.lastReliveTime
	clearTime := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeReliveTimesClearTime)
	//复活清空数据
	if elapse >= int64(clearTime) {
		m.culTime = 0
		m.lastReliveTime = now
		//发送刷新事件
		gameevent.Emit(battleeventtypes.EventTypeBattlePlayerReliveRefresh, m.p, nil)
		return true
	}
	return false
}

//获取累计复活次数
func (m *PlayerReliveManager) GetCulReliveTime() int32 {
	return m.culTime
}

func (m *PlayerReliveManager) GetLastReliveTime() int64 {
	return m.lastReliveTime
}

//心跳
func (m *PlayerReliveManager) Heartbeat() {
	m.runner.Heartbeat()
}

//复活
func (m *PlayerReliveManager) Relive() {
	now := global.GetGame().GetTimeService().Now()
	m.culTime += 1
	m.lastReliveTime = now
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerRelive, m.p, nil)
}

func (m *PlayerReliveManager) SyncRelive(reliveTime int32, lastReliveTime int64) {
	m.culTime = reliveTime
	m.lastReliveTime = lastReliveTime
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerReliveSync, m.p, nil)
}

func CreatePlayerReliveManager(p scene.Player, culTime int32, lastReliveTime int64) *PlayerReliveManager {
	m := &PlayerReliveManager{}
	m.p = p
	m.culTime = culTime
	m.lastReliveTime = lastReliveTime
	m.runner = heartbeat.NewHeartbeatTaskRunner()
	m.runner.AddTask(CreateReliveTask(p))
	return m
}

func CreatePlayerReliveManagerWithObj(p scene.Player, obj *relivecommon.PlayerReliveObject) *PlayerReliveManager {
	m := &PlayerReliveManager{}
	m.p = p
	m.culTime = obj.GetCulTime()
	m.lastReliveTime = obj.GetLastReliveTime()
	m.runner = heartbeat.NewHeartbeatTaskRunner()
	m.runner.AddTask(CreateReliveTask(p))
	return m
}
