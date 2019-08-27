package ai

import (
	"fgame/fgame/game/robot/robot"
	robottypes "fgame/fgame/game/robot/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	robot.RegisterActionFactory(robottypes.RobotTypeTeamCopy, robot.RobotPlayerStateIdle, robot.RobotActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*robot.DummyAction
}

//随机时间 追击
func (a *idleAction) OnEnter() {

}

func (a *idleAction) OnExit() {

}

func (a *idleAction) Action(p scene.RobotPlayer) {
	s := p.GetScene()
	if s == nil {
		return
	}
	e := scenelogic.FindHatestEnemy(p)

	if e != nil {
		p.SetAttackTarget(e.BattleObject)
		p.Trace()
		return
	}

	//查找默认目标
	bo := p.GetDefaultAttackTarget()
	if bo != nil {
		p.SetAttackTarget(bo)
		p.Trace()
		return
	}
	//判断是否正在移动
	if p.IsMove() {
		return
	}
	pos := s.MapTemplate().GetMap().RandomPosition()
	flag := p.SetDestPosition(pos)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Warn("teamcopy:机器人找不到路")
		return
	}

}

func newIdleAction() scene.RobotAction {
	a := &idleAction{}
	a.DummyAction = robot.NewDummyAction()
	return a
}
