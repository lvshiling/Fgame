package guaji

import (
	crossplayerlogic "fgame/fgame/cross/player/logic"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/game/battle/battle"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func init() {
	scene.RegisterGuaJiActionFactory(scenetypes.GuaJiTypeCross, battle.PlayerStateIdle, scene.GuaJiActionFactoryFunc(newIdleAction))
}

type idleAction struct {
	*scene.DummyGuaJiAction
}

//跨服挂机中
func (a *idleAction) GuaJi(p scene.Player) {
	//TODO 游荡
	pl, ok := p.(*player.Player)
	if !ok {
		return
	}
	//退出跨服
	crossplayerlogic.ExitCross(pl)
	return
}

func newIdleAction() scene.GuaJiAction {
	a := &idleAction{}
	a.DummyGuaJiAction = scene.NewDummyGuaJiAction()
	return a
}
