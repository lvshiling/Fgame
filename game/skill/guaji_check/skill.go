package guaji_check

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/guaji/guaji"
	playerguaji "fgame/fgame/game/guaji/player"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	skilllogic "fgame/fgame/game/skill/logic"
	playerskill "fgame/fgame/game/skill/player"
)

func init() {
	guaji.RegisterGuaJiCheckHandler(guajitypes.GuaJiCheckTypeSkill, guaji.GuaJiCheckHandlerFunc(skillGuaJiCheck))
}

func skillGuaJiCheck(pl player.Player) {

	//升级检查
	guaJiSkillUpgradeCheck(pl)

}

func guaJiSkillUpgradeCheck(pl player.Player) {
	skillManager := pl.GetPlayerDataManager(types.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
	guaJiManager := pl.GetPlayerDataManager(types.PlayerGuaJiManagerType).(*playerguaji.PlayerGuaJiManager)

	skillIdMap := skillManager.CanUpgradeRoleSkills()
	if skillIdMap == nil {
		return
	}

	//升级所有技能所需银量
	totalCostSilver := int64(0)
	for skillId, level := range skillIdMap {
		costSilver := skilllogic.UpgradeConsumeSilver(skillId, level+1)
		totalCostSilver += int64(costSilver)
	}
	totalCostSilver = totalCostSilver + guaJiManager.GetRemainSilver()
	//银量是否足够
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	flag := propertyManager.HasEnoughSilver(totalCostSilver)
	if !flag {
		return
	}
	skilllogic.HandleSkillUpgradeAll(pl)
	playerlogic.SendSystemMessage(pl, lang.GuaJiSkillUpgrade)
	return
}
