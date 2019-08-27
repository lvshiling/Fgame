package action

import (
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterDefaultGuaJiActionFactory(battle.PlayerStateAttacked, scene.GuaJiActionFactoryFunc(newAttackedAction))
}

type attackedAction struct {
	*scene.DummyGuaJiAction
}

func (a *attackedAction) GuaJi(p scene.Player) {
	log.Infof("玩家被攻击中")
	//TODO 停顿时间
}

func newAttackedAction() scene.GuaJiAction {
	a := &attackAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
