package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	commonlogic "fgame/fgame/game/common/logic"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	"fgame/fgame/game/goldequip/pbutil"
	playergoldequip "fgame/fgame/game/goldequip/player"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_GODCASTING_EQUIP_UPLEVEL_TYPE), dispatch.HandlerFunc(handleGodCastingEquipUplevel))
}

func handleGodCastingEquipUplevel(s session.Session, msg interface{}) (err error) {
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSGodCastingEquipUplevel)
	csBodyPos := csMsg.GetBodyPos()
	bodyPos := inventorytypes.BodyPositionType(csBodyPos)
	if !bodyPos.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bodyPos":  csBodyPos,
			}).Warn("goldequip:处理神铸装备升级请求失败，装备部位不合法")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	err = godCastingEquipUplevel(tpl, bodyPos)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("goldequip:处理神铸装备升级请求,错误")
		return err
	}
	return
}

func godCastingEquipUplevel(pl player.Player, bodyPos inventorytypes.BodyPositionType) (err error) {
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
			}).Warn("goldequip:处理神铸装备升级请求,装备未装上")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipEquipmentSlotNoEquip)
		return
	}

	//判断是否为神铸装备
	itemTemp := item.GetItemService().GetItem(int(equip.GetItemId()))
	if itemTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      bodyPos.String(),
			}).Warn("goldequip:处理神铸装备升级请求,物品模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	goldequipTemp := itemTemp.GetGoldEquipTemplate()
	if goldequipTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      bodyPos.String(),
			}).Warn("goldequip:处理神铸装备升级请求,元神金装模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	godcastingTemp := goldequipTemp.GetGodCastingEquipTemp()
	if !goldequipTemp.IsGodCastingEquip() && godcastingTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      bodyPos.String(),
			}).Warn("goldequip:处理神铸装备升级请求,不是神铸装备且不是十转橙装")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//神铸模板
	if godcastingTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      bodyPos.String(),
			}).Warn("goldequip:处理神铸装备升级请求,神铸模板不存在(可能满级了)")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	//判断物品够不够
	useItemMap := godcastingTemp.GetUseItemMap()
	useItemNumStr := ""
	useItemNameStr := ""
	for useItemId, useItemCnt := range useItemMap {
		curUseItemNum := inventoryManager.NumOfItems(useItemId)
		if curUseItemNum < useItemCnt {
			log.WithFields(
				log.Fields{
					"playerId":      pl.GetId(),
					"pos":           bodyPos.String(),
					"useItemId":     useItemId,
					"useItemCnt":    useItemCnt,
					"curUseItemNum": curUseItemNum,
				}).Warn("goldequip:处理神铸装备升级请求,神铸装备升级物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		useItemNumStr += "[" + string(useItemCnt) + "]"
		temp := item.GetItemService().GetItem(int(useItemId))
		useItemNameStr += "<" + temp.Name + ">"
	}

	//判断下一级装备物品存不存在
	nextItemId := godcastingTemp.ItemId
	nextItemTemp := item.GetItemService().GetItem(int(nextItemId))
	if nextItemTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"pos":        bodyPos.String(),
				"nextItemId": nextItemId,
			}).Warn("goldequip:处理神铸装备升级请求,神铸进阶装备物品错误")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//对象升级处理
	propertyData := equip.GetPropertyData()
	goldEquipPropertyData, _ := propertyData.(*goldequiptypes.GoldEquipPropertyData)

	updateRate := godcastingTemp.UpdateWfb
	addTimes := int32(1)
	curTimesNum := goldEquipPropertyData.GodCastingTimes
	curTimesNum += addTimes

	_, sucess := commonlogic.AdvancedStatusAndProgress(curTimesNum, 0, godcastingTemp.TimesMin, godcastingTemp.TimesMax, 0, updateRate, 0)
	finalItemId := int32(itemTemp.Id)
	isSuccess := int32(0)
	if sucess {
		finalItemId = nextItemId
		isSuccess = int32(1)
	}
	// flag := equip.UplevelGodCasting(nextItemId)
	flag := goldequipBag.UplevelGodCasting(bodyPos, finalItemId, sucess)
	if !flag {
		panic("goldequip:升级更新神铸装备信息应该是可以的")
	}

	//同步物品（删掉吃掉的物品）
	useReason := commonlog.InventoryLogReasonGodCastingUpLevel
	useReasonText := fmt.Sprintf(useReason.String(), useItemNameStr, useItemNumStr, bodyPos.String(), itemTemp.Name, nextItemTemp.Name)
	flag = inventoryManager.BatchRemove(useItemMap, useReason, useReasonText)
	if !flag {
		panic("inventory:移除物品应该是可以的")
	}

	goldequiplogic.GoldEquipPropertyChanged(pl)
	goldequiplogic.SnapInventoryGoldEquipChanged(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCGodCastingEquipUplevel(bodyPos, goldEquipPropertyData.GodCastingTimes, finalItemId, isSuccess)
	pl.SendMsg(scMsg)

	return
}
