package item

import (
	"fgame/fgame/common/lang"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"

	// playertypes"fgame/fgame/game/player/types"

	// playerproperty "fgame/fgame/game/property/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypePkValue, itemtypes.ItemDefaultSubTypeDefault, playerinventory.ItemUseHandleFunc(handlerPkValueReduce))
}

//红名值清除道具
func handlerPkValueReduce(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	if num <= 0 {
		return
	}

	if pl.GetPkValue() == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("pk:红名清除道具使用，红名值为0")
		playerlogic.SendSystemMessage(pl, lang.PlayerNotRedState)
		return
	}

	itemTemplate := item.GetItemService().GetItem(int(itemId))
	clearNum := itemTemplate.TypeFlag1 * num
	pl.ReducePkValue(clearNum)

	flag = true
	return
}
