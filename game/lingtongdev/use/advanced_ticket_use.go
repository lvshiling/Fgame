package use

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	lingtongdeveventtypes "fgame/fgame/game/lingtongdev/event/types"
	lingtongdevlogic "fgame/fgame/game/lingtongdev/logic"
	"fgame/fgame/game/lingtongdev/pbutil"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	lingtongdevtypes "fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeLingBing, playerinventory.ItemUseHandleFunc(handleLingTongDevAdvancedTicket))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeLingQi, playerinventory.ItemUseHandleFunc(handleLingTongDevAdvancedTicket))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeLingYi, playerinventory.ItemUseHandleFunc(handleLingTongDevAdvancedTicket))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeLingShen, playerinventory.ItemUseHandleFunc(handleLingTongDevAdvancedTicket))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeLingTongYu, playerinventory.ItemUseHandleFunc(handleLingTongDevAdvancedTicket))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeLingBao, playerinventory.ItemUseHandleFunc(handleLingTongDevAdvancedTicket))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeLingTi, playerinventory.ItemUseHandleFunc(handleLingTongDevAdvancedTicket))
}

func getAdvancedLingTongDevSysType(itemSubType itemtypes.ItemSubType) (classType lingtongdevtypes.LingTongDevSysType, flag bool) {
	flag = true
	switch itemSubType {
	case itemtypes.ItemAdvancedTicketSubTypeLingBing:
		classType = lingtongdevtypes.LingTongDevSysTypeLingBing
	case itemtypes.ItemAdvancedTicketSubTypeLingQi:
		classType = lingtongdevtypes.LingTongDevSysTypeLingQi
	case itemtypes.ItemAdvancedTicketSubTypeLingYi:
		classType = lingtongdevtypes.LingTongDevSysTypeLingYi
	case itemtypes.ItemAdvancedTicketSubTypeLingShen:
		classType = lingtongdevtypes.LingTongDevSysTypeLingShen
	case itemtypes.ItemAdvancedTicketSubTypeLingTongYu:
		classType = lingtongdevtypes.LingTongDevSysTypeLingYu
	case itemtypes.ItemAdvancedTicketSubTypeLingBao:
		classType = lingtongdevtypes.LingTongDevSysTypeLingBao
	case itemtypes.ItemAdvancedTicketSubTypeLingTi:
		classType = lingtongdevtypes.LingTongDevSysTypeLingTi
	default:
		flag = false
		return
	}
	return
}

//灵童养成类进阶丹
func handleLingTongDevAdvancedTicket(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return
	}
	itemSubType := itemTemplate.GetItemSubType()
	classType, isExist := getAdvancedLingTongDevSysType(itemSubType)
	if !isExist {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	lingTongDevInfo := manager.GetLingTongDevInfo(classType)
	// if lingTongDevInfo == nil {
	// 	lingTongDevInfo = manager.AdvancedInit(classType)
	// }
	curAdvancedId := lingTongDevInfo.GetAdvancedId()
	minAdvanceNum := itemTemplate.TypeFlag1
	maxAdvanceNum := itemTemplate.TypeFlag2
	if curAdvancedId < minAdvanceNum || curAdvancedId > maxAdvanceNum {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"curAdvancedId": curAdvancedId,
				"minAdvanceNum": itemTemplate.TypeFlag1,
				"maxAdvanceNum": itemTemplate.TypeFlag2,
			}).Warn("lingtongdev:进阶丹使用,灵童养成类阶数条件不符")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevAdvanceNotEqual, classType.String())
		return
	}

	addAdvancedNum := itemTemplate.TypeFlag3
	manager.LingTongDevAdvancedTicket(classType, addAdvancedNum)

	lingTongDevReason := commonlog.LingTongDevLogReasonAdvanced
	reasonText := fmt.Sprintf(lingTongDevReason.String(), commontypes.AdvancedTypeTicket.String())
	data := lingtongdeveventtypes.CreatePlayerLingTongDevAdvancedLogEventData(classType, curAdvancedId, addAdvancedNum, lingTongDevReason, reasonText)
	gameevent.Emit(lingtongdeveventtypes.EventTypeLingTongDevAdvancedLog, pl, data)

	//同步属性
	lingtongdevlogic.LingTongDevPropertyChanged(pl, classType)
	scLingTongDevAdavancedFinshed := pbutil.BuildSCLingTongDevAdavancedFinshed(int32(classType), lingTongDevInfo.GetAdvancedId(), lingTongDevInfo.GetSeqId(), commontypes.AdvancedTypeTicket)
	pl.SendMsg(scLingTongDevAdavancedFinshed)
	flag = true
	return
}
