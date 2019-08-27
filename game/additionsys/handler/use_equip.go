package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/additionsys/additionsys"
	additionsyslogic "fgame/fgame/game/additionsys/logic"
	"fgame/fgame/game/additionsys/pbutil"
	playeradditionsys "fgame/fgame/game/additionsys/player"
	additionsystypes "fgame/fgame/game/additionsys/types"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_USE_ADDITION_SYS_EQUIP_TYPE), dispatch.HandlerFunc(handleUseEquip))
}

//使用装备
func handleUseEquip(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理使用装备")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csUseAdditionSysEquip := msg.(*uipb.CSUseAdditionSysEquip)
	sysType := csUseAdditionSysEquip.GetSysType()
	index := csUseAdditionSysEquip.GetIndex()
	sysTypeId := additionsystypes.AdditionSysType(sysType)
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
	if index < 0 {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("additionsysequip:物品背包索引,错误")
		return
	}
	err = useEquip(tpl, sysTypeId, index)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("inventory:处理使用装备,错误")

		return err
	}
	log.Debug("inventory:处理使用装备,完成")
	return nil
}

//使用装备
func useEquip(pl player.Player, typ additionsystypes.AdditionSysType, index int32) (err error) {
	if !additionsyslogic.GetAdditionSysEquipFuncOpenByType(pl, typ) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("inventory:附加系统装备,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	additionsysManager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	it := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, index)
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
	bind := it.BindType
	needItemType, ok := typ.ConvertToItemType()
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("inventory:使用装备,没有对应的装备类型")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//判断物品是否可以装备
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate.GetItemType() != needItemType {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"typ":          typ.String(),
				"itemId":       itemId,
				"needItemType": needItemType,
			}).Warn("inventory:使用装备,此物品不是附加系统装备")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotEquip)
		return
	}

	// if itemTemplate.NeedProfession != 0 {
	// 	//角色
	// 	if itemTemplate.GetRole() != pl.GetRole() {
	// 		log.WithFields(
	// 			log.Fields{
	// 				"playerId": pl.GetId(),
	// 				"index":    index,
	// 			}).Warn("inventory:使用装备,角色不符")
	// 		playerlogic.SendSystemMessage(pl, lang.PlayerRoleWrong)
	// 		return
	// 	}
	// }
	// if itemTemplate.GetSex() != 0 {
	// 	//性别
	// 	if itemTemplate.GetSex() != pl.GetSex() {
	// 		log.WithFields(
	// 			log.Fields{
	// 				"playerId": pl.GetId(),
	// 				"index":    index,
	// 			}).Warn("inventory:使用装备,性别不符")
	// 		playerlogic.SendSystemMessage(pl, lang.PlayerSexWrong)
	// 		return
	// 	}
	// }

	//判断阶数
	curJieShu := additionsys.GetSystemAdvancedNum(pl, typ)
	if curJieShu > 0 && curJieShu < itemTemplate.TypeFlag2 {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"typ":       typ.String(),
				"itemId":    itemId,
				"curJieShu": curJieShu,
			}).Warn("inventory:使用装备,阶数不够")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
		return
	}

	//判断转数
	// if itemTemplate.NeedZhuanShu > propertyManager.GetZhuanSheng() {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": pl.GetId(),
	// 			"index":    index,
	// 		}).Warn("inventory:使用装备,转数不够")
	// 	playerlogic.SendSystemMessage(pl, lang.PlayerZhuanShengTooLow)
	// 	return
	// }

	//判断是否已经装备
	pos := additionsystypes.SlotPositionType(itemTemplate.GetItemSubType().SubType())
	if !pos.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
				"index":    index,
				"itemId":   itemId,
				"pos":      pos,
			}).Warn("inventory:使用装备,没有合适的装备位置")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	equipmentItem := additionsysManager.GetAdditionSysByArg(typ, pos)
	if equipmentItem != nil && !equipmentItem.IsEmpty() {
		flag := takeOffInternal(pl, typ, pos)
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

	flag = additionsysManager.PutOn(typ, pos, itemId, bind)
	if !flag {
		panic(fmt.Errorf("inventory:穿上位置 [%s]应该是可以的", pos.String()))
	}

	//更新属性
	additionsyslogic.UpdataAdditionSysPropertyByType(pl, typ)

	//同步改变
	additionsyslogic.SnapInventoryAdditionSysEquipChangedByType(pl, typ)
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	scInventoryUseEquip := pbutil.BuildSCUseAdditionSysEquip(int32(typ), index)
	pl.SendMsg(scInventoryUseEquip)
	return
}
