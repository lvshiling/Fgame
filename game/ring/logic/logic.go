package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	commomlogic "fgame/fgame/game/common/logic"
	droplogic "fgame/fgame/game/drop/logic"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fgame/fgame/game/ring/pbutil"
	playerring "fgame/fgame/game/ring/player"
	ringtemplate "fgame/fgame/game/ring/template"
	ringtypes "fgame/fgame/game/ring/types"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"

	log "github.com/Sirupsen/logrus"
)

// 特戒进阶判断
func RingAdvance(curTimesNum int32, curBless int32, temp *gametemplate.RingAdvanceTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := temp.TimesMin
	timesMax := temp.TimesMax
	updateRate := temp.UpdateWfb
	blessMax := temp.ZhufuMax
	addMin := temp.AddMin
	addMax := temp.AddMax + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

// 特戒强化升级判断
func RingStrengthen(curTimesNum int32, temp *gametemplate.RingStrengthenTemplate) (success bool) {
	timesMin := temp.TimesMin
	timesMax := temp.TimesMax
	updateRate := temp.UpdateWfb
	_, _, success = commomlogic.GetStatusAndProgress(curTimesNum, 0, timesMin, timesMax, 0, 0, updateRate, 1)
	return
}

// 特戒净灵升级判断
func RingJingLing(curTimesNum int32, temp *gametemplate.RingJingLingTemplate) (success bool) {
	timesMin := temp.TimesMin
	timesMax := temp.TimesMax
	updateRate := temp.UpdateWfb
	_, _, success = commomlogic.GetStatusAndProgress(curTimesNum, 0, timesMin, timesMax, 0, 0, updateRate, 1)
	return
}

// 属性变化
func RingPropertyChange(pl player.Player) {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeRing.Mask())
}

// 卸下装备
func RingUnload(pl player.Player, typ ringtypes.RingType) (err error) {
	flag := TakeOffInternal(pl, typ)
	if !flag {
		return
	}

	inventorylogic.SnapInventoryChanged(pl)

	// 推送属性变化
	RingPropertyChange(pl)
	propertylogic.SnapChangedProperty(pl)

	scRingUnload := pbutil.BuildSCRingUnload(int32(typ))
	pl.SendMsg(scRingUnload)
	return
}

// 脱装备
func TakeOffInternal(pl player.Player, typ ringtypes.RingType) (flag bool) {
	ringManager := pl.GetPlayerDataManager(playertypes.PlayerRingDataManagerType).(*playerring.PlayerRingDataManager)
	ringObj := ringManager.GetPlayerRingObject(typ)
	if ringObj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("ring: 玩家未穿戴该特戒")
		playerlogic.SendSystemMessage(pl, lang.RingNotEquip)
		return
	}

	itemId := ringObj.GetItemId()
	propertyData := ringObj.GetPropertyData()
	bindType := ringObj.GetBindType()

	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ.String(),
			}).Warn("ring: 模板不存在")
		playerlogic.SendSystemMessage(pl, lang.RingTempalteNotExist)
		return
	}

	level := itemTemplate.NeedLevel
	num := int32(1)

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//背包空间
	if !inventoryManager.HasEnoughSlotItemLevel(itemId, num, level, bindType) {
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	flag = ringManager.RingUnloadSuccess(typ)
	if !flag {
		panic("ring: 卸下特戒应该成功")
	}

	//添加物品
	reasonText := commonlog.InventoryLogReasonTakeOff.String()
	flag = inventoryManager.AddItemLevelWithPropertyData(itemId, num, level, bindType, propertyData, commonlog.InventoryLogReasonTakeOff, reasonText)
	if !flag {
		panic("ring: 添加特戒应该成功")
	}

	return true
}

func RingLuckyPointsTop(pl player.Player, typ ringtypes.BaoKuType) (flag bool, rewList []*droptemplate.DropItemData) {
	baoKuTemp := ringtemplate.GetRingTemplateService().GetRingBaoKuTemplate(typ)
	if baoKuTemp == nil {
		return
	}

	//掉落
	dropData := droptemplate.GetDropTemplateService().GetDropBaoKuItemLevel(baoKuTemp.ScriptXingYun)
	if dropData != nil {
		rewList = append(rewList, dropData)
	}

	var rewItemList []*droptemplate.DropItemData
	var resMap map[itemtypes.ItemAutoUseResSubType]int32
	if len(rewList) > 0 {
		rewItemList, resMap = droplogic.SeperateItemDatas(rewList)
	}

	//背包空间
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughSlotsOfItemLevel(rewItemList) {
		now := global.GetGame().GetTimeService().Now()
		title := lang.GetLangService().ReadLang(lang.RingLuckyBoxTitle)
		content := lang.GetLangService().ReadLang(lang.RingLuckyBoxContent)
		emaillogic.AddEmailItemLevel(pl, title, content, now, rewList)
		return
	}

	//增加掉落
	if len(resMap) > 0 {
		goldReason := commonlog.GoldLogReasonRingBaoKuLuckyBoxGet
		silverReason := commonlog.SilverLogReasonRingBaoKuLuckyBoxGet
		levelReason := commonlog.LevelLogReasonRingBaoKuLuckyBoxGet
		err := droplogic.AddRes(pl, resMap, goldReason, goldReason.String(), silverReason, silverReason.String(), levelReason, levelReason.String())
		if err != nil {
			return
		}
	}

	if len(rewItemList) > 0 {
		itemGetReason := commonlog.InventoryLogReasonRingBaoKuLuckyBoxGet
		flag := inventoryManager.BatchAddOfItemLevel(rewItemList, itemGetReason, itemGetReason.String())
		if !flag {
			panic("ring:增加物品应该成功")
		}
	}

	//同步
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	flag = true
	return
}
