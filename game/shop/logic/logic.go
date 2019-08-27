package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/common/common"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/shop/pbutil"
	playershop "fgame/fgame/game/shop/player"
	"fgame/fgame/game/shop/shop"
	shoptypes "fgame/fgame/game/shop/types"
	shopdiscountlogic "fgame/fgame/game/shopdiscount/logic"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

//商店购买道具的逻辑
func HandleShopBuy(pl player.Player, shopId int32, num int32) (err error) {
	shopTemplate := shop.GetShopService().GetShopTemplate(int(shopId))
	if shopTemplate == nil || num <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"num":      num,
			}).Warn("shop:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if num > shopTemplate.MaxCount {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"num":      num,
			}).Warn("shop:购买数量大于最大购买数量")
		playerlogic.SendSystemMessage(pl, lang.ShopBuyNumInvalid)
		return
	}

	//购买总数
	totalNum := int32(shopTemplate.BuyCount * num)

	shopManager := pl.GetPlayerDataManager(types.PlayerShopDataManagerType).(*playershop.PlayerShopDataManager)
	isLimitBuy, leftNum := shopManager.LeftDayCount(shopId)
	if isLimitBuy && leftNum < totalNum {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"num":      num,
			}).Warn("shop:购买次数，已达每日限购数量")
		playerlogic.SendSystemMessage(pl, lang.ShoBuyReacheLimit)
		return
	}

	//货币判断
	flag := true
	consumeType := shoptypes.ShopConsumeType(shopTemplate.ConsumeType)
	discountRatio := shopdiscountlogic.GetShopDiscount(pl, consumeType)
	consume := int32(math.Ceil(float64(shopTemplate.ConsumeData1*num) * discountRatio / float64(common.MAX_RATE)))
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	switch consumeType {
	case shoptypes.ShopConsumeTypeBindGold:
		flag = propertyManager.HasEnoughGold(int64(consume), true)
		break
	case shoptypes.ShopConsumeTypeGold:
		flag = propertyManager.HasEnoughGold(int64(consume), false)
		break
	case shoptypes.ShopConsumeTypeSliver:
		flag = propertyManager.HasEnoughSilver(int64(consume))
		break
	default:
		break
	}
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"num":      num,
				"consume":  consume,
			}).Warn("shop:元宝不足，无法完成购买")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	//判断背包空间
	itemId := shopTemplate.ItemId
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	flag = inventoryManager.HasEnoughSlot(itemId, totalNum)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"shopId":   shopId,
				"num":      num,
			}).Warn("shop:背包空间不足，请清理后再购买")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//消耗货币
	switch consumeType {
	case shoptypes.ShopConsumeTypeBindGold:
		goldReason := commonlog.GoldLogReasonShopBuyItem
		reasonText := fmt.Sprintf(goldReason.String(), shopId, shopTemplate.Name, num)
		flag = propertyManager.CostGold(int64(consume), true, goldReason, reasonText)
		break
	case shoptypes.ShopConsumeTypeGold:
		goldReason := commonlog.GoldLogReasonShopBuyItem
		reasonText := fmt.Sprintf(goldReason.String(), shopId, shopTemplate.Name, num)
		flag = propertyManager.CostGold(int64(consume), false, goldReason, reasonText)
		break
	case shoptypes.ShopConsumeTypeSliver:
		silverReason := commonlog.SilverLogReasonShopBuyItem
		reasonText := fmt.Sprintf(silverReason.String(), shopId, shopTemplate.Name, num)
		flag = propertyManager.CostSilver(int64(consume), silverReason, reasonText)
		break
	default:
		break
	}
	if !flag {
		panic("shop: costGold/Silver should be ok")
	}
	//同步元宝
	propertylogic.SnapChangedProperty(pl)

	//添加物品
	reasonText := commonlog.InventoryLogReasonShopBuy.String()
	flag = inventoryManager.AddItem(itemId, totalNum, commonlog.InventoryLogReasonShopBuy, reasonText)
	if !flag {
		panic(fmt.Errorf("shop: shopBuy add item should be ok"))
	}
	//同步物品
	inventorylogic.SnapInventoryChanged(pl)

	//更新当日购买次数
	dayCount := int32(0)
	shopManager.UpdateObject(shopId, totalNum, false)
	if shopTemplate.LimitCount != 0 {
		dayCount = shopManager.GetShopBuyByShopId(shopId).DayCount
	}
	scShopBuy := pbutil.BuildSCShopBuy(shopId, num, dayCount)
	pl.SendMsg(scShopBuy)
	return
}

