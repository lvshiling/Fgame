package use

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	commonlogic "fgame/fgame/game/common/logic"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	fabaoeventtypes "fgame/fgame/game/fabao/event/types"
	fabaologic "fgame/fgame/game/fabao/logic"
	"fgame/fgame/game/fabao/pbutil"
	playerfabao "fgame/fgame/game/fabao/player"
	fabaotemplate "fgame/fgame/game/fabao/template"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeFaBao, playerinventory.ItemUseHandleFunc(handleFaBaoBlessDan))
}

//法宝祝福丹
func handleFaBaoBlessDan(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	fabaoManager := pl.GetPlayerDataManager(playertypes.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	fabaoInfo := fabaoManager.GetFaBaoInfo()
	fabaoAdvancedId := fabaoInfo.GetAdvancedId()
	needAdvanceId := itemTemplate.TypeFlag1
	if fabaoAdvancedId != needAdvanceId {
		log.WithFields(
			log.Fields{
				"playerId":        pl.GetId(),
				"fabaoAdvancedId": fabaoAdvancedId,
				"needAdvanceId":   itemTemplate.TypeFlag1,
			}).Warn("fabao:祝福丹使用,法宝阶数不足")
		playerlogic.SendSystemMessage(pl, lang.FaBaoAdvanceNotEqual)
		return
	}

	nextFaBaoTemp := fabaotemplate.GetFaBaoTemplateService().GetFaBaoNumber(fabaoAdvancedId + 1)
	addBlessNum := itemTemplate.TypeFlag2
	addTimes := int32(0) //优化成没有次数效果
	curTimes := fabaoInfo.GetTimesNum()
	curTimes += addTimes
	// isAdvanced, pro := fabaologic.FaBaoEatZhuFuDan(pl, curTimes, fabaoInfo.GetBless(), addBlessNum, nextFaBaoTemp)
	isAdvanced, pro := commonlogic.AdvancedBlessDan(fabaoInfo.GetBless(), addBlessNum, nextFaBaoTemp.ZhufuMax)
	fabaoManager.FaBaoAdvanced(pro, addTimes, isAdvanced)

	if isAdvanced {
		//同步属性
		fabaologic.FaBaoPropertyChanged(pl)

		fabaoReason := commonlog.FaBaoLogReasonAdvanced
		reasonText := fmt.Sprintf(fabaoReason.String(), commontypes.AdvancedTypeBlessDan.String())
		data := fabaoeventtypes.CreatePlayerFaBaoAdvancedLogEventData(fabaoAdvancedId, 1, fabaoReason, reasonText)
		gameevent.Emit(fabaoeventtypes.EventTypeFaBaoAdvancedLog, pl, data)
	}

	scMsg := pbutil.BuildSCFaBaoAdavanced(fabaoInfo.GetAdvancedId(), fabaoInfo.GetFaBaoId(), addBlessNum, fabaoInfo.GetBless(), fabaoInfo.GetBlessTime(), commontypes.AdvancedTypeBlessDan, false)
	pl.SendMsg(scMsg)

	flag = true
	return
}
