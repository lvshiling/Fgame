package model

import (
	"fgame/fgame/game/robot/robot"
	robottypes "fgame/fgame/game/robot/types"
	"fgame/fgame/game/scene/scene"
)

func init() {
	robot.RegisterActionFactory(robottypes.RobotTypeModel, robot.RobotPlayerStateIdle, robot.RobotActionFactoryFunc(newRobbotModelAction))
	robot.RegisterActionFactory(robottypes.RobotTypeModel, robot.RobotPlayerStateTrace, robot.RobotActionFactoryFunc(newRobbotModelAction))
	robot.RegisterActionFactory(robottypes.RobotTypeModel, robot.RobotPlayerStateAttack, robot.RobotActionFactoryFunc(newRobbotModelAction))
	robot.RegisterActionFactory(robottypes.RobotTypeModel, robot.RobotPlayerStateAttacked, robot.RobotActionFactoryFunc(newRobbotModelAction))
	robot.RegisterActionFactory(robottypes.RobotTypeModel, robot.RobotPlayerStateRun, robot.RobotActionFactoryFunc(newRobbotModelAction))
	robot.RegisterActionFactory(robottypes.RobotTypeModel, robot.RobotPlayerStateDead, robot.RobotActionFactoryFunc(newRobbotModelAction))
}

func newRobbotModelAction() scene.RobotAction {
	return robot.NewDummyAction()
}
