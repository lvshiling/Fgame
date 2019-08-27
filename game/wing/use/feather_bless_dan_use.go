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
	"fgame/fgame/game/player/types"
	wingeventtypes "fgame/fgame/game/wing/event/types"
	winglogic "fgame/fgame/game/wing/logic"
	"fgame/fgame/game/wing/pbutil"
	playerwing "fgame/fgame/game/wing/player"
	"fgame/fgame/game/wing/wing"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeFeather, playerinventory.ItemUseHandleFunc(handleFeatherBlessDan))
}

//护体仙羽祝福丹
func handleFeatherBlessDan(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	wingManager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	wingInfo := wingManager.GetWingInfo()
	featherAdvancedId := int32(wingInfo.FeatherId)
	needAdvanceId := itemTemplate.TypeFlag1
	if featherAdvancedId != needAdvanceId {
		log.WithFields(
			log.Fields{
				"playerId":          pl.GetId(),
				"featherAdvancedId": featherAdvancedId,
				"needAdvanceId":     itemTemplate.TypeFlag1,
			}).Warn("feather:祝福丹使用,护体仙羽阶数不足")
		playerlogic.SendSystemMessage(pl, lang.FeatherAdvanceNotEqual)
		return
	}

	nextFeatherTemp := wing.GetWingService().GetFeather(featherAdvancedId + 1)
	addBlessNum := itemTemplate.TypeFlag2
	addTimes := int32(0) //优化成没有次数效果
	curTimes := wingInfo.FeatherNum
	curTimes += addTimes
	// isAdvanced, pro := winglogic.FeatherEatZhuFuDan(pl, curTimes, wingInfo.FeatherPro, addBlessNum, nextFeatherTemp)
	isAdvanced, pro := commonlogic.AdvancedBlessDan(wingInfo.FeatherPro, addBlessNum, nextFeatherTemp.NeedRate)
	wingManager.FeatherFeed(pro, addTimes, isAdvanced)

	if isAdvanced {
		featherReason := commonlog.WingLogReasonFeatherAdvanced
		reasonText := fmt.Sprintf(featherReason.String(), commontypes.AdvancedTypeBlessDan.String())
		data := wingeventtypes.CreatePlayerFeatherAdvancedLogEventData(featherAdvancedId, 1, featherReason, reasonText)
		gameevent.Emit(wingeventtypes.EventTypeFeatherAdvancedLog, pl, data)

		//同步属性
		winglogic.FeatherPropertyChanged(pl)
	}

	scFeatherAdvanced := pbutil.BuildSCFeatherAdvanced(wingInfo.FeatherId, wingInfo.FeatherPro, commontypes.AdvancedTypeBlessDan)
	pl.SendMsg(scFeatherAdvanced)

	flag = true
	return
}
