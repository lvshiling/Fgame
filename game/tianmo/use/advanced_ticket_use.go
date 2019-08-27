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
	playertypes "fgame/fgame/game/player/types"
	tianmoeventtypes "fgame/fgame/game/tianmo/event/types"
	tianmologic "fgame/fgame/game/tianmo/logic"
	"fgame/fgame/game/tianmo/pbutil"
	playertianmo "fgame/fgame/game/tianmo/player"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeTianMoTi, playerinventory.ItemUseHandleFunc(handleTianMoTiAdvancedTicket))
}

//天魔体进阶丹
func handleTianMoTiAdvancedTicket(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	tianmoManager := pl.GetPlayerDataManager(playertypes.PlayerTianMoDataManagerType).(*playertianmo.PlayerTianMoDataManager)
	tianmoInfo := tianmoManager.GetTianMoInfo()
	curAdvancedId := tianmoInfo.AdvanceId
	minAdvanceNum := itemTemplate.TypeFlag1
	maxAdvanceNum := itemTemplate.TypeFlag2
	if curAdvancedId < minAdvanceNum || curAdvancedId > maxAdvanceNum {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"curAdvancedId": curAdvancedId,
				"minAdvanceNum": itemTemplate.TypeFlag1,
				"maxAdvanceNum": itemTemplate.TypeFlag2,
			}).Warn("tianmo:进阶丹使用,天魔体阶数条件不符")
		playerlogic.SendSystemMessage(pl, lang.TianMoAdvanceNotEqual)
		return
	}

	addAdvancedNum := itemTemplate.TypeFlag3
	tianmoManager.TianMoAdvancedTicket(addAdvancedNum)

	tianmoReason := commonlog.TianMoLogReasonAdvanced
	reasonText := fmt.Sprintf(tianmoReason.String(), commontypes.AdvancedTypeTicket.String())
	data := tianmoeventtypes.CreatePlayerTianMoAdvancedLogEventData(curAdvancedId, addAdvancedNum, tianmoReason, reasonText)
	gameevent.Emit(tianmoeventtypes.EventTypeTianMoAdvancedLog, pl, data)

	//同步属性
	tianmologic.TianMoPropertyChanged(pl)
	scTianMoTiAdvanced := pbutil.BuildSCTianMoAdavancedFinshed(tianmoInfo.AdvanceId, commontypes.AdvancedTypeTicket)
	pl.SendMsg(scTianMoTiAdvanced)

	flag = true
	return
}
