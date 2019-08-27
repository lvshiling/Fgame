package use

import (
	"fgame/fgame/common/lang"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	weaponlogic "fgame/fgame/game/weapon/logic"
	"fgame/fgame/game/weapon/pbutil"
	playerweapon "fgame/fgame/game/weapon/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeSoul, itemtypes.ItemSoulSubTypeDingZhi, playerinventory.ItemUseHandleFunc(handleDingZhi))
}

func handleDingZhi(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	if num != 1 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("weapon:使用定制冰魂,使用物品数量错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	weaponManager := pl.GetPlayerDataManager(playertypes.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	weaponId := itemTemplate.TypeFlag1

	// 是否激活过
	if weaponManager.IfWeaponExist(weaponId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("weapon:使用定制冰魂,已激活")
		playerlogic.SendSystemMessage(pl, lang.WeaponRepeatActive)
		return
	}

	weaponManager.WeaponActive(weaponId, true)

	//同步属性
	weaponlogic.WeaponPropertyChanged(pl)

	scWeaponActive := pbutil.BuildSCWeaponActive(weaponId)
	pl.SendMsg(scWeaponActive)

	flag = true
	return
}
