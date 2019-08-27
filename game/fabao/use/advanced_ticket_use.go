package use

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	fabaoeventtypes "fgame/fgame/game/fabao/event/types"
	fabaologic "fgame/fgame/game/fabao/logic"
	"fgame/fgame/game/fabao/pbutil"
	playerfabao "fgame/fgame/game/fabao/player"
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
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeFaBao, playerinventory.ItemUseHandleFunc(handleFaBaoAdvancedTicket))
}

//法宝直升券
func handleFaBaoAdvancedTicket(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	fabaoManager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	fabaoInfo := fabaoManager.GetFaBaoInfo()
	curAdvancedId := fabaoInfo.GetAdvancedId()
	minAdvanceNum := itemTemplate.TypeFlag1
	maxAdvanceNum := itemTemplate.TypeFlag2
	if curAdvancedId < minAdvanceNum || curAdvancedId > maxAdvanceNum {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"curAdvancedId": curAdvancedId,
				"minAdvanceNum": itemTemplate.TypeFlag1,
				"maxAdvanceNum": itemTemplate.TypeFlag2,
			}).Warn("fabao:进阶丹使用,法宝阶数条件不符")
		playerlogic.SendSystemMessage(pl, lang.FaBaoAdvanceNotEqual)
		return
	}

	addAdvancedNum := itemTemplate.TypeFlag3
	fabaoManager.FaBaoAdvancedTicket(addAdvancedNum)

	fabaoReason := commonlog.FaBaoLogReasonAdvanced
	reasonText := fmt.Sprintf(fabaoReason.String(), commontypes.AdvancedTypeTicket.String())
	data := fabaoeventtypes.CreatePlayerFaBaoAdvancedLogEventData(curAdvancedId, addAdvancedNum, fabaoReason, reasonText)
	gameevent.Emit(fabaoeventtypes.EventTypeFaBaoAdvancedLog, pl, data)

	//同步属性
	fabaologic.FaBaoPropertyChanged(pl)
	scFaBaoAdvanced := pbutil.BuildSCFaBaoAdavancedFinshed(fabaoInfo.GetAdvancedId(), fabaoInfo.GetFaBaoId(), commontypes.AdvancedTypeTicket)
	pl.SendMsg(scFaBaoAdvanced)

	flag = true
	return
}
