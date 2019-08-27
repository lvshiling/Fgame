package boss_handler

import (
	"fgame/fgame/game/player"
	worldbosslogic "fgame/fgame/game/worldboss/logic"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/game/worldboss/worldboss"
)

func init() {
	worldboss.RegistKillBossHandler(worldbosstypes.BossTypeWorldBoss, worldboss.KillBossHandlerFunc(killWorldBoss))
}

func killWorldBoss(pl player.Player, typ worldbosstypes.BossType, biologyId int32) {
	worldbosslogic.HandleKillWorldBoss(pl, biologyId)
}
