package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/inventory/inventory"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	ringlogic "fgame/fgame/game/ring/logic"
	"fgame/fgame/game/ring/pbutil"
	playerring "fgame/fgame/game/ring/player"
	ringtemplate "fgame/fgame/game/ring/template"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_RING_EQUIP_TYPE), dispatch.HandlerFunc(handleRingEquip))
}

func handleRingEquip(s session.Session, msg interface{}) (err error) {
	log.Debug("ring: 开始处理装备特戒请求消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csRingEquip := msg.(*uipb.CSRingEquip)
	index := csRingEquip.GetIndex()
	if index < 0 {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = ringEquip(tpl, index)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("ring: 处理装备特戒请求消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
		}).Debug("ring: 处理装备特戒请求消息,成功")

	return
}

func ringEquip(pl player.Player, index int32) (err error) {
	playerId := pl.GetId()
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeRing) {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
			}).Warn("ring: 功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	ringManager := pl.GetPlayerDataManager(playertypes.PlayerRingDataManagerType).(*playerring.PlayerRingDataManager)

	// 从背包寻找物品
	it := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, index)
	if it == nil || it.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"index":    index,
			}).Warn("ring: 使用特戒,物品不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	// 获取物品数据
	itemId := it.ItemId
	propertyData := it.PropertyData
	bindType := it.BindType

	// 判断该物品是否是特戒
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if !itemTemplate.IsTeRing() {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"index":    index,
			}).Warn("ring: 使用特戒,此物品不是特戒")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotEquip)
		return
	}

	if itemTemplate.NeedProfession != 0 {
		//角色
		if itemTemplate.GetRole() != pl.GetRole() {
			log.WithFields(
				log.Fields{
					"playerId": playerId,
					"index":    index,
				}).Warn("ring: 使用特戒,角色不符")
			playerlogic.SendSystemMessage(pl, lang.PlayerRoleWrong)
			return
		}
	}

	if itemTemplate.GetSex() != 0 {
		//性别
		if itemTemplate.GetSex() != pl.GetSex() {
			log.WithFields(
				log.Fields{
					"playerId": playerId,
					"index":    index,
				}).Warn("ring: 使用特戒,性别不符")
			playerlogic.SendSystemMessage(pl, lang.PlayerSexWrong)
			return
		}
	}

	//判断级别
	if itemTemplate.NeedLevel > pl.GetLevel() {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"index":    index,
			}).Warn("ring: 使用特戒,等级不够")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
		return
	}

	//判断转数
	if itemTemplate.NeedZhuanShu > pl.GetZhuanSheng() {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"index":    index,
			}).Warn("ring: 使用特戒,转数不够")
		playerlogic.SendSystemMessage(pl, lang.PlayerZhuanShengTooLow)
		return
	}

	ringTemp := ringtemplate.GetRingTemplateService().GetRingTemplate(itemId)
	if ringTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": playerId,
				"index":    index,
			}).Warn("ring: 模板不存在")
		playerlogic.SendSystemMessage(pl, lang.RingTempalteNotExist)
		return
	}
	typ := ringTemp.GetRingType()

	// 判断是否已经装备
	ringObj := ringManager.GetPlayerRingObject(typ)
	if ringObj != nil {
		// 脱下装备
		flag := ringlogic.TakeOffInternal(pl, typ)
		if !flag {
			return
		}
	}

	//移除物品
	reasonText := commonlog.InventoryLogReasonPutOn.String()
	flag, _ := inventoryManager.RemoveIndex(inventorytypes.BagTypePrim, index, 1, commonlog.InventoryLogReasonPutOn, reasonText)
	if !flag {
		panic(fmt.Errorf("ring: 移除物品应该是可以的"))
	}

	if propertyData == nil {
		base := inventorytypes.CreateDefaultItemPropertyDataBase()
		propertyData = inventory.CreatePropertyDataInterface(itemTemplate.GetItemType(), base)
	}

	flag = ringManager.RingEquipSuccess(typ, itemId, propertyData, bindType)
	if !flag {
		panic(fmt.Errorf("ring: 装备特戒应该成功"))
	}

	inventorylogic.SnapInventoryChanged(pl)

	// 推送属性变化
	ringlogic.RingPropertyChange(pl)
	propertylogic.SnapChangedProperty(pl)

	scRingEquip := pbutil.BuildSCRingEquip(index)
	pl.SendMsg(scRingEquip)

	// 推送特戒槽位变化
	obj := ringManager.GetPlayerRingObject(typ)
	scRingSlotChanged := pbutil.BuildSCRingSlotChanged(obj)
	pl.SendMsg(scRingSlotChanged)
	return
}
