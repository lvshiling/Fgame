package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	"fgame/fgame/game/global"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	"fgame/fgame/game/inventory/pbutil"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fgame/fgame/pkg/mathutils"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

//推送物品改变
func SnapInventoryChanged(pl player.Player) {
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	for _, bag := range manager.GetShoulderBagMap() {
		itemChangedList := bag.GetChangedSlotAndReset()
		if len(itemChangedList) <= 0 {
			continue
		}
		inventoryChanged := pbutil.BuildSCInventoryChanged(itemChangedList)
		pl.SendMsg(inventoryChanged)
	}
	return
}

//推送装备物品改变
func SnapInventoryEquipChanged(pl player.Player) {
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	slotChangedList := manager.GetChangedEquipmentSlotAndReset()
	if len(slotChangedList) <= 0 {
		return
	}
	inventoryEquipmentChanged := pbutil.BuildSCInventoryEquipmentChanged(slotChangedList)
	pl.SendMsg(inventoryEquipmentChanged)
	return
}

//推送仓库物品改变
func SnapDepotChanged(pl player.Player) {
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	itemChangedList := manager.GetDepotChangedSlotAndReset()
	if len(itemChangedList) <= 0 {
		return
	}
	scDepotChanged := pbutil.BuildSCDepotChanged(itemChangedList)
	pl.SendMsg(scDepotChanged)
	return
}

//推送秘宝仓库物品改变
func SnapMiBaoDepotChanged(pl player.Player, typ equipbaokutypes.BaoKuType) {
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	itemChangedList := manager.GetMiBaoDepotChangedSlotAndReset(typ)
	if len(itemChangedList) <= 0 {
		return
	}
	scMiBaoDepotChanged := pbutil.BuildSCMiBaoDepotChanged(itemChangedList, int32(typ))
	pl.SendMsg(scMiBaoDepotChanged)
	return
}

func UpdateEquipmentProperty(pl player.Player) {
	pl.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeEquipment.Mask())
}

//稀有道具公告
func PrecioustemBroadcast(pl player.Player, itemId int32, num int32, msgCode lang.LangCode) {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	if itemTemplate.IsNotice() {
		itemName := coreutils.FormatColor(itemTemplate.GetQualityType().GetColor(), coreutils.FormatNoticeStrUnderline(itemTemplate.FormateItemNameOfNum(num)))
		linkArgs := []int64{int64(chattypes.ChatLinkTypeItem), int64(itemId)}
		itemNameLink := coreutils.FormatLink(itemName, linkArgs)
		content := fmt.Sprintf(lang.GetLangService().ReadLang(msgCode), playerName, itemNameLink)
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
		noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	}
}

//天机牌道具公告
func SecretCardPrecioustemBroadcast(pl player.Player, itemId int32, num int32) {
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	if itemTemplate.Quality > int32(itemtypes.ItemQualityTypeBlue) {
		itemName := coreutils.FormatColor(itemTemplate.GetQualityType().GetColor(), coreutils.FormatNoticeStrUnderline(itemTemplate.FormateItemNameOfNum(num)))
		linkArgs := []int64{int64(chattypes.ChatLinkTypeItem), int64(itemId)}
		itemNameLink := coreutils.FormatLink(itemName, linkArgs)
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.InventorySecretCardItemNotice), playerName, itemNameLink)
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
		noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)
	}
}

