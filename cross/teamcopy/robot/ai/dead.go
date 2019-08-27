package ai

import (
	teamcopylogic "fgame/fgame/cross/teamcopy/logic"
	teamscene "fgame/fgame/cross/teamcopy/scene"
	"fgame/fgame/cross/teamcopy/teamcopy"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/global"
	"fgame/fgame/game/robot/robot"
	robottypes "fgame/fgame/game/robot/types"
	"fgame/fgame/game/scene/scene"
	"math/rand"
)

func init() {
	robot.RegisterActionFactory(robottypes.RobotTypeTeamCopy, robot.RobotPlayerStateDead, robot.RobotActionFactoryFunc(newDeadAction))
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
	s := p.GetScene()
	if s == nil {
		return
	}
	sd, ok := s.SceneDelegate().(teamscene.TeamCopySceneData)
	if !ok {
		return
	}

	reliveTime := sd.GetReliveTime(p)

	//超过随机复活次数
	if reliveTime >= p.GetCanReliveTime() {
		//放弃
		teamcopy.GetTeamCopyService().TeamCopyMemeberGiveUp(p)
		//退出跨服
		p.BackLastScene()
		return
	}

	now := global.GetGame().GetTimeService().Now()
	elapse := now - p.GetDeadTime()
	if elapse >= int64(a.reliveWaitTime) {
		teamcopylogic.Reborn(sd, p)
	}
}

func newDeadAction() scene.RobotAction {
	a := &deadAction{}
	a.DummyAction = robot.NewDummyAction()
	return a
}