//银两商铺
func MaxBuyTimesSivler(pl player.Player, shopId int32, needNum int32, preCostSilver int64) (num int32, needSilver int64) {
	if needNum <= 0 {
		panic(fmt.Errorf("shoplogic: needNum 应该大于0"))
	}
	if preCostSilver < 0 {
		panic(fmt.Errorf("shoplogic: preCostSilver 应该大于等于0"))
	}
	shopTemplate := shop.GetShopService().GetShopTemplate(int(shopId))
	if shopTemplate == nil {
		return
	}
	if shopTemplate.GetShopConsumeType() != shoptypes.ShopConsumeTypeSliver {
		panic(fmt.Errorf("shoplogic: 应该在银两商铺"))
	}

	discountRatio := shopdiscountlogic.GetShopDiscount(pl, shopTemplate.GetShopConsumeType())
	costSilver := int64(math.Ceil(float64(shopTemplate.ConsumeData1) * discountRatio / float64(common.MAX_RATE)))
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	curSilver := propertyManager.GetSilver() - preCostSilver
	if curSilver < costSilver {
		return
	}

	shopManager := pl.GetPlayerDataManager(types.PlayerShopDataManagerType).(*playershop.PlayerShopDataManager)
	isLimit, leftNum := shopManager.LeftDayCount(shopId)
	if leftNum < 0 {
		panic(fmt.Errorf("shoplogic:silver leftNum 应该大于等于0"))
	}
	if isLimit && leftNum == 0 {
		return
	}

	buyNum := int32(curSilver / costSilver)
	if isLimit && buyNum > leftNum {
		buyNum = leftNum
	}
	if buyNum > needNum {
		buyNum = needNum
	}
	num = buyNum
	needSilver = int64(math.Ceil(float64(shopTemplate.ConsumeData1*num) * discountRatio / float64(common.MAX_RATE)))
	return
}

//绑元商铺
func MaxBuyTimesBindGold(pl player.Player, shopId int32, needNum int32, preCostBindGold int64, preCostGold int64, instead bool, complementGold int64) (num int32, needBindGold int64, needGold int64) {
	if needNum <= 0 {
		panic(fmt.Errorf("shoplogic: needNum 应该大于0"))
	}
	if preCostBindGold < 0 {
		panic(fmt.Errorf("shoplogic: preCostBindGold 应该大于等于0"))
	}
	if preCostGold < 0 {
		panic(fmt.Errorf("shoplogic: preCostGold 应该大于等于0"))
	}
	shopTemplate := shop.GetShopService().GetShopTemplate(int(shopId))
	if shopTemplate == nil {
		return
	}
	if shopTemplate.GetShopConsumeType() != shoptypes.ShopConsumeTypeBindGold {
		panic(fmt.Errorf("shoplogic: 应该在绑元商铺"))
	}

	discountRatio := shopdiscountlogic.GetShopDiscount(pl, shopTemplate.GetShopConsumeType())
	costBindGold := int64(math.Ceil(float64(shopTemplate.ConsumeData1) * discountRatio / float64(common.MAX_RATE)))
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	curBindGold := propertyManager.GetBindGlod() - preCostBindGold
	//zrc: 不能使用元宝代替绑元自动购买
	curGold := int64(0)
	if instead {
		curGold = propertyManager.GetGold() + complementGold - preCostGold
	}

	if curBindGold+curGold < costBindGold {
		return
	}

	shopManager := pl.GetPlayerDataManager(types.PlayerShopDataManagerType).(*playershop.PlayerShopDataManager)
	isLimit, leftNum := shopManager.LeftDayCount(shopId)
	if leftNum < 0 {
		panic(fmt.Errorf("shoplogic:bindGold leftNum 应该大于等于0"))
	}
	if isLimit && leftNum == 0 {
		return
	}

	buyNum := int32((curBindGold + curGold) / costBindGold)
	if isLimit && buyNum > leftNum {
		buyNum = leftNum
	}
	if buyNum > needNum {
		buyNum = needNum
	}
	num = buyNum
	totalBindGold := int64(math.Ceil(float64(shopTemplate.ConsumeData1*num) * discountRatio / float64(common.MAX_RATE)))
	needBindGold = totalBindGold
	if totalBindGold > curBindGold {
		needBindGold = curBindGold
		needGold = totalBindGold - needBindGold
	}
	return
}

