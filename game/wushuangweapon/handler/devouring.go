package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	constantservice "fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	itemservice "fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	wushuangweaponeventtypes "fgame/fgame/game/wushuangweapon/event/types"
	"fgame/fgame/game/wushuangweapon/pbutil"
	playerwushuangweapon "fgame/fgame/game/wushuangweapon/player"
	wushuangweapontypes "fgame/fgame/game/wushuangweapon/types"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WUSHUANGWEAPON_DEVOURING_TYPE), dispatch.HandlerFunc(handlerDevouring))
}

func handlerDevouring(s session.Session, msg interface{}) (err error) {
	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSWushuangWeaponDevouring)
	csBodyPos := csMsg.GetBodyPos()

	bodyPos := wushuangweapontypes.WushuangWeaponPart(csBodyPos)
	if !bodyPos.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bodyPos":  bodyPos,
			}).Warn("WushuangWeapon:部位吞噬请求，类型错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = devouring(tpl, bodyPos)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("WushuangWeapon:处理部位吞噬请求，错误")
		return
	}
	return
}

func devouring(pl player.Player, bodyPos wushuangweapontypes.WushuangWeaponPart) (err error) {
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	wushuangDataManager := pl.GetPlayerDataManager(playertypes.PlayerWushuangWeaponDataManagerType).(*playerwushuangweapon.PlayerWushuangWeaponDataManager)
	itemServer := itemservice.GetItemService()
	slotObj := wushuangDataManager.GetSlotObjectFromBodyPos(bodyPos)
	curlevel := slotObj.GetLevel()
	//判断身上是否有装备
	if !slotObj.IsEquip() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bodyPos":  bodyPos,
			}).Warn("wushuangWeapon:该部位没有装备，不能吞噬")
		playerlogic.SendSystemMessage(pl, lang.WushuangWeaponBodyPosEquipmentNotExist)
		return
	}
	itemTemp := itemServer.GetItem(int(slotObj.GetItemId()))
	if itemTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   slotObj.GetItemId(),
			}).Warn("wushuangWeapon:部位装备模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	baseTemplate := itemTemp.GetWushuangBaseTemplate()
	if baseTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bodyPos":  bodyPos,
				"curlevel": curlevel,
			}).Warn("wushuangWeapon:base模板不存在！")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	wushuangStrengthenTemplate := baseTemplate.GetStrengthTemplateByLevel(curlevel)
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
	//判断是否满级了
	if wushuangStrengthenTemplate.IsMaxLevel() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"level":    slotObj.GetLevel(),
			}).Warn("wushuangWeapon:满级了，不能吞噬")
		playerlogic.SendSystemMessage(pl, lang.WushuangWeaponLevelFull)
		return
	}

	//判断经验值是否满了
	curExperience := slotObj.GetExperience()
	nextTemplate := wushuangStrengthenTemplate.GetNextStrengthenTemplate()
	if nextTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bodyPos":  bodyPos,
				"curlevel": curlevel,
			}).Warn("wushuangWeapon:下一级strenth模板不存在！")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	maxExperience := nextTemplate.GetAllNeedExperience()
	if curExperience >= maxExperience {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"curExperience": curExperience,
			}).Warn("wushuangWeapon:经验值已满，不能吞噬")
		playerlogic.SendSystemMessage(pl, lang.WushuangWeaponExperienceFull)
		return
	}

	//计算吃掉物品后的经验值（判断是否超过最大值）
	itemId := constantservice.GetConstantService().GetConstant(constanttypes.ConstantTypeWushuangEssence)
	devouringItemTemplate := itemServer.GetItem(int(itemId))
	if devouringItemTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
			}).Warn("wushuangWeapon:物品模板不存在错误")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	curItemNum := inventoryManager.NumOfItems(itemId)
	if curItemNum <= 0 {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"curItemNum": curItemNum,
				"itemId":     devouringItemTemplate.Id,
				"itemName":   devouringItemTemplate.Name,
			}).Warn("wushuangWeapon:物品数量不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	itemEx := int64(devouringItemTemplate.TypeFlag1)
	needNum := int32(math.Ceil(float64(maxExperience-slotObj.GetExperience()) / float64(itemEx)))
	useItemNum := int32(0)
	addExperience := int64(0)
	if curItemNum < needNum {
		addExperience = int64(curItemNum) * itemEx
		useItemNum = curItemNum
	} else {
		addExperience = int64(needNum) * itemEx
		useItemNum = needNum
	}
	slotObj.AddExperience(addExperience)

	//同步物品（删掉吃掉的物品）
	useReason := commonlog.InventoryLogReasonWushuangWeaponEat
	useReasonText := fmt.Sprintf(useReason.String(), devouringItemTemplate.Name, useItemNum)
	flag := inventoryManager.UseItem(itemId, useItemNum, useReason, useReasonText)
	if !flag {
		panic("inventory:移除物品应该是可以的")
	}

	inventorylogic.SnapInventoryChanged(pl)

	eventdata := wushuangweaponeventtypes.CreatePlayerWushuangDevouringEventData(itemId, useItemNum)
	gameevent.Emit(wushuangweaponeventtypes.EventTypeWushuangDevouring, pl, eventdata)

	//发消息
	scMsg := pbutil.BuildSCWushuangWeaponDevouring(bodyPos, slotObj.GetLevel(), slotObj.GetExperience())
	pl.SendMsg(scMsg)
	return
}
