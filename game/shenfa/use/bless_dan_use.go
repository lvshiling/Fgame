package use

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	commonlogic "fgame/fgame/game/common/logic"
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
	shenfatemplate "fgame/fgame/game/shenfa/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeShenfa, playerinventory.ItemUseHandleFunc(handleShenfaBlessDan))
}

//身法祝福丹
func handleShenfaBlessDan(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	shenfaManager := pl.GetPlayerDataManager(types.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
	shenfaInfo := shenfaManager.GetShenfaInfo()
	shenfaAdvancedId := int32(shenfaInfo.AdvanceId)
	needAdvanceId := itemTemplate.TypeFlag1
	if shenfaAdvancedId != needAdvanceId {
		log.WithFields(
			log.Fields{
				"playerId":         pl.GetId(),
				"shenfaAdvancedId": shenfaAdvancedId,
				"needAdvanceId":    itemTemplate.TypeFlag1,
			}).Warn("shenfa:祝福丹使用,身法阶数不足")
		playerlogic.SendSystemMessage(pl, lang.ShenfaAdvanceNotEqual)
		return
	}

	nextShenfaTemp := shenfatemplate.GetShenfaTemplateService().GetShenfaByNumber(shenfaAdvancedId + 1)
	addBlessNum := itemTemplate.TypeFlag2
	addTimes := int32(0) //优化成没有次数效果
	curTimes := shenfaInfo.TimesNum
	curTimes += addTimes
	// isAdvanced, pro := shenfalogic.ShenFaEatZhuFuDan(pl, curTimes, shenfaInfo.Bless, addBlessNum, nextShenfaTemp)
	isAdvanced, pro := commonlogic.AdvancedBlessDan(shenfaInfo.Bless, addBlessNum, nextShenfaTemp.ZhufuMax)
	shenfaManager.ShenfaAdvanced(pro, addTimes, isAdvanced)

	if isAdvanced {
		shenfaReason := commonlog.ShenfaLogReasonAdvanced
		reasonText := fmt.Sprintf(shenfaReason.String(), commontypes.AdvancedTypeBlessDan.String())
		data := shenfaeventtypes.CreatePlayerShenfaAdvancedLogEventData(shenfaAdvancedId, 1, shenfaReason, reasonText)
		gameevent.Emit(shenfaeventtypes.EventTypeShenfaAdvancedLog, pl, data)

		//同步属性
		shenfalogic.ShenfaPropertyChanged(pl)
		scShenfaAdvanced := pbutil.BuildSCShenfaAdavancedFinshed(int32(shenfaInfo.AdvanceId), shenfaInfo.ShenfaId, commontypes.AdvancedTypeBlessDan)
		pl.SendMsg(scShenfaAdvanced)
	} else {
		//进阶不成功
		scShenfaAdvanced := pbutil.BuildSCShenfaAdavanced(int32(shenfaInfo.AdvanceId), shenfaInfo.ShenfaId, addBlessNum, shenfaInfo.Bless, shenfaInfo.BlessTime, commontypes.AdvancedTypeBlessDan, false)
		pl.SendMsg(scShenfaAdvanced)
	}

	flag = true
	return
}