//元宝商铺
func MaxBuyTimesGold(pl player.Player, shopId int32, needNum int32, preCostGold int64, complementGold int64) (num int32, needGold int64) {
	if needNum <= 0 {
		panic(fmt.Errorf("shoplogic: needNum 应该大于0"))
	}
	if preCostGold < 0 {
		panic(fmt.Errorf("shoplogic: preCostGold 应该大于等于0"))
	}
	shopTemplate := shop.GetShopService().GetShopTemplate(int(shopId))
	if shopTemplate == nil {
		return
	}
	if shopTemplate.GetShopConsumeType() != shoptypes.ShopConsumeTypeGold {
		panic(fmt.Errorf("shoplogic: 应该在元宝商铺"))
	}
	discountRatio := shopdiscountlogic.GetShopDiscount(pl, shopTemplate.GetShopConsumeType())
	costGold := int64(math.Ceil(float64(shopTemplate.ConsumeData1) * discountRatio / float64(common.MAX_RATE)))
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	curGold := propertyManager.GetGold() + complementGold - preCostGold
	if curGold < costGold {
		return
	}

	shopManager := pl.GetPlayerDataManager(types.PlayerShopDataManagerType).(*playershop.PlayerShopDataManager)
	isLimit, leftNum := shopManager.LeftDayCount(shopId)
	if leftNum < 0 {
		panic(fmt.Errorf("shoplogic:gold leftNum 应该大于等于0"))
	}
	if isLimit && leftNum == 0 {
		return
	}

	buyNum := int32(curGold / costGold)
	if isLimit && buyNum > leftNum {
		buyNum = leftNum
	}
	if buyNum > needNum {
		buyNum = needNum
	}
	num = buyNum
	needGold = int64(math.Ceil(float64(shopTemplate.ConsumeData1*num) * discountRatio / float64(common.MAX_RATE)))
	return
}

//银两商店配多种价格
func MaxBuyTimesForSilver(pl player.Player, needNum int32, shopItemMap map[shoptypes.ShopConsumeType][]*gametemplate.ShopTemplate) (maxNum int32, silver int64, shopIdMap map[int32]int32) {
	if len(shopItemMap) == 0 {
		return
	}
	shopSilverItemList, exist := shopItemMap[shoptypes.ShopConsumeTypeSliver]
	if !exist {
		return
	}
	silver = 0
	leftNum := needNum
	shopIdMap = make(map[int32]int32)
	for _, shopTemplate := range shopSilverItemList {
		num, needSilver := MaxBuyTimesSivler(pl, int32(shopTemplate.Id), leftNum, silver)
		if num > 0 {
			leftNum -= num
			maxNum += num
			silver += needSilver
			shopIdMap[int32(shopTemplate.Id)] = num
		}
		if leftNum == 0 {
			return
		}
	}
	return
}

//绑元商店配多种价格
func MaxBuyTimesForBindGold(pl player.Player, needNum int32, shopItemMap map[shoptypes.ShopConsumeType][]*gametemplate.ShopTemplate, preCostBindGold int64, preCostGold int64, instead bool, complementGold int64) (maxNum int32, bindGold int64, gold int64, shopIdMap map[int32]int32) {
	if len(shopItemMap) == 0 {
		return
	}
	shopBindGoldItemList, exist := shopItemMap[shoptypes.ShopConsumeTypeBindGold]
	if !exist {
		return
	}
	bindGold = preCostBindGold
	gold = preCostGold
	leftNum := needNum
	shopIdMap = make(map[int32]int32)
	for _, shopTemplate := range shopBindGoldItemList {
		num, needBindGold, needGold := MaxBuyTimesBindGold(pl, int32(shopTemplate.Id), leftNum, bindGold, gold, instead, complementGold)
		if num > 0 {
			leftNum -= num
			maxNum += num
			bindGold += needBindGold
			gold += needGold
			shopIdMap[int32(shopTemplate.Id)] = num
		}
		if leftNum == 0 {
			return
		}
	}
	return
}

//元宝商店配多种价格
func MaxBuyTimesForGold(pl player.Player, needNum int32, shopItemMap map[shoptypes.ShopConsumeType][]*gametemplate.ShopTemplate, preCostGold int64, complementGold int64) (maxNum int32, gold int64, shopIdMap map[int32]int32) {
	if len(shopItemMap) == 0 {
		return
	}
	shopBindGoldItemList, exist := shopItemMap[shoptypes.ShopConsumeTypeGold]
	if !exist {
		return
	}
	gold = preCostGold
	leftNum := needNum
	shopIdMap = make(map[int32]int32)
	for _, shopTemplate := range shopBindGoldItemList {
		num, needGold := MaxBuyTimesGold(pl, int32(shopTemplate.Id), leftNum, gold, complementGold)
		if num > 0 {
			leftNum -= num
			maxNum += num
			gold += needGold
			shopIdMap[int32(shopTemplate.Id)] = num
		}
		if leftNum == 0 {
			return
		}
	}
	return
}

