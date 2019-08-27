package use

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	anqieventtypes "fgame/fgame/game/anqi/event/types"
	anqilogic "fgame/fgame/game/anqi/logic"
	"fgame/fgame/game/anqi/pbutil"
	playeranqi "fgame/fgame/game/anqi/player"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeAnqi, playerinventory.ItemUseHandleFunc(handleAnqiAdvancedTicket))
}

//暗器进阶丹
func handleAnqiAdvancedTicket(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	anqiManager := pl.GetPlayerDataManager(types.PlayerAnqiDataManagerType).(*playeranqi.PlayerAnqiDataManager)
	anqiInfo := anqiManager.GetAnqiInfo()
	curAdvancedId := int32(anqiInfo.AdvanceId)
	minAdvanceNum := itemTemplate.TypeFlag1
	maxAdvanceNum := itemTemplate.TypeFlag2
	if curAdvancedId < minAdvanceNum || curAdvancedId > maxAdvanceNum {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"curAdvancedId": curAdvancedId,
				"minAdvanceNum": itemTemplate.TypeFlag1,
				"maxAdvanceNum": itemTemplate.TypeFlag2,
			}).Warn("anqi:进阶丹使用,暗器阶数条件不符")
		playerlogic.SendSystemMessage(pl, lang.AnqiAdvanceNotEqual)
		return
	}

	addAdvancedNum := itemTemplate.TypeFlag3
	anqiManager.AnqiAdvancedTicket(addAdvancedNum)

	anqiReason := commonlog.AnqiLogReasonAdvanced
	reasonText := fmt.Sprintf(anqiReason.String(), commontypes.AdvancedTypeTicket.String())
	data := anqieventtypes.CreatePlayerAnqiAdvancedLogEventData(curAdvancedId, addAdvancedNum, anqiReason, reasonText)
	gameevent.Emit(anqieventtypes.EventTypeAnqiAdvancedLog, pl, data)

	//同步属性
	anqilogic.AnqiPropertyChanged(pl)
	scAnqiAdvanced := pbutil.BuildSCAnqiAdavancedFinshed(int32(anqiInfo.AdvanceId), commontypes.AdvancedTypeTicket)
	pl.SendMsg(scAnqiAdvanced)

	flag = true
	return
}
