package use

import (
	"fgame/fgame/common/lang"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	boxlogic "fgame/fgame/game/treasurebox/logic"
	treasureboxtypes "fgame/fgame/game/treasurebox/types"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeGiftBag, itemtypes.ItemGiftBagSubTypeCostChoose, playerinventory.ItemUseHandleFunc(handleCostChooseBoxUse))
}

// 可选消耗宝箱
func handleCostChooseBoxUse(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	costTypeInt, err := strconv.ParseInt(args, 10, 64)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"args":     args,
				"err":      err,
			}).Warn("box:使用可选消耗宝箱，参数格式化错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	costType := treasureboxtypes.BoxCostType(int32(costTypeInt))
	if !costType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"args":     args,
			}).Warn("box:使用可选消耗宝箱，参数格式化错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	startBoxTemplate := itemTemplate.GetBoxTemplateByCostType(costType)
	if startBoxTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"args":     args,
			}).Warn("box:使用可选消耗宝箱，宝箱模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	return boxlogic.OpenBox(pl, itemId, num, chooseIndexList, startBoxTemplate)
}
