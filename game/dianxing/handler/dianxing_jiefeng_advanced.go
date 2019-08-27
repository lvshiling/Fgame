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
	processor.Register(codec.MessageType(uipb.MessageType_CS_DIANXING_JIEFENG_ADVANCED_TYPE), dispatch.HandlerFunc(handleDianXingJieFengAdvanced))
}

//处理点星系统解封信息
func handleDianXingJieFengAdvanced(s session.Session, msg interface{}) (err error) {
	log.Debug("dianxing:处理点星系统解封信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csDianXingJieFengAdvanced := msg.(*uipb.CSDianxingJiefengAdvanced)
	buyFlag := csDianXingJieFengAdvanced.GetBuyFlag()

	err = dianXingJieFengAdvanced(tpl, buyFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("dianxing:处理点星系统解封信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("dianxing:处理点星系统解封完成")
	return nil
}

//点星系统解封的逻辑
func dianXingJieFengAdvanced(pl player.Player, buyFlag bool) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeDianXingJieFeng) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("inventory:升级失败,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	dianXingManager := pl.GetPlayerDataManager(types.PlayerDianXingDataManagerType).(*playerdianxing.PlayerDianXingDataManager)
	dianXingInfo := dianXingManager.GetDianXingObject()
	nextTemplate := dianXingManager.GetNextDianXingJieFengTemplate()
	if nextTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("dianxing:点星系统已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.DianXingJieFengAdanvacedReachedLimit)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//解封需要消耗的银两
	costSilver := int64(nextTemplate.UseSilver)
	//解封需要的消耗的绑元
	costBindGold := int32(0)
	//解封需要消耗的元宝
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
			}).Warn("dianxing:点星系统物品不足,无法解封")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//自动解封
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
				}).Warn("dianxing:购买物品失败,自动解封已停止")
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
			}).Warn("dianxing:银两不足,无法解封")
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
			}).Warn("dianxing:元宝不足,无法解封")
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
			}).Warn("dianxing:元宝不足,无法解封")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗钱
	reasonGoldText := fmt.Sprintf(commonlog.GoldLogReasonDianXingJieFengAdvanced.String(), dianXingInfo.JieFengLev, dianXingInfo.JieFengBless, dianXingInfo.JieFengTimes)
	reasonSliverText := fmt.Sprintf(commonlog.SilverLogReasonDianXingJieFengAdvanced.String(), dianXingInfo.JieFengLev, dianXingInfo.JieFengBless, dianXingInfo.JieFengTimes)
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), commonlog.GoldLogReasonDianXingJieFengAdvanced, reasonGoldText, costSilver, commonlog.SilverLogReasonDianXingJieFengAdvanced, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("dianxing: dianXingAdvanced Cost should be ok"))
	}

	//消耗物品
	if itemCount != 0 {
		inventoryReason := commonlog.InventoryLogReasonDianXingJieFengAdvanced
		reasonText := fmt.Sprintf(inventoryReason.String(), dianXingInfo.JieFengLev, dianXingInfo.JieFengBless, dianXingInfo.JieFengTimes)
		flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("dianxing: dianXingJieFengAdvanced use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//同步元宝
	if costGold != 0 || costSilver != 0 || costBindGold != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	//解封判断
	sucess, pro, _, addTimes, isDouble := dianxinglogic.DianXingJieFengAdvanced(pl, dianXingInfo.JieFengTimes, dianXingInfo.JieFengBless, nextTemplate)
	beforeLev := dianXingInfo.JieFengLev
	dianXingManager.DianXingJieFengAdvanced(pro, addTimes, sucess)
	if sucess {
		//日志
		dianXingReason := commonlog.DianXingLogReasonJieFengAdvanced
		reasonText := fmt.Sprintf(dianXingReason.String(), commontypes.AdvancedTypeJinJieDan.String())
		data := dianxingeventtypes.CreatePlayerDianXingJieFengAdvancedLogEventData(beforeLev, dianXingReason, reasonText)
		gameevent.Emit(dianxingeventtypes.EventTypeDianXingJieFengAdvancedLog, pl, data)

		//同步属性
		dianxinglogic.DianXingPropertyChanged(pl)
		advancedJieFengFinish(pl, dianXingInfo)
	} else {
		//解封不成功
		advancedJieFengBless(pl, dianXingInfo, isDouble, buyFlag)
	}
	return
}

func advancedJieFengFinish(pl player.Player, info *playerdianxing.PlayerDianXingObject) (err error) {
	scDianXingJieFengAdvanced := pbutil.BuildSCDianXingJieFengAdavancedFinshed(info, commontypes.AdvancedTypeXingChen)
	pl.SendMsg(scDianXingJieFengAdvanced)
	return
}

func advancedJieFengBless(pl player.Player, info *playerdianxing.PlayerDianXingObject, isDouble bool, isAutoBuy bool) (err error) {
	scDianXingJieFengAdvanced := pbutil.BuildSCDianXingJieFengAdavanced(info, commontypes.AdvancedTypeXingChen, isDouble, isAutoBuy)
	pl.SendMsg(scDianXingJieFengAdvanced)
	return
}