func UseItemIndex(pl player.Player, bagType inventorytypes.BagType, index int32, num int32, chooseIndexList []int32, args string) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	it := manager.FindItemByIndex(inventorytypes.BagTypePrim, index)
	if it == nil {
		return
	}
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	effectiveTime := itemTemplate.GetEffectiveTime()
	if effectiveTime > 0 && now < effectiveTime {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"itemGetTime": it.ItemGetTime,
			}).Warn("inventory:物品生效时间未到")
		dateStr := timeutils.MillisecondToTime(effectiveTime).Format("2006年01月02日15点")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNotReachEffectiveTime, dateStr)
		return
	}

	if it.PropertyData.GetExpireType() != inventorytypes.NewItemLimitTimeTypeNone {
		if it.PropertyData.IsExpire() {
			log.WithFields(
				log.Fields{
					"playerId":    pl.GetId(),
					"itemGetTime": it.PropertyData.GetItemGetTime(),
				}).Warn("inventory:物品已过期")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemExpire)
			return
		}
	} else {
		if itemTemplate.IsExpire(it.PropertyData.GetItemGetTime(), now) {
			log.WithFields(
				log.Fields{
					"playerId":    pl.GetId(),
					"itemGetTime": it.PropertyData.GetItemGetTime(),
				}).Warn("inventory:物品已过期")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemExpire)
			return
		}
	}

	if itemTemplate.IsCDItem() {
		if manager.IsItemIndexUseCd(index) {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"index":    index,
					"num":      num,
				}).Warn("inventory:物品使用CD中")
			playerlogic.SendSystemMessage(pl, lang.InventoryUseItemInCd)
			return
		}
	}

	if itemTemplate.NeedGender != 0 {
		if pl.GetSex() != itemTemplate.GetSex() {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"itemId":    itemTemplate.Id,
					"curLevel":  pl.GetLevel(),
					"needLevel": itemTemplate.NeedLevel,
				}).Warn("inventory:物品使用，性别不符")
			playerlogic.SendSystemMessage(pl, lang.PlayerSexWrong)
			return
		}
	}

	if itemTemplate.NeedProfession != 0 {
		if pl.GetRole() != itemTemplate.GetRole() {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"itemId":    itemTemplate.Id,
					"curLevel":  pl.GetLevel(),
					"needLevel": itemTemplate.NeedLevel,
				}).Warn("inventory:物品使用，职业不符")
			playerlogic.SendSystemMessage(pl, lang.PlayerRoleWrong)
			return
		}
	}

	isEnoughLevel := pl.GetLevel() >= itemTemplate.NeedLevel
	if !isEnoughLevel {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"itemId":    itemTemplate.Id,
				"curLevel":  pl.GetLevel(),
				"needLevel": itemTemplate.NeedLevel,
			}).Warn("inventory:物品使用，等级不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
		return
	}

	isEnoughZhuanSheng := pl.GetZhuanSheng() >= itemTemplate.NeedZhuanShu
	if !isEnoughZhuanSheng {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemTemplate.Id,
				"curZs":    pl.GetZhuanSheng(),
				"needZs":   itemTemplate.NeedZhuanShu,
			}).Warn("inventory:物品使用,转生数不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerZhuanShengTooLow)
		return
	}

	dayUseTimes, totalUseTimes := manager.GetItemUseTimes(itemId)
	if itemTemplate.IsLimitDayUseTimes() {
		maxTodayItemUseTimes := itemTemplate.LimitTimeDay
		leftDayUseTimes := maxTodayItemUseTimes - dayUseTimes
		if leftDayUseTimes < 1 {
			log.WithFields(
				log.Fields{
					"playerId":        pl.GetId(),
					"index":           index,
					"num":             num,
					"leftDayUseTimes": leftDayUseTimes,
				}).Warn("inventory:今日使用次数已达上限，请明日再来")
			playerlogic.SendSystemMessage(pl, lang.InventoryTodayUseTimesNotEnough)
			return
		}
		if num > leftDayUseTimes {
			num = leftDayUseTimes
		}
	}

	if itemTemplate.IsLimitTotalUseTimes() {
		maxTotalItemUseTimes := itemTemplate.LimitTimeAll
		leftTotalUseTimes := maxTotalItemUseTimes - totalUseTimes
		if leftTotalUseTimes < 1 {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"index":    index,
					"num":      num,
				}).Warn("inventory:物品使用总次数已达上限")
			playerlogic.SendSystemMessage(pl, lang.InventoryTotalUseTimesNotEnough)
			return
		}
		if num > leftTotalUseTimes {
			num = leftTotalUseTimes
		}
	}

	reasonText := commonlog.InventoryLogReasonUse.String()
	flag, err := manager.UseIndex(inventorytypes.BagTypePrim, index, num, chooseIndexList, args, commonlog.InventoryLogReasonUse, reasonText)
	if err != nil {
		return err
	}
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"itemId":   itemId,
				"num":      num,
			}).Warn("inventory:使用失败")
		// playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotUse)
		return
	}

	SnapInventoryChanged(pl)

	scInventoryItemUse := pbutil.BuildSCInventoryItemUse(bagType, index, itemId, num)
	pl.SendMsg(scInventoryItemUse)

	return
}

