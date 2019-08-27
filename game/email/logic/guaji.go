package logic

import (
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	playeremail "fgame/fgame/game/email/player"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

//领取附件逻辑
func CheckPlayerIfCanGetEmailAttachement(pl player.Player, emailId int64) (flag bool) {
	emailManager := pl.GetPlayerDataManager(types.PlayerEmailDataManagerType).(*playeremail.PlayerEmailDataManager)
	//验证参数

	_, emailObj := emailManager.GetEmail(emailId)
	if emailObj == nil {
		return false
	}
	//是否存在附件
	if emailManager.HasNotOrReceiveAttachment(emailId) {
		return false
	}

	var newItemList []*droptemplate.DropItemData

	if len(emailObj.GetAttachmentInfo()) != 0 {
		newItemList, _ = droplogic.SeperateItemDatas(emailObj.GetAttachmentInfo())
	}
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//是否足够背包空间
	if len(newItemList) > 0 {
		if !inventoryManager.HasEnoughSlotsOfItemLevel(newItemList) {
			return false
		}

	}

	return true
}

//检查是否有足够的位置
func CheckPlayerIfCanGetEmailAttachementBatch(pl player.Player) (flag bool) {
	emailManager := pl.GetPlayerDataManager(types.PlayerEmailDataManagerType).(*playeremail.PlayerEmailDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	emailObjArr := emailManager.GetNotReceiveAttachmentEmails()
	var totalAttacheList []*droptemplate.DropItemData
	if len(emailObjArr) > 0 {
		//获取所有物品map、资源map
		for _, emailObj := range emailObjArr {
			itemList := emailObj.GetAttachmentInfo()
			totalAttacheList = append(totalAttacheList, itemList...)
		}

		var totalItemList []*droptemplate.DropItemData

		if len(totalAttacheList) != 0 {
			totalItemList, _ = droplogic.SeperateItemDatas(totalAttacheList)
		}
		//物品加入背包
		if len(totalItemList) > 0 {
			//是否足够背包空间
			if !inventoryManager.HasEnoughSlotsOfItemLevel(totalItemList) {
				return false
			}
		}

	}

	return
}
