package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	inventoryeventtypes "fgame/fgame/game/inventory/event/types"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	logmodel "fgame/fgame/logserver/model"
)

//物品变化
func playerItemChangedLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*inventoryeventtypes.PlayerInventoryChangedLogEventData)
	if !ok {
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	itemId := eventData.GetItemId()
	itemName := ""
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate != nil {
		itemName = itemTemplate.Name
	}
	beforeNum := eventData.GetBeforeItemNum()
	curNum := inventoryManager.NumOfItems(itemId)

	logItemChanged := &logmodel.PlayerItemChanged{}
	logItemChanged.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logItemChanged.ChangedItemId = itemId
	logItemChanged.ChangedItemName = itemName
	logItemChanged.ChangedItemNum = eventData.GetChangedNum()
	logItemChanged.BeforeItemNum = beforeNum
	logItemChanged.CurItemNum = curNum
	logItemChanged.Reason = int32(eventData.GetReason())
	logItemChanged.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logItemChanged)
	return
}

func init() {
	gameevent.AddEventListener(inventoryeventtypes.EventTypeInventoryChangedLog, event.EventListenerFunc(playerItemChangedLog))
}
