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
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeWing, playerinventory.ItemUseHandleFunc(handleWingBlessDan))
}

//战翼祝福丹
func handleWingBlessDan(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	wingManager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	wingInfo := wingManager.GetWingInfo()
	wingAdvancedId := int32(wingInfo.AdvanceId)
	needAdvanceId := itemTemplate.TypeFlag1
	if wingAdvancedId != needAdvanceId {
		log.WithFields(
			log.Fields{
				"playerId":       pl.GetId(),
				"wingAdvancedId": wingAdvancedId,
				"needAdvanceId":  itemTemplate.TypeFlag1,
			}).Warn("wing:祝福丹使用,战翼阶数不满足要求")
		playerlogic.SendSystemMessage(pl, lang.WingAdvanceNotEqual)
		return
	}

	nextWingTemp := wing.GetWingService().GetWingNumber(wingAdvancedId + 1)
	addBlessNum := itemTemplate.TypeFlag2
	addTimes := int32(0) //优化成没有次数效果
	curTimes := wingInfo.TimesNum
	curTimes += addTimes
	// isAdvanced, pro := winglogic.WingEatZhuFuDan(pl, curTimes, wingInfo.Bless, addBlessNum, nextWingTemp)
	isAdvanced, pro := commonlogic.AdvancedBlessDan(wingInfo.Bless, addBlessNum, nextWingTemp.ZhufuMax)
	wingManager.WingAdvanced(pro, addTimes, isAdvanced)
	if isAdvanced {
		wingReason := commonlog.WingLogReasonAdvanced
		reasonText := fmt.Sprintf(wingReason.String(), commontypes.AdvancedTypeBlessDan.String())
		data := wingeventtypes.CreatePlayerWingAdvancedLogEventData(wingAdvancedId, 1, wingReason, reasonText)
		gameevent.Emit(wingeventtypes.EventTypeWingAdvancedLog, pl, data)

		//同步属性
		winglogic.WingPropertyChanged(pl)
		scWingAdvanced := pbutil.BuildSCWingAdavancedFinshed(int32(wingInfo.AdvanceId), wingInfo.WingId, commontypes.AdvancedTypeBlessDan)
		pl.SendMsg(scWingAdvanced)
	} else {
		//进阶不成功
		scWingAdvanced := pbutil.BuildSCWingAdavanced(int32(wingInfo.AdvanceId), wingInfo.WingId, addBlessNum, wingInfo.Bless, wingInfo.BlessTime, commontypes.AdvancedTypeBlessDan, false)
		pl.SendMsg(scWingAdvanced)
	}

	flag = true
	return
}