//强化升级
func EquipmentSlotStrengthenUpgrade(pl player.Player, pos inventorytypes.BodyPositionType, ignoreErrorMsg bool) (result inventorytypes.EquipmentStrengthenResultType, flag bool) {
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	item := manager.GetEquipByPos(pos)
	//物品不存在
	if item == nil || item.IsEmpty() {
		if ignoreErrorMsg {
			return
		}
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("inventory:强化升级失败,装备不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotNoEquip)
		return
	}

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	equipmentBag := manager.GetEquipmentBag()

	//判断槽位是否可以升星
	nextEquipmentStrengthenTemplate := equipmentBag.GetNextUpgradeEquipStrengthenTemplate(pos)
	if nextEquipmentStrengthenTemplate == nil {
		if ignoreErrorMsg {
			return
		}
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("inventory:强化升级失败,已经满级")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotStarMax)
		return
	}
	maxLevel := int32(math.Floor(float64(pl.GetLevel()) / float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeEquipmentStrengthenLevelLimit))))
	//判断等级限制
	if nextEquipmentStrengthenTemplate.Level > maxLevel {
		if ignoreErrorMsg {
			return
		}
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("inventory:强化升级失败,达到极限")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotLevelExceedLevel)
		return
	}

	items := nextEquipmentStrengthenTemplate.GetNeedItemMap()
	if len(items) != 0 {
		if !manager.HasEnoughItems(items) {
			if ignoreErrorMsg {
				return
			}
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"pos":      pos.String(),
				}).Warn("inventory:强化升级失败,物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}
	//判断是否有足够的银两
	if nextEquipmentStrengthenTemplate.SilverNum > 0 {
		if !propertyManager.HasEnoughSilver(int64(nextEquipmentStrengthenTemplate.SilverNum)) {
			if ignoreErrorMsg {
				return
			}
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"pos":      pos.String(),
				}).Warn("inventory:强化升级失败,银两不足")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}
	}
	if len(items) != 0 {
		reasonText := commonlog.InventoryLogReasonEquipSlotStrengthUpgrade.String()
		flag := manager.BatchRemove(items, commonlog.InventoryLogReasonEquipSlotStrengthUpgrade, reasonText)
		if !flag {
			panic(fmt.Errorf("inventory:装备槽强化升级移除材料应该成功"))
		}
	}
	if nextEquipmentStrengthenTemplate.SilverNum > 0 {
		reasonText := commonlog.SilverLogReasonEquipSlotStrengthUpgrade.String()
		flag := propertyManager.CostSilver(int64(nextEquipmentStrengthenTemplate.SilverNum), commonlog.SilverLogReasonEquipSlotStrengthUpgrade, reasonText)
		if !flag {
			panic(fmt.Errorf("inventory:装备槽强化升级花费银两应该成功"))
		}
	}

	success := mathutils.RandomHit(common.MAX_RATE, int(nextEquipmentStrengthenTemplate.SuccessRate))
	if success {
		//升星
		flag = equipmentBag.StrengthLevel(pos)
		if !flag {
			panic(fmt.Errorf("inventory: 强化升级应该成功"))
		}

		UpdateEquipmentProperty(pl)

		result = inventorytypes.EquipmentStrengthenResultTypeSuccess
		return
	}

	//判断是否会回退
	if nextEquipmentStrengthenTemplate.GetFailEquipStrengthenTemplate() != nil {
		fail := mathutils.RandomHit(common.MAX_RATE, int(nextEquipmentStrengthenTemplate.ReturnRate))
		if fail {
			//回退
			flag = equipmentBag.StrengthLevelBack(pos)
			if !flag {
				panic(fmt.Errorf("inventory: strength star back should be ok"))
			}
			result = inventorytypes.EquipmentStrengthenResultTypeBack
			return
		}
	}

	flag = true
	result = inventorytypes.EquipmentStrengthenResultTypeFailed
	return
}

