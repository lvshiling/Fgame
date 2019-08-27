package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	coreutils "fgame/fgame/core/utils"
	chattypes "fgame/fgame/game/chat/types"
	commomlogic "fgame/fgame/game/common/logic"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	goldequipeventtypes "fgame/fgame/game/goldequip/event/types"
	"fgame/fgame/game/goldequip/pbutil"
	goldequippbutil "fgame/fgame/game/goldequip/pbutil"
	playergoldequip "fgame/fgame/game/goldequip/player"
	goldequiptemplate "fgame/fgame/game/goldequip/template"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	playerrealm "fgame/fgame/game/realm/player"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//推送金装物品改变
func SnapInventoryGoldEquipChanged(pl player.Player) (err error) {
	goldequipManager := pl.GetPlayerDataManager(playertypes.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	slotChangedList := goldequipManager.GetChangedEquipmentSlotAndReset()
	if len(slotChangedList) <= 0 {
		return
	}
	goldEquipmentChanged := pbutil.BuildSCGoldEquipSlotChanged(slotChangedList)
	pl.SendMsg(goldEquipmentChanged)
	return
}

//变更转生属性
func ZhuanShengPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeZhuanSheng.Mask())
	return
}

//金装属性
func GoldEquipPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeGoldequip.Mask())
	return
}

//金装开光判断
func GoldEquipOpenLight(pl player.Player, curTimesNum int32, openlightTemp *gametemplate.GoldEquipOpenLightTemplate) (sucess bool) {
	timesMin := openlightTemp.TimesMin
	timesMax := openlightTemp.TimesMax
	updateRate := openlightTemp.SuccessRate
	_, _, sucess = commomlogic.GetStatusAndProgress(curTimesNum, 0, timesMin, timesMax, 0, 0, updateRate, 1)
	return
}

//计算强化成功率
func CountGoldEquipStrengthenRate(curLevel int32, itemMap map[int32]int32) int32 {
	rate := int32(0)
	for itemId, num := range itemMap {
		itemTemp := item.GetItemService().GetItem(int(itemId))
		rateId := itemTemp.TypeFlag1
		rateTemp := goldequiptemplate.GetGoldEquipTemplateService().GetGoldEquipStrengthenRateTemplate(rateId)
		if rateTemp == nil {
			continue
		}

		rate += rateTemp.OfferRate(curLevel) * num
	}
	return rate
}

//使用元神金装
func HandleUseGoldEquip(pl player.Player, index int32) (err error) {
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	goldequipManager := pl.GetPlayerDataManager(playertypes.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	it := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, index)
	//物品不存在
	if it == nil || it.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("inventory:使用元神金装,物品不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}
	itemId := it.ItemId
	level := it.Level
	propertyData := it.PropertyData
	bind := it.BindType
	//判断物品是否可以装备
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if !itemTemplate.IsGoldEquip() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("inventory:使用元神金装,此物品不是元神金装")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotEquip)
		return
	}
	if itemTemplate.NeedProfession != 0 {
		//角色
		if itemTemplate.GetRole() != pl.GetRole() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"index":    index,
				}).Warn("inventory:使用元神金装,角色不符")
			playerlogic.SendSystemMessage(pl, lang.PlayerRoleWrong)
			return
		}
	}
	if itemTemplate.GetSex() != 0 {
		//性别
		if itemTemplate.GetSex() != pl.GetSex() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"index":    index,
				}).Warn("inventory:使用元神金装,性别不符")
			playerlogic.SendSystemMessage(pl, lang.PlayerSexWrong)
			return
		}
	}
	//判断级别
	if itemTemplate.NeedLevel > pl.GetLevel() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("inventory:使用元神金装,等级不够")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
		return
	}

	//判断转数
	if itemTemplate.NeedZhuanShu > propertyManager.GetZhuanSheng() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("inventory:使用元神金装,转数不够")
		playerlogic.SendSystemMessage(pl, lang.PlayerZhuanShengTooLow)
		return
	}

	//判断是否已经装备
	equipmentSubType := itemTemplate.GetItemSubType().(itemtypes.ItemGoldEquipSubType)
	pos := equipmentSubType.Position()
	equipmentItem := goldequipManager.GetGoldEquipByPos(pos)
	if equipmentItem != nil && !equipmentItem.IsEmpty() {
		flag := takeOffInternal(pl, pos)
		if !flag {
			return
		}
	}
	//移除物品
	reasonText := commonlog.InventoryLogReasonPutOn.String()
	flag, _ := inventoryManager.RemoveIndex(inventorytypes.BagTypePrim, index, 1, commonlog.InventoryLogReasonPutOn, reasonText)
	if !flag {
		panic("inventory:移除物品应该是可以的")
	}

	flag = goldequipManager.PutOn(pos, itemId, level, bind, propertyData)
	if !flag {
		panic(fmt.Errorf("inventory:穿上位置 [%s]应该是可以的", pos.String()))
	}

	GoldEquipPropertyChanged(pl)

	//同步改变
	SnapInventoryGoldEquipChanged(pl)
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	scInventoryUseEquip := pbutil.BuildSCUseGoldEquip(index)
	pl.SendMsg(scInventoryUseEquip)

	teShuSkillList := goldequipManager.GetTeShuSkillList()
	pl.ResetTeShuSkills(teShuSkillList)
	return
}

