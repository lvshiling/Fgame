package lingtong

import (
	"fgame/fgame/core/fsm"
	"fgame/fgame/game/scene/scene"
)

type LingTongStateManager struct {
	l         scene.LingTong
	actionMap map[fsm.State]scene.LingTongAction
	state     fsm.State
	action    scene.LingTongAction
}

func (m *LingTongStateManager) CurrentState() fsm.State {
	return m.state
}

func (m *LingTongStateManager) OnEnter(state fsm.State) {
	m.l.PauseMove()
	m.state = state
	//设置行为
	//通过脚本类型和状态获取行为
	m.action = m.getAction()
}

func (m *LingTongStateManager) OnExit(state fsm.State) {
	//清除行为
	m.action = nil
}

func (m *LingTongStateManager) getAction() scene.LingTongAction {
	action, ok := m.actionMap[m.state]
	if !ok {
		action = scene.GetLingTongAction(m.state)
		m.actionMap[m.state] = action
	}
	return action
}

func (m *LingTongStateManager) Attack() bool {
	flag := scene.GetLingTongStateMachine().Trigger(m, scene.EventLingTongAttack)
	if !flag {
		return false
	}

	return true
}

func (m *LingTongStateManager) Idle() bool {
	flag := scene.GetLingTongStateMachine().Trigger(m, scene.EventLingTongIdle)
	if !flag {
		return false
	}
	return true
}

//返回
func (m *LingTongStateManager) Trace() bool {
	flag := scene.GetLingTongStateMachine().Trigger(m, scene.EventLingTongTrace)
	if !flag {
		return false
	}

	return true
}

func (m *LingTongStateManager) GetCurrentAction() scene.LingTongAction {
	return m.action
}

func CreateLingTongStateManager(l scene.LingTong) *LingTongStateManager {
	m := &LingTongStateManager{}
	m.l = l
	m.actionMap = make(map[fsm.State]scene.LingTongAction)
	m.state = scene.LingTongStateInit
	m.action = m.getAction()
	return m
}
