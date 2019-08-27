package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fgame/fgame/game/skill/pbutil"
	playerskill "fgame/fgame/game/skill/player"
	skilltemplate "fgame/fgame/game/skill/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//变更技能属性
func SkillPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeSkill.Mask())
	return
}

//临时技能改变
func TempSkillChange(pl player.Player, oldSkillId int32, newSkillId int32) (err error) {
	if oldSkillId == newSkillId {
		return
	}

	flag := pl.ChangeSkill(oldSkillId, newSkillId)
	if !flag {
		return
	}

	SkillPropertyChanged(pl)

	return
}

//临时技能改变
func TempSkillChangeNoUpdate(pl player.Player, oldSkillId int32, newSkillId int32) (err error) {
	if oldSkillId == newSkillId {
		return
	}
	flag := pl.ChangeSkill(oldSkillId, newSkillId)
	if !flag {
		return
	}
	return
}

//升级技能消耗银量
func UpgradeConsumeSilver(skillId int32, level int32) (silver int32) {
	silver = 0
	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(skillId)
	if skillTemplate == nil {
		return silver
	}
	skillLevelTemplate := skillTemplate.GetSkillByLevel(level)
	silver = skillLevelTemplate.CostSilver
	return silver
}

//处理职业技能全部升级的逻辑
func HandleSkillUpgradeAll(pl player.Player) (err error) {
	skillManager := pl.GetPlayerDataManager(types.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
	skillIdMap := skillManager.CanUpgradeRoleSkills()
	if skillIdMap == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("skill:当前无可升级技能")
		playerlogic.SendSystemMessage(pl, lang.SKillNotHasUpgrade)
		return
	}

	//升级所有技能所需银量
	totalCostSilver := int64(0)
	for skillId, level := range skillIdMap {
		costSilver := UpgradeConsumeSilver(skillId, level+1)
		totalCostSilver += int64(costSilver)
	}

	//银量是否足够
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	flag := propertyManager.HasEnoughSilver(totalCostSilver)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("skill:银两不足，无法升级")
		playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
		return
	}

	//消耗银两
	reasonText := commonlog.SilverLogReasonSkillUpgradeAll.String()
	flag = propertyManager.CostSilver(totalCostSilver, commonlog.SilverLogReasonSkillUpgradeAll, reasonText)
	if !flag {
		panic(fmt.Errorf("skill: skillUpgrade cost sliver should be ok"))
	}
	propertylogic.SnapChangedProperty(pl)

	flag = skillManager.UpgradeSkillAll(skillIdMap)
	if !flag {
		panic(fmt.Errorf("skill: handleSkillUpgradeAll  should be ok"))
	}

	SkillPropertyChanged(pl)
	scSkillUpgradeAll := pbutil.BuildSCSkillUpgradeAll(skillIdMap)
	pl.SendMsg(scSkillUpgradeAll)
	return
}