//脱下
func HandleTakeOff(pl player.Player, pos inventorytypes.BodyPositionType) (err error) {
	flag := takeOffInternal(pl, pos)
	if !flag {
		return
	}
	GoldEquipPropertyChanged(pl)
	//同步改变
	SnapInventoryGoldEquipChanged(pl)
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	//脱下成功
	scInventoryTakeOffEquip := pbutil.BuildTakeOffGoldEquip(pos)
	pl.SendMsg(scInventoryTakeOffEquip)
	goldequipManager := pl.GetPlayerDataManager(playertypes.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	teShuSkillList := goldequipManager.GetTeShuSkillList()
	pl.ResetTeShuSkills(teShuSkillList)
	return nil
}

func takeOffInternal(pl player.Player, pos inventorytypes.BodyPositionType) (flag bool) {
	goldequipManager := pl.GetPlayerDataManager(playertypes.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//没有东西
	item := goldequipManager.GetGoldEquipByPos(pos)
	if item == nil || item.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos,
			}).Warn("goldequip:脱下金装,金装不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipCanNotTakeOff)
		return
	}

	level := item.GetLevel()
	num := int32(1)
	bind := item.GetBindType()
	propertyData := item.GetPropertyData()

	//背包空间
	if !inventoryManager.HasEnoughSlotItemLevel(item.GetItemId(), num, level, bind) {
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	itemId := goldequipManager.TakeOff(pos)
	if itemId == 0 {
		panic(fmt.Errorf("goldequip:take off should more than 0"))
	}

	//添加物品
	reasonText := commonlog.InventoryLogReasonTakeOff.String()
	flag = inventoryManager.AddItemLevelWithPropertyData(itemId, num, level, bind, propertyData, commonlog.InventoryLogReasonTakeOff, reasonText)
	if !flag {
		panic(fmt.Errorf("goldequip:add item should be success"))
	}
	return
}

//吞噬
func HandleGoldEquipEat(pl player.Player, isAuto int32, itemIndexList []int32) (err error) {
	if len(itemIndexList) < 1 {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"indexList": itemIndexList,
			}).Warn("inventory:处理吞噬元神金装,没有装备")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipEquipmentNotItemEat)
		return
	}

	if coreutils.IfRepeatElementInt32(itemIndexList) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"indexList": itemIndexList,
			}).Warn("inventory:处理吞噬元神金装,索引重复")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	// 计算分解物品
	totalExp, returnItemMap, fenJieItemIdList, flag := countEatEquip(pl, itemIndexList)
	if !flag {
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	slotEnoughFlag := inventoryManager.HasEnoughSlots(returnItemMap)
	if !slotEnoughFlag {
		if isAuto == 0 {
			log.WithFields(
				log.Fields{
					"playerId":      pl.GetId(),
					"ItemIndex":     itemIndexList,
					"isAuto":        isAuto,
					"returnItemMap": returnItemMap,
				}).Warn("goldequip:吞噬失败,背包不足")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}
	}

	//消耗装备
	useReason := commonlog.InventoryLogReasonGoldEquipEat
	useReasonText := fmt.Sprintf(useReason.String(), isAuto)
	flag, err = inventoryManager.BatchRemoveIndex(inventorytypes.BagTypePrim, itemIndexList, useReason, useReasonText)
	if err != nil {
		return
	}
	if !flag {
		panic(fmt.Errorf("goldEquip:消耗物品应该成功"))
	}

	// 返还物品
	if len(returnItemMap) > 0 {
		if !slotEnoughFlag {
			title := lang.GetLangService().ReadLang(lang.GoldEquipAutoFenJieMailTitle)
			content := lang.GetLangService().ReadLang(lang.GoldEquipAutoFenJieSlotNotEnoughMailContent)
			emaillogic.AddEmail(pl, title, content, returnItemMap)
		} else {
			addItemReason := commonlog.InventoryLogReasonGoldEquipTunShiReturn
			addItemReasonText := fmt.Sprintf(addItemReason.String(), isAuto)
			flag := inventoryManager.BatchAdd(returnItemMap, addItemReason, addItemReasonText)
			if !flag {
				panic(fmt.Errorf("goldEquip:添加物品应该成功"))
			}
		}
	}

	if totalExp > 0 {
		goldLevelReason := commonlog.GoldYuanLevelLogReasonEatEquip
		propertyManager.AddGoldYuanExp(totalExp, goldLevelReason, goldLevelReason.String())
	}

	//玩家分解日志
	goldequipManager := pl.GetPlayerDataManager(playertypes.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	goldequipManager.AddGoldEquipLog(fenJieItemIdList, ConverToLogStr(returnItemMap))

	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)
	//发送事件
	gameevent.Emit(goldequipeventtypes.EventTypeGoldEquipResolve, pl, int32(len(itemIndexList)))

	scMsg := goldequippbutil.BuildSCEatGoldEquip(propertyManager.GetGoldYuanLevel(), propertyManager.GetGoldYuanExp())
	pl.SendMsg(scMsg)
	return
}

