package use

import (
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/template"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gametemplate "fgame/fgame/game/template"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeUpLevelDan, itemtypes.ItemUpLevelDanSubTypeUpLevelDan, playerinventory.ItemUseHandleFunc(handleUplevelDan))
}

//等级直升丹
func handleUplevelDan(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	if num <= 0 {
		return
	}
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	beforeLevel := itemTemplate.TypeFlag3
	curLevel := pl.GetLevel()
	totalExp := int64(0)
	addLevel := false
	// TODO:xzk:修改计算方法
	for beforeLevel > curLevel {
		addLevel = true
		// 直接升级
		addLevel := itemTemplate.TypeFlag1
		for lev := int32(1); lev <= addLevel; lev++ {
			curLevel += 1
			tempLevelTemplate := template.GetTemplateService().Get(int(curLevel), (*gametemplate.CharacterLevelTemplate)(nil)).(*gametemplate.CharacterLevelTemplate)
			totalExp += int64(tempLevelTemplate.Experience)
		}

		num -= 1
		if num == 0 {
			break
		}
	}

	if addLevel {
		totalExp -= propertyManager.GetExp()
	}

	// 添加固定经验
	for i := int32(0); i < num; i++ {
		addExp := int64(itemTemplate.TypeFlag2)
		totalExp += addExp
	}

	reason := commonlog.LevelLogReasonEatUplevelDan
	reasonText := reason.String()
	propertyManager.AddExp(totalExp, reason, reasonText)

	propertylogic.SnapChangedProperty(pl)

	flag = true
	return
}
