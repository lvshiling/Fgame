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
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/tulongequip/pbutil"
	tulongequiptemplate "fgame/fgame/game/tulongequip/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TULONG_EQUIP_RONGHE_TYPE), dispatch.HandlerFunc(handleTuLongEquipRongHe))
}

//处理融合
func handleTuLongEquipRongHe(s session.Session, msg interface{}) (err error) {
	log.Debug("tulongequip:处理屠龙装备融合")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSTuLongEquipRongHe)
	itemIndexList := csMsg.GetItemIndex()
	args := csMsg.GetArgs()

	err = tuLongEquipRongHe(tpl, itemIndexList, args)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"ItemIndex": itemIndexList,
				"error":     err,
			}).Error("tulongequip:处理屠龙装备融合,错误")
		return err
	}
	log.Debug("tulongequip:处理屠龙装备融合,完成")
	return nil
}

//融合
func tuLongEquipRongHe(pl player.Player, itemIndexList []int32, args int32) (err error) {
	if coreutils.IfRepeatElementInt32(itemIndexList) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"indexList": itemIndexList,
			}).Warn("inventory:处理融合失败,索引重复")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//材料数量
	needItemNum := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeTuLongEquipRongHeNum))
	if len(itemIndexList)%needItemNum != 0 {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"ItemIndex": itemIndexList,
			}).Warn("inventory:融合失败,材料不足，无法进行融合")
		playerlogic.SendSystemMessage(pl, lang.TuLongEquipRongHeItemNotEnough)
		return
	}

	rongHeTimes := len(itemIndexList) / needItemNum
	initJieShu := int32(0)
	needQuality := int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeTuLongEquipRongHeQuality))
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	for index, itemIndex := range itemIndexList {
		it := inventoryManager.FindItemByIndex(inventorytypes.BagTypeTuLongEquip, itemIndex)
		if it == nil || it.IsEmpty() {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"ItemIndex": itemIndexList,
				}).Warn("tulongequip:融合失败,材料不存在")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
			return
		}

		itemTemplate := item.GetItemService().GetItem(int(it.ItemId))
		if !itemTemplate.IsTuLongEquip() {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"ItemIndex": itemIndexList,
				}).Warn("tulongequip:融合失败,材料不是屠龙装备")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotUse)
			return
		}
		tulongEquipTemp := itemTemplate.GetTuLongEquipTemplate()
		if tulongEquipTemp == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"ItemId":   it.ItemId,
				}).Warn("tulongequip:融合失败,屠龙装备模板不存在")
			playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
			return
		}

		//品质等于橙色
		if itemTemplate.Quality < needQuality {
			log.WithFields(
				log.Fields{
					"playerId":    pl.GetId(),
					"needQuality": needQuality,
					"quality":     itemTemplate.Quality,
				}).Warn("tulongequip:融合失败,非橙色装备无法进行融合")
			playerlogic.SendSystemMessage(pl, lang.TuLongEquipRongHeQualityNotEnough)
			return
		}

		if index == 0 {
			initJieShu = tulongEquipTemp.Number
		}

		//材料阶数相同
		if index != 0 && initJieShu != tulongEquipTemp.Number {
			log.WithFields(
				log.Fields{
					"playerId":    pl.GetId(),
					"needQuality": needQuality,
					"quality":     itemTemplate.Quality,
				}).Warn("tulongequip:融合失败,屠龙装备阶数不同")
			playerlogic.SendSystemMessage(pl, lang.TuLongEquipRongHeLevelNotEqual)
			return
		}
	}

	rongHeJieShu := initJieShu
	rongheTemp := tulongequiptemplate.GetTuLongEquipTemplateService().GetTuLongEquipRongHeTemplat(rongHeJieShu)
	if rongheTemp == nil {
		log.WithFields(
			log.Fields{
				"rongHeJieShu": rongHeJieShu,
			}).Warn("tulongequip:融合失败,融合生成的物品模板不存在")
		playerlogic.SendSystemMessage(pl, lang.TuLongEquipRongHeFailed)
		return
	}

	var dropItemDataList []*droptemplate.DropItemData
	for i := rongHeTimes; i > 0; i-- {
		dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(rongheTemp.DropId)
		if dropData == nil {
			panic(fmt.Errorf("tulongequip: ronghe GetDropItem should be ok"))
		}
		dropItemDataList = append(dropItemDataList, dropData)
	}

	// 特殊处理BUG7334
	// // 背包空间
	// if !inventoryManager.HasEnoughSlotsOfItemLevel(dropItemDataList) {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId":         pl.GetId(),
	// 			"dropItemDataList": dropItemDataList,
	// 		}).Warn("tulongequip:融合失败,背包不足")
	// 	playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
	// 	return
	// }

	//消耗材料
	if len(itemIndexList) != 0 {
		itemUseReason := commonlog.InventoryLogReasonTuLongEquipRongHeUse
		flag, _ := inventoryManager.BatchRemoveIndex(inventorytypes.BagTypeTuLongEquip, itemIndexList, itemUseReason, itemUseReason.String())
		if !flag {
			panic(fmt.Errorf("tulongequip:装备融合移除材料应该成功"))
		}
	}

	//添加物品
	itemGetReason := commonlog.InventoryLogReasonTuLongEquipRongHeGet
	flag := inventoryManager.BatchAddOfItemLevel(dropItemDataList, itemGetReason, itemGetReason.String())
	if !flag {
		panic(fmt.Errorf("tulongequip:add item should be success"))
	}

	//同步改变
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCTuLongEquipRongHe(dropItemDataList, args)
	pl.SendMsg(scMsg)
	return
}