// 吞噬物品发送邮件
func GodCastingInheritSendEmail(pl player.Player, itemIndexList []int32, level int32, openLightLevel int32) (err error) {
	if len(itemIndexList) < 1 {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"indexList": itemIndexList,
			}).Warn("inventory:处理吞噬元神金装,没有装备")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipEquipmentNotItemEat)
		return
	}

	if coreutils.IfRepeatElementInt32(itemIndexList) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"indexList": itemIndexList,
			}).Warn("inventory:处理吞噬元神金装,索引重复")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	// 计算分解物品
	totalExp, returnItemMap, _, flag := countEatEquip(pl, itemIndexList)
	if !flag {
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//消耗装备
	useReason := commonlog.InventoryLogReasonGoldEquipEat
	useReasonText := fmt.Sprintf(useReason.String(), int32(0))
	flag, err = inventoryManager.BatchRemoveIndex(inventorytypes.BagTypePrim, itemIndexList, useReason, useReasonText)
	if err != nil {
		return
	}
	if !flag {
		panic(fmt.Errorf("goldEquip:消耗物品应该成功"))
	}

	if len(returnItemMap) > 0 {
		title := lang.GetLangService().ReadLang(lang.GoldEquipGodCastingInheritMailTitle)
		levelStr := coreutils.FormatColor(chattypes.ColorTypeEmailKeyWord, fmt.Sprintf("%d", level))
		openLightLevelStr := coreutils.FormatColor(chattypes.ColorTypeEmailKeyWord, fmt.Sprintf("%d", openLightLevel))
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.GoldEquipGodCastingInheritMailContent), levelStr, openLightLevelStr)
		emaillogic.AddEmail(pl, title, content, returnItemMap)
	}

	if totalExp > 0 {
		goldLevelReason := commonlog.GoldYuanLevelLogReasonEatEquip
		propertyManager.AddGoldYuanExp(totalExp, goldLevelReason, goldLevelReason.String())
	}

	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)
	return
}

