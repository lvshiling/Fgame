package use

import (
	"fgame/fgame/common/lang"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	winglogic "fgame/fgame/game/wing/logic"
	"fgame/fgame/game/wing/pbutil"
	playerwing "fgame/fgame/game/wing/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeWing, itemtypes.ItemWingSubTypeTrialCard, playerinventory.ItemUseHandleFunc(handleWingUseTrialCard))
}

func handleWingUseTrialCard(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	itemSubType := itemTemplate.GetItemSubType()
	if itemSubType != itemtypes.ItemWingSubTypeTrialCard {
		return
	}
	//参数不对
	if num != 1 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"num":      num,
			}).Warn("wing:使用战翼试用卡,使用物品数量错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	wingManager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	wingInfo := wingManager.GetWingInfo()
	wingAdvancedId := wingInfo.AdvanceId
	if wingAdvancedId != 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("wing:您的战翼阶数,无法使用战翼试用卡")
		playerlogic.SendSystemMessage(pl, lang.WingTrialCardUseAdvancedIsZero)
		return
	}

	wingTrialInfo := wingManager.GetWingTrialInfo()
	trialOrderId := wingTrialInfo.TrialOrderId
	if trialOrderId != 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("wing:已获得战翼试用,期间无法重复获取")
		playerlogic.SendSystemMessage(pl, lang.WingTrialCardUseIsExist)
		return
	}
	flag = true

	wingManager.WingTrialOrder()
	winglogic.WingPropertyChanged(pl)
	scWingTrialCard := pbutil.BuildSCWingTrialCard(wingTrialInfo.TrialOrderId, wingTrialInfo.ActiveTime)
	pl.SendMsg(scWingTrialCard)
	return
}
