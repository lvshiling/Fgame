package use

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	bodyshieldeventtypes "fgame/fgame/game/bodyshield/event/types"
	bodyshieldlogic "fgame/fgame/game/bodyshield/logic"
	"fgame/fgame/game/bodyshield/pbutil"
	playerbodyshield "fgame/fgame/game/bodyshield/player"
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
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeBodyShield, playerinventory.ItemUseHandleFunc(handleBodyShieldAdvancedTicket))
}

//护体盾进阶丹
func handleBodyShieldAdvancedTicket(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	bodyshieldManager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbodyshield.PlayerBodyShieldDataManager)
	bodyshieldInfo := bodyshieldManager.GetBodyShiedInfo()
	curAdvancedId := int32(bodyshieldInfo.AdvanceId)
	minAdvanceNum := itemTemplate.TypeFlag1
	maxAdvanceNum := itemTemplate.TypeFlag2
	if curAdvancedId < minAdvanceNum || curAdvancedId > maxAdvanceNum {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"curAdvancedId": curAdvancedId,
				"minAdvanceNum": itemTemplate.TypeFlag1,
				"maxAdvanceNum": itemTemplate.TypeFlag2,
			}).Warn("bodyshield:进阶丹使用,护体盾阶数不足")
		playerlogic.SendSystemMessage(pl, lang.BodyShieldAdvanceNotEqual)
		return
	}

	addAdvancedNum := itemTemplate.TypeFlag3
	bodyshieldManager.BodyShieldAdvancedTicket(addAdvancedNum)

	shieldReason := commonlog.ShieldLogReasonBodyAdvanced
	reasonText := fmt.Sprintf(shieldReason.String(), commontypes.AdvancedTypeTicket.String())
	data := bodyshieldeventtypes.CreatePlayerBodyShieldAdvancedLogEventData(curAdvancedId, addAdvancedNum, shieldReason, reasonText)
	gameevent.Emit(bodyshieldeventtypes.EventTypeBodyShieldAdvancedLog, pl, data)

	//同步属性
	bodyshieldlogic.BodyShieldPropertyChanged(pl)
	scBodyShieldAdvanced := pbutil.BuildSCBodyShieldAdavancedFinshed(int32(bodyshieldInfo.AdvanceId), commontypes.AdvancedTypeTicket)
	pl.SendMsg(scBodyShieldAdvanced)

	flag = true
	return
}
