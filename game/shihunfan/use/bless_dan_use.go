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
	shihunfaneventtypes "fgame/fgame/game/shihunfan/event/types"
	shihunfanlogic "fgame/fgame/game/shihunfan/logic"
	"fgame/fgame/game/shihunfan/pbutil"
	playershihunfan "fgame/fgame/game/shihunfan/player"
	shihunfantemplate "fgame/fgame/game/shihunfan/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeShiHunFan, playerinventory.ItemUseHandleFunc(handleShiHunFanBlessDan))
}

//噬魂幡祝福丹
func handleShiHunFanBlessDan(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	shihunfanManager := pl.GetPlayerDataManager(playertypes.PlayerShiHunFanDataManagerType).(*playershihunfan.PlayerShiHunFanDataManager)
	shihunfanInfo := shihunfanManager.GetShiHunFanInfo()
	shihunfanAdvancedId := int32(shihunfanInfo.AdvanceId)
	needAdvanceId := itemTemplate.TypeFlag1
	if shihunfanAdvancedId != needAdvanceId {
		log.WithFields(
			log.Fields{
				"playerId":            pl.GetId(),
				"shihunfanAdvancedId": shihunfanAdvancedId,
				"needAdvanceId":       itemTemplate.TypeFlag1,
			}).Warn("shihunfan:祝福丹使用,噬魂幡阶数不足")
		playerlogic.SendSystemMessage(pl, lang.ShiHunFanAdvanceNotEqual)
		return
	}

	nextShiHunFanTemp := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanNumber(shihunfanAdvancedId + 1)
	addBlessNum := itemTemplate.TypeFlag2
	addTimes := int32(0) //优化成没有次数效果
	curTimes := shihunfanInfo.TimesNum
	curTimes += addTimes
	// isAdvanced, pro := shihunfanlogic.ShiHunFanEatZhuFuDan(pl, curTimes, shihunfanInfo.Bless, addBlessNum, nextShiHunFanTemp)
	isAdvanced, pro := commonlogic.AdvancedBlessDan(shihunfanInfo.Bless, addBlessNum, nextShiHunFanTemp.ZhufuMax)
	shihunfanManager.ShiHunFanAdvanced(pro, addTimes, isAdvanced)

	if isAdvanced {
		//同步属性
		shihunfanlogic.ShiHunFanPropertyChanged(pl)

		shihunfanReason := commonlog.ShiHunFanLogReasonAdvanced
		reasonText := fmt.Sprintf(shihunfanReason.String(), commontypes.AdvancedTypeBlessDan.String())
		data := shihunfaneventtypes.CreatePlayerShiHunFanAdvancedLogEventData(shihunfanAdvancedId, 1, shihunfanReason, reasonText)
		gameevent.Emit(shihunfaneventtypes.EventTypeShiHunFanAdvancedLog, pl, data)
	}

	scMsg := pbutil.BuildSCShiHunFanAdavanced(shihunfanInfo, commontypes.AdvancedTypeBlessDan, false, false)
	pl.SendMsg(scMsg)

	flag = true
	return
}
