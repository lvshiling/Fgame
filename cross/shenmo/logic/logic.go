package logic

import (
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/shenmo/pbutil"
	"fgame/fgame/cross/shenmo/shenmo"
	"fgame/fgame/game/scene/scene"
)

func BroadShenMoLineUpChanged(pos int32, lineList []int64) {
	for index, playerId := range lineList {
		if int32(index) < pos {
			continue
		}
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl == nil {
			continue
		}
		scShenMoLineUpChanged := pbutil.BuildSCShenMoLineUpChanged(int32(index))
		pl.SendMsg(scShenMoLineUpChanged)
	}
}

func BroadShenMoFinishToLineUpCancle(lineList []int64) {
	for _, playerId := range lineList {
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl == nil {
			continue
		}
		isShenMoFinshToLineUp := pbutil.BuildISShenMoFinishLineUpCancle()
		pl.SendMsg(isShenMoFinshToLineUp)
	}
}

func JiFenChangedAllianceBroadcast(s scene.Scene, allianceId int64) {
	if s == nil {
		return
	}

	jiFenNum := shenmo.GetShenMoService().GetJiFenNum(allianceId)
	if jiFenNum == 0 {
		return
	}

	scShenMoJiFenNumChanged := pbutil.BuildSCShenMoJiFenNumChanged(jiFenNum)
	allPlayers := s.GetAllPlayers()
	for _, pl := range allPlayers {
		if pl == nil {
			continue
		}
		if pl.GetAllianceId() != allianceId {
			continue
		}
		pl.SendMsg(scShenMoJiFenNumChanged)
	}
}
