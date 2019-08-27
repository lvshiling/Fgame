package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	additionsyslogic "fgame/fgame/game/additionsys/logic"
	"fgame/fgame/game/additionsys/pbutil"
	playeradditionsys "fgame/fgame/game/additionsys/player"
	additionsystypes "fgame/fgame/game/additionsys/types"
	droptemplate "fgame/fgame/game/drop/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_TAKE_OFF_ADDITION_SYS_EQUIP_TYPE), dispatch.HandlerFunc(handleTakeOffEquip))
}

//处理脱下装备
func handleTakeOffEquip(s session.Session, msg interface{}) (err error) {
	log.Debug("additionsysequip:处理脱下装备")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	cSTakeOffAdditionSysEquip := msg.(*uipb.CSTakeOffAdditionSysEquip)
	sysType := cSTakeOffAdditionSysEquip.GetSysType()
	slotId := cSTakeOffAdditionSysEquip.GetSlotId()
	sysTypeId := additionsystypes.AdditionSysType(sysType)
	slotPosition := additionsystypes.SlotPositionType(slotId)
	//参数不对
	if !sysTypeId.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"sysTypeId": sysTypeId.String(),
			}).Warn("additionsysequip:系统类型,错误")
		return
	}
	if !slotPosition.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"slotPosition": slotPosition.String(),
			}).Warn("additionsysequip:穿戴位置,错误")
		return
	}
	err = takeOff(tpl, sysTypeId, slotPosition)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("additionsysequip:处理脱下装备,错误")

		return err
	}
	log.Debug("additionsysequip:处理脱下装备,完成")
	return nil
}

//脱下
func takeOff(pl player.Player, typ additionsystypes.AdditionSysType, pos additionsystypes.SlotPositionType) (err error) {
	flag := takeOffInternal(pl, typ, pos)
	if !flag {
		return
	}
	//更新属性
	additionsyslogic.UpdataAdditionSysPropertyByType(pl, typ)

	//同步改变
	additionsyslogic.SnapInventoryAdditionSysEquipChangedByType(pl, typ)
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	//脱下成功
	scTakeOffEquip := pbutil.BuildSCTakeOffAdditionSysEquip(int32(typ), int32(pos))
	pl.SendMsg(scTakeOffEquip)

	return nil
}

func takeOffInternal(pl player.Player, typ additionsystypes.AdditionSysType, pos additionsystypes.SlotPositionType) (flag bool) {
	additionsysManager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	item := additionsysManager.GetAdditionSysByArg(typ, pos)
	level := int32(0)
	num := int32(1)
	bind := item.GetBindType()
	//没有东西
	if item == nil || item.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"pos":      pos.String(),
			}).Warn("additionsysequip:脱下装备,装备不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipCanNotTakeOff)
		return
	}

	//背包空间
	if !inventoryManager.HasEnoughSlotItemLevel(item.GetItemId(), num, level, bind) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"pos":      pos.String(),
			}).Warn("additionsysequip:脱下装备,背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	itemId := additionsysManager.TakeOff(typ, pos)
	if itemId == 0 {
		panic(fmt.Errorf("additionsysequip:take off should more than 0"))
	}

	//添加物品
	itemData := droptemplate.CreateItemData(itemId, num, level, bind)
	reasonText := commonlog.InventoryLogReasonTakeOff.String()
	flag = inventoryManager.AddItemLevel(itemData, commonlog.InventoryLogReasonTakeOff, reasonText)
	if !flag {
		panic(fmt.Errorf("additionsysequip:add item should be success"))
	}

	return
}
