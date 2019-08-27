package ai

import (
	"fgame/fgame/game/global"
	"fgame/fgame/game/robot/robot"
	"fgame/fgame/game/scene/scene"
)

func init() {
	robot.RegisterDefaultActionFactory(robot.RobotPlayerStateAttack, robot.RobotActionFactoryFunc(newAttackAction))
}

type attackAction struct {
	*robot.DummyAction
}

func (a *attackAction) Action(p scene.RobotPlayer) {
	now := global.GetGame().GetTimeService().Now()
	elapse := now - p.GetSkillTime()
	//没到时间
	if elapse < p.GetSkillActionTime() {
		return
	}
	p.Trace()
}

func newAttackAction() scene.RobotAction {
	a := &attackAction{}
	a.DummyAction = robot.NewDummyAction()
	return a
}
