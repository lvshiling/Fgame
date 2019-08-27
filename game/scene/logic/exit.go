package logic

import (
	"fgame/fgame/game/scene/scene"

	log "github.com/Sirupsen/logrus"
)

func PlayerExitScene(pl scene.Player, active bool) bool {
	s := pl.GetScene()
	if s == nil {
		return true
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Info("player:正在退出场景")
	flag := pl.LeaveScene()
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("scene:玩家退出场景失败")
		return false
	}

	// //同步一下视野 防止下个场景出现看不到玩家
	// PlayerSyncNeighbors(pl)
	// //同步加载过的玩家
	// PlayerSyncLoadedPlayers(pl)
	s.RemoveSceneObject(pl, active)
	return true
}
