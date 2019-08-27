package logic

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/common/common"
	commomlogic "fgame/fgame/game/common/logic"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	consttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/massacre/pbutil"
	playermassacre "fgame/fgame/game/massacre/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fgame/fgame/game/shop/shop"
	gametemplate "fgame/fgame/game/template"
	viplogic "fgame/fgame/game/vip/logic"
	viptypes "fgame/fgame/game/vip/types"
	"fmt"
	"math"

	commonlog "fgame/fgame/common/log"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	massacreeventtypes "fgame/fgame/game/massacre/event/types"

	massacretemplate "fgame/fgame/game/massacre/template"

	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"

	shoplogic "fgame/fgame/game/shop/logic"

	log "github.com/Sirupsen/logrus"
)

const (
	MassacreAdvancedResTypeLevSucceed  int32 = 1 + iota //升级成功
	MassacreAdvancedResTypeStarSucceed                  //升星成功
	MassacreAdvancedResTypeDefeated                     //升阶失败
)

//戮仙刃进阶
func HandleMassacreAdvanced(pl player.Player, autoFlag bool) (err error) {
	massacreManager := pl.GetPlayerDataManager(playertypes.PlayerMassacreDataManagerType).(*playermassacre.PlayerMassacreDataManager)
	massacreInfo := massacreManager.GetMassacreInfo()
	beforeNum := int32(massacreInfo.AdvanceId)
	beforeLev := massacreInfo.CurrLevel
	beforeShaQiNum := massacreInfo.ShaQiNum
	nextAdvancedId := massacreInfo.AdvanceId + 1
	massacreTemplate := massacretemplate.GetMassacreTemplateService().GetMassacre(nextAdvancedId)
	if massacreTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("massacre:戮仙刃已达最高阶")

		playerlogic.SendSystemMessage(pl, lang.MassacreAdanvacedReachedLimit)
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//需要消耗的杀气
	needSqNum := int64(massacreTemplate.UseGas)

	//进阶需要消耗的银两
	costSilver := int64(massacreTemplate.UseMoney)
	//进阶需要消耗的元宝
	costGold := int32(0)
	//进阶需要消耗的绑元
	costBindGold := int32(0)

	//需要消耗物品
	itemCount := int32(0)
	totalNum := int32(0)
	useItem := massacreTemplate.UseItem
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	useItemTemplate := massacreTemplate.GetUseItemTemplate()
	if useItemTemplate != nil {
		itemCount = massacreTemplate.ItemCount
		totalNum = inventoryManager.NumOfItems(int32(useItem))
	}
	if totalNum < itemCount {
		if autoFlag == false {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("massacre:戮仙刃进阶丹不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		//自动进阶
		needBuyNum := itemCount - totalNum
		itemCount = totalNum
		//获取价格
		// shopTemplate := shop.GetShopService().GetShopTemplateByItem(useItem)
		// if shopTemplate == nil {
		// 	log.WithFields(log.Fields{
		// 		"playerId": pl.GetId(),
		// 	}).Warn("massacre:商铺没有该道具,无法自动购买")
		// 	playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
		// 	return
		// }
		// shopNeedGold, shopNeedBindGold, shopNeedSilver := shopTemplate.GetConsumeData(needBuyNum)
		// costGold += shopNeedGold
		// costBindGold += shopNeedBindGold
		// costSilver += shopNeedSilver

		if needBuyNum > 0 {
			if !shop.GetShopService().ShopIsSellItem(useItem) {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("massacre:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItem, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("massacre:购买物品失败,自动进阶已停止")
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
				"autoFlag": autoFlag,
			}).Warn("massacre:银两不足,无法进阶")
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
				"autoFlag": autoFlag,
			}).Warn("massacre:元宝不足,无法进阶")
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
				"autoFlag": autoFlag,
			}).Warn("massacre:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	if needSqNum > 0 && needSqNum > massacreInfo.ShaQiNum {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"autoFlag": autoFlag,
		}).Warn("massacre:杀气不足,无法进阶")
		playerlogic.SendSystemMessage(pl, lang.MassacreAdvanceNotGen)
		return
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗钱
	goldUseReason := commonlog.GoldLogReasonMassacreAdvanced
	silverUseReason := commonlog.SilverLogReasonMassacreAdvanced
	goldUseReasonText := fmt.Sprintf(commonlog.GoldLogReasonMassacreAdvanced.String(), massacreInfo.CurrLevel, massacreInfo.CurrStar)
	silverUseReasonText := fmt.Sprintf(commonlog.SilverLogReasonMassacreAdvanced.String(), massacreInfo.CurrLevel, massacreInfo.CurrStar)
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), goldUseReason, goldUseReasonText, costSilver, silverUseReason, silverUseReasonText)
	if !flag {
		panic(fmt.Errorf("massacre: massacreAdvanced Cost should be ok"))
	}
	//同步元宝
	if costGold != 0 || costSilver != 0 || costBindGold != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	//消耗物品
	if itemCount != 0 {
		inventoryReason := commonlog.InventoryLogReasonMassacreAdvanced
		reasonText := fmt.Sprintf(inventoryReason.String(), massacreInfo.CurrLevel, massacreInfo.CurrStar)
		flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("massacre: massacreAdvanced use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//消耗杀气
	if needSqNum > 0 {
		flag := massacreManager.SubShaQiNum(needSqNum)
		if !flag {
			panic(fmt.Errorf("massacre: massacreAdvanced use shaqinum should be ok"))
		}
	}

	//进阶判断
	sucess, addTimes := MassacreAdvanced(pl, massacreInfo.TimesNum, massacreTemplate)
	massacreManager.MassacreAdvanced(addTimes, sucess)

	res := MassacreAdvancedResTypeDefeated //用于通知前段结果
	if sucess {
		//同步属性
		MassacrePropertyChanged(pl)
		if beforeLev != massacreInfo.CurrLevel {
			res = MassacreAdvancedResTypeLevSucceed
		} else {
			res = MassacreAdvancedResTypeStarSucceed
		}

		massacreReason := commonlog.MassacreLogReasonAdvanced
		reasonText := fmt.Sprintf(massacreReason.String(), commontypes.AdvancedTypeShaQi.String())
		data := massacreeventtypes.CreatePlayerMassacreChangedLogEventData(beforeNum, 1, beforeShaQiNum, massacreReason, reasonText)
		gameevent.Emit(massacreeventtypes.EventTypeMassacreChangedLog, pl, data)
		// //是否激活兵魂
		// weaponLev := massacretemplate.GetMassacreTemplateService().GetMassacreeWeaponLev()
		// if beforeLev != massacreInfo.CurrLevel && massacreInfo.CurrLevel == weaponLev {
		// 	data := massacreeventtypes.CreatePlayerMassacreWeaponEventData(massacreTemplate.WeaponId, true)
		// 	gameevent.Emit(massacreeventtypes.EventTypeMassacreWeapon, pl, data)
		// }
	}
	advancedFinish(pl, int32(massacreInfo.AdvanceId), massacreInfo.ShaQiNum, commontypes.AdvancedTypeShaQi, res)

	return
}

func advancedFinish(pl player.Player, advancedId int32, sqNum int64, typ commontypes.AdvancedType, resultType int32) {
	scMassacreAdvanced := pbutil.BuildSCMassacreAdavanced(advancedId, sqNum, commontypes.AdvancedTypeShaQi, resultType)
	pl.SendMsg(scMassacreAdvanced)
	return
}

//变更戮仙刃属性
func MassacrePropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeMassacre.Mask())
	return
}

