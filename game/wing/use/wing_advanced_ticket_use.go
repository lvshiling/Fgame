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
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeWing, playerinventory.ItemUseHandleFunc(handleWingAdvancedTicket))
}

//战翼进阶丹
func handleWingAdvancedTicket(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	wingManager := pl.GetPlayerDataManager(types.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	wingInfo := wingManager.GetWingInfo()
	curAdvancedId := int32(wingInfo.AdvanceId)
	minAdvanceNum := itemTemplate.TypeFlag1
	maxAdvanceNum := itemTemplate.TypeFlag2
	if curAdvancedId < minAdvanceNum || curAdvancedId > maxAdvanceNum {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"curAdvancedId": curAdvancedId,
				"minAdvanceNum": itemTemplate.TypeFlag1,
				"maxAdvanceNum": itemTemplate.TypeFlag2,
			}).Warn("wing:进阶丹使用,战翼阶数条件不符")
		playerlogic.SendSystemMessage(pl, lang.WingAdvanceNotEqual)
		return
	}

	addAdvancedNum := itemTemplate.TypeFlag3
	wingManager.WingAdvancedTicket(addAdvancedNum)

	wingReason := commonlog.WingLogReasonAdvanced
	reasonText := fmt.Sprintf(wingReason.String(), commontypes.AdvancedTypeTicket.String())
	data := wingeventtypes.CreatePlayerWingAdvancedLogEventData(curAdvancedId, addAdvancedNum, wingReason, reasonText)
	gameevent.Emit(wingeventtypes.EventTypeWingAdvancedLog, pl, data)

	//同步属性
	winglogic.WingPropertyChanged(pl)
	scWingAdvanced := pbutil.BuildSCWingAdavancedFinshed(int32(wingInfo.AdvanceId), wingInfo.WingId, commontypes.AdvancedTypeTicket)
	pl.SendMsg(scWingAdvanced)

	flag = true
	return
}
