package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/game/treasurebox/pbutil"
	boxtemplate "fgame/fgame/game/treasurebox/template"
	treasureboxtypes "fgame/fgame/game/treasurebox/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

var (
	minZhuanShu = int(0)
	maxZhuanShu = int(4)

	minLevel = int(1)
	maxLevel = int(500)
)

//获取宝箱id
func GetBoxTemplate(pl player.Player, startBoxTemplate *gametemplate.BoxTemplate) (boxTemplate *gametemplate.BoxTemplate) {
	if startBoxTemplate.NextId == 0 {
		return startBoxTemplate
	}
	zhuanShu := pl.GetZhuanSheng()
	level := pl.GetLevel()

	nextTemplate := boxtemplate.GetBoxTemplateService().GetBoxTemplate(startBoxTemplate.NextId)
	lastTemplate := startBoxTemplate
	for nextTemplate != nil {
		if zhuanShu < nextTemplate.ZhuanshuMin || level < nextTemplate.LevelMin {
			return lastTemplate
		}
		lastTemplate = nextTemplate
		nextTemplate = boxtemplate.GetBoxTemplateService().GetBoxTemplate(nextTemplate.NextId)
	}
	return lastTemplate
}

func GetRandomBoxTemplate(startBoxTemplate *gametemplate.BoxTemplate) (boxTemplate *gametemplate.BoxTemplate) {
	if startBoxTemplate.NextId == 0 {
		return startBoxTemplate
	}
	zhuanShu := int32(mathutils.RandomRange(minZhuanShu, maxZhuanShu))
	level := int32(mathutils.RandomRange(minLevel, maxLevel))

	nextTemplate := boxtemplate.GetBoxTemplateService().GetBoxTemplate(startBoxTemplate.NextId)
	lastTemplate := startBoxTemplate
	for nextTemplate != nil {
		if zhuanShu < nextTemplate.ZhuanshuMin || level < nextTemplate.LevelMin {
			return lastTemplate
		}
		lastTemplate = nextTemplate
		nextTemplate = boxtemplate.GetBoxTemplateService().GetBoxTemplate(nextTemplate.NextId)
	}
	return lastTemplate
}

//获取乾坤袋配置
func GetQianKunBoxTemplate(pl player.Player, startBoxTemplate *gametemplate.BoxTemplate, useTimes int32) (boxTemplate *gametemplate.BoxTemplate, isUseOut bool) {
	if startBoxTemplate.NextId == 0 {
		return startBoxTemplate, false
	}

	curOpenTimes := useTimes + 1
	nextTemplate := startBoxTemplate
	for nextTemplate != nil {
		if nextTemplate.Times == curOpenTimes {
			return nextTemplate, false
		}
		if nextTemplate.NextId == 0 {
			return nextTemplate, true
		}
		nextTemplate = boxtemplate.GetBoxTemplateService().GetBoxTemplate(nextTemplate.NextId)
	}

	return nil, false
}

