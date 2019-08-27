package logic

import (
	gamelianyupbutil"fgame/fgame/game/lianyu/pbutil"
	crosslianyupbutil "fgame/fgame/cross/lianyu/pbutil"
	"fgame/fgame/cross/player/player"
)

func BroadLianYuLineUpChanged(pos int32, lineList []int64) {
	for index, playerId := range lineList {
		if int32(index) < pos {
			continue
		}
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl == nil {
			continue
		}
		scLianYuLineUpChanged := gamelianyupbutil.BuildSCLianYuLineUpChanged(int32(index))
		pl.SendMsg(scLianYuLineUpChanged)
	}
}

func BroadLineYuFinishToLineUp(lineList []int64) {
	for _, playerId := range lineList {
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl == nil {
			continue
		}
		isLianYuFinshToLineUp := crosslianyupbutil.BuildISLianYuFinshToLineUp()
		pl.SendMsg(isLianYuFinshToLineUp)
	}
}
