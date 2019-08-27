package use

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
	boxlogic "fgame/fgame/game/treasurebox/logic"
	"fgame/fgame/game/treasurebox/pbutil"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeGiftBag, itemtypes.ItemGiftBagSubTypeQianKunDai, playerinventory.ItemUseHandleFunc(handleItemUse))
}

func handleItemUse(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	itemId := it.ItemId
	_, totalUseTimes := inventoryManager.GetItemUseTimes(itemId)
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	starBoxTemplate := itemTemplate.GetBoxTemplate()
	parentBindType := itemTemplate.GetBindType()
	boxTemplate, isUseOut := boxlogic.GetQianKunBoxTemplate(pl, starBoxTemplate, totalUseTimes)

	if isUseOut {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"itemId":        itemId,
				"totalUseTimes": totalUseTimes,
			}).Warn("box:使用宝箱,乾坤袋开启完毕，无法再次使用")
		playerlogic.SendSystemMessage(pl, lang.InventoryQianKunDaiUseOut)
		return
	}

	if boxTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"itemId":        itemId,
				"totalUseTimes": totalUseTimes,
			}).Warn("box:使用宝箱,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//判断银两是否足够
	needSilver := int64(boxTemplate.UseSilver * num)
	if needSilver > 0 {
		if !propertyManager.HasEnoughSilver(needSilver) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"itemId":   itemId,
					"num":      num,
				}).Warn("box:使用宝箱,银两不足，无法使用")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}
	//判断元宝是否足够
	needGold := boxTemplate.UseGold * num
	needBindGold := boxTemplate.UseBindgold * num
	needTotalGold := needGold + needBindGold
	if needGold > 0 {
		if !propertyManager.HasEnoughGold(int64(needGold), false) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"itemId":   itemId,
					"num":      num,
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
					"num":      num,
				}).Warn("box:使用宝箱,绑元不足，无法使用")
			playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
			return
		}
	}

	//判断消耗道具是否足够
	needItemMap := boxTemplate.GetUseItemMap(pl.GetRole(), num)
	if len(needItemMap) > 0 {
		if !inventoryManager.HasEnoughItems(needItemMap) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"itemId":   itemId,
					"num":      num,
				}).Warn("box:使用宝箱,道具不足，无法使用")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	//计算掉落id
	dropIdList := boxTemplate.GetDropIdList(pl.GetRole(), pl.GetSex())
	dropItemList := droptemplate.GetDropTemplateService().GetDropListItemLevelList(dropIdList)

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
				"num":      num,
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

	//消耗银两元宝
	if needBindGold > 0 || needGold > 0 || needSilver > 0 {
		silverUseReason := commonlog.SilverLogReasonBoxUse
		goldUseReason := commonlog.GoldLogReasonBoxUse
		flag = propertyManager.Cost(int64(needBindGold), int64(needGold), goldUseReason, goldUseReason.String(), needSilver, silverUseReason, silverUseReason.String())
		if !flag {
			panic(fmt.Errorf("box: box use silver or gold should be ok"))
		}
	}

	//获得物品
	itemGetReason := commonlog.InventoryLogReasonBoxRew
	flag = inventoryManager.BatchAddOfItemLevel(rewItemList, itemGetReason, itemGetReason.String())
	if !flag {
		panic("box: box add item should be ok")
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
	propertylogic.SnapChangedProperty(pl)

	//消息
	scInventoryBoxDropInfo := pbutil.BuildSCInventoryBoxDropInfo(dropItemList)
	pl.SendMsg(scInventoryBoxDropInfo)

	flag = true
	return
}
