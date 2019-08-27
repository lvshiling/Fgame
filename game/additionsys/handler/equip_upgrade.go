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
	"fgame/fgame/game/item/item"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_ADDITION_SYS_UPGRADE_TYPE), dispatch.HandlerFunc(handleAdditionSysUpgrade))
}

//处理升阶
func handleAdditionSysUpgrade(s session.Session, msg interface{}) (err error) {
	log.Debug("additionsys:处理系统装备升阶")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csAdditionSysUpgrade := msg.(*uipb.CSAdditionSysUpgrade)
	slotId := csAdditionSysUpgrade.GetSlotId()
	slotPosition := additionsystypes.SlotPositionType(slotId)

	//参数不对
	if !slotPosition.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      slotPosition.String(),
			}).Warn("additionsys:强化升阶,参数不对")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	sysType := additionsystypes.AdditionSysType(csAdditionSysUpgrade.GetSysType())
	//参数不对
	if !sysType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"sysType":  sysType.String(),
			}).Warn("additionsys:升级系统类型,错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = additionSysUpgrade(tpl, sysType, slotPosition)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("additionsys:处理装备升阶,错误")

		return err
	}
	log.Debug("additionsys:处理装备升阶,完成")
	return nil
}

//升阶
func additionSysUpgrade(pl player.Player, sysType additionsystypes.AdditionSysType, pos additionsystypes.SlotPositionType) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	slotItem := manager.GetAdditionSysByArg(sysType, pos)
	equipBag := manager.GetAdditionSysEquipBagByType(sysType)
	//物品不存在
	if slotItem == nil || slotItem.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("additionsys:装备升阶,装备不存在")
		playerlogic.SendSystemMessage(pl, lang.AdditionSysSlotNoEquip)
		return
	}

	//物品不存在
	itemTemplate := item.GetItemService().GetItem(int(slotItem.ItemId))
	if itemTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("additionsys:装备升阶,装备不存在")
		return nil
	}
	systemEquipTemplate := itemTemplate.GetSystemEquipTemplate()
	if systemEquipTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("additionsys:装备升阶,装备不存在")
		return nil
	}
	nextItemTemplate := systemEquipTemplate.GetNextItemTemplate()
	if nextItemTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("additionsys:装备升阶,已经满级")
		playerlogic.SendSystemMessage(pl, lang.AdditionSysEquipmentUpgradeMax)
		return nil
	}

	nextSysEquipTemplate := nextItemTemplate.GetSystemEquipTemplate()
	if nextSysEquipTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("additionsys:装备升阶,下一阶为空")
		return nil
	}
	//判断消耗条件
	items := nextSysEquipTemplate.GetNeedItemMap()
	if !inventoryManager.HasEnoughItems(items) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos.String(),
			}).Warn("additionsys:装备升阶,物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	reasonText := commonlog.InventoryLogReasonEquipUpgrade.String()
	flag := inventoryManager.BatchRemove(items, commonlog.InventoryLogReasonEquipUpgrade, reasonText)
	if !flag {
		panic(fmt.Errorf("additionsys:装备升阶移除材料应该成功"))
	}
	result := inventorytypes.EquipmentStrengthenResultTypeFailed
	//判断是否可以升阶
	success := mathutils.RandomHit(common.MAX_RATE, int(nextSysEquipTemplate.GetSuccessRate()))
	if success {
		result = inventorytypes.EquipmentStrengthenResultTypeSuccess
		//升级
		flag = equipBag.Upgrade(pos)
		if !flag {
			panic(fmt.Errorf("additionsys: 装备升阶移除材料应该成功"))
		}
		//更新属性
		additionsyslogic.UpdataAdditionSysPropertyByType(pl, sysType)
		eventData := additionsyseventtypes.CreatePlayerAdditionUpgradeEventData(sysType, pos)
		gamevent.Emit(additionsyseventtypes.EventTypeAdditionSysUpgrade, pl, eventData)
	}

	//同步改变
	additionsyslogic.SnapInventoryAdditionSysEquipChangedByType(pl, sysType)
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	//强化成功
	scAdditionSysUpgrade := pbutil.BuildSCAdditionSysUpgrade(sysType, pos, int32(result))
	pl.SendMsg(scAdditionSysUpgrade)

	return
}
