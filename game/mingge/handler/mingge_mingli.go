package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	minggeeventtypes "fgame/fgame/game/mingge/event/types"
	minggelogic "fgame/fgame/game/mingge/logic"
	"fgame/fgame/game/mingge/pbutil"
	playermingge "fgame/fgame/game/mingge/player"
	minggetypes "fgame/fgame/game/mingge/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	shoplogic "fgame/fgame/game/shop/logic"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_MINGGE_MINGLI_BAPTIZE_TYPE), dispatch.HandlerFunc(handleMingGeMingLiBaptize))
}

//处理命理洗炼信息
func handleMingGeMingLiBaptize(s session.Session, msg interface{}) (err error) {
	log.Debug("mingge:处理命理洗炼信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMingGeMingLiBaptize := msg.(*uipb.CSMingGeMingLiBaptize)
	autoBuy := csMingGeMingLiBaptize.GetAutoBuy()
	mingGongType := csMingGeMingLiBaptize.GetMingGongType()
	posTag := csMingGeMingLiBaptize.GetPosTag()
	slotList := csMingGeMingLiBaptize.GetSlotList()

	err = mingGeMingLiBaptize(tpl, autoBuy, minggetypes.MingGongType(mingGongType), minggetypes.MingGongAllSubType(posTag), slotList)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"autoBuy":      autoBuy,
				"mingGongType": mingGongType,
				"posTag":       posTag,
				"slotList":     slotList,
				"error":        err,
			}).Error("mingge:处理命理洗炼信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("mingge:处理命理洗炼信息完成")
	return nil
}

//处理命理洗炼信息逻辑
func mingGeMingLiBaptize(pl player.Player, autoBuy bool,
	mingGongType minggetypes.MingGongType, mingGongSubType minggetypes.MingGongAllSubType, slotList []int32) (err error) {

	if !mingGongType.Valid() || !mingGongSubType.Valid() || len(slotList) == 0 {
		log.WithFields(log.Fields{
			"playerId":        pl.GetId(),
			"autoBuy":         autoBuy,
			"mingGongType":    mingGongType,
			"mingGongSubType": mingGongSubType,
			"slotList":        slotList,
		}).Warn("mingge:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	var mingLiSlotTypeList []minggetypes.MingLiSlotType
	for _, slot := range slotList {
		slotType := minggetypes.MingLiSlotType(slot)
		if !slotType.Vaild() {
			log.WithFields(log.Fields{
				"playerId":        pl.GetId(),
				"autoBuy":         autoBuy,
				"mingGongType":    mingGongType,
				"mingGongSubType": mingGongSubType,
				"slotList":        slotList,
			}).Warn("mingge:参数无效")
			playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
			return
		}
		mingLiSlotTypeList = append(mingLiSlotTypeList, slotType)
	}

	manager := pl.GetPlayerDataManager(types.PlayerMingGeDataManagerType).(*playermingge.PlayerMingGeDataManager)
	mingLiObj := manager.GetMingGeMingLiByTypeAndSubType(mingGongType, mingGongSubType)
	if mingLiObj == nil {
		log.WithFields(log.Fields{
			"playerId":        pl.GetId(),
			"autoBuy":         autoBuy,
			"mingGongType":    mingGongType,
			"mingGongSubType": mingGongSubType,
			"slotList":        slotList,
		}).Warn("mingge:参数无效")
		playerlogic.SendSystemMessage(pl, lang.MingGeMingLiBaptize)
		return
	}

	//物品判断
	needItemMap, flag := manager.GetMingLiBaptizeNeedAllNum(mingGongType, mingGongSubType, mingLiSlotTypeList)
	if !flag {
		panic("mingge: mingGeMingLiBaptize GetMingLiBaptizeNeedAllNum should be ok")
	}
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(needItemMap) != 0 {
		if !inventoryManager.HasEnoughItems(needItemMap) && !autoBuy {
			log.WithFields(log.Fields{
				"playerId":        pl.GetId(),
				"autoBuy":         autoBuy,
				"mingGongType":    mingGongType,
				"mingGongSubType": mingGongSubType,
				"slotList":        slotList,
			}).Warn("mingge:命格物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	//获取背包物品和需要购买物品
	items, buyItems := inventoryManager.GetItemsAndNeedBuy(needItemMap)
	//计算需要元宝等
	if len(buyItems) != 0 {
		bindGold := int32(0)
		gold := int32(0)
		sliver := int64(0)
		propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

		isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayerMap(pl, buyItems)
		if !isEnoughBuyTimes {
			log.WithFields(log.Fields{
				"playerId":        pl.GetId(),
				"autoBuy":         autoBuy,
				"mingGongType":    mingGongType,
				"mingGongSubType": mingGongSubType,
				"slotList":        slotList,
			}).Warn("mingge:购买物品失败,自动洗练已停止")
			playerlogic.SendSystemMessage(pl, lang.ShopMingLiAutoBuyItemFail)
			return
		}

		shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
		gold += int32(shopNeedGold)
		bindGold += int32(shopNeedBindGold)
		sliver += shopNeedSilver

		flag = propertyManager.HasEnoughCost(int64(bindGold), int64(gold), sliver)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":        pl.GetId(),
				"autoBuy":         autoBuy,
				"mingGongType":    mingGongType,
				"mingGongSubType": mingGongSubType,
				"slotList":        slotList,
			}).Warn("mingge:元宝不足，无法洗练")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}

		reasonGoldText := commonlog.GoldLogReasonMingGeBaptizeCost.String()
		reasonSliverText := commonlog.SilverLogReasonMingGeMingLiBaptizeCost.String()
		flag = propertyManager.Cost(int64(bindGold), int64(gold), commonlog.GoldLogReasonWeapUpstar, reasonGoldText, sliver, commonlog.SilverLogReasonWeapUpstar, reasonSliverText)
		if !flag {
			panic(fmt.Errorf("mingge: MingLiBaptize Cost should be ok"))
		}
		propertylogic.SnapChangedProperty(pl)
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗物品
	if len(items) != 0 {
		reasonText := commonlog.InventoryLogReasonMingGeMingLiBaptizeUse.String()
		flag := inventoryManager.BatchRemove(items, commonlog.InventoryLogReasonMingGeMingLiBaptizeUse, reasonText)
		if !flag {
			panic(fmt.Errorf("mingge: MingLiBaptize use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	mingGongTypeMap, flag := manager.MingLiBaptize(mingGongType, mingGongSubType, mingLiSlotTypeList)
	if !flag {
		panic(fmt.Errorf("mingge: MingLiBaptize use item should be ok"))
	}
	//命宫激活
	if len(mingGongTypeMap) != 0 {
		mingLiMap := manager.GetMingLiMap()
		scMingGeMingGongActivate := pbutil.BuildSCMingGeMingGongActivate(mingLiMap, mingGongTypeMap)
		pl.SendMsg(scMingGeMingGongActivate)
	}

	eventdata := minggeeventtypes.CreatePlayerMingGeMingLiEventData(needItemMap)
	gameevent.Emit(minggeeventtypes.EventTypeMingGeMingLi, pl, eventdata)

	// 属性变化
	minggelogic.MingGePropertyChanged(pl)
	obj := manager.GetMingGeMingLiByTypeAndSubType(mingGongType, mingGongSubType)
	scMingGeMingLiBaptize := pbutil.BuildSCMingGeMingLiBaptize(int32(mingGongType), int32(mingGongSubType), obj, slotList)
	pl.SendMsg(scMingGeMingLiBaptize)
	return
}
