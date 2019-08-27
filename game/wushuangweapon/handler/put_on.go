package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	itemservice "fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	gamesession "fgame/fgame/game/session"
	wushuangweaponlogic "fgame/fgame/game/wushuangweapon/logic"
	"fgame/fgame/game/wushuangweapon/pbutil"
	playerwushuangweapon "fgame/fgame/game/wushuangweapon/player"
	wushuangweapontypes "fgame/fgame/game/wushuangweapon/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_WUSHUANGWEAPON_PUT_ON_TYPE), dispatch.HandlerFunc(handlerPutOn))
}

func handlerPutOn(s session.Session, msg interface{}) (err error) {
	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSWushuangWeaponPutOn)
	index := csMsg.GetIndex()

	err = putOn(tpl, index)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"err":      err,
			}).Error("WushuangWeapon:处理穿上无双神器，错误")
		return
	}
	return
}

func putOn(pl player.Player, index int32) (err error) {
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	itemObj := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, index)

	//物品不存在
	if itemObj == nil || itemObj.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("inventory:使用装备,物品不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	itemId := itemObj.ItemId
	bindType := itemObj.BindType
	itemTemplate := itemservice.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
			}).Warn("wushuangWeapon:物品模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	baseTemp := itemTemplate.GetWushuangBaseTemplate()
	if baseTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemId,
			}).Warn("wushuangWeapon:无双神器模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}
	//判断传入物品类型
	if itemTemplate.GetItemType() != itemtypes.ItemTypeWushuangWeapon {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemType": itemTemplate.GetItemType(),
			}).Warn("wushuangWeapon:物品类型错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	itemSubType := itemTemplate.GetItemSubType()
	itemWushuangSubType, ok := itemSubType.(itemtypes.ItemWushuangWeaponSubType)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemType": itemTemplate.GetItemType(),
			}).Warn("wushuangWeapon:应该能获得无双神器物品的部位类型")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	bodyPos := wushuangweapontypes.GetBodyPosFromItemSubType(itemWushuangSubType)
	wushuangDataManager := pl.GetPlayerDataManager(playertypes.PlayerWushuangWeaponDataManagerType).(*playerwushuangweapon.PlayerWushuangWeaponDataManager)
	slotObj := wushuangDataManager.GetSlotObjectFromBodyPos(bodyPos)

	if itemTemplate.NeedProfession != 0 {
		//角色
		if itemTemplate.GetRole() != pl.GetRole() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"index":    index,
				}).Warn("wushuangWeapon:使用无双神器,角色不符")
			playerlogic.SendSystemMessage(pl, lang.PlayerRoleWrong)
			return
		}
	}
	if itemTemplate.NeedGender != 0 {
		//性别
		if itemTemplate.GetSex() != pl.GetSex() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"index":    index,
				}).Warn("wushuangWeapon:使用无双神器,性别不符")
			playerlogic.SendSystemMessage(pl, lang.PlayerSexWrong)
			return
		}
	}
	//判断级别
	if itemTemplate.NeedLevel > pl.GetLevel() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("wushuangWeapon:使用无双神器,等级不够")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
		return
	}

	//判断转数
	if itemTemplate.NeedZhuanShu > pl.GetZhuanSheng() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("wushuangWeapon:使用无双神器,转数不够")
		playerlogic.SendSystemMessage(pl, lang.PlayerZhuanShengTooLow)
		return
	}

	//判断是否已经装备
	if slotObj.IsEquip() {
		flag := wushuangweaponlogic.TakeOffLogic(pl, bodyPos)
		if !flag {
			return
		}
	}

	//从物品配置读取到槽位上
	settingsObj := wushuangDataManager.GetWushuangSettings(itemId)
	settingsLevel := int32(0)
	if settingsObj != nil {
		settingsLevel = settingsObj.GetLevel()
	}
	flag := slotObj.PutOn(itemId, bindType, settingsLevel)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"itemId":   itemTemplate.Id,
			}).Error("wushuangWeapon:使用装备失败")
		playerlogic.SendSystemMessage(pl, lang.WushuangWeaponItemNotExist)
		return
	}
	// flag = wushuangweaponlogic.PutOnEquipmentChangeLevel(slotObj)
	// if !flag {
	// 	panic("WushuangWeapon:穿上装备情况下应该是能更新等级的")
	// }

	//删除装备
	reasonText := commonlog.InventoryLogReasonPutOn.String()
	flag, _ = inventoryManager.RemoveIndex(inventorytypes.BagTypePrim, index, 1, commonlog.InventoryLogReasonPutOn, reasonText)
	if !flag {
		panic("inventory:移除物品应该是可以的")
	}
	//更改属性
	wushuangweaponlogic.WushuangWeaponPropertyChanged(pl)

	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	//发消息
	scMsg := pbutil.BuildSCWushuangWeaponPutOn(bodyPos, slotObj.GetLevel(), slotObj.GetExperience())
	pl.SendMsg(scMsg)
	return
}