//升阶
func EquipSlotUpgrade(pl player.Player, pos inventorytypes.BodyPositionType) {
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	equipBag := manager.GetEquipmentBag()
	item := manager.GetEquipByPos(pos)
	//物品不存在
	if item == nil || item.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("inventory:强化升阶,装备不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotNoEquip)
		return
	}

	nextItemTemplate := equipBag.GetNextEquipment(pos)
	if nextItemTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("inventory:强化升阶,已经满级")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentUpgradeMax)
		return
	}
	nextEquipTemplate := nextItemTemplate.GetEquipmentTemplate()

	//判断消耗条件
	items := nextEquipTemplate.GetNeedItemMap()
	if !manager.HasEnoughItems(items) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("inventory:强化升阶,物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	reasonText := commonlog.InventoryLogReasonEquipUpgrade.String()
	flag := manager.BatchRemove(items, commonlog.InventoryLogReasonEquipUpgrade, reasonText)
	if !flag {
		panic(fmt.Errorf("inventory:装备升阶移除材料应该成功"))
	}
	result := inventorytypes.EquipmentStrengthenResultTypeFailed
	//判断是否可以升阶
	success := mathutils.RandomHit(common.MAX_RATE, int(nextEquipTemplate.SuccessRate))
	if success {
		result = inventorytypes.EquipmentStrengthenResultTypeSuccess
		//升级
		flag = equipBag.Upgrade(pos)
		if !flag {
			panic(fmt.Errorf("inventory: 装备升阶移除材料应该成功"))
		}

		UpdateEquipmentProperty(pl)

	}

	//同步改变
	SnapInventoryChanged(pl)
	SnapInventoryEquipChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	//强化成功
	scInventoryEquipmentUpgrade := pbutil.BuildSCInventoryEquipmentUpgrade(pos, result)
	pl.SendMsg(scInventoryEquipmentUpgrade)
	return
}

//脱下
func TakeOff(pl player.Player, pos inventorytypes.BodyPositionType) (err error) {
	flag := takeOffInternal(pl, pos)
	if !flag {
		return
	}

	//同步改变
	UpdateEquipmentProperty(pl)
	SnapInventoryEquipChanged(pl)
	SnapInventoryChanged(pl)

	//脱下成功
	scInventoryTakeOffEquip := pbutil.BuildSCInventoryTakeOffEquip(pos)
	pl.SendMsg(scInventoryTakeOffEquip)

	return nil
}

func takeOffInternal(pl player.Player, pos inventorytypes.BodyPositionType) (flag bool) {
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	item := manager.GetEquipByPos(pos)
	//没有东西
	if item == nil || item.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos,
			}).Warn("inventory:脱下装备,装备不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipCanNotTakeOff)
		return
	}
	itemId := item.ItemId
	bindType := item.BindType
	defaultNum := int32(1)
	defaultLevel := int32(0)

	//背包空间不足
	if !manager.HasEnoughSlotItemLevel(itemId, defaultNum, defaultLevel, bindType) {
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	itemId = manager.TakeOff(pos)
	if itemId == 0 {
		panic(fmt.Errorf("inventory:take off should more than 0"))
	}

	//添加物品
	itemData := droptemplate.CreateItemData(itemId, defaultNum, defaultLevel, bindType)
	reasonText := commonlog.InventoryLogReasonTakeOff.String()
	flag = manager.AddItemLevel(itemData, commonlog.InventoryLogReasonTakeOff, reasonText)
	if !flag {
		panic(fmt.Errorf("inventory:add item should be success"))
	}
	return
}