//开启宝箱
func OpenBox(pl player.Player, itemId, useNum int32, chooseIndexList []int32, startBoxTemplate *gametemplate.BoxTemplate) (flag bool, err error) {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	boxTemplate := GetBoxTemplate(pl, startBoxTemplate)
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	parentBindType := itemTemplate.GetBindType()

	//判断银两是否足够
	needSilver := int64(boxTemplate.UseSilver * useNum)
	if needSilver > 0 {
		if !propertyManager.HasEnoughSilver(needSilver) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"itemId":   itemId,
					"useNum":   useNum,
				}).Warn("box:使用宝箱,银两不足，无法使用")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}

	//判断元宝是否足够
	needGold := boxTemplate.UseGold * useNum
	needBindGold := boxTemplate.UseBindgold * useNum
	needTotalGold := needGold + needBindGold
	if needGold > 0 {
		if !propertyManager.HasEnoughGold(int64(needGold), false) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"itemId":   itemId,
					"needGold": needGold,
					"useNum":   useNum,
				}).Warn("box:使用宝箱,元宝不足，无法使用")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}
	if needTotalGold > 0 {
		if !propertyManager.HasEnoughGold(int64(needTotalGold), true) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"itemId":   itemId,
					"useNum":   useNum,
				}).Warn("box:使用宝箱,绑元不足，无法使用")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//判断消耗道具是否足够
	needItemMap := boxTemplate.GetUseItemMap(pl.GetRole(), useNum)
	if len(needItemMap) > 0 {
		if !inventoryManager.HasEnoughItems(needItemMap) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"itemId":   itemId,
					"useNum":   useNum,
				}).Warn("box:使用宝箱,道具不足，无法使用")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//计算掉落id
	dropIdList := boxTemplate.GetDropIdList(pl.GetRole(), pl.GetSex())
	//自由选择
	if boxTemplate.GetBoxType() == treasureboxtypes.ChooseRes {
		if len(chooseIndexList) < 1 {
			log.WithFields(
				log.Fields{
					"playerId":        pl.GetId(),
					"itemId":          itemId,
					"useNum":          useNum,
					"chooseIndexList": chooseIndexList,
				}).Warn("box:使用自选宝箱,自由选择数量错误")
			playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
			return
		}

		var newDropIdList []int32
		for _, chooseIndex := range chooseIndexList {
			targetPos := chooseIndex + 1
			lenOfDrop := int32(len(dropIdList))
			if targetPos < 1 || targetPos > lenOfDrop {
				log.WithFields(
					log.Fields{
						"playerId":        pl.GetId(),
						"targetPos":       targetPos,
						"lenOfDrop":       lenOfDrop,
						"chooseIndexList": chooseIndexList,
					}).Warn("box:自由选择超出宝箱物品数量")
				playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
				return
			}

			newDropIdList = append(newDropIdList, dropIdList[chooseIndex])
		}
		dropIdList = newDropIdList
	}

	var dropItemList []*droptemplate.DropItemData
	for index := int32(0); index < useNum; index++ {
		dropList := droptemplate.GetDropTemplateService().GetDropListItemLevelList(dropIdList)
		dropItemList = append(dropItemList, dropList...)
	}

	var rewItemList []*droptemplate.DropItemData
	var rewResMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(dropItemList) != 0 {
		rewItemList, rewResMap = droplogic.SeperateItemDatas(dropItemList)
	}
	//继承绑定属性
	for _, itemData := range rewItemList {
		itemData.BindType = parentBindType
	}

	//背包空间
	if !inventoryManager.HasEnoughSlotsOfItemLevel(rewItemList) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
				"useNum":   useNum,
			}).Warn("box:使用宝箱,背包空间不足，请清理后再来")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//消耗物品
	if len(needItemMap) > 0 {
		itemUsereason := commonlog.InventoryLogReasonBoxUse
		if flag := inventoryManager.BatchRemove(needItemMap, itemUsereason, itemUsereason.String()); !flag {
			panic(fmt.Errorf("box: box use item should be ok"))
		}
	}

	//合成消耗银两元宝
	if needBindGold > 0 || needGold > 0 || needSilver > 0 {
		silverUseReason := commonlog.SilverLogReasonBoxUse
		goldUseReason := commonlog.GoldLogReasonBoxUse
		flag = propertyManager.Cost(int64(needBindGold), int64(needGold), goldUseReason, goldUseReason.String(), needSilver, silverUseReason, silverUseReason.String())
		if !flag {
			panic(fmt.Errorf("box: box use silver or gold should be ok"))
		}
	}

	//获得物品
	if len(rewItemList) > 0 {
		itemGetReason := commonlog.InventoryLogReasonBoxRew
		flag = inventoryManager.BatchAddOfItemLevel(rewItemList, itemGetReason, itemGetReason.String())
		if !flag {
			panic("box: box add item should be ok")
		}
	}

	//获得资源
	if len(rewResMap) > 0 {
		goldGetReason := commonlog.GoldLogReasonBoxGet
		silverGetReason := commonlog.SilverLogReasonBoxGet
		levelGetReason := commonlog.LevelLogReasonBoxGet
		err = droplogic.AddRes(pl, rewResMap, goldGetReason, goldGetReason.String(), silverGetReason, silverGetReason.String(), levelGetReason, levelGetReason.String())
		if err != nil {
			return
		}
	}

	//同步背包
	inventorylogic.SnapInventoryChanged(pl)

	//同步资源
	propertylogic.SnapChangedProperty(pl)

	//消息
	scInventoryBoxDropInfo := pbutil.BuildSCInventoryBoxDropInfo(dropItemList)
	pl.SendMsg(scInventoryBoxDropInfo)

	flag = true
	return
}
