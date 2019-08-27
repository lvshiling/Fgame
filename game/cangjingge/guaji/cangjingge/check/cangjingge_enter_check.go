package check

import (
	"fgame/fgame/game/cangjingge/cangjingge"
	cangjinggelogic "fgame/fgame/game/cangjingge/logic"
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
)

func init() {
	guaji.RegisterGuaJiEnterCheckHandler(guajitypes.GuaJiTypeCangJingGe, guaji.GuaJiEnterCheckHandlerFunc(cangJingGeEnterCheck))
}

func cangJingGeEnterCheck(pl player.Player) bool {
	//获取藏经阁boss
	bossList := cangjingge.GetCangJingGeService().GetGuaiJiCangJingGeBossList(pl.GetForce())
	lenOfBossList := len(bossList)
	if lenOfBossList <= 0 {
		return false
	}

	for _, wb := range bossList {
		if cangjinggelogic.CheckPlayerIfCanCangJingGeBossChallenge(pl, int32(wb.GetBiologyTemplate().TemplateId())) {
			return true
		}
	}
	return false
}
