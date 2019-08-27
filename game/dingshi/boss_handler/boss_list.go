package boss_handler

import (
	"fgame/fgame/game/dingshi/dingshi"
	"fgame/fgame/game/player"
	"fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/game/worldboss/worldboss"
)

func init() {
	worldboss.RegistBossListHandler(worldbosstypes.BossTypeDingShi, worldboss.BossListHandlerFunc(dingShiBossList))
}

func dingShiBossList(pl player.Player, typ worldbosstypes.BossType) {
	bossList := dingshi.GetDingShiService().GetDingShiBossList()
	scMsg := pbutil.BuildSCWorldBossList(bossList, int32(typ))
	pl.SendMsg(scMsg)
}
