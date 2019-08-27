package guaji_check

// func init() {
// 	guaji.RegisterGuaJiCheckHandler(guajitypes.GuaJiCheckTypeMount, guaji.GuaJiCheckHandlerFunc(mountGuaJiCheck))
// }

// func mountGuaJiCheck(pl player.Player) {
// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Info("mount_guaji:挂机检查坐骑")
// 	//升级检查
// 	guaJiMountUpgradeCheck(pl)
// 	//进阶检查
// 	guaJiMountAdvancedCheck(pl)

// }

// func guaJiMountUpgradeCheck(pl player.Player) {

// }

// func guaJiMountAdvancedCheck(pl player.Player) {
// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Info("mount_guaji:挂机检查坐骑进阶")
// 	//TODO 荣昌 添加自动购买
// 	for {
// 		mountManager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
// 		mountInfo := mountManager.GetMountInfo()

// 		nextAdvancedId := mountInfo.AdvanceId + 1
// 		mountTemplate := mount.GetMountService().GetMountNumber(int32(nextAdvancedId))
// 		if mountTemplate == nil {
// 			return
// 		}

// 		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
// 		propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
// 		guaJiManager := pl.GetPlayerDataManager(types.PlayerGuaJiManagerType).(*playerguaji.PlayerGuaJiManager)

// 		//进阶需要消耗的元宝
// 		costGold := mountTemplate.UseMoney
// 		//进阶需要消耗的银两
// 		costSilver := int64(mountTemplate.UseYinliang)
// 		//进阶需要的消耗的绑元
// 		costBindGold := int32(0)

// 		//需要消耗物品
// 		itemCount := int32(0)
// 		totalNum := int32(0)
// 		useItem := mountTemplate.UseItem

// 		useItemTemplate := mountTemplate.GetUseItemTemplate()
// 		if useItemTemplate != nil {
// 			itemCount = mountTemplate.ItemCount
// 			totalNum = inventoryManager.NumOfItems(int32(useItem))
// 		}

// 		if totalNum < itemCount {
// 			return
// 		}

// 		//是否足够银两
// 		if costSilver != 0 {
// 			totalSilver := costSilver + guaJiManager.GetRemainSilver()
// 			flag := propertyManager.HasEnoughSilver(int64(totalSilver))
// 			if !flag {
// 				return
// 			}
// 		}
// 		//是否足够元宝
// 		if costGold != 0 {
// 			flag := propertyManager.HasEnoughGold(int64(costGold), false)
// 			if !flag {
// 				return
// 			}
// 		}

// 		//是否足够绑元
// 		needBindGold := costBindGold + costGold
// 		if needBindGold != 0 {
// 			flag := propertyManager.HasEnoughGold(int64(needBindGold), true)
// 			if !flag {
// 				return
// 			}
// 		}
// 		log.WithFields(
// 			log.Fields{
// 				"playerId":  pl.GetId(),
// 				"advanceId": nextAdvancedId,
// 			}).Info("mount_guaji:挂机检查坐骑,准备进阶")
// 		mountlogic.HandleMountAdvanced(pl, false)
// 		playerlogic.SendSystemMessage(pl, lang.GuaJiMountAdvanced)

// 	}
// }
