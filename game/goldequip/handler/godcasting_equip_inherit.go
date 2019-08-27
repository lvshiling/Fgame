package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	"fgame/fgame/game/goldequip/pbutil"
	playergoldequip "fgame/fgame/game/goldequip/player"
	goldequiptemplate "fgame/fgame/game/goldequip/template"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GODCASTING_EQUIP_INHERIT_TYPE), dispatch.HandlerFunc(handleGodCastingEquipInherit))
}

func handleGodCastingEquipInherit(s session.Session, msg interface{}) (err error) {
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSGodCastingEquipInherit)
	csBodyPos := csMsg.GetBodyPos()
	bodyPos := inventorytypes.BodyPositionType(csBodyPos)
	index := csMsg.GetEquipIndex()
	if !bodyPos.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bodyPos":  csBodyPos,
			}).Warn("goldequip:处理神铸装备继承请求失败，装备部位不合法")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	err = godCastingEquipInherit(tpl, bodyPos, index)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("goldequip:处理神铸装备继承请求,错误")
		return err
	}
	return
}

func godCastingEquipInherit(pl player.Player, bodyPos inventorytypes.BodyPositionType, index int32) (err error) {
	goldequipManager := pl.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	goldequipBag := goldequipManager.GetGoldEquipBag()

	//判断是否有装备
	equip := goldequipBag.GetByPosition(bodyPos)
	if equip == nil || equip.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      bodyPos.String(),
			}).Warn("goldequip:处理神铸装备继承请求,装备未装上")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipEquipmentSlotNoEquip)
		return
	}

	//判断身上穿的是否为神铸装备
	itemTemp := item.GetItemService().GetItem(int(equip.GetItemId()))
	if itemTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      bodyPos.String(),
			}).Warn("goldequip:处理神铸装备继承请求,物品模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	goldequipTemp := itemTemp.GetGoldEquipTemplate()
	if goldequipTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      bodyPos.String(),
			}).Warn("goldequip:处理神铸装备继承请求,元神金装模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	godcastingTemp := goldequipTemp.GetGodCastingEquipTemp()
	if !goldequipTemp.IsGodCastingEquip() && godcastingTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      bodyPos.String(),
			}).Warn("goldequip:处理神铸装备继承请求,不是神铸装备或者十转橙装")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//判断传入的物品是否为神铸装备
	changeItem := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, index)
	if changeItem == nil || changeItem.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("goldequip:物品错误，继承物品不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}
	changeItemTemp := item.GetItemService().GetItem(int(changeItem.ItemId))
	if changeItemTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      bodyPos.String(),
				"itemId":   changeItem.ItemId,
			}).Warn("goldequip:处理神铸装备继承请求,继承物品模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	changeGoldequipTemp := changeItemTemp.GetGoldEquipTemplate()
	if changeGoldequipTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      bodyPos.String(),
				"itemId":   changeItem.ItemId,
			}).Warn("goldequip:处理神铸装备继承请求,继承元神金装模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	changeGodcastingTemp := changeGoldequipTemp.GetGodCastingEquipTemp()
	if !changeGoldequipTemp.IsGodCastingEquip() && changeGodcastingTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      bodyPos.String(),
				"itemId":   changeItem.ItemId,
			}).Warn("goldequip:处理神铸装备继承请求,继承不是神铸装备或者十转橙装")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	targetLevel := changeGoldequipTemp.GetGodCastingEquipLevel()
	if targetLevel <= goldequipTemp.GetGodCastingEquipLevel() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"pos":         bodyPos.String(),
				"targetLevel": targetLevel,
				"itemId":      changeItem.ItemId,
			}).Warn("goldequip:处理神铸装备继承请求,继承装备神铸等级小于等于身上装备神铸等级")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipGodCastingInheritLevelTooLow)
		return
	}

	jichengTemp := goldequiptemplate.GetGoldEquipTemplateService().GetGodcastingJiChengTemplate(targetLevel)
	if jichengTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"pos":         bodyPos.String(),
				"targetLevel": targetLevel,
			}).Warn("goldequip:处理神铸装备继承请求,继承模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//判断物品够不够
	useItemId := jichengTemp.NeedItemId
	useItemCnt := jichengTemp.NeedItemCount
	curUseItemNum := inventoryManager.NumOfItems(useItemId)
	if curUseItemNum < useItemCnt {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"pos":        bodyPos.String(),
				"useItemId":  useItemId,
				"useItemCnt": useItemCnt,
			}).Warn("goldequip:处理神铸装备继承请求,继承物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//部位物品对象继承（转数对应神铸装备和propertyData）
	finalItemId, flag := goldequiptemplate.GetGoldEquipTemplateService().GetGodCastingInheritTargetLevelItemId(goldequipTemp, targetLevel)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"pos":         bodyPos.String(),
				"targetLevel": targetLevel,
			}).Warn("goldequip:处理神铸装备继承请求,继承阶数神铸装备模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	// propertyData := changeItem.PropertyData
	// flag = equip.GodCastingInherit(finalItemId, propertyData)
	goldequipBag.GodCastingInherit(bodyPos, finalItemId)
	indexList := []int32{index}
	data := changeItem.PropertyData.(*goldequiptypes.GoldEquipPropertyData)
	goldequiplogic.GodCastingInheritSendEmail(pl, indexList, changeItem.Level, data.OpenLightLevel)

	//同步物品（删掉吃掉的物品）
	useReason := commonlog.InventoryLogReasonGodCastingInherit
	useItemTemp := item.GetItemService().GetItem(int(useItemId))
	useReasonText := fmt.Sprintf(useReason.String(), useItemTemp.Name, useItemCnt, bodyPos.String(), itemTemp.Name, changeItemTemp.Name)
	flag = inventoryManager.UseItem(useItemId, useItemCnt, useReason, useReasonText)
	if !flag {
		panic("inventory:移除物品应该是可以的")
	}

	// changeItemCnt := int32(1)
	// removeReasonText := fmt.Sprintf(useReason.String(), changeItemTemp.Name, changeItemCnt, bodyPos.String(), itemTemp.Name, changeItemTemp.Name)
	// flag, _ = inventoryManager.RemoveIndex(inventorytypes.BagTypePrim, index, changeItemCnt, useReason, removeReasonText)
	// if !flag {
	// 	panic("inventory:移除物品应该是可以的")
	// }

	goldequiplogic.GoldEquipPropertyChanged(pl)
	goldequiplogic.SnapInventoryGoldEquipChanged(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCGodCastingEquipInherit()
	pl.SendMsg(scMsg)
	return
}
