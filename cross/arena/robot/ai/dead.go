package ai

import (
	"fgame/fgame/cross/arena/arena"
	arenalogic "fgame/fgame/cross/arena/logic"
	arenatemplate "fgame/fgame/game/arena/template"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/global"
	"fgame/fgame/game/robot/robot"
	robottypes "fgame/fgame/game/robot/types"
	"fgame/fgame/game/scene/scene"
	"math/rand"
)

func init() {
	robot.RegisterActionFactory(robottypes.RobotTypeArena, robot.RobotPlayerStateDead, robot.RobotActionFactoryFunc(newDeadAction))
}

var (
	minTimeRelive = 3 * common.SECOND
	maxTimeRelive = 10 * common.SECOND
)

type deadAction struct {
	*robot.DummyAction
	reliveWaitTime int64
}

func (a *deadAction) OnEnter() {
	a.reliveWaitTime = rand.Int63n(int64(maxTimeRelive-minTimeRelive)) + int64(minTimeRelive)
	return
}

func (a *deadAction) OnExit() {
	return
}

func (a *deadAction) Action(p scene.RobotPlayer) {

	remainReliveTime := arenatemplate.GetArenaTemplateService().GetArenaConstantTemplate().RebornAmountMax - p.GetArenaReliveTime()
	//剩余复活次数
	if remainReliveTime <= 0 {
		//放弃
		arena.GetArenaService().ArenaMemeberGiveUp(p)
		//退出跨服
		p.BackLastScene()
		return
	}

	//超过随机复活次数
	if p.GetArenaReliveTime() >= p.GetCanReliveTime() {
		//放弃
		arena.GetArenaService().ArenaMemeberGiveUp(p)
		//退出跨服
		p.BackLastScene()
		return
	}

	now := global.GetGame().GetTimeService().Now()
	elapse := now - p.GetDeadTime()
	if elapse >= int64(a.reliveWaitTime) {
		arenalogic.Reborn(p)
	}
}

func newDeadAction() scene.RobotAction {
	a := &deadAction{}
	a.DummyAction = robot.NewDummyAction()
	return a
}