func MaxBuyTimesForPlayer(pl player.Player, itemId int32, needNum int32) (isEnoughBuyTimes bool, shopIdMap map[int32]int32) {
	return MaxBuyTimesForPlayerComplementGold(pl, itemId, needNum, 0)
}

//自动购买单个物品
func MaxBuyTimesForPlayerComplementGold(pl player.Player, itemId int32, needNum int32, complementGold int64) (isEnoughBuyTimes bool, shopIdMap map[int32]int32) {
	if needNum < 0 {
		panic(fmt.Errorf("shoplogic: needNum 应该大于等于0"))
	}
	if !shop.GetShopService().ShopIsSellItem(itemId) {
		return
	}
	needSilver := int64(0)
	preCostBindGold := int64(0)
	preCostGold := int64(0)
	leftNum := needNum
	shopIdMap = make(map[int32]int32)
	shopItemMap := shop.GetShopService().GetShopItemMap(itemId)
	//判断银两商铺
	silverNum, silver, tempSilverShopIdMap := MaxBuyTimesForSilver(pl, leftNum, shopItemMap)
	if silverNum > 0 {
		leftNum -= silverNum
		needSilver += silver
		for shopId, num := range tempSilverShopIdMap {
			shopIdMap[shopId] += num
		}
	}
	if leftNum == 0 {
		isEnoughBuyTimes = true
		return
	}

	//判断绑元商铺
	bindGoldNum, bindGold, preGold, tempBindGoldShopIdMap := MaxBuyTimesForBindGold(pl, leftNum, shopItemMap, preCostBindGold, preCostGold, false, complementGold)
	if bindGoldNum > 0 {
		leftNum -= bindGoldNum
		preCostBindGold = bindGold
		preCostGold = preGold
		for shopId, num := range tempBindGoldShopIdMap {
			shopIdMap[shopId] += num
		}
	}
	if leftNum == 0 {
		isEnoughBuyTimes = true
		return
	}

	//判断元宝商铺
	goldNum, gold, tempGoldShopIdMap := MaxBuyTimesForGold(pl, leftNum, shopItemMap, preCostGold, complementGold)
	if goldNum > 0 {
		leftNum -= goldNum
		preCostGold = gold
		for shopId, num := range tempGoldShopIdMap {
			shopIdMap[shopId] += num
		}
	}
	if leftNum == 0 {
		isEnoughBuyTimes = true
		return
	}

	//判断绑元商铺
	bindGoldNum, bindGold, preGold, tempBindGoldShopIdMap = MaxBuyTimesForBindGold(pl, leftNum, shopItemMap, preCostBindGold, preCostGold, true, complementGold)
	if bindGoldNum > 0 {
		leftNum -= bindGoldNum
		preCostBindGold = bindGold
		preCostGold = preGold
		for shopId, num := range tempBindGoldShopIdMap {
			shopIdMap[shopId] += num
		}
	}
	if leftNum == 0 {
		isEnoughBuyTimes = true
		return
	}

	//通知前端停止自动购买
	if !isEnoughBuyTimes {
		scShopStopAutoBuy := pbutil.BuildSCShopStopAutoBuy()
		pl.SendMsg(scShopStopAutoBuy)
	}
	return
}

//自动购买多个物品的
func MaxBuyTimesForPlayerMap(pl player.Player, itemMap map[int32]int32) (isEnoughBuyTimes bool, shopIdMap map[int32]int32) {
	return MaxBuyTimesForPlayerMapComplementGold(pl, 0, itemMap)
}

