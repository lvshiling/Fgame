package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	additionsyseventtypes "fgame/fgame/game/additionsys/event/types"
	additionsyslogic "fgame/fgame/game/additionsys/logic"
	"fgame/fgame/game/additionsys/pbutil"
	playeradditionsys "fgame/fgame/game/additionsys/player"
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fgame/fgame/game/common/common"
	gamevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/pkg/mathutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ADDITION_SYS_STRENGTHEN_BODY_TYPE), dispatch.HandlerFunc(handleEquipStrengthenBody))
}

//处理强化
func handleEquipStrengthenBody(s session.Session, msg interface{}) (err error) {
	log.Debug("additionsysequip:处理装备槽强化")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csEquipStrengthenBody := msg.(*uipb.CSAdditionSysStrengthenBody)
	targetSysType := csEquipStrengthenBody.GetSysType()
	targetSlotId := csEquipStrengthenBody.GetSlotId()
	auto := csEquipStrengthenBody.GetAuto()
	isProtect := csEquipStrengthenBody.GetIsProtect()
	sysType := additionsystypes.AdditionSysType(targetSysType)
	slotId := additionsystypes.SlotPositionType(targetSlotId)

	//参数不对
	if !sysType.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"sysType":  sysType.String(),
			}).Warn("additionsyse:系统类型,错误")
		return
	}
	//参数不对
	if !slotId.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"slotId":   slotId.String(),
			}).Warn("additionsysequip:强化位置,错误")
		return
	}

	err = equipStrengthen(tpl, sysType, slotId, auto, isProtect)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"targetSysType": targetSysType,
				"targetSlotId":  targetSlotId,
				"error":         err,
			}).Error("additionsys:处理装备槽装备强化,错误")

		return err
	}
	log.Debug("additionsys:处理装备槽装备强化,完成")
	return nil
}

//强化
func equipStrengthen(pl player.Player, typ additionsystypes.AdditionSysType, pos additionsystypes.SlotPositionType, auto bool, isProtect bool) (err error) {
	var respond bool
	var result inventorytypes.EquipmentStrengthenResultType
	result, respond = equipmentSlotStrengthenUpgrade(pl, typ, pos, isProtect, false)
	//是否回应
	if !respond {
		return
	}
	//同步改变
	additionsyslogic.SnapInventoryAdditionSysEquipChangedByType(pl, typ)
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	//强化成功
	scEquipmentSlotStrength := pbutil.BuildSCAdditionSysStrengthenBody(int32(typ), int32(pos), int32(result), auto, isProtect)
	pl.SendMsg(scEquipmentSlotStrength)

	return
}

