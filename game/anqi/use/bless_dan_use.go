package use

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	anqieventtypes "fgame/fgame/game/anqi/event/types"
	anqilogic "fgame/fgame/game/anqi/logic"
	"fgame/fgame/game/anqi/pbutil"
	playeranqi "fgame/fgame/game/anqi/player"
	anqitemplate "fgame/fgame/game/anqi/template"
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
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeAnqi, playerinventory.ItemUseHandleFunc(handleAnqiBlessDan))
}

//暗器祝福丹
func handleAnqiBlessDan(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	anqiManager := pl.GetPlayerDataManager(types.PlayerAnqiDataManagerType).(*playeranqi.PlayerAnqiDataManager)
	anqiInfo := anqiManager.GetAnqiInfo()
	anqiAdvancedId := int32(anqiInfo.AdvanceId)
	needAdvanceId := itemTemplate.TypeFlag1
	if anqiAdvancedId != needAdvanceId {
		log.WithFields(
			log.Fields{
				"playerId":       pl.GetId(),
				"anqiAdvancedId": anqiAdvancedId,
				"needAdvanceId":  itemTemplate.TypeFlag1,
			}).Warn("anqi:祝福丹使用,暗器阶数不足")
		playerlogic.SendSystemMessage(pl, lang.AnqiAdvanceNotEqual)
		return
	}

	nextAnqiTemp := anqitemplate.GetAnqiTemplateService().GetAnqiNumber(anqiAdvancedId + 1)
	addBlessNum := itemTemplate.TypeFlag2
	addTimes := int32(0) //优化成没有次数效果
	curTimes := anqiInfo.TimesNum
	curTimes += addTimes
	isAdvanced, pro := commonlogic.AdvancedBlessDan(anqiInfo.Bless, addBlessNum, nextAnqiTemp.ZhufuMax)
	// isAdvanced, pro := anqilogic.AnqiEatZhuFuDan(pl, curTimes, anqiInfo.Bless, addBlessNum, nextAnqiTemp)
	anqiManager.AnqiAdvanced(pro, addTimes, isAdvanced)

	if isAdvanced {
		//同步属性
		anqilogic.AnqiPropertyChanged(pl)
		scAnqiAdvanced := pbutil.BuildSCAnqiAdavancedFinshed(int32(anqiInfo.AdvanceId), commontypes.AdvancedTypeBlessDan)
		pl.SendMsg(scAnqiAdvanced)

		anqiReason := commonlog.AnqiLogReasonAdvanced
		reasonText := fmt.Sprintf(anqiReason.String(), commontypes.AdvancedTypeBlessDan.String())
		data := anqieventtypes.CreatePlayerAnqiAdvancedLogEventData(anqiAdvancedId, 1, anqiReason, reasonText)
		gameevent.Emit(anqieventtypes.EventTypeAnqiAdvancedLog, pl, data)
	} else {
		//进阶不成功
		scAnqiAdvanced := pbutil.BuildSCAnqiAdavanced(int32(anqiInfo.AdvanceId), addBlessNum, anqiInfo.Bless, anqiInfo.BlessTime, commontypes.AdvancedTypeBlessDan, false)
		pl.SendMsg(scAnqiAdvanced)
	}

	flag = true
	return
}
