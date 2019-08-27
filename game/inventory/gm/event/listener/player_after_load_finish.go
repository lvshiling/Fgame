package listener

// //加载完成后
// func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
// 	//设置
// 	if !global.PRESSURE {
// 		return
// 	}
// 	pl, ok := target.(player.Player)
// 	if !ok {
// 		return
// 	}
// 	bagTotalNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeBagTotalNum)
// 	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
// 	slots := manager.GetSlots()
// 	if slots < bagTotalNum {
// 		flag := manager.AddSlots(bagTotalNum - slots)
// 		if !flag {
// 			panic(fmt.Errorf("inventory:add slots should be ok"))
// 		}
// 	}

// 	primBagId := int32(105)
// 	needComplement := manager.RemainSlotForItem(primBagId)
// 	if needComplement > 0 {
// 		manager.AddItem(primBagId, needComplement, commonlog.InventoryLogReasonGM, commonlog.InventoryLogReasonGM.String())
// 	}
// 	gemBagId := int32(2101)
// 	gemNeedComplement := manager.RemainSlotForItem(gemBagId)
// 	if gemNeedComplement > 0 {
// 		manager.AddItem(gemBagId, gemNeedComplement, commonlog.InventoryLogReasonGM, commonlog.InventoryLogReasonGM.String())
// 	}
// 	kunBagId := int32(8001)
// 	kunNeedComplement := manager.RemainSlotForItem(kunBagId)
// 	if kunNeedComplement > 0 {
// 		manager.AddItem(kunBagId, kunNeedComplement, commonlog.InventoryLogReasonGM, commonlog.InventoryLogReasonGM.String())
// 	}
// 	needDepotComplement := manager.RemainDepotSlotForItem(primBagId, 0, itemtypes.ItemBindTypeUnBind)
// 	if needDepotComplement > 0 {
// 		manager.AddItemInDepot(primBagId, needDepotComplement, 0, itemtypes.ItemBindTypeUnBind, commonlog.InventoryLogReasonGM, commonlog.InventoryLogReasonGM.String())
// 	}
// 	return
// }

// func init() {
// 	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
// }
