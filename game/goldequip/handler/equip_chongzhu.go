package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	funcopentypes "fgame/fgame/game/funcopen/types"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	"fgame/fgame/game/goldequip/pbutil"
	goldequiptemplate "fgame/fgame/game/goldequip/template"
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
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_CHONGZHU_TYPE), dispatch.HandlerFunc(handleGoldEquipChongzhu))
}

//处理重铸
func handleGoldEquipChongzhu(s session.Session, msg interface{}) (err error) {
	log.Debug("goldequip:处理元神金装重铸")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csGoldEquipChongzhu := msg.(*uipb.CSGoldEquipChongzhu)
	itemIndexList := csGoldEquipChongzhu.GetItemIndex()

	err = goldEquipChongzhu(tpl, itemIndexList)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"ItemIndex": itemIndexList,
				"error":     err,
			}).Error("goldequip:处理元神金装重铸,错误")
		return err
	}
	log.Debug("goldequip:处理元神金装重铸,完成")
	return nil
}

//重铸
func goldEquipChongzhu(pl player.Player, itemIndexList []int32) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeYuanGodRecasting) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"ItemIndex": itemIndexList,
			}).Warn("inventory:重铸失败,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//材料数量
	needItemNum := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeGoldEquipChongzhuItemUseNeed))
	if len(itemIndexList) != needItemNum {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"ItemIndex": itemIndexList,
			}).Warn("inventory:重铸失败,材料不足，无法进行重铸")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipChongzhuItemNotEnough)
		return
	}

	//材料的最低质量和转生数
	badQuality := int32(0)
	minEquipZhuansheng := int32(0)
	bindType := itemtypes.ItemBindTypeUnBind
	for index, itemIndex := range itemIndexList {
		bagItem := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, itemIndex)
		if bagItem == nil || bagItem.IsEmpty() {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"ItemIndex": itemIndexList,
				}).Warn("goldequip:重铸失败,材料不存在")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
			return
		}

		itemTemplate := item.GetItemService().GetItem(int(bagItem.ItemId))
		if !itemTemplate.IsGoldEquip() {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"ItemIndex": itemIndexList,
				}).Warn("goldequip:重铸失败,材料不是元神金装")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotUse)
			return
		}

		if index == 0 {
			// 初始化条件
			badQuality = itemTemplate.Quality
			minEquipZhuansheng = itemTemplate.NeedZhuanShu
			continue
		}

		if badQuality > itemTemplate.Quality {
			badQuality = itemTemplate.Quality
		}
		if minEquipZhuansheng > itemTemplate.NeedZhuanShu {
			minEquipZhuansheng = itemTemplate.NeedZhuanShu
		}

		// 绑定属性
		if bagItem.BindType == itemtypes.ItemBindTypeBind {
			bindType = itemtypes.ItemBindTypeBind
		}
	}

	//品质大于等于橙色
	needQuality := int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeGoldEquipChongzhuQualityNeed))
	if badQuality < needQuality {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"badQuality": badQuality,
				"ItemIndex":  itemIndexList,
			}).Warn("goldequip:重铸失败,非橙色装备无法进行重铸")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipChongzhuQualityNotEnough)
		return
	}

	chongzhuTemp := goldequiptemplate.GetGoldEquipTemplateService().GetGoldEquipChongzhuTemplat(pl.GetRole(), pl.GetSex(), badQuality, minEquipZhuansheng)
	if chongzhuTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":           pl.GetId(),
				"role":               pl.GetRole(),
				"sex":                pl.GetSex(),
				"badQuality":         badQuality,
				"minEquipZhuansheng": minEquipZhuansheng,
				"ItemIndex":          itemIndexList,
			}).Warn("goldequip:重铸失败,重铸生成的物品模板不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(chongzhuTemp.DropId)
	if dropData == nil {
		panic(fmt.Errorf("goldequip: chongzhu GetDropItem should be ok"))
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
			}).Warn("goldequip:重铸失败,背包不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//消耗材料
	if len(itemIndexList) != 0 {
		itemUseReason := commonlog.InventoryLogReasonGoldEquipChongzhuUse
		flag, _ := inventoryManager.BatchGoldEquipRemoveIndex(itemIndexList, itemUseReason, itemUseReason.String())
		if !flag {
			panic(fmt.Errorf("goldequip:装备重铸移除材料应该成功"))
		}
	}

	//添加物品
	itemGetReason := commonlog.InventoryLogReasonGoldEquipChongzhuGet
	flag := inventoryManager.AddItemLevel(dropData, itemGetReason, itemGetReason.String())
	if !flag {
		panic(fmt.Errorf("goldequip:add item should be success"))
	}

	//同步改变
	goldequiplogic.SnapInventoryGoldEquipChanged(pl)
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	scGoldEquipChongzhu := pbutil.BuildSCGoldEquipChongzhu(itemId, level)
	pl.SendMsg(scGoldEquipChongzhu)

	return
}
