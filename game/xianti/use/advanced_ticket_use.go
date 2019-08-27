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
	xiantieventtypes "fgame/fgame/game/xianti/event/types"
	xiantilogic "fgame/fgame/game/xianti/logic"
	"fgame/fgame/game/xianti/pbutil"
	playerxianti "fgame/fgame/game/xianti/player"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeXianTi, playerinventory.ItemUseHandleFunc(handleXianTiAdvancedTicket))
}

//仙体直升券
func handleXianTiAdvancedTicket(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	xiantiManager := pl.GetPlayerDataManager(types.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	xiantiInfo := xiantiManager.GetXianTiInfo()
	curAdvancedId := int32(xiantiInfo.AdvanceId)
	minAdvanceNum := itemTemplate.TypeFlag1
	maxAdvanceNum := itemTemplate.TypeFlag2
	if curAdvancedId < minAdvanceNum || curAdvancedId > maxAdvanceNum {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"curAdvancedId": curAdvancedId,
				"minAdvanceNum": itemTemplate.TypeFlag1,
				"maxAdvanceNum": itemTemplate.TypeFlag2,
			}).Warn("xianti:进阶丹使用,仙体阶数条件不符")
		playerlogic.SendSystemMessage(pl, lang.XianTiAdvanceNotEqual)
		return
	}

	addAdvancedNum := itemTemplate.TypeFlag3
	xiantiManager.XianTiAdvancedTicket(addAdvancedNum)

	xiantiReason := commonlog.XianTiLogReasonAdvanced
	reasonText := fmt.Sprintf(xiantiReason.String(), commontypes.AdvancedTypeTicket.String())
	data := xiantieventtypes.CreatePlayerXianTiAdvancedLogEventData(curAdvancedId, addAdvancedNum, xiantiReason, reasonText)
	gameevent.Emit(xiantieventtypes.EventTypeXianTiAdvancedLog, pl, data)

	//同步属性
	xiantilogic.XianTiPropertyChanged(pl)
	scXianTiAdvanced := pbutil.BuildSCXianTiAdavancedFinshed(int32(xiantiInfo.AdvanceId), int32(xiantiInfo.XianTiId), commontypes.AdvancedTypeTicket)
	pl.SendMsg(scXianTiAdvanced)

	flag = true
	return
}
