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
	lingtongdeveventtypes "fgame/fgame/game/lingtongdev/event/types"
	lingtongdevlogic "fgame/fgame/game/lingtongdev/logic"
	"fgame/fgame/game/lingtongdev/pbutil"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	lingtongdevtemplate "fgame/fgame/game/lingtongdev/template"
	lingtongdevtypes "fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeLingBing, playerinventory.ItemUseHandleFunc(handleLingTongDevBlessDan))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeLingQi, playerinventory.ItemUseHandleFunc(handleLingTongDevBlessDan))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeLingYi, playerinventory.ItemUseHandleFunc(handleLingTongDevBlessDan))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeLingShen, playerinventory.ItemUseHandleFunc(handleLingTongDevBlessDan))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeLingTongYu, playerinventory.ItemUseHandleFunc(handleLingTongDevBlessDan))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeLingBao, playerinventory.ItemUseHandleFunc(handleLingTongDevBlessDan))
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBlessDan, itemtypes.ItemBlessDanSubTypeLingTi, playerinventory.ItemUseHandleFunc(handleLingTongDevBlessDan))
}

func getBlessDanLingTongDevSysType(itemSubType itemtypes.ItemSubType) (classType lingtongdevtypes.LingTongDevSysType, flag bool) {
	flag = true
	switch itemSubType {
	case itemtypes.ItemBlessDanSubTypeLingBing:
		classType = lingtongdevtypes.LingTongDevSysTypeLingBing
	case itemtypes.ItemBlessDanSubTypeLingQi:
		classType = lingtongdevtypes.LingTongDevSysTypeLingQi
	case itemtypes.ItemBlessDanSubTypeLingYi:
		classType = lingtongdevtypes.LingTongDevSysTypeLingYi
	case itemtypes.ItemBlessDanSubTypeLingShen:
		classType = lingtongdevtypes.LingTongDevSysTypeLingShen
	case itemtypes.ItemBlessDanSubTypeLingTongYu:
		classType = lingtongdevtypes.LingTongDevSysTypeLingYu
	case itemtypes.ItemBlessDanSubTypeLingBao:
		classType = lingtongdevtypes.LingTongDevSysTypeLingBao
	case itemtypes.ItemBlessDanSubTypeLingTi:
		classType = lingtongdevtypes.LingTongDevSysTypeLingTi
	default:
		flag = false
		return
	}
	return
}

//灵童养成类祝福丹
func handleLingTongDevBlessDan(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	itemSubType := itemTemplate.GetItemSubType()
	classType, isExist := getBlessDanLingTongDevSysType(itemSubType)
	if !isExist {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	lingTongDevInfo := manager.GetLingTongDevInfo(classType)
	// if lingTongDevInfo == nil {
	// 	lingTongDevInfo = manager.AdvancedInit(classType)
	// }
	curAfdvacedId := lingTongDevInfo.GetAdvancedId()
	needAdvanceId := itemTemplate.TypeFlag1
	if curAfdvacedId != needAdvanceId {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"curAfdvacedId": curAfdvacedId,
				"needAdvanceId": itemTemplate.TypeFlag1,
			}).Warn("lingtongdev:祝福丹使用,灵童养成类阶数不足")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevAdvanceNotEqual, classType.String())
		return
	}

	nextLingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevNumber(classType, curAfdvacedId+1)
	if nextLingTongDevTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"curAfdvacedId": curAfdvacedId,
				"needAdvanceId": itemTemplate.TypeFlag1,
			}).Warn("lingtongdev:祝福丹使用,灵童养成类阶数不足")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevAdanvacedReachedLimit, classType.String())
		return
	}
	addBlessNum := itemTemplate.TypeFlag2
	addTimes := int32(0) //优化成没有次数效果
	curTimes := lingTongDevInfo.GetTimesNum()
	curTimes += addTimes
	// isAdvanced, pro := lingtongdevlogic.LingTongDevEatZhuFuDan(pl, curTimes, lingTongDevInfo.GetBless(), addBlessNum, nextLingTongDevTemplate)
	isAdvanced, pro := commonlogic.AdvancedBlessDan(lingTongDevInfo.GetBless(), addBlessNum, nextLingTongDevTemplate.GetZhuFuMax())
	manager.LingTongDevAdvanced(classType, pro, addTimes, isAdvanced)

	if isAdvanced {
		lingTongDevReason := commonlog.LingTongDevLogReasonAdvanced
		reasonText := fmt.Sprintf(lingTongDevReason.String(), commontypes.AdvancedTypeBlessDan.String())
		data := lingtongdeveventtypes.CreatePlayerLingTongDevAdvancedLogEventData(classType, curAfdvacedId, 1, lingTongDevReason, reasonText)
		gameevent.Emit(lingtongdeveventtypes.EventTypeLingTongDevAdvancedLog, pl, data)

		//同步属性
		lingtongdevlogic.LingTongDevPropertyChanged(pl, classType)
		scLingTongDevAdavancedFinshed := pbutil.BuildSCLingTongDevAdavancedFinshed(int32(classType), lingTongDevInfo.GetAdvancedId(), lingTongDevInfo.GetSeqId(), commontypes.AdvancedTypeBlessDan)
		pl.SendMsg(scLingTongDevAdavancedFinshed)
	} else {
		//进阶不成功
		scLingTongDevAdavanced := pbutil.BuildSCLingTongDevAdavanced(int32(classType), lingTongDevInfo.GetAdvancedId(), lingTongDevInfo.GetSeqId(), addBlessNum, lingTongDevInfo.GetBless(), lingTongDevInfo.GetBlessTime(), commontypes.AdvancedTypeBlessDan, false)
		pl.SendMsg(scLingTongDevAdavanced)
	}

	flag = true
	return
}
