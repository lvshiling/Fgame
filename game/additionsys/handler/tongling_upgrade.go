package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	additionsyslogic "fgame/fgame/game/additionsys/logic"
	"fgame/fgame/game/additionsys/pbutil"
	playeradditionsys "fgame/fgame/game/additionsys/player"
	additionsystemplate "fgame/fgame/game/additionsys/template"
	additionsystypes "fgame/fgame/game/additionsys/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ADDITION_SYS_TONGLING_UPGRADE_TYPE), dispatch.HandlerFunc(handleAdditionSysTongLingUpgrade))
}

//处理附加系统通灵升级
func handleAdditionSysTongLingUpgrade(s session.Session, msg interface{}) (err error) {
	log.Debug("tongling:处理附加系统通灵升级")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csAdditionSysTongLingUpgrade := msg.(*uipb.CSAdditionSysTongLingUpgrade)
	num := csAdditionSysTongLingUpgrade.GetNum()
	sysTypeId := csAdditionSysTongLingUpgrade.GetSysType()
	sysType := additionsystypes.AdditionSysType(sysTypeId)

	//参数不对
	if num <= 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"num":      num,
		}).Warn("tongling:参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	if !sysType.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"sysType":  sysType.String(),
			}).Warn("tongling:系统通灵升级,类型错误")
		return
	}

	err = additionSysTongLingUpgrade(tpl, sysType, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"sysType":  sysType.String(),
				"num":      num,
				"error":    err,
			}).Warn("tongling:处理附加系统通灵升级,错误")
		return
	}

	log.Debug("tongling:处理附加系统通灵升级,完成")
	return nil

}

//系统通灵升级逻辑
// func additionSysTongLingUpgrade(pl player.Player, typ additionsystypes.AdditionSysType, autoFlag bool) (err error) {
// 	if !additionsyslogic.GetAdditionSysTongLingFuncOpenByType(pl, typ) {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"typ":      typ.String(),
// 			}).Warn("tongling:通灵升级失败,功能未开启")
// 		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
// 		return
// 	}

// 	additionsysManager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
// 	tongLingInfo := additionsysManager.GetAdditionSysTongLingInfoByType(typ)

// 	//判断槽位是否可以升
// 	nextTemplate := tongLingInfo.GetNextTongLingTemplate()
// 	if nextTemplate == nil {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"typ":      typ.String(),
// 				"autoFlag": autoFlag,
// 			}).Warn("tongling:升级通灵失败,已经满级")
// 		playerlogic.SendSystemMessage(pl, lang.AdditionSysLevelHighest)
// 		return
// 	}

// 	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
// 	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

// 	//进阶需要消耗的银两
// 	costSilver := int64(0)
// 	//进阶需要消耗的元宝
// 	costGold := int32(0)
// 	//进阶需要消耗的绑元
// 	costBindGold := int32(0)

// 	//需要消耗物品
// 	useTemp := additionsystemplate.GetAdditionSysTemplateService().GetTongLingUseByType(typ)
// 	itemCount := int32(0)
// 	totalNum := int32(0)
// 	useItem := useTemp.UseItem
// 	isEnoughBuyTimes := true
// 	shopIdMap := make(map[int32]int32)
// 	itemCount = nextTemplate.GetItemCount()
// 	totalNum = inventoryManager.NumOfItems(int32(useItem))
// 	if totalNum < itemCount {
// 		if autoFlag == false {
// 			log.WithFields(log.Fields{
// 				"playerId": pl.GetId(),
// 				"typ":      typ.String(),
// 				"autoFlag": autoFlag,
// 			}).Warn("tongling:升级通灵,物品不足无法升级")
// 			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
// 			return
// 		}
// 		//自动进阶
// 		needBuyNum := itemCount - totalNum
// 		itemCount = totalNum

// 		if needBuyNum > 0 {
// 			if !shop.GetShopService().ShopIsSellItem(useItem) {
// 				log.WithFields(log.Fields{
// 					"playerId": pl.GetId(),
// 					"typ":      typ.String(),
// 					"autoFlag": autoFlag,
// 				}).Warn("tongling:升级通灵,商铺没有该道具无法自动购买")
// 				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
// 				return
// 			}

// 			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItem, needBuyNum)
// 			if !isEnoughBuyTimes {
// 				log.WithFields(log.Fields{
// 					"playerId": pl.GetId(),
// 					"typ":      typ.String(),
// 					"autoFlag": autoFlag,
// 				}).Warn("tongling:升级通灵,购买物品失败自动升级已停止")
// 				playerlogic.SendSystemMessage(pl, lang.ShopAdvancedAutoBuyItemFail)
// 				return
// 			}

// 			shopNeedBindGold, shopNeedGold, shopNeedSilver := shoplogic.ShopCostData(pl, shopIdMap)
// 			costGold += int32(shopNeedGold)
// 			costBindGold += int32(shopNeedBindGold)
// 			costSilver += shopNeedSilver
// 		}
// 	}

// 	//是否足够银两
// 	if costSilver != 0 {
// 		flag := propertyManager.HasEnoughSilver(int64(costSilver))
// 		if !flag {
// 			log.WithFields(log.Fields{
// 				"playerId": pl.GetId(),
// 				"sysType":  typ.String(),
// 				"autoFlag": autoFlag,
// 			}).Warn("tongling:升级通灵,银两不足无法升级")
// 			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
// 			return
// 		}
// 	}
// 	//是否足够元宝
// 	if costGold != 0 {
// 		flag := propertyManager.HasEnoughGold(int64(costGold), false)
// 		if !flag {
// 			log.WithFields(log.Fields{
// 				"playerId": pl.GetId(),
// 				"sysType":  typ.String(),
// 				"autoFlag": autoFlag,
// 			}).Warn("tongling:升级通灵,元宝不足无法升级")
// 			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
// 			return
// 		}
// 	}
// 	//是否足够绑元
// 	needBindGold := costBindGold + costGold
// 	if needBindGold != 0 {
// 		flag := propertyManager.HasEnoughGold(int64(needBindGold), true)
// 		if !flag {
// 			log.WithFields(log.Fields{
// 				"playerId": pl.GetId(),
// 				"sysType":  typ.String(),
// 				"autoFlag": autoFlag,
// 			}).Warn("tongling:升级通灵,元宝不足无法升级")
// 			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
// 			return
// 		}
// 	}