//戮仙刃进阶判断
func MassacreAdvanced(pl player.Player, curTimesNum int32, massacreTemplate *gametemplate.MassacreTemplate) (sucess bool, addTimes int32) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeMassacre, massacreTemplate.TimesMin, massacreTemplate.TimesMax)
	updatePercent := massacreTemplate.UpdatePercent

	addTimes = int32(1)
	curTimesNum += addTimes

	_, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, 0, timesMin, timesMax, 0, updatePercent, 0)
	return
}

// func MassacreProcessDrop(pl player.Player, attackId int64, attackName string) (itemId int32, itemNum, dropNum int64) {
// 	manager := pl.GetPlayerDataManager(types.PlayerMassacreDataManagerType).(*playermassacre.PlayerMassacreDataManager)
// 	//是否掉落冷却中
// 	now := global.GetGame().GetTimeService().Now()
// 	dropCd := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypePlayerDropSqCd))
// 	if now < manager.GetMassacreInfo().LastTime+dropCd {
// 		log.WithFields(
// 			log.Fields{
// 				"playerId": pl.GetId(),
// 				"nextTime": now + dropCd,
// 			}).Warn("massacre:处理获取戮仙刃掉落,掉落冷却中")
// 		return
// 	}

// 	bagDropNum := int64(0) //传言用
// 	costStar := int32(0)   //传言用
// 	//掉落储存的杀气
// 	itemId = int32(consttypes.ShaQiItem)
// 	itemNum = manager.GetMassacreInfo().ShaQiNum
// 	dropNum = int64(0)
// 	minDropPercent := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeBagDropSqMinPer))
// 	maxDropPercent := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeBagDropSqMaxPer))
// 	dropSqRate := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeDropSqRate))
// 	if itemNum > 0 && mathutils.RandomHit(common.MAX_RATE, dropSqRate) {
// 		bagDropPercent := float64(mathutils.RandomRange(minDropPercent, maxDropPercent)) / float64(common.MAX_RATE)
// 		backBuyXiShu := float64(1 - float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeBagDropSqXiShu))/float64(common.MAX_RATE))
// 		dropNum = int64(math.Ceil(float64(itemNum) * bagDropPercent * backBuyXiShu))
// 		bagDropNum = int64(math.Ceil(float64(itemNum) * bagDropPercent))
// 		itemNum -= bagDropNum
// 	}

