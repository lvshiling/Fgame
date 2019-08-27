package use

import (
	"fgame/fgame/common/lang"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	lingtonglogic "fgame/fgame/game/lingtong/logic"
	"fgame/fgame/game/lingtong/pbutil"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeLingTongFashion, itemtypes.ItemLingTongFashionSubTypeTrialCard, playerinventory.ItemUseHandleFunc(handleLingTongFashionTrialCard))
}

func handleLingTongFashionTrialCard(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	if num != 1 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("lingtong:使用时装试用卡,使用物品数量错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongObj := manager.GetLingTong()
	if lingTongObj == nil || !lingTongObj.IsActivateSys() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("lingtong:请先激活灵童时装激活系统")
		playerlogic.SendSystemMessage(pl, lang.LingTongActiveSystem)
		return
	}

	itemTemplate := item.GetItemService().GetItem(int(itemId))
	trialFashionId := itemTemplate.TypeFlag1

	isBorn := lingtongtemplate.GetLingTongTemplateService().IsBornFashion(trialFashionId)
	if isBorn {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("lingtong:使用时装试用卡,出生时装无法试用")
		return
	}

	if len(manager.GetActivateFashionMap()) != 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("lingtong:使用时装试用卡,已激活永久时装")
		playerlogic.SendSystemMessage(pl, lang.LingTongFashionTrialHadActivate)
		return
	}
	// 已激活永久
	lingTongFashionInfoObject := manager.GetFashionInfoById(trialFashionId)
	if lingTongFashionInfoObject != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("lingtong:使用时装试用卡,已激活永久时装")
		playerlogic.SendSystemMessage(pl, lang.LingTongFashionTrialHadActivate)
		return
	}

	// 已获取体验
	if manager.IsFashionTrial(trialFashionId) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
			}).Warn("lingtong:已获得时装试用,期间无法重复获取")
		playerlogic.SendSystemMessage(pl, lang.LingTongFashionTrialCardUseIsExist)
		return
	}

	expireTime := manager.UseFashionTrialCard(itemId)
	lingtonglogic.LingTongFashionPropertyChanged(pl)

	scMsg := pbutil.BuildSCLingTongFashionTrialNotice(trialFashionId, expireTime)
	pl.SendMsg(scMsg)
	flag = true
	return
}
