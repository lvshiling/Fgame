package boss_handler

import (
	"fgame/fgame/game/outlandboss/outlandboss"
	playeroutlandboss "fgame/fgame/game/outlandboss/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/game/worldboss/worldboss"
)

func init() {
	worldboss.RegistBossListHandler(worldbosstypes.BossTypeOutlandBoss, worldboss.BossListHandlerFunc(outBossList))
}

func outBossList(pl player.Player, typ worldbosstypes.BossType) {
	bossList := outlandboss.GetOutlandBossService().GetOutlandBossList()
	outlandbossManager := pl.GetPlayerDataManager(playertypes.PlayerOutlandBossDataManagerType).(*playeroutlandboss.PlayerOutlandBossDataManager)
	outlandbossManager.RefreshZhuoQi()

	curZhuoQi := outlandbossManager.GetCurZhuoQiNum()
	scMsg := pbutil.BuildSCWorldBossListOutBoss(bossList, curZhuoQi, int32(typ))
	pl.SendMsg(scMsg)
}
