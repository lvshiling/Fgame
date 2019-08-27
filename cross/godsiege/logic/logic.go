package logic

import (
	"fgame/fgame/cross/godsiege/pbutil"
	"fgame/fgame/cross/player/player"
	gamegodsiegepbutil "fgame/fgame/game/godsiege/pbutil"
)

func BroadGodSiegeLineUpChanged(godType int32, pos int32, lineList []int64) {
	for index, playerId := range lineList {
		if int32(index) < pos {
			continue
		}
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl == nil {
			continue
		}
		scGodSiegeLineUpChanged := gamegodsiegepbutil.BuildSCGodSiegeLineUp(godType, int32(index))
		pl.SendMsg(scGodSiegeLineUpChanged)
	}
}

func BroadGodSiegeFinishToLineUp(godType int32, lineList []int64) {
	for _, playerId := range lineList {
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl == nil {
			continue
		}
		isGodSiegeFinshToLineUp := pbutil.BuildISGodSiegeFinshToLineUp(godType)
		pl.SendMsg(isGodSiegeFinshToLineUp)
	}
}
