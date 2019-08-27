package check

import (
	"fgame/fgame/game/guaji/guaji"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	towerlogic "fgame/fgame/game/tower/logic"
	playertower "fgame/fgame/game/tower/player"
	towertempalte "fgame/fgame/game/tower/template"
)

func init() {
	guaji.RegisterGuaJiEnterCheckHandler(guajitypes.GuaJiTypeDaBaoTa, guaji.GuaJiEnterCheckHandlerFunc(daBaoTaEnterCheck))
}

const (
	defaultFloor = 1
)

func daBaoTaEnterCheck(pl player.Player) bool {
	//TODO 检查功能开启
	towerManager := pl.GetPlayerDataManager(playertypes.PlayerTowerDataManagerType).(*playertower.PlayerTowerDataManager)

	remainTime := towerManager.GetRemainTime()
	if remainTime <= 0 {
		return false
	}

	//检查 打宝时间
	towerTemplate := towertempalte.GetTowerTemplateService().GetRecommentTower(pl.GetLevel())
	if towerTemplate == nil {
		return false
	}
	floor := int32(towerTemplate.TemplateId())
	if towerlogic.CheckIfCanEnterTower(pl, floor) {
		return true
	}
	return towerlogic.CheckIfCanEnterTower(pl, defaultFloor)
}
