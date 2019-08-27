package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	emaillogic "fgame/fgame/game/email/logic"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fmt"
)

//变更神龙属性
func DragonPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeDragon.Mask())

	return
}

//喂养物品奖励
func GiveDragonFeedReward(pl player.Player, rewItemId int32, num int32) {
	if rewItemId <= 0 {
		return
	}
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughSlot(rewItemId, num) {
		//写邮件
		rewItemMap := make(map[int32]int32)
		rewItemMap[rewItemId] = 1
		emailTitle := lang.GetLangService().ReadLang(lang.DragonAdvancedRewTitle)
		emailContent := lang.GetLangService().ReadLang(lang.DragonAdvancedRewContent)
		emaillogic.AddEmail(pl, emailTitle, emailContent, rewItemMap)
		return
	}
	reasonText := commonlog.InventoryLogReasonDragonAdvancedRew.String()
	if !inventoryManager.AddItem(rewItemId, num, commonlog.InventoryLogReasonDragonAdvancedRew, reasonText) {
		panic(fmt.Errorf("dragon: GiveDragonFeedReward AddItem should be ok"))
	}
	//同步物品
	inventorylogic.SnapInventoryChanged(pl)
	return
}
