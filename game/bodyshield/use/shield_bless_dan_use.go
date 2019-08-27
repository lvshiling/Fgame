package use

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/bodyshield/bodyshield"
	bodyshieldeventtypes "fgame/fgame/game/bodyshield/event/types"
	bodyshieldlogic "fgame/fgame/game/bodyshield/logic"
	"fgame/fgame/game/bodyshield/pbutil"
	playerbodyshield "fgame/fgame/game/bodyshield/player"
	commonlogic "fgame/fgame/game/common/logic"
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
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeShield, playerinventory.ItemUseHandleFunc(handleShieldBlessDan))
}

//神盾尖刺祝福丹
func handleShieldBlessDan(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	shieldManager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbodyshield.PlayerBodyShieldDataManager)
	shieldInfo := shieldManager.GetBodyShiedInfo()
	shieldAdvancedId := int32(shieldInfo.ShieldId)
	needAdvanceId := itemTemplate.TypeFlag1
	if shieldAdvancedId != needAdvanceId {
		log.WithFields(
			log.Fields{
				"playerId":         pl.GetId(),
				"shieldAdvancedId": shieldAdvancedId,
				"needAdvanceId":    itemTemplate.TypeFlag1,
			}).Warn("shield:祝福丹使用,神盾尖刺阶数不足")
		playerlogic.SendSystemMessage(pl, lang.ShieldAdvanceNotEqual)
		return
	}

	nextShieldTemp := bodyshield.GetBodyShieldService().GetShield(shieldAdvancedId + 1)
	addBlessNum := itemTemplate.TypeFlag2
	addTimes := int32(0) //优化成没有次数效果
	curTimes := shieldInfo.ShieldNum
	curTimes += addTimes
	// isAdvanced, pro := bodyshieldlogic.ShieldEatZhuFuDan(pl, curTimes, shieldInfo.Bless, addBlessNum, nextShieldTemp)
	isAdvanced, pro := commonlogic.AdvancedBlessDan(shieldInfo.Bless, addBlessNum, nextShieldTemp.NeedRate)
	shieldManager.ShieldFeed(pro, addTimes, isAdvanced)

	if isAdvanced {
		//同步属性
		bodyshieldlogic.ShieldPropertyChanged(pl)

		shieldReason := commonlog.ShieldLogReasonAdvanced
		reasonText := fmt.Sprintf(shieldReason.String(), commontypes.AdvancedTypeBlessDan.String())
		data := bodyshieldeventtypes.CreatePlayerShieldAdvancedLogEventData(shieldAdvancedId, 1, shieldReason, reasonText)
		gameevent.Emit(bodyshieldeventtypes.EventTypeShieldAdvancedLog, pl, data)
	}

	scShieldAdvanced := pbutil.BuildSCShieldAdvanced(shieldInfo.ShieldId, shieldInfo.ShieldPro, commontypes.AdvancedTypeBlessDan, false)
	pl.SendMsg(scShieldAdvanced)

	flag = true
	return
}
