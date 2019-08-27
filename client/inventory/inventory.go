package inventory

import (
	"fgame/fgame/client/player"

	log "github.com/Sirupsen/logrus"
)

//加载请求
func InventoryGet(pl *player.Player, page int32) (err error) {
	log.WithFields(
		log.Fields{
			"playerId": pl.Id(),
			"page":     page,
		}).Debug("inventory:加载物品")
	cmd := buildInventoryGet(page)
	pl.SendMessage(cmd)
	return nil
}

//使用物品
func InventoryUse(pl *player.Player, index int32, num int32) (err error) {
	log.WithFields(
		log.Fields{
			"playerId": pl.Id(),
			"index":    index,
			"num":      num,
		}).Debug("inventory:使用物品")

	cmd := buildInventoryUse(index, num)
	pl.SendMessage(cmd)
	return nil
}

//合并物品
func InventoryMerge(pl *player.Player) (err error) {
	log.WithFields(
		log.Fields{
			"playerId": pl.Id(),
		}).Debug("inventory:合并物品")

	cmd := buildInventoryMerge()
	pl.SendMessage(cmd)
	return nil
}

//获取物品
func OnInventoryGet(pl *player.Player, page int32, items []*ItemObject) {
	pidm := pl.GetManager(player.PlayerDataKeyInventory).(*PlayerInventoryDataManager)
	pidm.OnLoadItems(page, items)
}

//物品改变
func OnInventoryChanged(pl *player.Player, items []*ItemObject) {
	pidm := pl.GetManager(player.PlayerDataKeyInventory).(*PlayerInventoryDataManager)
	pidm.OnItemsChange(items)
	pl.GetStrategy().OnItemChanged()
}
