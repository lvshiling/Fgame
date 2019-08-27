package use

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	lingyueventtypes "fgame/fgame/game/lingyu/event/types"
	lingyulogic "fgame/fgame/game/lingyu/logic"
	"fgame/fgame/game/lingyu/pbutil"
	playerlingyu "fgame/fgame/game/lingyu/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeLingyu, playerinventory.ItemUseHandleFunc(handleLingyuAdvancedTicket))
}

//领域进阶丹
func handleLingyuAdvancedTicket(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	lingyuManager := pl.GetPlayerDataManager(types.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	lingyuInfo := lingyuManager.GetLingyuInfo()
	curAdvancedId := int32(lingyuInfo.AdvanceId)
	minAdvanceNum := itemTemplate.TypeFlag1
	maxAdvanceNum := itemTemplate.TypeFlag2
	if curAdvancedId < minAdvanceNum || curAdvancedId > maxAdvanceNum {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"curAdvancedId": curAdvancedId,
				"minAdvanceNum": itemTemplate.TypeFlag1,
				"maxAdvanceNum": itemTemplate.TypeFlag2,
			}).Warn("lingyu:进阶丹使用,领域阶数不符")
		playerlogic.SendSystemMessage(pl, lang.LingyuAdvanceNotEqual)
		return
	}

	addAdvancedNum := itemTemplate.TypeFlag3
	lingyuManager.LingyuAdvancedTicket(addAdvancedNum)

	lingyuReason := commonlog.LingyuLogReasonAdvanced
	reasonText := fmt.Sprintf(lingyuReason.String(), commontypes.AdvancedTypeTicket.String())
	data := lingyueventtypes.CreatePlayerLingyuAdvancedLogEventData(curAdvancedId, addAdvancedNum, lingyuReason, reasonText)
	gameevent.Emit(lingyueventtypes.EventTypeLingyuAdvancedLog, pl, data)

	//同步属性
	lingyulogic.LingyuPropertyChanged(pl)
	scLingyuAdvanced := pbutil.BuildSCLingyuAdavancedFinshed(int32(lingyuInfo.AdvanceId), lingyuInfo.LingyuId, commontypes.AdvancedTypeTicket)
	pl.SendMsg(scLingyuAdvanced)

	flag = true
	return
}
