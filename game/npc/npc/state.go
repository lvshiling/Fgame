package npc

import (
	"fgame/fgame/core/fsm"
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//状态管理器
type NPCStateManager struct {
	//npc
	n scene.NPC
	//状态
	state fsm.State
	//定时器
	hbRunner heartbeat.HeartbeatTaskRunner
	//当前行为
	action scene.NPCAction
}

func NewNPCStateManager(n scene.NPC, scriptType scenetypes.BiologyScriptType) *NPCStateManager {
	npob := &NPCStateManager{}
	npob.n = n
	npob.state = scene.NPCStateInit
	npob.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	npob.action = scene.GetAction(scriptType, scene.NPCStateInit)
	npob.hbRunner.AddTask(CreateNPCTask(npob.n))
	return npob
}

func (m *NPCStateManager) Heartbeat() {
	m.hbRunner.Heartbeat()
}

//当前状态
func (m *NPCStateManager) CurrentState() fsm.State {
	return m.state
}

func (m *NPCStateManager) OnEnter(state fsm.State) {
	m.n.PauseMove()
	m.state = state
	//设置行为
	//通过脚本类型和状态获取行为
	m.action = scene.GetAction(m.n.GetBiologyTemplate().GetBiologyScriptType(), state)
}

func (m *NPCStateManager) OnExit(state fsm.State) {
	//清除行为
	m.action = nil
}

func (m *NPCStateManager) Trace() bool {
	flag := scene.GetNPCStateMachine().Trigger(m, scene.EventNPCTrace)
	if !flag {
		return false
	}

	return true
}

func (m *NPCStateManager) Dead() bool {
	flag := scene.GetNPCStateMachine().Trigger(m, scene.EventNPCDead)
	if !flag {
		return false
	}

	return true
}

func (m *NPCStateManager) Attack() bool {
	flag := scene.GetNPCStateMachine().Trigger(m, scene.EventNPCAttack)
	if !flag {
		return false
	}

	return true
}

func (m *NPCStateManager) Idle() bool {
	flag := scene.GetNPCStateMachine().Trigger(m, scene.EventNPCIdle)
	if !flag {
		return false
	}
	return true
}

//返回
func (m *NPCStateManager) Back() bool {
	flag := scene.GetNPCStateMachine().Trigger(m, scene.EventNPCBack)
	if !flag {
		return false
	}

	return true
}

//被攻击时位移
func (m *NPCStateManager) AttackedMove() bool {
	flag := scene.GetNPCStateMachine().Trigger(m, scene.EventNPCAttacked)
	if !flag {
		return false
	}

	return true
}

//获取主人
func (m *NPCStateManager) GetCurrentAction() scene.NPCAction {
	return m.action
}
