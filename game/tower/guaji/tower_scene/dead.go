package tower_scene

import (
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/global"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"math/rand"
)

func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeTowerScene, battle.PlayerStateDead, scene.GuaJiActionFactoryFunc(newDeadAction))
}

type deadAction struct {
	*scene.DummyGuaJiAction
	deadTime int64
}

const (
	exitDeadTime = 1
	minDeadTime  = int64(5 * common.SECOND)
	maxDeadTIme  = int64(10 * common.SECOND)
)

func (a *deadAction) OnEnter() {
	a.deadTime = rand.Int63n(maxDeadTIme-minDeadTime) + minDeadTime
	return
}

func (a *deadAction) OnExit() {
	return
}

//打宝塔副本挂机中
func (a *deadAction) GuaJi(p scene.Player) {
	if p.GetGuaJiDeadTimes() >= exitDeadTime {
		p.BackLastScene()
		return
	}
	//做延长时间
	now := global.GetGame().GetTimeService().Now()
	if now-p.GetDeadTime() < a.deadTime {
		return
	}
	//TODO 自动复活不了
	scenelogic.AutoReborn(p)
	return
}

func newDeadAction() scene.GuaJiAction {
	a := &deadAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
