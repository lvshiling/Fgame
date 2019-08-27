package boss_handler

import (
	"fgame/fgame/game/player"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/game/worldboss/worldboss"
	zhenxilogic "fgame/fgame/game/zhenxi/logic"
)

func init() {
	worldboss.RegistKillBossHandler(worldbosstypes.BossTypeZhenXi, worldboss.KillBossHandlerFunc(killZhenXiBoss))
}

func killZhenXiBoss(pl player.Player, typ worldbosstypes.BossType, biologyId int32) {
	zhenxilogic.HandleKillZhenXiBoss(pl, biologyId)
}