// 	old_advanceId := manager.GetMassacreInfo().AdvanceId
// 	old_massacreTemplate := massacretemplate.GetMassacreTemplateService().GetMassacre(old_advanceId)

// 	//是否掉戮仙刃等级,0阶0星的不掉落杀气
// 	if old_massacreTemplate != nil && mathutils.RandomHit(common.MAX_RATE, int(old_massacreTemplate.GasPercent)) {
// 		//掉落几颗星数
// 		new_advanceId := old_advanceId
// 		befLev := old_massacreTemplate.Type
// 		newLev := befLev
// 		subSrar := int32(mathutils.RandomRange(int(old_massacreTemplate.GasMin), int(old_massacreTemplate.GasMax)))
// 		if subSrar > 0 {
// 			tempTemplate := old_massacreTemplate
// 			for subSrar > 0 {
// 				new_advanceId--
// 				subSrar--
// 				costStar++
// 				dropNum += int64(tempTemplate.StarCount)
// 				tempTemplate = massacretemplate.GetMassacreTemplateService().GetMassacre(new_advanceId)
// 				if tempTemplate == nil {
// 					if new_advanceId > 0 {
// 						panic("massacre: 戮仙刃数据表配置错误")
// 					}
// 					new_advanceId = 0
// 					newLev = 0
// 					break
// 				}
// 				newLev = tempTemplate.Type
// 			}

// 			manager.SetMassacreObjInfo(new_advanceId)
// 			//告诉前端兵魂变化了
// 			scMassacreGet := pbutil.BuildSCMassacreGet(manager.GetMassacreInfo())
// 			pl.SendMsg(scMassacreGet)
// 			//同步属性
// 			MassacrePropertyChanged(pl)
// 			//是否下发兵魂消失事件
// 			new_massacreTemplate := massacretemplate.GetMassacreTemplateService().GetMassacre(new_advanceId)
// 			if befLev != newLev && old_massacreTemplate.WeaponId > 0 && (new_massacreTemplate == nil || new_massacreTemplate.WeaponId == 0) {

// 				data := massacreeventtypes.CreatePlayerMassacreWeaponEventData(old_massacreTemplate.WeaponId, false)
// 				gameevent.Emit(massacreeventtypes.EventTypeMassacreWeapon, pl, data)
// 				//告诉前端兵魂变化了
// 				scMassacreWeaponLose := pbutil.BuildSCMassacreWeaponLose(int32(old_advanceId), int32(new_advanceId), attackName)
// 				pl.SendMsg(scMassacreWeaponLose)
// 			}
// 		}
// 	}

// 	if bagDropNum > 0 || costStar > 0 {
// 		//告诉前端掉落杀气了
// 		scMassacreShaQiDrop := pbutil.BuildSCMassacreShaQiDrop(costStar, int32(bagDropNum), attackName)
// 		pl.SendMsg(scMassacreShaQiDrop)
// 	}
// 	return
// }

func MassacreProcessDrop(pl player.Player, attackId int64, attackName string) (itemId int32, dropNum int64) {
	manager := pl.GetPlayerDataManager(types.PlayerMassacreDataManagerType).(*playermassacre.PlayerMassacreDataManager)

	if !manager.IfCanDrop() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("massacre:处理获取戮仙刃掉落,掉落冷却中")
		return
	}
	//掉落储存的杀气
	itemId = int32(consttypes.ShaQiItem)
	//获取掉落的
	flag, bagDropNum, costStar := manager.MassacreDrop(attackName)
	if !flag {
		return
	}

	if bagDropNum > 0 {
		backBuyXiShu := float64(1 - float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeBagDropSqXiShu))/float64(common.MAX_RATE))
		//掉落数量
		dropNum = int64(math.Ceil(float64(bagDropNum) * backBuyXiShu))
	}

	if costStar > 0 {
		//同步属性
		MassacrePropertyChanged(pl)
	}

	if bagDropNum > 0 || costStar > 0 {
		//告诉前端掉落杀气了
		scMassacreShaQiDrop := pbutil.BuildSCMassacreShaQiDrop(costStar, int32(bagDropNum), attackName)
		pl.SendMsg(scMassacreShaQiDrop)
	}

	return
}
