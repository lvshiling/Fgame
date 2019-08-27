package listener

import (
	"fgame/fgame/core/event"
	dropeventtypes "fgame/fgame/game/drop/event/types"
	gameevent "fgame/fgame/game/event"
)

//物品获得
func fubenDropItemsGet(target event.EventTarget, data event.EventData) (err error) {

	// pl, ok := target.(player.Player)
	// if !ok {
	// 	return
	// }

	// //发送事件
	// s := pl.GetScene()
	// if s == nil {
	// 	return
	// }
	// itemList := make([]*droptemplate.DropItemData, 0, 4)
	// for _, item := range s.GetAllItems() {
	// 	newItem := droptemplate.CreateItemData(item.GetItemId(), item.GetItemNum(), item.GetLevel(), itemtypes.ItemBindTypeUnBind)
	// 	itemList = append(itemList, newItem)
	// }

	// var rewItemList []*droptemplate.DropItemData
	// var resMap map[itemtypes.ItemAutoUseResSubType]int32
	// if len(itemList) > 0 {
	// 	rewItemList, resMap = droplogic.SeperateItemDatas(itemList)
	// }

	// inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	// if !inventoryManager.HasEnoughSlotsOfItemLevel(rewItemList) {
	// 	return
	// }

	// goldLog := commonlog.GoldLogReasonMonsterKilled
	// goldReasonText := goldLog.String()
	// silverLog := commonlog.SilverLogReasonMonsterKilled
	// silverReasonText := silverLog.String()
	// inventoryLog := commonlog.InventoryLogReasonMonsterKilled
	// inventoryReasonText := inventoryLog.String()
	// levelLog := commonlog.LevelLogReasonMonsterKilled
	// levelReasonText := levelLog.String()

	// flag, err := droplogic.AddItem(pl, itemId, num, level, bind, goldLog, goldReasonText, silverLog, silverReasonText, inventoryLog, inventoryReasonText, levelLog, levelReasonText)
	// if err != nil {
	// 	return
	// }

	// if !flag {
	// 	playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
	// 	return
	// }

	// //同步属性
	// propertylogic.SnapChangedProperty(pl)
	// inventorylogic.SnapInventoryChanged(pl)

	// s.OnPlayerGetItem(pl, dropItem.GetItemId(), dropItem.GetItemNum(), dropItem.GetLevel())
	// dropItem.GetScene().RemoveSceneObject(dropItem, false)
	return
}

func init() {
	gameevent.AddEventListener(dropeventtypes.EventTypeFubenDropItemsGet, event.EventListenerFunc(fubenDropItemsGet))
}
