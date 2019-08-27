package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	commomlogic "fgame/fgame/game/common/logic"
	commontypes "fgame/fgame/game/common/types"
	gameevent "fgame/fgame/game/event"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	mounteventtypes "fgame/fgame/game/mount/event/types"
	"fgame/fgame/game/mount/mount"
	"fgame/fgame/game/mount/pbutil"
	playermount "fgame/fgame/game/mount/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	shoplogic "fgame/fgame/game/shop/logic"
	"fgame/fgame/game/shop/shop"
	gametemplate "fgame/fgame/game/template"
	viplogic "fgame/fgame/game/vip/logic"
	viptypes "fgame/fgame/game/vip/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//坐骑进阶的逻辑
func HandleMountAdvanced(pl player.Player, autoFlag bool) (err error) {
	isFuncOpen := pl.IsFuncOpen(funcopentypes.FuncOpenTypeMount)
	if !isFuncOpen {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("mount:坐骑升阶功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	mountManager := pl.GetPlayerDataManager(types.PlayerMountDataManagerType).(*playermount.PlayerMountDataManager)
	mountInfo := mountManager.GetMountInfo()
	beforeNum := int32(mountInfo.AdvanceId)
	nextAdvancedId := mountInfo.AdvanceId + 1
	mountTemplate := mount.GetMountService().GetMountNumber(int32(nextAdvancedId))
	if mountTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"autoFlag": autoFlag,
		}).Warn("mount:坐骑已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.MountAdanvacedReachedLimit)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//进阶需要消耗的元宝
	costGold := mountTemplate.UseMoney
	//进阶需要消耗的银两
	costSilver := int64(mountTemplate.UseYinliang)
	//进阶需要的消耗的绑元
	costBindGold := int32(0)

	//需要消耗物品
	itemCount := int32(0)
	totalNum := int32(0)
	useItem := mountTemplate.UseItem
	isEnoughBuyTimes := true
	shopIdMap := make(map[int32]int32)
	useItemTemplate := mountTemplate.GetUseItemTemplate()
	if useItemTemplate != nil {
		itemCount = mountTemplate.ItemCount
		totalNum = inventoryManager.NumOfItems(int32(useItem))
	}
	if totalNum < itemCount {
		if autoFlag == false {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
			}).Warn("mount:坐骑进阶丹不足,无法进阶")
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
		// 		"autoFlag": autoFlag,
		// 	}).Warn("mount:商铺没有该道具,无法自动购买")
		// 	playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
		// 	return
		// }

		if needBuyNum > 0 {
			if !shop.GetShopService().ShopIsSellItem(useItem) {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("mount:商铺没有该道具,无法自动购买")
				playerlogic.SendSystemMessage(pl, lang.ShopBuyNotItem)
				return
			}

			isEnoughBuyTimes, shopIdMap = shoplogic.MaxBuyTimesForPlayer(pl, useItem, needBuyNum)
			if !isEnoughBuyTimes {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"autoFlag": autoFlag,
				}).Warn("mount:购买物品失败,自动进阶已停止")
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
			}).Warn("mount:银两不足,无法进阶")
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
			}).Warn("mount:元宝不足,无法进阶")
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
			}).Warn("mount:元宝不足,无法进阶")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//更新自动购买每日限购次数
	if len(shopIdMap) != 0 {
		shoplogic.ShopDayCountChanged(pl, shopIdMap)
	}

	//消耗钱
	reasonGoldText := commonlog.GoldLogReasonMountAdvanced.String()
	reasonSliverText := commonlog.SilverLogReasonMountAdvanced.String()
	flag := propertyManager.Cost(int64(costBindGold), int64(costGold), commonlog.GoldLogReasonMountAdvanced, reasonGoldText, costSilver, commonlog.SilverLogReasonMountAdvanced, reasonSliverText)
	if !flag {
		panic(fmt.Errorf("mount: mountAdvanced Cost should be ok"))
	}

	//消耗物品
	if itemCount != 0 {
		inventoryReason := commonlog.InventoryLogReasonMountAdvanced
		reasonText := fmt.Sprintf(inventoryReason.String(), mountInfo.AdvanceId)
		flag := inventoryManager.UseItem(int32(useItem), itemCount, inventoryReason, reasonText)
		if !flag {
			panic(fmt.Errorf("mount: mountAdvanced use item should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	if useItemTemplate != nil {
		//消耗事件
		gameevent.Emit(mounteventtypes.EventTypeMountAdvancedCost, pl, int32(mountInfo.AdvanceId))
	}

	//同步元宝
	if costGold != 0 || costSilver != 0 || costBindGold != 0 {
		propertylogic.SnapChangedProperty(pl)
	}

	//进阶判断
	sucess, pro, randBless, addTimes, isDouble := MountAdvanced(pl, mountInfo.TimesNum, mountInfo.Bless, mountTemplate)
	mountManager.MountAdvanced(pro, addTimes, sucess)
	if sucess {
		mountReason := commonlog.MountLogReasonAdvanced
		reasonText := fmt.Sprintf(mountReason.String(), commontypes.AdvancedTypeJinJieDan.String())
		data := mounteventtypes.CreatePlayerMountAdvancedLogEventData(beforeNum, 1, mountReason, reasonText)
		gameevent.Emit(mounteventtypes.EventTypeMountAdvancedLog, pl, data)

		//同步属性
		MountPropertyChanged(pl)
		advancedFinish(pl, int32(mountInfo.AdvanceId), mountInfo.MountId)
	} else {
		//进阶不成功
		advancedBless(pl, mountInfo.AdvanceId, mountInfo.MountId, randBless, mountInfo.Bless, mountInfo.BlessTime, isDouble)
	}
	return
}

func advancedFinish(pl player.Player, advancedId int32, mountId int32) (err error) {
	scMountAdvanced := pbutil.BuildSCMountAdavancedFinshed(advancedId, mountId, commontypes.AdvancedTypeJinJieDan)
	pl.SendMsg(scMountAdvanced)
	return
}

func advancedBless(pl player.Player, advancedId int, mountId int32, bless int32, totalBless int32, blessTime int64, isDouble bool) (err error) {
	scMountAdvanced := pbutil.BuildSCMountAdavanced(int32(advancedId), mountId, bless, totalBless, blessTime, commontypes.AdvancedTypeJinJieDan, isDouble)
	pl.SendMsg(scMountAdvanced)
	return
}

//变更坐骑属性
func MountPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeMount.Mask())
	return
}