//计算分解装备
func countEatEquip(pl player.Player, itemIndexList []int32) (totalExp int64, returnItemMap map[int32]int32, fenJieItemIdList []int32, flag bool) {
	returnItemMap = make(map[int32]int32)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	for _, itemIndex := range itemIndexList {
		it := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, itemIndex)
		if it == nil {
			log.WithFields(
				log.Fields{
					"index": itemIndex,
				}).Warn("goldequip:格子不存在")
			return
		}
		if it.ItemId == 0 {
			log.WithFields(
				log.Fields{
					"index": itemIndex,
				}).Warn("goldequip:格子没有物品")
			return
		}

		itemTemp := item.GetItemService().GetItem(int(it.ItemId))
		//只分解元神金装和系统装备
		if !itemTemp.IfCanFenJie() {
			log.WithFields(
				log.Fields{
					"index":       itemIndex,
					"itemId":      it.ItemId,
					"itemType":    itemTemp.GetItemType(),
					"itemSubType": itemTemp.GetItemSubType(),
				}).Warn("goldequip:分解失败，该物品不允许分解")
			return
		}

		if !coreutils.ContainInt32(fenJieItemIdList, it.ItemId) {
			fenJieItemIdList = append(fenJieItemIdList, it.ItemId)
		}

		// 分解金装
		if itemTemp.IsGoldEquip() {
			// 计算装备提供的经验
			goldEquipTemp := itemTemp.GetGoldEquipTemplate()
			totalExp += int64(goldEquipTemp.TunshiExp)

			// 返还物品
			data := it.PropertyData.(*goldequiptypes.GoldEquipPropertyData)
			if data.OpenLightLevel >= 0 {
				openLightTemp := goldEquipTemp.GetOpenLightTemplate(data.OpenLightLevel)
				if openLightTemp == nil {
					goto UpstarLevel
				}
				for itemId, num := range openLightTemp.GetReturnItemMap() {
					_, ok := returnItemMap[itemId]
					if !ok {
						returnItemMap[itemId] = num
					} else {
						returnItemMap[itemId] += num
					}
				}
			}
		UpstarLevel:
			if data.UpstarLevel >= 0 {
				upstarTemp := goldEquipTemp.GetUpstarTemplate(data.UpstarLevel)
				if upstarTemp == nil {
					goto Level
				}
				for itemId, num := range upstarTemp.GetReturnItemMap() {
					_, ok := returnItemMap[itemId]
					if !ok {
						returnItemMap[itemId] = num
					} else {
						returnItemMap[itemId] += num
					}
				}
			}
		Level:
			if it.Level >= 0 {
				strengthenTemp := goldEquipTemp.GetStrengthenTemplate(it.Level)
				if strengthenTemp == nil {
					goto Done
				}
				for itemId, num := range strengthenTemp.GetReturnItemMap() {
					_, ok := returnItemMap[itemId]
					if !ok {
						returnItemMap[itemId] = num
					} else {
						returnItemMap[itemId] += num
					}
				}
			}
		Done:
			//神器精华和神铸结晶
			wushuangDropItemTemp := droptemplate.GetDropTemplateService().GetDropFromGroup(int32(goldEquipTemp.TunshiDrop))
			if wushuangDropItemTemp != nil {
				itemId := wushuangDropItemTemp.ItemId
				num := int32(mathutils.RandomRange(int(wushuangDropItemTemp.MinCount), int(wushuangDropItemTemp.MaxCount)))
				_, ok := returnItemMap[itemId]
				if !ok {
					returnItemMap[itemId] = num
				} else {
					returnItemMap[itemId] += num
				}
			}
		}

		// 分解系统装备
		sysEquipTemp := itemTemp.GetSystemEquipTemplate()
		if sysEquipTemp != nil {
			totalExp += int64(sysEquipTemp.GetTushiExp())
			for itemId, num := range sysEquipTemp.GetReturnItemMap() {
				_, ok := returnItemMap[itemId]
				if !ok {
					returnItemMap[itemId] = num
				} else {
					returnItemMap[itemId] += num
				}
			}
		}
	}
	flag = true
	return
}

//分解奖励日志
func ConverToLogStr(rewMap map[int32]int32) string {
	itemNameLinkStr := ""
	for itemId, num := range rewMap {
		itemTemplate := item.GetItemService().GetItem(int(itemId))
		itemName := coreutils.FormatColor(itemTemplate.GetQualityType().GetColor(), coreutils.FormatNoticeStr(itemTemplate.FormateItemNameOfNum(num)))
		if len(itemNameLinkStr) == 0 {
			itemNameLinkStr += itemName
		} else {
			itemNameLinkStr += ", " + itemName
		}
	}
	return itemNameLinkStr
}

