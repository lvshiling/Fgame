package boss_handler

import (
	dingshilogic "fgame/fgame/game/dingshi/logic"
	"fgame/fgame/game/player"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/game/worldboss/worldboss"
)

func init() {
	worldboss.RegistKillBossHandler(worldbosstypes.BossTypeDingShi, worldboss.KillBossHandlerFunc(killDingShiBoss))
}

func killDingShiBoss(pl player.Player, typ worldbosstypes.BossType, biologyId int32) {
	dingshilogic.HandleKillDingShiBoss(pl, biologyId)
}
