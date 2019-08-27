package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/tulongequip/pbutil"
	tulongequiptemplate "fgame/fgame/game/tulongequip/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TULONG_EQUIP_ZHUANHUA_TYPE), dispatch.HandlerFunc(handleTuLongEquipZhuanHua))
}

//处理转化
func handleTuLongEquipZhuanHua(s session.Session, msg interface{}) (err error) {
	log.Debug("tulongequip:处理屠龙装备转化")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSTuLongEquipZhuanHua)
	itemIndexList := csMsg.GetItemIndex()
	slotId := csMsg.GetSlotId()

	posType := inventorytypes.BodyPositionType(slotId)
	if !posType.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = tuLongEquipZhuanHua(tpl, itemIndexList, posType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"ItemIndex": itemIndexList,
				"error":     err,
			}).Error("tulongequip:处理屠龙装备转化,错误")
		return err
	}
	log.Debug("tulongequip:处理屠龙装备转化,完成")
	return nil
}

//转化
func tuLongEquipZhuanHua(pl player.Player, itemIndexList []int32, posType inventorytypes.BodyPositionType) (err error) {
	if coreutils.IfRepeatElementInt32(itemIndexList) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"indexList": itemIndexList,
			}).Warn("inventory:处理转化失败,索引重复")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//材料数量
	needItemNum := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeTuLongEquipZhuanHuaNum))
	if len(itemIndexList) != needItemNum {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"ItemIndex": itemIndexList,
			}).Warn("inventory:转化失败,材料不足，无法进行转化")
		playerlogic.SendSystemMessage(pl, lang.TuLongEquipZhuanHuaItemNotEnough)
		return
	}

	initJieShu := int32(0)
	bindType := itemtypes.ItemBindTypeUnBind
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	for index, itemIndex := range itemIndexList {
		it := inventoryManager.FindItemByIndex(inventorytypes.BagTypeTuLongEquip, itemIndex)
		if it == nil || it.IsEmpty() {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"ItemIndex": itemIndexList,
				}).Warn("tulongequip:转化失败,材料不存在")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
			return
		}

		itemTemplate := item.GetItemService().GetItem(int(it.ItemId))
		if !itemTemplate.IsTuLongEquip() {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"ItemIndex": itemIndexList,
				}).Warn("tulongequip:转化失败,材料不是屠龙装备")
			playerlogic.SendSystemMessage(pl, lang.TuLongEquipReachStrengthenMax)
			return
		}

		// //部位相同
		// equipPos := itemTemplate.GetTuLongEquipTemplate().GetPosType()
		// if equipPos != posType {
		// 	log.WithFields(
		// 		log.Fields{
		// 			"playerId":  pl.GetId(),
		// 			"ItemIndex": itemIndexList,
		// 		}).Warn("tulongequip:转化失败,材料装备部位不一致")
		// 	playerlogic.SendSystemMessage(pl, lang.TuLongEquipZhuanHuaPosNotEqual)
		// 	return
		// }

		// 绑定属性
		if it.BindType == itemtypes.ItemBindTypeBind {
			bindType = itemtypes.ItemBindTypeBind
		}

		equipJieShu := itemTemplate.GetTuLongEquipTemplate().Number
		if index == 0 {
			initJieShu = equipJieShu
		}

		//材料阶数相同
		if index != 0 && initJieShu != equipJieShu {
			log.WithFields(
				log.Fields{
					"playerId":   pl.GetId(),
					"initJieShu": initJieShu,
					"curLevel":   equipJieShu,
				}).Warn("tulongequip:转化失败,屠龙装备阶数不同")
			playerlogic.SendSystemMessage(pl, lang.TuLongEquipZhuanHuaLevelNotEqual)
			return
		}
	}

	zhuanHuaLevel := initJieShu
	zhuanHuaTemp := tulongequiptemplate.GetTuLongEquipTemplateService().GetTuLongEquipZhuanHuaTemplat(posType, zhuanHuaLevel)
	if zhuanHuaTemp == nil {
		log.WithFields(
			log.Fields{
				"itemIndexList": itemIndexList,
				"zhuanHuaLevel": zhuanHuaLevel,
			}).Warn("tulongequip:转化失败,转化生成的物品模板不存在")
		playerlogic.SendSystemMessage(pl, lang.TuLongEquipZhuanHuaFailed)
		return
	}

	dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(zhuanHuaTemp.DropId)
	if dropData == nil {
		panic(fmt.Errorf("tulongequip: zhuanHua GetDropItem should be ok"))
	}
	itemId := dropData.GetItemId()
	num := dropData.GetNum()
	level := dropData.GetLevel()

	// 背包空间
	if !inventoryManager.HasEnoughSlotItemLevel(itemId, num, level, bindType) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"ItemIndex": itemIndexList,
			}).Warn("tulongequip:转化失败,背包不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//消耗材料
	if len(itemIndexList) != 0 {
		itemUseReason := commonlog.InventoryLogReasonTuLongEquipZhuanHuaUse
		flag, _ := inventoryManager.BatchRemoveIndex(inventorytypes.BagTypeTuLongEquip, itemIndexList, itemUseReason, itemUseReason.String())
		if !flag {
			panic(fmt.Errorf("tulongequip:装备转化移除材料应该成功"))
		}
	}

	//添加物品
	itemGetReason := commonlog.InventoryLogReasonTuLongEquipZhuanHuaGet
	flag := inventoryManager.AddItemLevel(dropData, itemGetReason, itemGetReason.String())
	if !flag {
		panic(fmt.Errorf("tulongequip:add item should be success"))
	}

	//同步改变
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	scMsg := pbutil.BuildSCTuLongEquipZhuanHua(itemId, int32(posType))
	pl.SendMsg(scMsg)
	return
}
