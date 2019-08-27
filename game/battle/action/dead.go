package action

import (
	"fgame/fgame/game/battle/battle"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func init() {
	scene.RegisterDefaultGuaJiActionFactory(battle.PlayerStateDead, scene.GuaJiActionFactoryFunc(newDeadAction))
}

type deadAction struct {
	*scene.DummyGuaJiAction
}

func (a *deadAction) GuaJi(p scene.Player) {
	log.Infof("玩家死亡")
	s := p.GetScene()
	if s == nil {
		return
	}
	log.Infof("玩家死亡[%d]", s.MapId())
	//TODO 随机时间 复活策略
	scenelogic.AutoReborn(p)
}

func newDeadAction() scene.GuaJiAction {
	a := &deadAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
