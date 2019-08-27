package use

import (
	"fgame/fgame/common/lang"
	fashionlogic "fgame/fgame/game/fashion/logic"
	"fgame/fgame/game/fashion/pbutil"
	playerfashion "fgame/fgame/game/fashion/player"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeFashion, itemtypes.ItemFashionSubTypeTrialCard, playerinventory.ItemUseHandleFunc(handleFashionTrialCard))
}

func handleFashionTrialCard(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	if num != 1 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("fashion:使用时装试用卡,使用物品数量错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	fashionManager := pl.GetPlayerDataManager(playertypes.PlayerFashionDataManagerType).(*playerfashion.PlayerFashionDataManager)
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	trialFashionId := itemTemplate.TypeFlag1
	expireTime := int64(itemTemplate.TypeFlag2)

	// 已激活永久
	if fashionManager.IfFashionExist(trialFashionId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("fashion:使用时装试用卡,已激活永久时装")
		playerlogic.SendSystemMessage(pl, lang.FashionTrialHadActivate)
		return
	}

	// 已获取体验
	if fashionManager.IsFashionTrial(trialFashionId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
			}).Warn("fashion:已获得时装试用,期间无法重复获取")
		playerlogic.SendSystemMessage(pl, lang.FashionTrialCardUseIsExist)
		return
	}

	fashionManager.UseFashionTrialCard(itemId)
	fashionlogic.FashionPropertyChanged(pl)

	scMsg := pbutil.BuildSCFashionTrialNotice(trialFashionId, expireTime)
	pl.SendMsg(scMsg)

	flag = true
	return
}