// 	//更新自动购买每日限购次数
// 	if len(shopIdMap) != 0 {
// 		shoplogic.ShopDayCountChanged(pl, shopIdMap)
// 	}

// 	//消耗钱
// 	goldUseReason := commonlog.GoldLogReasonAdditionSysTongLingCost
// 	silverUseReason := commonlog.SilverLogReasonAdditionSysTongLingCost
// 	goldUseReasonText := fmt.Sprintf(goldUseReason.String(), typ.String(), tongLingInfo.TongLingLev, tongLingInfo.TongLingPro, tongLingInfo.TongLingNum)
// 	silverUseReasonText := fmt.Sprintf(silverUseReason.String(), typ.String(), tongLingInfo.TongLingLev, tongLingInfo.TongLingPro, tongLingInfo.TongLingNum)
// 	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), goldUseReason, goldUseReasonText, costSilver, silverUseReason, silverUseReasonText)
// 	if !flag {
// 		panic(fmt.Errorf("tongling:additionsys tongling Cost should be ok"))
// 	}

// 	//消耗物品
// 	if itemCount != 0 {
// 		inventoryReason := commonlog.InventoryLogReasonAdditionSysTongLingCost
// 		reasonText := fmt.Sprintf(inventoryReason.String(), typ.String(), tongLingInfo.TongLingLev, tongLingInfo.TongLingPro, tongLingInfo.TongLingNum)
// 		flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
// 		if !flag {
// 			panic(fmt.Errorf("tongling: tongling use item should be ok"))
// 		}
// 		inventorylogic.SnapInventoryChanged(pl)
// 	}

// 	//同步元宝
// 	if costGold != 0 || costSilver != 0 || costBindGold != 0 {
// 		propertylogic.SnapChangedProperty(pl)
// 	}

// 	beforeLev := tongLingInfo.TongLingLev
// 	//进阶判断
// 	sucess, pro, _, addTimes := additionsyslogic.AdditionSysUpgradeCommonJudge(pl, tongLingInfo.TongLingNum, tongLingInfo.TongLingPro, nextTemplate)
// 	tongLingInfo.TongLingUpgrade(pro, addTimes, sucess)

// 	if sucess {
// 		//更新属性
// 		additionsyslogic.UpdataAdditionSysPropertyByType(pl, typ)
// 		//日志
// 		additionsysReason := commonlog.AdditionSysLogReasonTongLing
// 		reasonText := fmt.Sprintf(additionsysReason.String(), typ.String(), commontypes.AdvancedTypeJinJieDan.String())
// 		data := additionsyseventtypes.CreatePlayerAdditionSysTongLingLogEventData(typ, beforeLev, additionsysReason, reasonText)
// 		gamevent.Emit(additionsyseventtypes.EventTypeAdditionSysTongLingLog, pl, data)
// 	}

// 	scMsg := pbutil.BuildSCAdditionSysTongLingUpgrade(typ, tongLingInfo.TongLingLev, tongLingInfo.TongLingPro)
// 	pl.SendMsg(scMsg)
// 	return
// }

//系统通灵升级逻辑
func additionSysTongLingUpgrade(pl player.Player, typ additionsystypes.AdditionSysType, num int32) (err error) {
	if !additionsyslogic.GetAdditionSysTongLingFuncOpenByType(pl, typ) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("tongling:通灵升级失败,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	additionsysManager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	tongLingInfo := additionsysManager.GetAdditionSysTongLingInfoByType(typ)
	culLevel := tongLingInfo.TongLingLev

	//判断槽位是否可以升
	nextTemplate := tongLingInfo.GetNextTongLingTemplate()
	if nextTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"num":      num,
			}).Warn("tongling:升级通灵失败,已经满级")
		playerlogic.SendSystemMessage(pl, lang.AdditionSysLevelHighest)
		return
	}

	reachTemplate, flag := additionsystemplate.GetAdditionSysTemplateService().GetTongLingReachByArg(typ, culLevel, num)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ.String(),
			"num":      num,
		}).Warn("tongling:升级通灵失败,参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//需要消耗物品
	useTemp := additionsystemplate.GetAdditionSysTemplateService().GetTongLingUseByType(typ)
	useItem := useTemp.UseItem
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"num":      num,
			}).Warn("tongling:当前系统通灵食用丹数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		inventoryReason := commonlog.InventoryLogReasonAdditionSysTongLingCost
		reasonText := fmt.Sprintf(inventoryReason.String(), typ.String(), tongLingInfo.TongLingLev, tongLingInfo.TongLingPro, tongLingInfo.TongLingNum)
		flag = inventoryManager.UseItem(useItem, num, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("tongling: tongling use item should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	tongLingInfo.TongLingUpgrade(reachTemplate.GetLevel())
	//更新属性
	additionsyslogic.UpdataAdditionSysPropertyByType(pl, typ)

	scMsg := pbutil.BuildSCAdditionSysTongLingUpgrade(typ, tongLingInfo.TongLingLev, tongLingInfo.TongLingPro)
	pl.SendMsg(scMsg)
	return
}