//强化升级
func equipmentSlotStrengthenUpgrade(pl player.Player, typ additionsystypes.AdditionSysType, pos additionsystypes.SlotPositionType, isProtect bool, ignoreErrorMsg bool) (result inventorytypes.EquipmentStrengthenResultType, flag bool) {
	manager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	equipBag := manager.GetAdditionSysEquipBagByType(typ)
	item := equipBag.GetByPosition(pos)
	//物品不存在
	if item == nil || item.IsEmpty() {
		if ignoreErrorMsg {
			return
		}
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"pos":      pos.String(),
			}).Warn("inventory:强化升级失败,装备不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotNoEquip)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//判断槽位是否可以升
	nextEquipmentStrengthenTemplate := equipBag.GetNextStrengthenTemplate(pos)
	if nextEquipmentStrengthenTemplate == nil {
		if ignoreErrorMsg {
			return
		}
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"pos":      pos.String(),
			}).Warn("inventory:强化升级失败,已经满级")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotStarMax)
		return
	}
	// maxLevel := int32(math.Floor(float64(pl.GetLevel()) / float64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeEquipmentStrengthenLevelLimit))))
	// //判断等级限制
	// if nextEquipmentStrengthenTemplate.Level > maxLevel {
	// 	if ignoreErrorMsg {
	// 		return
	// 	}
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 			"typ":      typ.String(),
	// 			"pos":      pos.String(),
	// 		}).Warn("inventory:强化升级失败,达到极限")
	// 	playerlogic.SendSystemMessage(pl, lang.InventoryEquipmentSlotLevelExceedLevel)
	// 	return
	// }

	items := nextEquipmentStrengthenTemplate.GetNeedItemMap()
	//是否使用防爆符
	if isProtect && nextEquipmentStrengthenTemplate.ProtectItemId != 0 {
		items[nextEquipmentStrengthenTemplate.ProtectItemId] = nextEquipmentStrengthenTemplate.ProtectItemCount
	}

	if len(items) != 0 {
		if !inventoryManager.HasEnoughItems(items) {
			if ignoreErrorMsg {
				return
			}
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"typ":      typ.String(),
					"pos":      pos.String(),
				}).Warn("inventory:强化升级失败,物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
	}

	if len(items) != 0 {
		reasonText := commonlog.InventoryLogReasonEquipSlotStrengthUpgrade.String()
		flag := inventoryManager.BatchRemove(items, commonlog.InventoryLogReasonEquipSlotStrengthUpgrade, reasonText)
		if !flag {
			panic(fmt.Errorf("inventory:装备槽强化升级移除材料应该成功"))
		}
	}

	useItemEventData := additionsyseventtypes.CreatePlayerAdditionSysUseItemEventData(items)
	gamevent.Emit(additionsyseventtypes.EventTypeAdditionSysUseItem, pl, useItemEventData)

	//是否消耗了保级物品
	useProtect := false
	if isProtect && nextEquipmentStrengthenTemplate.ProtectItemId != 0 {
		useProtect = true
	}

	beflev := item.Level //强化前等级
	success := mathutils.RandomHit(common.MAX_RATE, int(nextEquipmentStrengthenTemplate.SuccessRate))
	if success {
		//升星
		flag = equipBag.StrengthLevel(pos)
		if !flag {
			panic(fmt.Errorf("inventory: 强化升级应该成功"))
		}
		//日志
		additionsysReason := commonlog.AdditionSysLogReasonStrengthenUpgrade
		reasonText := fmt.Sprintf(additionsysReason.String(), typ.String(), pos.String())
		data := additionsyseventtypes.CreatePlayerAdditionSysStrengthenLevLogEventData(typ, pos, beflev, additionsysReason, reasonText)
		gamevent.Emit(additionsyseventtypes.EventTypeAdditionSysStrengthenLevLog, pl, data)
		//更新属性
		additionsyslogic.UpdataAdditionSysPropertyByType(pl, typ)
		result = inventorytypes.EquipmentStrengthenResultTypeSuccess
		return
	}

	//判断是否会回退
	if !useProtect && nextEquipmentStrengthenTemplate.GetFailTemplate() != nil {
		fail := mathutils.RandomHit(common.MAX_RATE, int(nextEquipmentStrengthenTemplate.FailBackRate))
		if fail {
			//回退
			flag = equipBag.StrengthLevelBack(pos)
			if !flag {
				panic(fmt.Errorf("inventory: strength star back should be ok"))
			}
			//日志
			additionsysReason := commonlog.AdditionSysLogReasonStrengthenBackLev
			reasonText := fmt.Sprintf(additionsysReason.String(), typ.String(), pos.String())
			data := additionsyseventtypes.CreatePlayerAdditionSysStrengthenLevLogEventData(typ, pos, beflev, additionsysReason, reasonText)
			gamevent.Emit(additionsyseventtypes.EventTypeAdditionSysStrengthenLevLog, pl, data)
			//更新属性
			additionsyslogic.UpdataAdditionSysPropertyByType(pl, typ)
			result = inventorytypes.EquipmentStrengthenResultTypeBack
			return
		}
	}

	flag = true
	result = inventorytypes.EquipmentStrengthenResultTypeFailed
	return

}
