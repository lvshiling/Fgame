package boss_handler

import (
	cangjinggelogic "fgame/fgame/game/cangjingge/logic"
	"fgame/fgame/game/player"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/game/worldboss/worldboss"
)

func init() {
	worldboss.RegistKillBossHandler(worldbosstypes.BossTypeCangJingGe, worldboss.KillBossHandlerFunc(killCangJingGeBoss))
}

func killCangJingGeBoss(pl player.Player, typ worldbosstypes.BossType, biologyId int32) {
	cangjinggelogic.HandleKillCangJingGeBoss(pl, biologyId)
}
