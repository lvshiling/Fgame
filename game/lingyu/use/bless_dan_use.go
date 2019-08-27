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
	lingyueventtypes "fgame/fgame/game/lingyu/event/types"
	lingyulogic "fgame/fgame/game/lingyu/logic"
	"fgame/fgame/game/lingyu/pbutil"
	playerlingyu "fgame/fgame/game/lingyu/player"
	lingyutemplate "fgame/fgame/game/lingyu/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeLingyu, playerinventory.ItemUseHandleFunc(handleLingyuBlessDan))
}

//领域祝福丹
func handleLingyuBlessDan(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	lingyuManager := pl.GetPlayerDataManager(types.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	lingyuInfo := lingyuManager.GetLingyuInfo()
	lingyuAdvancedId := int32(lingyuInfo.AdvanceId)
	needAdvanceId := itemTemplate.TypeFlag1
	if lingyuAdvancedId != needAdvanceId {
		log.WithFields(
			log.Fields{
				"playerId":         pl.GetId(),
				"lingyuAdvancedId": lingyuAdvancedId,
				"needAdvanceId":    itemTemplate.TypeFlag1,
			}).Warn("lingyu:祝福丹使用,领域阶数不足")
		playerlogic.SendSystemMessage(pl, lang.LingyuAdvanceNotEqual)
		return
	}

	nextLingyuTemp := lingyutemplate.GetLingyuTemplateService().GetLingyuByNumber(lingyuAdvancedId + 1)
	addBlessNum := itemTemplate.TypeFlag2
	addTimes := int32(0) //优化成没有次数效果
	curTimes := lingyuInfo.TimesNum
	curTimes += addTimes
	// isAdvanced, pro := lingyulogic.LingEatZhuFuDan(pl, curTimes, lingyuInfo.Bless, addBlessNum, nextLingyuTemp)
	isAdvanced, pro := commonlogic.AdvancedBlessDan(lingyuInfo.Bless, addBlessNum, nextLingyuTemp.ZhufuMax)
	lingyuManager.LingyuAdvanced(pro, addTimes, isAdvanced)

	if isAdvanced {
		lingyuReason := commonlog.LingyuLogReasonAdvanced
		reasonText := fmt.Sprintf(lingyuReason.String(), commontypes.AdvancedTypeBlessDan.String())
		data := lingyueventtypes.CreatePlayerLingyuAdvancedLogEventData(lingyuAdvancedId, 1, lingyuReason, reasonText)
		gameevent.Emit(lingyueventtypes.EventTypeLingyuAdvancedLog, pl, data)

		//同步属性
		lingyulogic.LingyuPropertyChanged(pl)
		scLingyuAdvanced := pbutil.BuildSCLingyuAdavancedFinshed(int32(lingyuInfo.AdvanceId), lingyuInfo.LingyuId, commontypes.AdvancedTypeBlessDan)
		pl.SendMsg(scLingyuAdvanced)
	} else {
		//进阶不成功
		scLingyuAdvanced := pbutil.BuildSCLingyuAdavanced(int32(lingyuInfo.AdvanceId), lingyuInfo.LingyuId, addBlessNum, lingyuInfo.Bless, lingyuInfo.BlessTime, commontypes.AdvancedTypeBlessDan, false)
		pl.SendMsg(scLingyuAdvanced)
	}

	flag = true
	return
}
