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
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeAdvancedTicket, itemtypes.ItemAdvancedTicketSubTypeShield, playerinventory.ItemUseHandleFunc(handleShieldAdvancedTicket))
}

//神盾尖刺进阶丹
func handleShieldAdvancedTicket(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	shieldManager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbodyshield.PlayerBodyShieldDataManager)
	shieldInfo := shieldManager.GetBodyShiedInfo()
	curAdvancedId := int32(shieldInfo.ShieldId)
	minAdvanceNum := itemTemplate.TypeFlag1
	maxAdvanceNum := itemTemplate.TypeFlag2
	if curAdvancedId < minAdvanceNum || curAdvancedId > maxAdvanceNum {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"curAdvancedId": curAdvancedId,
				"minAdvanceNum": itemTemplate.TypeFlag1,
				"maxAdvanceNum": itemTemplate.TypeFlag2,
			}).Warn("shield:进阶丹使用,神盾尖刺阶数不足")
		playerlogic.SendSystemMessage(pl, lang.ShieldAdvanceNotEqual)
		return
	}

	addAdvancedNum := itemTemplate.TypeFlag3
	shieldManager.ShieldFeedTicket(addAdvancedNum)

	shieldReason := commonlog.ShieldLogReasonAdvanced
	reasonText := fmt.Sprintf(shieldReason.String(), commontypes.AdvancedTypeTicket.String())
	data := bodyshieldeventtypes.CreatePlayerShieldAdvancedLogEventData(curAdvancedId, addAdvancedNum, shieldReason, reasonText)
	gameevent.Emit(bodyshieldeventtypes.EventTypeShieldAdvancedLog, pl, data)

	//同步属性
	bodyshieldlogic.BodyShieldPropertyChanged(pl)

	scShieldAdvanced := pbutil.BuildSCShieldAdvanced(shieldInfo.ShieldId, shieldInfo.ShieldPro, commontypes.AdvancedTypeTicket, false)
	pl.SendMsg(scShieldAdvanced)

	flag = true
	return
}
