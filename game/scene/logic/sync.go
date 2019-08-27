package logic

import (
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func PlayerSyncNeighbors(pl scene.Player) {

	enterScope, monsterList := pbutil.BuildEnterScope(pl)
	if enterScope != nil {
		//内挂不发节省流量
		if !pl.IsGuaJiPlayer() {
			pl.SendMsg(enterScope)
		}
	}

	//客户端要求再发一遍
	for _, n := range monsterList {
		if n.GetBiologyTemplate().GetBiologyType() != scenetypes.BiologyTypeBoss {
			continue
		}
		isEnemy := pl.IsEnemy(n)
		scMonsterCampChanged := pbutil.BuildSCMonsterCampChanged(n, isEnemy)
		pl.SendMsg(scMonsterCampChanged)
	}

	exitScope := pbutil.BuildExitScope(pl)
	if exitScope != nil {
		if !pl.IsGuaJiPlayer() {
			pl.SendMsg(exitScope)
		}
	}
}

func PlayerSyncLoadedPlayers(pl scene.Player) {
	scExitScene := pbutil.BuildExitScene(pl.GetId())
	for _, loadedPlayer := range pl.GetLoadedPlayers() {
		// log.WithFields(
		// 	log.Fields{
		// 		"玩家":    loadedPlayer.GetId(),
		// 		"退出的玩家": pl.GetId(),
		// 		"goId":  runtimeutils.Goid(),
		// 	}).Infoln("玩家退出场景")

		loadedPlayer.SendMsg(scExitScene)
	}
}

func LingTongSyncLoadedPlayers(pl scene.LingTong) {
	scExitScene := pbutil.BuildExitScene(pl.GetId())
	for _, loadedPlayer := range pl.GetLoadedPlayers() {
		loadedPlayer.SendMsg(scExitScene)
	}
}