//自动购买多个物品的
func MaxBuyTimesForPlayerMapComplementGold(pl player.Player, complementGold int64, itemMap map[int32]int32) (isEnoughBuyTimes bool, shopIdMap map[int32]int32) {
	if len(itemMap) == 0 {
		panic(fmt.Errorf("shoplogic: itemMap 应该大于0"))
	}

	costBindGold := int64(0)
	costGold := int64(0)
	costSilver := int64(0)
	shopIdMap = make(map[int32]int32)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	curBindGold := propertyManager.GetBindGlod()
	curGold := propertyManager.GetGold() + complementGold
	curSilver := propertyManager.GetSilver()
	isEnoughMoney := true
	for itemId, needNum := range itemMap {
		isEnoughBuy, curShopIdMap := MaxBuyTimesForPlayerComplementGold(pl, itemId, needNum, complementGold)
		if !isEnoughBuy {
			return
		}
		needBindGold, needGold, needSilver := ShopCostData(pl, curShopIdMap)
		costBindGold += needBindGold
		costGold += needGold
		costSilver += needSilver
		if curSilver < costSilver {
			isEnoughMoney = false
			break
		}
		if curGold < costGold {
			isEnoughMoney = false
			break
		}
		if curBindGold < costBindGold && ((costBindGold - curBindGold) > (curGold - costGold)) {
			isEnoughMoney = false
			break
		}
		for shopId, num := range curShopIdMap {
			shopIdMap[shopId] += num
		}
	}

	if isEnoughMoney {
		isEnoughBuyTimes = true
	}
	//通知前端停止自动购买
	if !isEnoughBuyTimes {
		scShopStopAutoBuy := pbutil.BuildSCShopStopAutoBuy()
		pl.SendMsg(scShopStopAutoBuy)
	}

	return
}

//计算花费
func ShopCostData(pl player.Player, shopIdMap map[int32]int32) (needBindGold int64, needGold int64, needSilver int64) {
	if len(shopIdMap) == 0 {
		panic(fmt.Errorf("shoplogic: shopIdMap 应该大于0"))
	}

	for shopId, num := range shopIdMap {
		shopTemplate := shop.GetShopService().GetShopTemplate(int(shopId))
		costGold, costBindGold, costSilver := shopTemplate.GetConsumeData(num)
		needBindGold += int64(costBindGold)
		needGold += int64(costGold)
		needSilver += costSilver
	}

	if needBindGold > 0 {
		discountRatio := shopdiscountlogic.GetShopDiscount(pl, shoptypes.ShopConsumeTypeBindGold)
		needBindGold = int64(math.Ceil(float64(needBindGold) * discountRatio / float64(common.MAX_RATE)))
	}

	if needGold > 0 {
		discountRatio := shopdiscountlogic.GetShopDiscount(pl, shoptypes.ShopConsumeTypeGold)
		needGold = int64(math.Ceil(float64(needGold) * discountRatio / float64(common.MAX_RATE)))

	}
	if needSilver > 0 {
		discountRatio := shopdiscountlogic.GetShopDiscount(pl, shoptypes.ShopConsumeTypeSliver)
		needSilver = int64(math.Ceil(float64(needSilver) * discountRatio / float64(common.MAX_RATE)))
	}

	return
}

//更新每日限购数量
func ShopDayCountChanged(pl player.Player, shopIdMap map[int32]int32) {
	if len(shopIdMap) == 0 {
		return
	}
	shopManager := pl.GetPlayerDataManager(types.PlayerShopDataManagerType).(*playershop.PlayerShopDataManager)
	for shopId, num := range shopIdMap {
		shopManager.UpdateObject(shopId, num, true)
	}
	scShopAutoBuyList := pbutil.BuildSCShopAutoBuyList(pl, shopIdMap)
	pl.SendMsg(scShopAutoBuyList)
}

// 修改:zrc
//获取银两消耗
func GetPlayerShopCostByComsumeType(pl player.Player, maxMoney int64, shopConsumeType shoptypes.ShopConsumeType, itemId int32, needNum int32) (maxNum int32, shopIdMap map[int32]int32, totalCost int64) {
	shopItemMap := shop.GetShopService().GetShopItemMap(itemId)
	shopSilverItemList, exist := shopItemMap[shopConsumeType]
	if !exist {
		return
	}

	leftNum := needNum

	shopIdMap = make(map[int32]int32)
	shopManager := pl.GetPlayerDataManager(types.PlayerShopDataManagerType).(*playershop.PlayerShopDataManager)

	for _, shopTemplate := range shopSilverItemList {
		cost := int64(shopTemplate.ConsumeData1)
		playerLeftNum := leftNum
		if shopTemplate.LimitCount != 0 {
			_, playerLeftNum = shopManager.LeftDayCount(int32(shopTemplate.TemplateId()))
			//没有剩余次数
			if playerLeftNum <= 0 {
				continue
			}
		}
		leftMoney := maxMoney - totalCost
		buyNum := int32(leftMoney / cost)
		if buyNum >= leftNum {
			buyNum = leftNum
		}
		//购买所有剩余次数
		if buyNum <= playerLeftNum {
			maxNum += buyNum
			totalCost += cost * int64(buyNum)
			shopIdMap[int32(shopTemplate.Id)] = buyNum
			return
		}
		maxNum += playerLeftNum
		totalCost += cost * int64(playerLeftNum)
		shopIdMap[int32(shopTemplate.Id)] = playerLeftNum
		leftNum -= playerLeftNum

	}
	return
}

