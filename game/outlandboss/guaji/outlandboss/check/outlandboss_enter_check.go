package check

import (
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	outlandbosslogic "fgame/fgame/game/outlandboss/logic"
	"fgame/fgame/game/outlandboss/outlandboss"
	"fgame/fgame/game/player"
)

func init() {
	guaji.RegisterGuaJiEnterCheckHandler(guajitypes.GuaJiTypeOutlandBoss, guaji.GuaJiEnterCheckHandlerFunc(outbossEnterCheck))
}

func outbossEnterCheck(pl player.Player) bool {
	//获取幻境boss
	outlandBossList := outlandboss.GetOutlandBossService().GetGuaiJiOutlandBossList(pl.GetForce())
	lenOfOutlandbossList := len(outlandBossList)
	if lenOfOutlandbossList <= 0 {
		return false
	}

	for _, wb := range outlandBossList {
		if outlandbosslogic.CheckPlayerIfCanOutlandbossChallenge(pl, int32(wb.GetBiologyTemplate().TemplateId())) {
			return true
		}
	}
	return false
}
