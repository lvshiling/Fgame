package action

import (
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterDefaultGuaJiActionFactory(battle.PlayerStateRun, scene.GuaJiActionFactoryFunc(newRunAction))
}

type runAction struct {
	*scene.DummyGuaJiAction
}

func (a *runAction) GuaJi(p scene.Player) {
	//TODO 回到出生点
	log.Infof("玩家逃跑中")
	return
}

func newRunAction() scene.GuaJiAction {
	a := &runAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
