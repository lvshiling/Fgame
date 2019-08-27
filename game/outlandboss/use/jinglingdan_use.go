package use

import (
	"fgame/fgame/common/lang"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/outlandboss/pbutil"
	playeroutlandboss "fgame/fgame/game/outlandboss/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBossItem, itemtypes.ItemBossSubTypeJingLingDan, playerinventory.ItemUseHandleFunc(handleUseBossItemJingLingDan))
}

// 使用净灵丹
func handleUseBossItemJingLingDan(pl player.Player, it *playerinventory.PlayerItemObject, itemNum int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	zhuoQi := int32(itemTemplate.TypeFlag1)

	outLandBossManager := pl.GetPlayerDataManager(playertypes.PlayerOutlandBossDataManagerType).(*playeroutlandboss.PlayerOutlandBossDataManager)
	// 判断浊气值
	if !outLandBossManager.CanUseJingLingDan() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("jinglingdan: 浊气值为零，无需使用")
		playerlogic.SendSystemMessage(pl, lang.OutlandBossZhuoQiNumIsZero)
		return
	}
	// 设置浊气值
	outLandBossManager.SetZhuoQiByJingLingDan(zhuoQi)

	scMsg := pbutil.BuildSCOutlandBossZhuoqiInfo(outLandBossManager.GetCurZhuoQiNum())
	pl.SendMsg(scMsg)

	flag = true
	return
}
