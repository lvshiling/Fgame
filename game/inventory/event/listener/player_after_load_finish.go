package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/inventory/pbutil"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	//同步背包的格子数、仓库格子数
	p := target.(player.Player)
	inventory := p.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	slots := inventory.GetSlots()
	depotSlots := inventory.GetDepotSlots()
	scInventorySlots := pbutil.BuildSCInventorySlots(slots, depotSlots)
	p.SendMsg(scInventorySlots)

	//退还装备宝石
	slotList := inventory.GetEquipmentSlots()
	returnItemMap := make(map[int32]int32)
	for _, slot := range slotList {
		if len(slot.GemInfo) < 0 {
			continue
		}

		for _, itemId := range slot.GemInfo {
			_, ok := returnItemMap[itemId]
			if ok {
				returnItemMap[itemId] += 1
			} else {
				returnItemMap[itemId] = 1
			}
		}
	}

	if len(returnItemMap) > 0 {
		inventory.ClearAllEquipmentGemInfo()

		title := lang.GetLangService().ReadLang(lang.InventoryClearEquipmentGemReurnMailTitle)
		content := lang.GetLangService().ReadLang(lang.InventoryClearEquipmentGemReurnMailContent)
		emaillogic.AddEmail(p, title, content, returnItemMap)
	}

	//推送所有物品
	items := inventory.GetAll()
	depotList := inventory.GetDepotAll()
	itemUseMap := inventory.GetItemUseAll()
	miBaoDepotList := inventory.GetMiBaoDepotAll()
	materialDepotList := inventory.GetMaterialDepotAll()
	scInventoryGetAll := pbutil.BuildSCInventoryGetAll(items, depotList, slotList, itemUseMap, miBaoDepotList, materialDepotList)
	p.SendMsg(scInventoryGetAll)

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
