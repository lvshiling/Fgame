package use

import (
	"fgame/fgame/common/lang"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/xuechi/pbutil"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLifeOrigin, itemtypes.ItemDefaultSubTypeDefault, playerinventory.ItemUseHandleFunc(handleLifeBottleUse))
}

func handleLifeBottleUse(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	//参数不对
	if num < 1 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("xueChi:使用血药,使用物品数量错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate.GetItemType() != itemtypes.ItemTypeLifeOrigin {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"itemId":   itemId,
			"num":      num,
		}).Warn("xueChi:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	addBlood := int64(itemTemplate.TypeFlag1) * int64(num)
	pl.AddBlood(addBlood)
	scXueChiBlood := pbutil.BuildSCXueChiBlood(pl.GetBlood())
	pl.SendMsg(scXueChiBlood)

	flag = true
	return
}
