package use

import (
	"fgame/fgame/common/lang"
	playerinventory "fgame/fgame/game/inventory/player"
	item "fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	pbutil "fgame/fgame/game/itemskill/pbutil"
	playeritemskill "fgame/fgame/game/itemskill/player"
	itemskilltemplate "fgame/fgame/game/itemskill/template"
	itemskilltypes "fgame/fgame/game/itemskill/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	types "fgame/fgame/game/player/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeItemSkill, itemtypes.ItemDefaultSubTypeDefault, playerinventory.ItemUseHandleFunc(handleItemSkillBook))
}

// 物品技能书
func handleItemSkillBook(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	typ := itemskilltypes.ItemSkillType(itemTemplate.TypeFlag1)
	if !typ.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("itemskill:使用物品技能书,参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerItemSkillDataManagerType).(*playeritemskill.PlayerItemSkillDataManager)
	obj := manager.GetItemSkillObjctByTyp(typ)
	if obj == nil {
		flag = itemSkillActive(pl, typ, num)
	} else {
		flag = itemSkillUpgrade(pl, typ, num)
	}

	return
}

func itemSkillActive(pl player.Player, typ itemskilltypes.ItemSkillType, num int32) (flag bool) {
	skTemplate := itemskilltemplate.GetItemSkillTemplateService().GetItemSkillTemplateByTypeAndLevel(typ, num)
	if skTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("itemskill:使用物品技能书,参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerItemSkillDataManagerType).(*playeritemskill.PlayerItemSkillDataManager)
	obj, flag := manager.ItemSkillActive(typ, num)
	if !flag {
		return
	}

	scMsg := pbutil.BuildSCItemSkillActive(obj)
	pl.SendMsg(scMsg)
	return
}

func itemSkillUpgrade(pl player.Player, typ itemskilltypes.ItemSkillType, num int32) (flag bool) {
	manager := pl.GetPlayerDataManager(types.PlayerItemSkillDataManagerType).(*playeritemskill.PlayerItemSkillDataManager)
	level := manager.GetItemSkillLevelByTyp(typ)
	skTemplate := itemskilltemplate.GetItemSkillTemplateService().GetItemSkillTemplateByTypeAndLevel(typ, level+num)
	if skTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("itemskill:使用物品技能书,参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	obj, flag := manager.Upgrade(typ, num)
	if !flag {
		return
	}

	scMsg := pbutil.BuildSCItemSkillUpgrade(obj)
	pl.SendMsg(scMsg)
	return
}
