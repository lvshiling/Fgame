package robot

import (
	"fgame/fgame/core/fsm"
	robottypes "fgame/fgame/game/robot/types"
	"fgame/fgame/game/scene/scene"
	"fmt"
)

type RobotActionFactory interface {
	CreateAction() scene.RobotAction
}

type RobotActionFactoryFunc func() scene.RobotAction

func (f RobotActionFactoryFunc) CreateAction() scene.RobotAction {
	return f()
}

type DummyAction struct {
}

func (a *DummyAction) OnEnter() {
	return
}

func (a *DummyAction) Action(p scene.RobotPlayer) {
	return
}

func (a *DummyAction) OnExit() {
	return
}

func NewDummyAction() *DummyAction {
	return &DummyAction{}
}

var (
	dummyActionInstance = NewDummyAction()
)

var (
	robotStateMap   = make(map[robottypes.RobotType]map[fsm.State]RobotActionFactory)
	defaultStateMap = make(map[fsm.State]RobotActionFactory)
)

func RegisterDefaultActionFactory(state fsm.State, action RobotActionFactory) {
	_, ok := defaultStateMap[state]
	if ok {
		panic(fmt.Errorf("repeate register default action  %d state script action", state))
	}
	defaultStateMap[state] = action
}

func RegisterActionFactory(typ robottypes.RobotType, state fsm.State, action RobotActionFactory) {
	stateMap, ok := robotStateMap[typ]
	if !ok {
		stateMap = make(map[fsm.State]RobotActionFactory)
		robotStateMap[typ] = stateMap
	}
	_, ok = stateMap[state]
	if ok {
		panic(fmt.Errorf("repeate register %d typ, %d state script action", typ, state))
	}
	stateMap[state] = action
}

func GetAction(typ robottypes.RobotType, state fsm.State) scene.RobotAction {
	stateMap, ok := robotStateMap[typ]
	if !ok {
		return getDefaultAction(state)
	}
	stateActionFactory, ok := stateMap[state]
	if !ok {
		return getDefaultAction(state)
	}
	action := stateActionFactory.CreateAction()
	return action
}

func getDefaultAction(state fsm.State) scene.RobotAction {
	stateActionFactory, ok := defaultStateMap[state]
	if !ok {
		return dummyActionInstance
	}
	action := stateActionFactory.CreateAction()
	return action
}
