package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	itemservice "fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	wushuangweaponlogic "fgame/fgame/game/wushuangweapon/logic"
	"fgame/fgame/game/wushuangweapon/pbutil"
	playerwushuangweapon "fgame/fgame/game/wushuangweapon/player"
	wushuangweapontypes "fgame/fgame/game/wushuangweapon/types"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WUSHUANGWEAPON_BREAKTHROUGH_TYPE), dispatch.HandlerFunc(handlerBreakthrough))
}

func handlerBreakthrough(s session.Session, msg interface{}) (err error) {
	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSWushuangWeaponBreakthrough)
	csBodyPos := csMsg.GetBodyPos()

	bodyPos := wushuangweapontypes.WushuangWeaponPart(csBodyPos)
	if !bodyPos.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bodyPos":  bodyPos,
			}).Warn("WushuangWeapon:部位突破请求，类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = breakthrough(tpl, bodyPos, csMsg.GetIndexList())
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("WushuangWeapon:处理部位突破请求，错误")
		return
	}
	return
}

func breakthrough(pl player.Player, bodyPos wushuangweapontypes.WushuangWeaponPart, indexList []int32) (err error) {

	// if coreutils.IfRepeatElementInt32(indexList) {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId":  pl.GetId(),
	// 			"indexList": indexList,
	// 		}).Warn("wushuangWeapon:处理突破,索引重复")
	// 	playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
	// 	return
	// }

	allItemId := make([]int32, 0, len(indexList))
	useItemMap := make(map[int32]int32)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	itemTemplate := itemservice.GetItemService()

	for _, index := range indexList {
		item := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, index)
		if item == nil || item.IsEmpty() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"index":    index,
				}).Warn("wushuangWeapon:物品错误，物品不存在")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
			return
		}
		itemId := item.ItemId
		allItemId = append(allItemId, itemId)
		useItemMap[itemId]++
	}

	for itemId, itemNum := range useItemMap {
		if inventoryManager.NumOfItems(itemId) < itemNum {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"itemId":   itemId,
					"itemNum":  itemNum,
				}).Warn("wushuangWeapon:突破物品数量不够")
			playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		}
	}

	wushuangDataManager := pl.GetPlayerDataManager(playertypes.PlayerWushuangWeaponDataManagerType).(*playerwushuangweapon.PlayerWushuangWeaponDataManager)
	slotObj := wushuangDataManager.GetSlotObjectFromBodyPos(bodyPos)

	//判断身上是否有装备
	if !slotObj.IsEquip() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bodyPos":  bodyPos,
			}).Warn("wushuangWeapon:该部位没有装备，不能突破")
		playerlogic.SendSystemMessage(pl, lang.WushuangWeaponBodyPosEquipmentNotExist)
		return
	}
	curlevel := slotObj.GetLevel()
	itemTemp := itemTemplate.GetItem(int(slotObj.GetItemId()))
	if itemTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   slotObj.GetItemId(),
			}).Warn("wushuangWeapon:部位装备模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	baseTemp := itemTemp.GetWushuangBaseTemplate()
	if baseTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bodyPos":  bodyPos,
				"curlevel": curlevel,
			}).Warn("wushuangWeapon:base模板不存在！")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	wushuangStrengthenTemplate := baseTemp.GetStrengthTemplateByLevel(curlevel)
	if wushuangStrengthenTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bodyPos":  bodyPos,
				"curlevel": curlevel,
			}).Warn("wushuangWeapon:strenth模板不存在！")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//判断是否满级
	if wushuangStrengthenTemplate.IsMaxLevel() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bodyPos":  bodyPos,
				"level":    curlevel,
			}).Warn("wushuangWeapon:满级了，不能突破")
		playerlogic.SendSystemMessage(pl, lang.WushuangWeaponLevelFull)
		return
	}

	//判断经验值是否满了
	curExperience := slotObj.GetExperience()
	nextTemplate := wushuangStrengthenTemplate.GetNextStrengthenTemplate()
	if nextTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":                   pl.GetId(),
				"wushuangStrengthenTemplate": wushuangStrengthenTemplate.Id,
			}).Warn("wushuangWeapon:下一级模板不存在！")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	maxExperience := nextTemplate.GetAllNeedExperience()
	if curExperience < maxExperience {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"curExperience": curExperience,
			}).Warn("wushuangWeapon:经验值未满，不能突破")
		playerlogic.SendSystemMessage(pl, lang.WushuangWeaponExperienceNotFull)
		return
	}

	//判断是否需要突破
	if nextTemplate.IsNeedBreakThrough() {
		//判断物品转数以及物品数目是否足够
		needEquipNum := int(nextTemplate.TupoNeedCount)
		if len(allItemId) != needEquipNum {
			log.WithFields(
				log.Fields{
					"playerId":     pl.GetId(),
					"allItemIdLen": len(allItemId),
					"needEquipNum": needEquipNum,
				}).Warn("wushuangWeapon:吞噬物品数量不够")
			playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
			return
		}

		minZhuangshu := nextTemplate.TupoZhuanshu
		minQuality := nextTemplate.GetMinNeedQuality()
		for itemId, _ := range useItemMap {
			itemTemp := itemTemplate.GetItem(int(itemId))
			if itemTemp == nil {
				log.WithFields(
					log.Fields{
						"playerId": pl.GetId(),
						"itemId":   itemId,
					}).Warn("wushuangWeapon:吞噬物品模板不存在")
				playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
				return
			}
			if itemTemp.GetItemType() != itemtypes.ItemTypeWushuangWeaponEssence || itemTemp.GetItemSubType() != itemtypes.ItemWushuangWeaponEssenceSubTypeMaterial {
				log.WithFields(
					log.Fields{
						"playerId":    pl.GetId(),
						"itemType":    itemTemp.GetItemType(),
						"itemSubType": itemTemp.GetItemSubType(),
					}).Warn("wushuangWeapon:吞噬物品类型错误")
				playerlogic.SendSystemMessage(pl, lang.WushuangWeaponBreakthroughItemTypeWrong)
				return
			}
			if itemTemp.NeedZhuanShu < minZhuangshu {
				log.WithFields(
					log.Fields{
						"playerId":      pl.GetId(),
						"itemZhuangShu": itemTemp.NeedZhuanShu,
					}).Warn("wushuangWeapon:吞噬物品转数不够")
				playerlogic.SendSystemMessage(pl, lang.WushuangWeaponBreakthroughItemZhuanshuWrong)
				return
			}

			if itemTemp.GetQualityType() < minQuality {
				log.WithFields(
					log.Fields{
						"playerId":   pl.GetId(),
						"minQuality": minQuality,
					}).Warn("wushuangWeapon:吞噬物品品质不足")
				playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
				return
			}
		}

		//突破（成功率）
		// isBreakthrough := false
		successRate := nextTemplate.TupoRate
		isBreakthrough := mathutils.RandomHit(common.MAX_RATE, int(successRate))

		if isBreakthrough {
			slotObj.Uplevel()
			wushuangweaponlogic.WushuangWeaponPropertyChanged(pl)
		}

		//同步物品（删除5件物品）
		useReason := commonlog.InventoryLogReasonWushuangWeaponBreakthrough
		useReasonText := fmt.Sprintf(commonlog.InventoryLogReasonWushuangWeaponBreakthrough.String(), bodyPos.String(), slotObj.GetLevel())
		flag := inventoryManager.BatchRemove(useItemMap, useReason, useReasonText)
		if !flag {
			panic("inventory:移除物品应该是可以的")
		}

		inventorylogic.SnapInventoryChanged(pl)

	} else {
		//普通的直接升级
		slotObj.Uplevel()
		wushuangweaponlogic.WushuangWeaponPropertyChanged(pl)
	}

	//等级够了就自动穿戴无双神器外观
	if baseTemp.IsCanActiveShow(slotObj.GetLevel()) {
		wushuangweaponlogic.ActiveShow(pl, bodyPos, slotObj.GetItemId())
	}

	//发消息
	scMsg := pbutil.BuildSCWushuangWeaponBreakthrough(bodyPos, slotObj.GetLevel())
	pl.SendMsg(scMsg)
	return
}
