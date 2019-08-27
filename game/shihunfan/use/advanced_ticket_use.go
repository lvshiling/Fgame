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
	shihunfaneventtypes "fgame/fgame/game/shihunfan/event/types"
	shihunfanlogic "fgame/fgame/game/shihunfan/logic"
	"fgame/fgame/game/shihunfan/pbutil"
	playershihunfan "fgame/fgame/game/shihunfan/player"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeShiHunFan, playerinventory.ItemUseHandleFunc(handleShiHunFanAdvancedTicket))
}

//噬魂幡直升券
func handleShiHunFanAdvancedTicket(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	shihunfanManager := pl.GetPlayerDataManager(types.PlayerShiHunFanDataManagerType).(*playershihunfan.PlayerShiHunFanDataManager)
	shihunfanInfo := shihunfanManager.GetShiHunFanInfo()
	curAdvancedId := int32(shihunfanInfo.AdvanceId)
	minAdvanceNum := itemTemplate.TypeFlag1
	maxAdvanceNum := itemTemplate.TypeFlag2
	if curAdvancedId < minAdvanceNum || curAdvancedId > maxAdvanceNum {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"curAdvancedId": curAdvancedId,
				"minAdvanceNum": itemTemplate.TypeFlag1,
				"maxAdvanceNum": itemTemplate.TypeFlag2,
			}).Warn("shihunfan:进阶丹使用,噬魂幡阶数条件不符")
		playerlogic.SendSystemMessage(pl, lang.ShiHunFanAdvanceNotEqual)
		return
	}

	addAdvancedNum := itemTemplate.TypeFlag3
	shihunfanManager.ShiHunFanAdvancedTicket(addAdvancedNum)

	shihunfanReason := commonlog.ShiHunFanLogReasonAdvanced
	reasonText := fmt.Sprintf(shihunfanReason.String(), commontypes.AdvancedTypeTicket.String())
	data := shihunfaneventtypes.CreatePlayerShiHunFanAdvancedLogEventData(curAdvancedId, addAdvancedNum, shihunfanReason, reasonText)
	gameevent.Emit(shihunfaneventtypes.EventTypeShiHunFanAdvancedLog, pl, data)

	//同步属性
	shihunfanlogic.ShiHunFanPropertyChanged(pl)
	scShiHunFanAdvanced := pbutil.BuildSCShiHunFanAdavancedFinshed(shihunfanInfo, commontypes.AdvancedTypeTicket)
	pl.SendMsg(scShiHunFanAdvanced)

	flag = true
	return
}
