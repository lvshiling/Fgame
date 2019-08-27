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
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GODCASTING_FORGESOUL_UPLEVEL_TYPE), dispatch.HandlerFunc(handleGodCastingForgeSoulUplevel))
}

func handleGodCastingForgeSoulUplevel(s session.Session, msg interface{}) (err error) {
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSGodCastingForgeSoulUplevel)
	bodyPos := inventorytypes.BodyPositionType(csMsg.GetBodyPos())
	soulType := goldequiptypes.ForgeSoulType(csMsg.GetSoulType())
	if !bodyPos.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bodyPos":  bodyPos,
			}).Warn("goldequip:处理神铸锻魂升级请求失败，装备部位不合法")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	if !soulType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"soulType": soulType,
			}).Warn("goldequip:处理神铸锻魂升级请求失败，锻魂类型不合法")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	err = godCastingForgeSoulUplevel(tpl, bodyPos, soulType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("goldequip:处理神铸锻魂升级请求,错误")
		return err
	}
	return
}

func godCastingForgeSoulUplevel(pl player.Player, bodyPos inventorytypes.BodyPositionType, soulType goldequiptypes.ForgeSoulType) (err error) {
	goldequipManager := pl.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	forgeSoulTemplate := goldequiptemplate.GetGoldEquipTemplateService().GetForgeSoulTemplate(bodyPos, soulType)
	goldequipBag := goldequipManager.GetGoldEquipBag()

	//判断是否有装备
	equip := goldequipBag.GetByPosition(bodyPos)
	if equip == nil || equip.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      bodyPos.String(),
			}).Warn("goldequip:处理神铸锻魂升级请求,装备未装上")
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
			}).Warn("goldequip:处理神铸锻魂升级请求,物品模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	goldequipTemp := itemTemp.GetGoldEquipTemplate()
	if goldequipTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      bodyPos.String(),
			}).Warn("goldequip:处理神铸锻魂升级请求,元神金装模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	if !goldequipTemp.IsGodCastingEquip() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      bodyPos.String(),
			}).Warn("goldequip:处理神铸锻魂升级请求,不是神铸装备")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//判断锻魂升级模板存不存在
	soulInfo := equip.GetForgeSoulInfo(soulType)
	soulNextLevelTemp := forgeSoulTemplate.GetLevelTemplate(soulInfo.Level + 1)
	if soulNextLevelTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"pos":       bodyPos.String(),
				"soulType":  soulType,
				"soulLevel": soulInfo.Level,
			}).Warn("goldequip:处理神铸锻魂升级请求,锻魂升级模板不存在（可能最大等级上限）")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//锻魂和神铸等级挂钩，判断锻魂等级是否为当前阶层最大级
	maxLevel := goldequipTemp.GetGodCastingForgeSoulMaxLevel()
	if soulNextLevelTemp.IsMaxLevel(maxLevel) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"pos":       bodyPos.String(),
				"soulType":  soulType,
				"soulLevel": soulInfo.Level,
			}).Warn("goldequip:处理神铸锻魂升级请求,锻魂已经满级")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipGodCastingForgeSoulLevelFull)
		return
	}

	//判断物品够不够
	useItemId := forgeSoulTemplate.UseItemId
	useItemCnt := soulNextLevelTemp.UseItemCount
	curUseItemNum := inventoryManager.NumOfItems(useItemId)
	if curUseItemNum < useItemCnt {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"pos":        bodyPos.String(),
				"soulType":   soulType,
				"useItemId":  useItemId,
				"useItemCnt": useItemCnt,
			}).Warn("goldequip:处理神铸锻魂升级请求,锻魂升级物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//部位对象升级处理
	updateRate := soulNextLevelTemp.UpdateWfb
	addTimes := int32(1)
	curTimesNum := soulInfo.Times
	curTimesNum += addTimes

	_, sucess := commonlogic.AdvancedStatusAndProgress(curTimesNum, 0, soulNextLevelTemp.TimesMin, soulNextLevelTemp.TimesMax, 0, updateRate, 0)
	isSuccess := int32(0)
	if sucess {
		isSuccess = int32(1)
	}
	// equip.UplevelSoul(soulType)
	goldequipManager.UplevelSoul(bodyPos, soulType, sucess)

	//同步物品（删掉吃掉的物品）
	useReason := commonlog.InventoryLogReasonForgeSoulUplevel
	useItemTemp := item.GetItemService().GetItem(int(useItemId))
	useReasonText := fmt.Sprintf(useReason.String(), useItemTemp.Name, useItemCnt, bodyPos.String(), soulType.String())
	flag := inventoryManager.UseItem(useItemId, useItemCnt, useReason, useReasonText)
	if !flag {
		panic("inventory:移除物品应该是可以的")
	}

	goldequiplogic.GoldEquipPropertyChanged(pl)
	goldequiplogic.SnapInventoryGoldEquipChanged(pl)
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCGodCastingForgeSoulUpLevel(bodyPos, soulType, soulInfo, isSuccess)
	pl.SendMsg(scMsg)
	teShuSkillList := goldequipManager.GetTeShuSkillList()
	pl.ResetTeShuSkills(teShuSkillList)
	return
}
