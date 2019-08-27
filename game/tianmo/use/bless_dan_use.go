package use

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	commonlogic "fgame/fgame/game/common/logic"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	tianmoeventtypes "fgame/fgame/game/tianmo/event/types"
	tianmologic "fgame/fgame/game/tianmo/logic"
	"fgame/fgame/game/tianmo/pbutil"
	playertianmo "fgame/fgame/game/tianmo/player"
	tianmotemplate "fgame/fgame/game/tianmo/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeTianMoTi, playerinventory.ItemUseHandleFunc(handleTianMoTiBlessDan))
}

//天魔体祝福丹
func handleTianMoTiBlessDan(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	tianmoManager := pl.GetPlayerDataManager(playertypes.PlayerTianMoDataManagerType).(*playertianmo.PlayerTianMoDataManager)
	tianmoInfo := tianmoManager.GetTianMoInfo()
	tianmoAdvancedId := tianmoInfo.AdvanceId
	needAdvanceId := itemTemplate.TypeFlag1
	if tianmoAdvancedId != needAdvanceId {
		log.WithFields(
			log.Fields{
				"playerId":         pl.GetId(),
				"tianmoAdvancedId": tianmoAdvancedId,
				"needAdvanceId":    itemTemplate.TypeFlag1,
			}).Warn("tianmo:祝福丹使用,天魔体阶数不足")
		playerlogic.SendSystemMessage(pl, lang.TianMoAdvanceNotEqual)
		return
	}

	nextTianMoTiTemp := tianmotemplate.GetTianMoTemplateService().GetTianMoNumber(tianmoAdvancedId + 1)
	addBlessNum := itemTemplate.TypeFlag2
	addTimes := int32(0) //优化成没有次数效果
	curTimes := tianmoInfo.TimesNum
	curTimes += addTimes
	// isAdvanced, pro := tianmologic.TianMoEatZhuFuDan(pl, curTimes, tianmoInfo.Bless, addBlessNum, nextTianMoTiTemp)
	isAdvanced, pro := commonlogic.AdvancedBlessDan(tianmoInfo.Bless, addBlessNum, nextTianMoTiTemp.ZhufuMax)
	tianmoManager.TianMoAdvanced(pro, addTimes, isAdvanced)

	if isAdvanced {
		//同步属性
		tianmologic.TianMoPropertyChanged(pl)

		tianmoReason := commonlog.TianMoLogReasonAdvanced
		reasonText := fmt.Sprintf(tianmoReason.String(), commontypes.AdvancedTypeBlessDan.String())
		data := tianmoeventtypes.CreatePlayerTianMoAdvancedLogEventData(tianmoAdvancedId, 1, tianmoReason, reasonText)
		gameevent.Emit(tianmoeventtypes.EventTypeTianMoAdvancedLog, pl, data)
	}

	scMsg := pbutil.BuildSCTianMoAdavanced(tianmoInfo.AdvanceId, addBlessNum, tianmoInfo.Bless, tianmoInfo.BlessTime, commontypes.AdvancedTypeBlessDan, false)
	pl.SendMsg(scMsg)

	flag = true
	return
}
