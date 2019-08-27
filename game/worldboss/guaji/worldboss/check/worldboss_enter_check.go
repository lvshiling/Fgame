package check

import (
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	worldbosslogic "fgame/fgame/game/worldboss/logic"
	"fgame/fgame/game/worldboss/worldboss"
)

func init() {
	guaji.RegisterGuaJiEnterCheckHandler(guajitypes.GuaJiTypeWorldboss, guaji.GuaJiEnterCheckHandlerFunc(worldbossEnterCheck))
}

func worldbossEnterCheck(pl player.Player) bool {
	return false
	//获取幻境boss
	worldBossList := worldboss.GetWorldBossService().GetGuaiJiWorldBossList(pl.GetForce())
	lenOfWorldBossList := len(worldBossList)
	if lenOfWorldBossList <= 0 {
		return false
	}

	for _, wb := range worldBossList {
		if worldbosslogic.CheckPlayerIfCanKillBoss(pl, int32(wb.GetBiologyTemplate().TemplateId())) {
			return true
		}
	}
	return false
}
