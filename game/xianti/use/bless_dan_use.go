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
	playertypes "fgame/fgame/game/player/types"
	xiantieventtypes "fgame/fgame/game/xianti/event/types"
	xiantilogic "fgame/fgame/game/xianti/logic"
	"fgame/fgame/game/xianti/pbutil"
	playerxianti "fgame/fgame/game/xianti/player"
	"fgame/fgame/game/xianti/xianti"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeXianTi, playerinventory.ItemUseHandleFunc(handleXianTiBlessDan))
}

//仙体祝福丹
func handleXianTiBlessDan(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	xiantiManager := pl.GetPlayerDataManager(playertypes.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	xiantiInfo := xiantiManager.GetXianTiInfo()
	xiantiAdvancedId := int32(xiantiInfo.AdvanceId)
	needAdvanceId := itemTemplate.TypeFlag1
	if xiantiAdvancedId != needAdvanceId {
		log.WithFields(
			log.Fields{
				"playerId":         pl.GetId(),
				"xiantiAdvancedId": xiantiAdvancedId,
				"needAdvanceId":    itemTemplate.TypeFlag1,
			}).Warn("xianti:祝福丹使用,仙体阶数不足")
		playerlogic.SendSystemMessage(pl, lang.XianTiAdvanceNotEqual)
		return
	}

	nextXianTiTemp := xianti.GetXianTiService().GetXianTiNumber(xiantiAdvancedId + 1)
	addBlessNum := itemTemplate.TypeFlag2
	addTimes := int32(0) //优化成没有次数效果
	curTimes := xiantiInfo.TimesNum
	curTimes += addTimes
	// isAdvanced, pro := xiantilogic.XianTiEatZhuFuDan(pl, curTimes, xiantiInfo.Bless, addBlessNum, nextXianTiTemp)
	isAdvanced, pro := commonlogic.AdvancedBlessDan(xiantiInfo.Bless, addBlessNum, nextXianTiTemp.ZhufuMax)
	xiantiManager.XianTiAdvanced(pro, addTimes, isAdvanced)

	if isAdvanced {
		//同步属性
		xiantilogic.XianTiPropertyChanged(pl)

		xiantiReason := commonlog.XianTiLogReasonAdvanced
		reasonText := fmt.Sprintf(xiantiReason.String(), commontypes.AdvancedTypeBlessDan.String())
		data := xiantieventtypes.CreatePlayerXianTiAdvancedLogEventData(xiantiAdvancedId, 1, xiantiReason, reasonText)
		gameevent.Emit(xiantieventtypes.EventTypeXianTiAdvancedLog, pl, data)
	}

	scMsg := pbutil.BuildSCXianTiAdavanced(int32(xiantiInfo.AdvanceId), xiantiInfo.XianTiId, addBlessNum, xiantiInfo.Bless, xiantiInfo.BlessTime, commontypes.AdvancedTypeBlessDan, false)
	pl.SendMsg(scMsg)

	flag = true
	return
}