//使用装备
func UseEquip(pl player.Player, index int32) (err error) {

	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	it := manager.FindItemByIndex(inventorytypes.BagTypePrim, index)

	//物品不存在
	if it == nil || it.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("inventory:使用装备,物品不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}
	itemId := it.ItemId
	bindType := it.BindType

	//判断物品是否可以装备
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if !itemTemplate.IsEquipment() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("inventory:使用装备,此物品不是装备")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotEquip)
		return
	}

	if itemTemplate.NeedProfession != 0 {
		//TODO 判断等级 转生 性别
		if itemTemplate.GetRole() != pl.GetRole() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"index":    index,
				}).Warn("inventory:使用装备,角色不符")
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
				}).Warn("inventory:使用装备,性别不符")
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
			}).Warn("inventory:使用装备,等级不够")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
		return
	}

	//判断转数
	if itemTemplate.NeedZhuanShu > propertyManager.GetZhuanSheng() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("inventory:使用装备,转数不够")
		playerlogic.SendSystemMessage(pl, lang.PlayerZhuanShengTooLow)
		return
	}

	equipmentSubType := itemTemplate.GetItemSubType().(itemtypes.ItemEquipmentSubType)
	pos := equipmentSubType.Position()

	//判断是否已经装备了
	equipmentItem := manager.GetEquipByPos(pos)
	if equipmentItem != nil && !equipmentItem.IsEmpty() {
		flag := takeOffInternal(pl, pos)
		if !flag {
			return
		}
	}
	reasonText := commonlog.InventoryLogReasonPutOn.String()
	//移除物品
	flag, _ := manager.RemoveIndex(inventorytypes.BagTypePrim, index, 1, commonlog.InventoryLogReasonPutOn, reasonText)
	if !flag {
		panic(fmt.Errorf("inventory:脱下物品应该是可以的"))
	}

	flag = manager.PutOn(pos, itemId, bindType)
	if !flag {
		//发送错误信息
		panic(fmt.Errorf("inventory:穿上位置 [%s]应该是可以的", pos.String()))
	}

	UpdateEquipmentProperty(pl)

	//同步改变
	SnapInventoryEquipChanged(pl)
	SnapInventoryChanged(pl)

	scInventoryUseEquip := pbutil.BuildSCInventoryUseEquip(index)
	scInventoryUseEquip.Index = &index
	pl.SendMsg(scInventoryUseEquip)
	return nil
}

//转成物品掉落
func ConverToGoldEquipItemData(itemId, num, level int32, bindType itemtypes.ItemBindType, propertyData inventorytypes.ItemPropertyData) *droptemplate.DropItemData {
	expireType := propertyData.GetExpireType()
	expireTime := propertyData.GetExpireTime()
	itemGetTime := propertyData.GetItemGetTime()
	goldequip, ok := propertyData.(*goldequiptypes.GoldEquipPropertyData)
	if !ok {
		return droptemplate.CreateItemDataWithExpire(itemId, num, level, bindType, expireType, expireTime, itemGetTime)
	} else {
		upstar := goldequip.UpstarLevel
		attrList := goldequip.AttrList
		isRandom := goldequip.IsHadCountAttr
		return droptemplate.CreateItemDataWithPropertyData(itemId, num, level, bindType, upstar, attrList, isRandom, expireType, expireTime, itemGetTime)
	}
}

//转成物品掉落
func ConverToItemData(itemId, num, level int32, bindType itemtypes.ItemBindType, propertyData inventorytypes.ItemPropertyData) *droptemplate.DropItemData {

	expireType := propertyData.GetExpireType()
	expireTime := propertyData.GetExpireTime()
	itemGetTime := propertyData.GetItemGetTime()
	goldequip, ok := propertyData.(*goldequiptypes.GoldEquipPropertyData)
	if !ok {
		return droptemplate.CreateItemDataWithExpire(itemId, num, level, bindType, expireType, expireTime, itemGetTime)
	} else {
		upstar := goldequip.UpstarLevel
		attrList := goldequip.AttrList
		isRandom := goldequip.IsHadCountAttr
		openLightLevel := goldequip.OpenLightLevel
		openTimes := goldequip.OpenTimes
		return droptemplate.CreateItemDataWithGoldPropertyData(itemId, num, level, bindType, upstar, attrList, openLightLevel, openTimes, isRandom, expireType, expireTime, itemGetTime)
	}
}
