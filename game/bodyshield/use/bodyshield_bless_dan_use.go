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
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeBodyShield, playerinventory.ItemUseHandleFunc(handleBodyShieldBlessDan))
}

//护体盾祝福丹
func handleBodyShieldBlessDan(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	bodyshieldManager := pl.GetPlayerDataManager(types.PlayerBShieldDataManagerType).(*playerbodyshield.PlayerBodyShieldDataManager)
	bodyshieldInfo := bodyshieldManager.GetBodyShiedInfo()
	bodyshieldAdvancedId := int32(bodyshieldInfo.AdvanceId)
	needAdvanceId := itemTemplate.TypeFlag1
	if bodyshieldAdvancedId != needAdvanceId {
		log.WithFields(
			log.Fields{
				"playerId":             pl.GetId(),
				"bodyshieldAdvancedId": bodyshieldAdvancedId,
				"needAdvanceId":        itemTemplate.TypeFlag1,
			}).Warn("bodyshield:祝福丹使用,护体盾阶数不足")
		playerlogic.SendSystemMessage(pl, lang.BodyShieldAdvanceNotEqual)
		return
	}

	nextBodyshieldTemp := bodyshield.GetBodyShieldService().GetBodyShieldNumber(bodyshieldAdvancedId + 1)
	addBlessNum := itemTemplate.TypeFlag2
	addTimes := int32(0) //优化成没有次数效果
	curTimes := bodyshieldInfo.TimesNum
	curTimes += addTimes
	// isAdvanced, pro := bodyshieldlogic.BodyShieldEatZhuFuDan(pl, curTimes, bodyshieldInfo.Bless, addBlessNum, nextBodyshieldTemp)
	isAdvanced, pro := commonlogic.AdvancedBlessDan(bodyshieldInfo.Bless, addBlessNum, nextBodyshieldTemp.ZhufuMax)
	bodyshieldManager.BodyShieldAdvanced(pro, addTimes, isAdvanced)
	if isAdvanced {
		shieldReason := commonlog.ShieldLogReasonBodyAdvanced
		reasonText := fmt.Sprintf(shieldReason.String(), commontypes.AdvancedTypeBlessDan.String())
		data := bodyshieldeventtypes.CreatePlayerBodyShieldAdvancedLogEventData(bodyshieldAdvancedId, 1, shieldReason, reasonText)
		gameevent.Emit(bodyshieldeventtypes.EventTypeBodyShieldAdvancedLog, pl, data)

		//同步属性
		bodyshieldlogic.BodyShieldPropertyChanged(pl)
		scBodyShieldAdvanced := pbutil.BuildSCBodyShieldAdavancedFinshed(int32(bodyshieldInfo.AdvanceId), commontypes.AdvancedTypeBlessDan)
		pl.SendMsg(scBodyShieldAdvanced)
	} else {
		//进阶不成功
		scBodyShieldAdvanced := pbutil.BuildSCBodyShieldAdavanced(int32(bodyshieldInfo.AdvanceId), addBlessNum, bodyshieldInfo.Bless, bodyshieldInfo.BlessTime, commontypes.AdvancedTypeBlessDan, false)
		pl.SendMsg(scBodyShieldAdvanced)
	}
	flag = true
	return
}
