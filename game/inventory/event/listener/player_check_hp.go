package listener

// //玩家血量变化
// func playerHPCheck(target event.EventTarget, data event.EventData) (err error) {
// 	//同步背包的格子数、仓库格子数
// 	p := target.(player.Player)
// 	if p.IsDead() {
// 		return
// 	}
// 	//自动使用血药
// 	rate := float64(p.GetHP()) / float64(p.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP))
// 	if rate <= 0.7 {
// 		inventory := p.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
// 		flag := inventory.HasEnoughItem(11, 1)
// 		if !flag {
// 			return
// 		}
// 		if inventory.IsItemUseCd(11) {
// 			return
// 		}
// 		//查找血药
// 		reason := commonlog.InventoryLogReasonRecover
// 		reasonText := commonlog.InventoryLogReasonRecover.String()
// 		flag = inventory.UseItem2(11, 1, reason, reasonText)
// 		if !flag {
// 			panic(fmt.Errorf("inventory:使用血药错误"))
// 		}
// 		inventorylogic.SnapInventoryChanged(p)
// 	}
// 	// inventory := p.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
// 	// slots := inventory.GetSlots()
// 	// depotSlots := inventory.GetDepotSlots()
// 	// scInventorySlots := pbutil.BuildSCInventorySlots(slots, depotSlots)
// 	// p.SendMsg(scInventorySlots)

// 	// //推送所有物品
// 	// items := inventory.GetAll()
// 	// depotList := inventory.GetDepotAll()
// 	// slotList := inventory.GetEquipmentSlots()
// 	// itemUseMap := inventory.GetItemUseAll()
// 	// scInventoryGetAll := pbutil.BuildSCInventoryGetAll(items, depotList, slotList, itemUseMap)
// 	// p.SendMsg(scInventoryGetAll)
// 	return
// }

// func init() {
// 	gameevent.AddEventListener(battleeventtypes.EventTypePlayerCheckHP, event.EventListenerFunc(playerHPCheck))
// }
