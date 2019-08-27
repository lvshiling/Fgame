package boss_handler

import (
	"fgame/fgame/game/cangjingge/cangjingge"
	"fgame/fgame/game/player"
	"fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/game/worldboss/worldboss"
)

func init() {
	worldboss.RegistBossListHandler(worldbosstypes.BossTypeCangJingGe, worldboss.BossListHandlerFunc(cangJingGeBossList))
}

func cangJingGeBossList(pl player.Player, typ worldbosstypes.BossType) {
	bossList := cangjingge.GetCangJingGeService().GetCangJingGeBossList()
	scMsg := pbutil.BuildSCWorldBossList(bossList, int32(typ))
	pl.SendMsg(scMsg)
}
