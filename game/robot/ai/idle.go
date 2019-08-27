package ai

import (
	"fgame/fgame/game/robot/robot"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

func init() {
	robot.RegisterDefaultActionFactory(robot.RobotPlayerStateIdle, robot.RobotActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*robot.DummyAction
}

func (a *idleAction) Action(p scene.RobotPlayer) {
	//查找敌人
	e := scenelogic.FindHatestEnemy(p)

	if e == nil {
		//查找默认目标
		bo := p.GetDefaultAttackTarget()
		if bo == nil {
			return
		}
		p.SetAttackTarget(bo)
		p.Trace()
		return
	}
	p.SetAttackTarget(e.BattleObject)
	p.Trace()

	return
}

func newIdleAction() scene.RobotAction {
	a := &idleAction{}
	a.DummyAction = robot.NewDummyAction()
	return a
}
