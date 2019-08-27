package use

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
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
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeFeather, playerinventory.ItemUseHandleFunc(handleFeatherAdvancedTicket))
}

//护体仙羽进阶丹
func handleFeatherAdvancedTicket(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	wingManager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	wingInfo := wingManager.GetWingInfo()
	curAdvancedId := int32(wingInfo.FeatherId)
	minAdvanceNum := itemTemplate.TypeFlag1
	maxAdvanceNum := itemTemplate.TypeFlag2
	if curAdvancedId < minAdvanceNum || curAdvancedId > maxAdvanceNum {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"curAdvancedId": curAdvancedId,
				"minAdvanceNum": itemTemplate.TypeFlag1,
				"maxAdvanceNum": itemTemplate.TypeFlag2,
			}).Warn("wing:进阶丹使用,护体仙羽阶数不符")
		playerlogic.SendSystemMessage(pl, lang.FeatherAdvanceNotEqual)
		return
	}

	addAdvancedNum := itemTemplate.TypeFlag3
	wingManager.FeatherFeedTicket(addAdvancedNum)

	featherReason := commonlog.WingLogReasonFeatherAdvanced
	reasonText := fmt.Sprintf(featherReason.String(), commontypes.AdvancedTypeTicket.String())
	data := wingeventtypes.CreatePlayerFeatherAdvancedLogEventData(curAdvancedId, addAdvancedNum, featherReason, reasonText)
	gameevent.Emit(wingeventtypes.EventTypeFeatherAdvancedLog, pl, data)

	//同步属性
	winglogic.WingPropertyChanged(pl)

	scFeatherAdvanced := pbutil.BuildSCFeatherAdvanced(wingInfo.FeatherId, wingInfo.FeatherPro, commontypes.AdvancedTypeTicket)
	pl.SendMsg(scFeatherAdvanced)

	flag = true
	return
}
