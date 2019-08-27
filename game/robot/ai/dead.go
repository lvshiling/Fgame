package ai

import (
	robotlogic "fgame/fgame/game/robot/logic"
	"fgame/fgame/game/robot/robot"
	"fgame/fgame/game/scene/scene"
)

func init() {
	robot.RegisterDefaultActionFactory(robot.RobotPlayerStateDead, robot.RobotActionFactoryFunc(newDeadAction))
}

type deadAction struct {
	*robot.DummyAction
}

func (a *deadAction) Action(p scene.RobotPlayer) {
	robotlogic.RemoveRobot(p)
}

func newDeadAction() scene.RobotAction {
	a := &deadAction{}
	a.DummyAction = robot.NewDummyAction()
	return a
}
