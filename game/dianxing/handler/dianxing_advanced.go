package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	commontypes "fgame/fgame/game/common/types"
	dianxingeventtypes "fgame/fgame/game/dianxing/event/types"
	dianxinglogic "fgame/fgame/game/dianxing/logic"
	"fgame/fgame/game/dianxing/pbutil"
	playerdianxing "fgame/fgame/game/dianxing/player"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_DIANXING_ADVANCED_TYPE), dispatch.HandlerFunc(handleDianXingAdvanced))
}

//处理点星系统进阶信息
func handleDianXingAdvanced(s session.Session, msg interface{}) (err error) {
	log.Debug("dianxing:处理点星系统进阶信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csDianXingAdvanced := msg.(*uipb.CSDianxingAdvanced)
	buyFlag := csDianXingAdvanced.GetBuyFlag()
	fuFlag := csDianXingAdvanced.GetFuFlag()

	err = dianXingAdvanced(tpl, buyFlag, fuFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("dianxing:处理点星系统进阶信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("dianxing:处理点星系统进阶完成")
	return nil
}

//点星系统进阶的逻辑
func dianXingAdvanced(pl player.Player, buyFlag bool, fuFlag bool) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeDianXing) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("inventory:升级失败,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	dianXingManager := pl.GetPlayerDataManager(types.PlayerDianXingDataManagerType).(*playerdianxing.PlayerDianXingDataManager)
	dianXingInfo := dianXingManager.GetDianXingObject()
	nextTemplate := dianXingManager.GetNextDianXingTemplate()
	if nextTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("dianxing:点星系统已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.DianXingAdanvacedReachedLimit)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//需要消耗的星尘值
	needXcNum := int64(nextTemplate.UseXingChen)
	//进阶需要消耗的银两
	costSilver := int64(nextTemplate.UseSilver)
	//进阶需要的消耗的绑元
	costBindGold := int32(0)
	//进阶需要消耗的元宝
	costGold := int32(0)

	//需要消耗物品
	itemCount := int32(0)
	totalNum := int32(0)
	useItem := nextTemplate.UseItem
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	useItemTemplate := nextTemplate.GetUseItemTemplate()
	if useItemTemplate != nil {
		itemCount = nextTemplate.ItemCount
		totalNum = inventoryManager.NumOfItems(int32(useItem))
	}
	if totalNum < itemCount {
		if buyFlag == false {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"buyFlag":  buyFlag,
			}).Warn("dianxing:点星系统进物品不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//自动进阶
		needBuyNum := itemCount - totalNum
		itemCount = totalNum

		if needBuyNum > 0 {
			if !shop.GetShopService().ShopIsSellItem(useItem) {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"buyFlag":  buyFlag,
				}).Warn("dianxing:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItem, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"buyFlag":  buyFlag,
				}).Warn("dianxing:购买物品失败,自动进阶已停止")
				playerlogic.SendSystemMessage(pl, lang.ShopAdvancedAutoBuyItemFail)
				return
			}

			shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
			costGold += int32(shopNeedGold)
			costBindGold += int32(shopNeedBindGold)
			costSilver += shopNeedSilver
		}
	}

	//是否足够银两
	if costSilver != 0 {
		flag := propertyManager.HasEnoughSilver(int64(costSilver))
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"buyFlag":  buyFlag,
			}).Warn("dianxing:银两不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}
	//是否足够元宝
	if costGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(costGold), false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"buyFlag":  buyFlag,
			}).Warn("dianxing:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//是否足够绑元
	needBindGold := costBindGold + costGold
	if needBindGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(needBindGold), true)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"buyFlag":  buyFlag,
			}).Warn("dianxing:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//是否足够星尘值
	if needXcNum > 0 && needXcNum > dianXingInfo.XingChenNum {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"autoFlag": buyFlag,
		}).Warn("dianxing:星尘值不足,无法进阶")
		playerlogic.SendSystemMessage(pl, lang.DianXingAdvanceNotXingChen)
		return
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗星尘值
	if needXcNum > 0 {
		flag := dianXingManager.SubXingChenNum(needXcNum)
		if !flag {
			panic(fmt.Errorf("dianxing: dianxingAdvanced use xingchenNum should be ok"))
		}
	}

	//消耗钱
	reasonGoldText := fmt.Sprintf(commonlog.GoldLogReasonDianXingAdvanced.String(), dianXingInfo.CurrType, dianXingInfo.CurrLevel, dianXingInfo.DianXingBless, dianXingInfo.DianXingTimes)
	reasonSliverText := fmt.Sprintf(commonlog.SilverLogReasonDianXingAdvanced.String(), dianXingInfo.CurrType, dianXingInfo.CurrLevel, dianXingInfo.DianXingBless, dianXingInfo.DianXingTimes)
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), commonlog.GoldLogReasonDianXingAdvanced, reasonGoldText, costSilver, commonlog.SilverLogReasonDianXingAdvanced, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("dianxing: dianXingAdvanced Cost should be ok"))
	}

	//消耗物品
	if itemCount != 0 {
		inventoryReason := commonlog.InventoryLogReasonDianXingAdvanced
		reasonText := fmt.Sprintf(inventoryReason.String(), dianXingInfo.CurrType, dianXingInfo.CurrLevel, dianXingInfo.DianXingBless, dianXingInfo.DianXingTimes)
		flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("dianxing: dianXingAdvanced use item should be ok"))
		}

	}

	autoItem := item.GetItemService().GetItemTemplate(itemtypes.ItemTypeAutoUseRes, itemtypes.ItemAutoUseResSubTypeXingChen)
	eventdata := dianxingeventtypes.CreatePlayerDianXingUseItemData(int32(autoItem.Id), nextTemplate.UseXingChen)
	gameevent.Emit(dianxingeventtypes.EventTypeDianXingUseItem, pl, eventdata)

	if itemCount > 0 {
		inventorylogic.SnapInventoryChanged(pl)
	}

	//同步元宝
	if costGold != 0 || costSilver != 0 || costBindGold != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	//进阶判断
	sucess, pro, _, addTimes, isDouble := dianxinglogic.DianXingAdvanced(pl, dianXingInfo.DianXingTimes, dianXingInfo.DianXingBless, 0, nextTemplate)
	beforeXingPu, beforeLev := dianXingInfo.CurrType, dianXingInfo.CurrLevel
	dianXingManager.DianXingAdvanced(pro, addTimes, sucess)
	if sucess {
		//日志
		dianXingReason := commonlog.DianXingLogReasonAdvanced
		reasonText := fmt.Sprintf(dianXingReason.String(), commontypes.AdvancedTypeXingChen.String())
		data := dianxingeventtypes.CreatePlayerDianXingAdvancedLogEventData(beforeXingPu, beforeLev, dianXingReason, reasonText)
		gameevent.Emit(dianxingeventtypes.EventTypeDianXingAdvancedLog, pl, data)

		//同步属性
		dianxinglogic.DianXingPropertyChanged(pl)
		advancedFinish(pl, dianXingInfo)
	} else {
		//进阶不成功
		advancedBless(pl, dianXingInfo, isDouble, buyFlag, fuFlag)
	}
	return
}

func advancedFinish(pl player.Player, info *playerdianxing.PlayerDianXingObject) (err error) {
	scDianXingAdvanced := pbutil.BuildSCDianXingAdavancedFinshed(info, commontypes.AdvancedTypeXingChen)
	pl.SendMsg(scDianXingAdvanced)
	return
}

func advancedBless(pl player.Player, info *playerdianxing.PlayerDianXingObject, isDouble bool, isAutoBuy bool, isFu bool) (err error) {
	scDianXingAdvanced := pbutil.BuildSCDianXingAdavanced(info, commontypes.AdvancedTypeXingChen, isDouble, isAutoBuy, isFu)
	pl.SendMsg(scDianXingAdvanced)
	return
}
