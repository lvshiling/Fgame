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
	shenfaeventtypes "fgame/fgame/game/shenfa/event/types"
	shenfalogic "fgame/fgame/game/shenfa/logic"
	"fgame/fgame/game/shenfa/pbutil"
	playershenfa "fgame/fgame/game/shenfa/player"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeShenfa, playerinventory.ItemUseHandleFunc(handleShenfaAdvancedTicket))
}

//身法进阶丹
func handleShenfaAdvancedTicket(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	shenfaManager := pl.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
	shenfaInfo := shenfaManager.GetShenfaInfo()
	curAdvancedId := int32(shenfaInfo.AdvanceId)
	minAdvanceNum := itemTemplate.TypeFlag1
	maxAdvanceNum := itemTemplate.TypeFlag2
	if curAdvancedId < minAdvanceNum || curAdvancedId > maxAdvanceNum {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"curAdvancedId": curAdvancedId,
				"minAdvanceNum": itemTemplate.TypeFlag1,
				"maxAdvanceNum": itemTemplate.TypeFlag2,
			}).Warn("shenfa:进阶丹使用,身法阶数不足")
		playerlogic.SendSystemMessage(pl, lang.ShenfaAdvanceNotEqual)
		return
	}

	addAdvancedNum := itemTemplate.TypeFlag3
	shenfaManager.ShenfaAdvancedTicket(addAdvancedNum)

	shenfaReason := commonlog.ShenfaLogReasonAdvanced
	reasonText := fmt.Sprintf(shenfaReason.String(), commontypes.AdvancedTypeTicket.String())
	data := shenfaeventtypes.CreatePlayerShenfaAdvancedLogEventData(curAdvancedId, addAdvancedNum, shenfaReason, reasonText)
	gameevent.Emit(shenfaeventtypes.EventTypeShenfaAdvancedLog, pl, data)

	//同步属性
	shenfalogic.ShenfaPropertyChanged(pl)
	scShenfaAdvanced := pbutil.BuildSCShenfaAdavancedFinshed(int32(shenfaInfo.AdvanceId), shenfaInfo.ShenfaId, commontypes.AdvancedTypeTicket)
	pl.SendMsg(scShenfaAdvanced)

	flag = true
	return
}
