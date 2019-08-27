package use

import (
	"fgame/fgame/common/lang"
	pbutil "fgame/fgame/game/inventory/pbutil"
	playerinventory "fgame/fgame/game/inventory/player"
	item "fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	types "fgame/fgame/game/player/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeExpendBagSlotCard, itemtypes.ItemDefaultSubTypeDefault, playerinventory.ItemUseHandleFunc(handleExpandSlotCard))
}

// 背包扩充符
func handleExpandSlotCard(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	openNum := itemTemplate.TypeFlag1 * num
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	lockNum := manager.NumOfRemainBuySlots()
	if lockNum <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("inventory:使用背包扩充符,不能扩充槽位了")
		playerlogic.SendSystemMessage(pl, lang.InventoryCanNotAddSlot)
		return
	}

	if lockNum < openNum {
		diffNum := openNum - lockNum
		if itemTemplate.TypeFlag1 < diffNum {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"itemId":   itemId,
					"num":      num,
				}).Warn("inventory:使用背包扩充符,使用物品数量错误")
			playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
			return
		}
		openNum = lockNum
	}

	ok := manager.AddSlots(openNum)
	if !ok {
		panic(fmt.Errorf("inventory:add slots should be ok"))
	}
	numStr := fmt.Sprintf("%d", num)
	openNumStr := fmt.Sprintf("%d", openNum)
	playerlogic.SendSystemMessage(pl, lang.InventoryUseExpendSlotCardSucceed, numStr, itemTemplate.Name, openNumStr)
	slots := manager.GetSlots()
	depotSlots := manager.GetDepotSlots()
	scInventorySlots := pbutil.BuildSCInventorySlots(slots, depotSlots)
	pl.SendMsg(scInventorySlots)
	flag = true
	return
}