//坐骑进阶判断
func MountAdvanced(pl player.Player, curTimesNum int32, curBless int32, mountTemplate *gametemplate.MountTemplate) (sucess bool, pro, randBless, addTimes int32, isDouble bool) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeMount, mountTemplate.TimesMin, mountTemplate.TimesMax)
	updateRate := mountTemplate.UpdateWfb
	blessMax := mountTemplate.ZhufuMax
	addMin := mountTemplate.AddMin
	addMax := mountTemplate.AddMax + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	addTimes = int32(1)
	isDouble, extralAddTimes := welfarelogic.IsCanAdvancedBlessCrit(welfaretypes.AdvancedTypeMount)
	if isDouble {
		addTimes += extralAddTimes
		randBless += mountTemplate.AddMax
	}
	curTimesNum += addTimes
	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//坐骑祝福丹判断
func MountEatZhuFuDan(pl player.Player, curTimesNum, curBless, randBless int32, mountTemplate *gametemplate.MountTemplate) (sucess bool, pro int32) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeMount, mountTemplate.TimesMin, mountTemplate.TimesMax)
	updateRate := mountTemplate.UpdateWfb
	blessMax := mountTemplate.ZhufuMax

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//坐骑草料喂养判断
func MountCaoLiaoFeed(curTimesNum int32, curBless int32, caoLiaoTemplate *gametemplate.MountCaoLiaoTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := caoLiaoTemplate.TimesMin
	timesMax := caoLiaoTemplate.TimesMax
	updateRate := caoLiaoTemplate.UpdateWfb
	blessMax := caoLiaoTemplate.ZhufuMax
	addMin := caoLiaoTemplate.AddMin
	addMax := caoLiaoTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//坐骑幻化丹培养判断
func MountHunaHuaFeed(curTimesNum int32, curBless int32, huanHuaTemplate *gametemplate.MountHuanHuaTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := huanHuaTemplate.TimesMin
	timesMax := huanHuaTemplate.TimesMax
	updateRate := huanHuaTemplate.UpdateWfb
	blessMax := huanHuaTemplate.ZhufuMax
	addMin := huanHuaTemplate.AddMin
	addMax := huanHuaTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}

//坐骑皮肤升星判断
func MountSkinUpstar(curTimesNum int32, curBless int32, upstarTemplate *gametemplate.MountUpstarTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := upstarTemplate.TimesMin
	timesMax := upstarTemplate.TimesMax
	updateRate := upstarTemplate.UpstarRate
	blessMax := upstarTemplate.ZhufuMax
	addMin := upstarTemplate.AddMin
	addMax := upstarTemplate.AddMax + 1
	pro, randBless, sucess = commomlogic.GetStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, addMin, addMax, updateRate, blessMax)
	return
}
