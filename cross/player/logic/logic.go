package logic

import (
	"fgame/fgame/cross/player/pbutil"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

func ExitCross(pl scene.Player) {
	scenelogic.PlayerExitScene(pl, true)
	isPlayerExitCross := pbutil.BuildISPlayerExitCross()
	pl.SendMsg(isPlayerExitCross)
	pl.Close(nil)
}
