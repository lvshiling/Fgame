package logic

import (
	crosslogic "fgame/fgame/game/cross/logic"
	crosstypes "fgame/fgame/game/cross/types"
	"fgame/fgame/game/lineup/pbutil"
	"fgame/fgame/game/player"
)

//跨服排队
func SendCrossLineup(pl player.Player, crossType crosstypes.CrossType, sceneId int64) {
	siMsg := pbutil.BuildSILineupAttend(int32(crossType), sceneId)
	pl.SendCrossMsg(siMsg)
}

//取消排队
func CancelCrossLineup(pl player.Player) {
	if pl.IsLineUp() {
		pl.SetLineUp(false)

		//退出跨服
		crosslogic.PlayerExitCross(pl)
	}
}
