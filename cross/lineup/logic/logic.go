package logic

import (
	"fgame/fgame/cross/lineup/pbutil"
	"fgame/fgame/cross/player/player"
)

func BroadLineUpChanged(pos int32, lineList []int64, crossType int32) {
	for index, playerId := range lineList {
		if int32(index) < pos {
			continue
		}
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl == nil {
			continue
		}
		scMsg := pbutil.BuildSCLineupNotice(int32(index), crossType)
		pl.SendMsg(scMsg)
	}
}

func BroadLineUpFinishToCancel(crossType int32, lineList []int64) {
	for _, playerId := range lineList {
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl == nil {
			continue
		}
		scMsg := pbutil.BuildSCLineupSceneFinishToCancel(crossType)
		pl.SendMsg(scMsg)

		isMsg := pbutil.BuildISLineupSceneFinishToCancel(crossType)
		pl.SendMsg(isMsg)

	}
}