// 自动分解装备
func AutoFenJieGoldEquip(pl player.Player) (err error) {
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	goldequipManager := pl.GetPlayerDataManager(playertypes.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	setiingObj := goldequipManager.GetGoldEquipSetting()
	if !setiingObj.IsFenJieAuto() {
		return
	}

	var fenJieItemIndex []int32
	itemObjList := inventoryManager.GetBagAll(inventorytypes.BagTypePrim)
	for _, ito := range itemObjList {
		if ito.IsEmpty() {
			continue
		}

		itemTemp := item.GetItemService().GetItem(int(ito.ItemId))
		// 不是金装
		if !itemTemp.IsGoldEquip() {
			continue
		}

		// 超过品质
		if itemTemp.GetQualityType() > setiingObj.GetFenJieQuality() {
			continue
		}
		// 超过品质
		if itemTemp.NeedZhuanShu > setiingObj.GetFenJieZhuanShu() {
			continue
		}
		fenJieItemIndex = append(fenJieItemIndex, ito.Index)
	}

	if len(fenJieItemIndex) > 0 {
		err = HandleGoldEquipEat(pl, setiingObj.GetFenJieIsAuto(), fenJieItemIndex)
		if err != nil {
			return
		}
	}

	return
}

//返回老玩家宝石槽解锁
func RestitutionGemUnlock(pl player.Player) {
	goldequipManager := pl.GetPlayerDataManager(playertypes.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	goldequipBag := goldequipManager.GetGoldEquipBag()
	slotList := goldequipBag.GetAll()
	for _, slot := range slotList {
		for order, itemId := range slot.GemInfo {
			_, ok := slot.GemUnlockInfo[order]
			if itemId > 0 && !ok {
				goldequipBag.UnlockGem(slot.GetSlotId(), order)
			}
		}
	}
}

//镶嵌宝石时检测宝石槽位是否解锁
func CheckGemUnlockByUse(pl player.Player, slotPosition inventorytypes.BodyPositionType, order int32) (flag bool) {
	goldequipManager := pl.GetPlayerDataManager(playertypes.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	realmManager := pl.GetPlayerDataManager(playertypes.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
	goldequipBag := goldequipManager.GetGoldEquipBag()

	if goldequipBag.IsUnlockGem(slotPosition, order) {
		flag = true
		return
	}

	//等级不够
	gemSlotTemplate := item.GetItemService().GetEquipGemSlotTemplate(slotPosition, order)
	if gemSlotTemplate.NeedLevel > propertyManager.GetLevel() {
		return
	}

	//天截塔等级不够
	if gemSlotTemplate.NeedLayer > realmManager.GetTianJieTaLevel() {
		return
	}

	items := gemSlotTemplate.GetNeedItemMap()
	if len(items) != 0 {
		return
	}
	goldequipBag.UnlockGem(slotPosition, order)
	flag = true
	return
}

//检测老玩家金装强化等级迁移
func CheckStrengthenLevelMove(pl player.Player) {
	goldequipManager := pl.GetPlayerDataManager(playertypes.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	if goldequipManager.GetGoldEquipSetting().IsCheckOldStLev() {
		return
	}
	goldequipManager.GetGoldEquipSetting().SetIsCheckOldSt()
	goldequipBag := goldequipManager.GetGoldEquipBag()
	slotList := goldequipBag.GetAll()
	for _, slot := range slotList {
		if slot.IsEmpty() {
			continue
		}
		propertyData, ok := slot.GetPropertyData().(*goldequiptypes.GoldEquipPropertyData)
		if !ok {
			continue
		}
		posType := slot.GetSlotId()
		level := propertyData.UpstarLevel
		temp := goldequiptemplate.GetGoldEquipTemplateService().GetGoldEquipStrengthenBuWeiTemplate(posType, level)
		if temp == nil {
			continue
		}
		if level > 0 {
			goldequipBag.SetStrengthBuWeiLevel(posType, level)
		}
	}
}
