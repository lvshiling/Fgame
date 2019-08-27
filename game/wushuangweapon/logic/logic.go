package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	droptemplate "fgame/fgame/game/drop/template"
	fashionservice "fgame/fgame/game/fashion/fashion"
	fashionlogic "fgame/fgame/game/fashion/logic"
	fashionpbutil "fgame/fgame/game/fashion/pbutil"
	playerfashion "fgame/fgame/game/fashion/player"
	playerinventory "fgame/fgame/game/inventory/player"
	itemservice "fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	weaponlogic "fgame/fgame/game/weapon/logic"
	weaponpbutil "fgame/fgame/game/weapon/pbutil"
	playerweapon "fgame/fgame/game/weapon/player"
	weaponservice "fgame/fgame/game/weapon/weapon"
	playerwushuangweapon "fgame/fgame/game/wushuangweapon/player"
	wushuangweapontypes "fgame/fgame/game/wushuangweapon/types"

	"fmt"

	log "github.com/Sirupsen/logrus"
)

// 无双神器属性变更
func WushuangWeaponPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeWushuangWeapon.Mask())
	return
}

// 脱掉装备处理逻辑
func TakeOffLogic(pl player.Player, bodyPos wushuangweapontypes.WushuangWeaponPart) (flag bool) {
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	wushuangDataManager := pl.GetPlayerDataManager(playertypes.PlayerWushuangWeaponDataManagerType).(*playerwushuangweapon.PlayerWushuangWeaponDataManager)
	slotObj := wushuangDataManager.GetSlotObjectFromBodyPos(bodyPos)
	if !slotObj.IsEquip() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bodyPos":  bodyPos,
			}).Warn("wushuangweapon:脱下无双神器，未装备")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipCanNotTakeOff)
		return
	}
	//背包空间
	itemId := slotObj.GetItemId()
	num := int32(1)
	if !inventoryManager.HasEnoughSlot(itemId, num) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bodyPos":  bodyPos,
			}).Warn("wushuangweapon:脱下无双神器，背包不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	wushuangDataManager.SaveWushuangSettings(slotObj.GetItemId(), slotObj.GetLevel())
	slotObj.TakeOff()

	//添加物品
	itemTemplate := itemservice.GetItemService().GetItem(int(itemId))
	reason := commonlog.InventoryLogReasonWushuangTakeOff
	reasonText := fmt.Sprintf(reason.String(), itemTemplate.Name, bodyPos.String())
	itemData := droptemplate.CreateItemData(itemId, num, 0, slotObj.GetBindType())
	flag = inventoryManager.AddItemLevel(itemData, reason, reasonText)
	if !flag {
		panic(fmt.Errorf("wushuangweapon:add item should be success"))
	}
	return
}

// 激活外观
func ActiveShow(pl player.Player, bodyPos wushuangweapontypes.WushuangWeaponPart, itemId int32) {
	itemTemp := itemservice.GetItemService().GetItem(int(itemId))
	if itemTemp == nil {
		return
	}
	itemName := itemTemp.Name
	linkArgs := []int64{int64(chattypes.ChatLinkTypeItem), int64(itemId)}
	nameLink := coreutils.FormatLink(itemName, linkArgs)
	template := itemTemp.GetWushuangBaseTemplate()
	weartype := template.GetWearType()
	if !weartype.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bodyPos":  bodyPos,
			}).Warn("wushuangWeapon:穿戴类型错误")
		return
	}
	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, fmt.Sprintf("%s", pl.GetName()))
	switch weartype {
	case wushuangweapontypes.BodyPosWearTypeWeapon:
		//武器
		weaponDataManager := pl.GetPlayerDataManager(playertypes.PlayerWeaponDataManagerType).(*playerweapon.PlayerWeaponDataManager)
		ok := weaponDataManager.WeaponActiveTemp(template.WaiguanId)
		if !ok {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"bodyPos":  bodyPos,
					"weaponId": template.WaiguanId,
				}).Warn("wushuangWeapon:兵魂激活失败")
			return
		}
		ok = weaponDataManager.Wear(template.WaiguanId)
		if !ok {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"bodyPos":  bodyPos,
					"weaponId": template.WaiguanId,
				}).Warn("wushuangWeapon:兵魂穿戴失败")
			return
		}
		//同步属性
		weaponlogic.WeaponPropertyChanged(pl)
		scWeaponActive := weaponpbutil.BuildSCWeaponActive(template.WaiguanId)
		pl.SendMsg(scWeaponActive)
		//全服广播
		weaponTemp := weaponservice.GetWeaponService().GetWeaponTemplate(int(template.WaiguanId))
		power := propertylogic.CulculateForce(weaponTemp.GetBattleAttrTemplate().GetAllBattleProperty())
		powerStr := coreutils.FormatColor(chattypes.ColorTypeEmailRedWord, fmt.Sprintf(lang.GetLangService().ReadLang(lang.WushuangWeaponPowerUp), power))
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.WushuangWeaponBreakthroughTopLevel), playerName, nameLink, powerStr)
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	case wushuangweapontypes.BodyPosWearTypeCloths:
		//时装
		fashionDataManager := pl.GetPlayerDataManager(playertypes.PlayerFashionDataManagerType).(*playerfashion.PlayerFashionDataManager)
		activeTime, ok := fashionDataManager.FashionActive(template.WaiguanId, true)
		if !ok {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"bodyPos":   bodyPos,
					"fashionId": template.WaiguanId,
				}).Warn("wushuangWeapon:时装激活失败")
			return
		}
		ok = fashionDataManager.FashionWear(template.WaiguanId)
		if !ok {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"bodyPos":   bodyPos,
					"fashionId": template.WaiguanId,
				}).Warn("wushuangWeapon:时装穿戴失败")
			return
		}
		//同步属性
		fashionlogic.FashionPropertyChanged(pl)
		scFashionActive := fashionpbutil.BuildSCFashionActive(template.WaiguanId, activeTime)
		pl.SendMsg(scFashionActive)
		//全服广播
		fashionTemp := fashionservice.GetFashionService().GetFashionTemplate(int(template.WaiguanId))
		power := propertylogic.CulculateForce(fashionTemp.GetBattleAttrTemplate().GetAllBattleProperty())
		powerStr := coreutils.FormatColor(chattypes.ColorTypeEmailRedWord, fmt.Sprintf(lang.GetLangService().ReadLang(lang.WushuangWeaponPowerUp), power))
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.WushuangWeaponBreakthroughTopLevel), playerName, nameLink, powerStr)
		chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	}
}

//穿上装备计算等级
// func PutOnEquipmentChangeLevel(slotObj *playerwushuangweapon.PlayerWushuangWeaponSlotObject) bool {
// 	if !slotObj.IsEquip() {
// 		return false
// 	}
// 	itemTemp := itemservice.GetItemService().GetItem(int(slotObj.GetItemId()))
// 	if itemTemp == nil {
// 		return false
// 	}
// 	baseTemplate := itemTemp.GetWushuangBaseTemplate()
// 	cntEx := slotObj.GetExperience()
// 	totalLevel := baseTemplate.GetLevel(cntEx)

// 	slotObj.ChangeLevel(totalLevel)
// 	return true
// }
