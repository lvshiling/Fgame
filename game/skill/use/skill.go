package use

import (
	"fgame/fgame/common/lang"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fgame/fgame/game/skill/pbutil"
	playerskill "fgame/fgame/game/skill/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeSkill, itemtypes.ItemDefaultSubTypeDefault, playerinventory.ItemUseHandleFunc(handleItemUse))
}

func handleItemUse(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	//参数不对
	itemId := it.ItemId
	if num != 1 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("skill:添加技能,使用物品数量错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	skillTemplate := itemTemplate.GetSkillTemplate()
	skillManager := pl.GetPlayerDataManager(playertypes.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if skillTemplate.GetRoleType() != pl.GetRole() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("skill:添加技能,职业不符合")
		playerlogic.SendSystemMessage(pl, lang.PlayerRoleWrong)
		return
	}
	skillId := int32(skillTemplate.TemplateId())

	exist := skillManager.IfSkillExist(skillId)
	if exist {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"skillId":  skillId,
			}).Warn("skill:添加技能,技能存在")
		playerlogic.SendSystemMessage(pl, lang.SkillHasExist)
		return
	}
	flag = skillManager.AddSkill(skillId)
	if !flag {
		panic("skill:添加技能应该成功")
	}
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeSkill.Mask())
	scSkillLearn := pbutil.BuildSCSkillLearn(skillId, skillTemplate.Lev)
	pl.SendMsg(scSkillLearn)
	return
}
