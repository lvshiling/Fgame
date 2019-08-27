package guaji

import (
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeWorld, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*scene.DummyGuaJiAction
}

//世界地图挂机中
func (a *idleAction) GuaJi(p scene.Player) {
	//TODO 游荡
	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