//购买所需花费
func GetPlayerShopCost(pl player.Player, itemId int32, needNum int32) (isEnoughBuyTimes bool, shopIdMap map[int32]int32) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	maxSilver := propertyManager.GetSilver()
	maxBindGold := propertyManager.GetBindGlod()
	maxGold := propertyManager.GetGold()
	return GetPlayerShopCostForMaxMoney(pl, itemId, needNum, maxGold, maxBindGold, maxSilver)

}

//购买所需花费
func GetPlayerShopCostForMaxMoney(pl player.Player, itemId int32, needNum int32, maxGold int64, maxBindGold int64, maxSilver int64) (isEnoughBuyTimes bool, shopIdMap map[int32]int32) {
	if needNum < 0 {
		panic(fmt.Errorf("shoplogic: needNum 应该大于等于0"))
	}
	if !shop.GetShopService().ShopIsSellItem(itemId) {
		return
	}

	shopIdMap = make(map[int32]int32)
	leftNum := needNum

	silverMaxNum, silverShopMap, _ := GetPlayerShopCostByComsumeType(pl, maxSilver, shoptypes.ShopConsumeTypeSliver, itemId, needNum)
	if silverMaxNum > 0 {
		leftNum -= silverMaxNum

		for shopId, num := range silverShopMap {
			shopIdMap[shopId] = num
		}
	}
	if leftNum < 0 {
		panic(fmt.Errorf("剩余数量应该不小于0"))
	}

	if leftNum == 0 {
		isEnoughBuyTimes = true
		return
	}

	bindGoldMaxNum, bindGoldShopMap, bindGoldCost := GetPlayerShopCostByComsumeType(pl, maxBindGold+maxGold, shoptypes.ShopConsumeTypeBindGold, itemId, leftNum)
	if bindGoldMaxNum > 0 {
		leftNum -= bindGoldMaxNum

		for shopId, num := range bindGoldShopMap {
			shopIdMap[shopId] = num
		}
	}
	if leftNum < 0 {
		panic(fmt.Errorf("剩余数量应该不小于0"))
	}

	if leftNum == 0 {
		isEnoughBuyTimes = true
		return
	}
	remainGold := maxGold
	if bindGoldCost > maxBindGold {
		remainGold -= (bindGoldCost - maxBindGold)
	}
	goldMaxNum, goldShopMap, _ := GetPlayerShopCostByComsumeType(pl, remainGold, shoptypes.ShopConsumeTypeGold, itemId, leftNum)
	if goldMaxNum > 0 {
		leftNum -= goldMaxNum

		for shopId, num := range goldShopMap {
			shopIdMap[shopId] = num
		}
	}
	if leftNum < 0 {
		panic(fmt.Errorf("剩余数量应该不小于0"))
	}

	if leftNum == 0 {
		isEnoughBuyTimes = true
		return
	}
	return
}

//自动购买多个物品的
func GetPlayerShopCostForItemMap(pl player.Player, itemMap map[int32]int32) (isEnoughBuyTimes bool, shopIdMap map[int32]int32) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	maxSilver := propertyManager.GetSilver()
	maxBindGold := propertyManager.GetBindGlod()
	maxGold := propertyManager.GetGold()
	shopIdMap = make(map[int32]int32)
	for itemId, needNum := range itemMap {
		isEnoughBuyTimes, tempShopMap := GetPlayerShopCostForMaxMoney(pl, itemId, needNum, maxGold, maxBindGold, maxSilver)
		if !isEnoughBuyTimes {
			return false, nil
		}
		needBindGold, needGold, needSiver := ShopCostData(pl, tempShopMap)
		maxSilver -= needSiver
		maxBindGold -= needBindGold
		if maxBindGold < 0 {
			maxGold += maxBindGold
			maxBindGold = 0
		}
		maxGold -= needGold
		for shopId, num := range tempShopMap {
			shopIdMap[shopId] += num
		}
	}
	return true, shopIdMap
}
