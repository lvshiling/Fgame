package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	commomlogic "fgame/fgame/game/common/logic"
	droptemplate "fgame/fgame/game/drop/template"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	playershenqi "fgame/fgame/game/shenqi/player"
	shenqitypes "fgame/fgame/game/shenqi/types"
	gametemplate "fgame/fgame/game/template"
	viplogic "fgame/fgame/game/vip/logic"
	viptypes "fgame/fgame/game/vip/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//计算分解器灵
func CountResolveQiLing(pl player.Player, itemIndexList []int32) (lingQiVal int64, flag bool) {
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	for _, itemIndex := range itemIndexList {
		it := inventoryManager.FindItemByIndex(inventorytypes.BagTypeQiLing, itemIndex)
		if it == nil {
			log.WithFields(
				log.Fields{
					"index": itemIndex,
				}).Warn("qilingresolve:格子不存在")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
			return
		}
		if it.ItemId == 0 || it.Num == 0 {
			log.WithFields(
				log.Fields{
					"index": itemIndex,
				}).Warn("qilingresolve:格子没有物品")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
			return
		}

		itemTemp := item.GetItemService().GetItem(int(it.ItemId))
		if itemTemp == nil {
			log.WithFields(
				log.Fields{
					"index": itemIndex,
				}).Warn("qilingresolve:物品模板错误")
			playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
			return
		}
		qiLingTemplate := itemTemp.GetShenQiQiLingTemplate()
		if qiLingTemplate == nil {
			log.WithFields(
				log.Fields{
					"index": itemIndex,
				}).Warn("qilingresolve:不是器灵物品")
			playerlogic.SendSystemMessage(pl, lang.ShenQiBagItemNotQiLing)
			return
		}
		lingQiVal += int64(qiLingTemplate.FenJieGet * it.Num)
	}
	flag = true
	return
}

//卸下器灵
func TakeOffQiLing(pl player.Player, typ shenqitypes.ShenQiType, subType shenqitypes.QiLingType, pos shenqitypes.QiLingSubType) (flag bool) {
	shenQiManager := pl.GetPlayerDataManager(playertypes.PlayerShenQiDataManagerType).(*playershenqi.PlayerShenQiDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	item := shenQiManager.GetShenQiQiLingMapByArg(typ, subType, pos)
	level := int32(0)
	num := int32(1)
	bind := item.BindType
	//没有东西
	if item == nil || item.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"subType":  int32(subType),
				"pos":      pos.String(),
			}).Warn("takeoffqiling:脱下器灵,器灵不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	//背包空间
	if !inventoryManager.HasEnoughSlotItemLevel(item.ItemId, num, level, bind) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"subType":  int32(subType),
				"pos":      pos.String(),
			}).Warn("takeoffqiling:脱下器灵,背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	itemId := shenQiManager.QiLingTakeOff(typ, subType, pos)
	if itemId == 0 {
		panic(fmt.Errorf("takeoffqiling:take off should more than 0"))
	}

	//添加物品
	itemData := droptemplate.CreateItemData(itemId, num, level, bind)
	reasonText := commonlog.InventoryLogReasonTakeOff.String()
	flag = inventoryManager.AddItemLevel(itemData, commonlog.InventoryLogReasonTakeOff, reasonText)
	if !flag {
		panic(fmt.Errorf("takeoffqiling:add item should be success"))
	}

	return
}

//变更神器属性
func ShenQiPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeShenQi.Mask())
	return
}

//神器碎片升级判断
func DebrisUpJudge(pl player.Player, curTimesNum int32, curBless int32, template *gametemplate.ShenQiLevelTemplate) (sucess bool, pro, randBless, addTimes int32, isDouble bool) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeShenQi, template.TimesMin, template.TimesMax)
	updateRate := template.UpdateWfb
	blessMax := template.ZhufuMax
	addMin := template.AddMin
	addMax := template.AddMax + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	addTimes = int32(1)

	curTimesNum += addTimes
	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//神器淬炼升级判断
func SmeltUpJudge(pl player.Player, curTimesNum int32, curBless int32, template *gametemplate.ShenQiCuiLianLevelTemplate) (sucess bool, pro, randBless, addTimes int32, isDouble bool) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeShenQi, template.TimesMin, template.TimesMax)
	updateRate := template.UpdateWfb
	blessMax := template.ZhufuMax
	addMin := template.AddMin
	addMax := template.AddMax + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	addTimes = int32(1)

	curTimesNum += addTimes
	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//神器器灵注灵判断
func ZhuLingJudge(pl player.Player, curTimesNum int32, curBless int32, template *gametemplate.ShenQiZhuLingTemplate) (sucess bool, pro, randBless, addTimes int32, isDouble bool) {
	timesMin, timesMax := viplogic.CountWithCostLevel(pl, viptypes.CostLevelRuleTypeShenQi, template.TimesMin, template.TimesMax)
	updateRate := template.UpdateWfb
	blessMax := template.ZhufuMax
	addMin := template.AddMin
	addMax := template.AddMax + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	addTimes = int32(1)

	curTimesNum += addTimes
	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}
