package logic

import (
	"fgame/fgame/game/additionsys/pbutil"
	playeradditionsys "fgame/fgame/game/additionsys/player"
	additionsystemplate "fgame/fgame/game/additionsys/template"
	additionsystypes "fgame/fgame/game/additionsys/types"
	commomlogic "fgame/fgame/game/common/logic"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	skilllogic "fgame/fgame/game/skill/logic"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"
)

//根据系统类型功能开启
func GetAdditionSysFuncOpenByType(pl player.Player, typ additionsystypes.AdditionSysType) bool {
	funcOpenType, ok := typ.ConvertToFuncOpenType()
	if !ok {
		return false
	}
	return pl.IsFuncOpen(funcOpenType)
}

//根据系统类型 化灵功能开启
func GetAdditionSysHuaLingFuncOpenByType(pl player.Player, typ additionsystypes.AdditionSysType) bool {
	funcOpenType, ok := typ.ConvertToHuaLingFuncOpenType()
	if !ok {
		return false
	}
	return pl.IsFuncOpen(funcOpenType)
}

//根据系统类型 觉醒功能开启
func GetAdditionSysAwakeFuncOpenByType(pl player.Player, typ additionsystypes.AdditionSysType) bool {
	funcOpenType, ok := typ.ConvertToAwakeFuncOpenType()
	if !ok {
		return false
	}
	return pl.IsFuncOpen(funcOpenType)
}

//根据系统类型 神铸功能开启
func GetAdditionSysShenZhuFuncOpenByType(pl player.Player, typ additionsystypes.AdditionSysType) bool {
	funcOpenType, ok := typ.ConvertToShenZhuFuncOpenType()
	if !ok {
		return false
	}
	return pl.IsFuncOpen(funcOpenType)
}

//根据系统类型 通灵功能开启
func GetAdditionSysTongLingFuncOpenByType(pl player.Player, typ additionsystypes.AdditionSysType) bool {
	funcOpenType, ok := typ.ConvertToTongLingFuncOpenType()
	if !ok {
		return false
	}
	return pl.IsFuncOpen(funcOpenType)
}

//根据系统类型 升级功能开启
func GetAdditionSysShengJiFuncOpenByType(pl player.Player, typ additionsystypes.AdditionSysType) bool {
	funcOpenType, ok := typ.ConvertToShengJiFuncOpenType()
	if !ok {
		return false
	}
	return pl.IsFuncOpen(funcOpenType)
}

//根据系统类型 装备功能开启
func GetAdditionSysEquipFuncOpenByType(pl player.Player, typ additionsystypes.AdditionSysType) bool {
	funcOpenType, ok := typ.ConvertToEquipFuncOpenType()
	if !ok {
		return false
	}
	return pl.IsFuncOpen(funcOpenType)
}

//根据系统类型 灵珠功能开启
func GetAdditionSysLingZhuFuncOpenByType(pl player.Player, typ additionsystypes.AdditionSysType) bool {
	funcOpenType, ok := typ.ConvertToLingZhuOpenType()
	if !ok {
		return false
	}
	id, ok := typ.ConvertAdditionSysTypeToLingTongId()
	if !ok {
		return false
	}
	lingtongTemp := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(int32(id))
	if lingtongTemp == nil {
		return false
	}
	needLevel := lingtongTemp.LingzhuOpenLevel
	manager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	obj, ok := manager.GetLingTongInfo(int32(id))
	if !ok {
		return false
	}
	curLevel := obj.GetLevel()
	if curLevel < needLevel {
		return false
	}
	return pl.IsFuncOpen(funcOpenType)
}

func LingTongLingZhuActiveSkill(pl player.Player, typ additionsystypes.AdditionSysType) bool {
	id, ok := typ.ConvertAdditionSysTypeToLingTongId()
	if !ok {
		return false
	}
	additionsysManager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	totalLevel := int32(0)
	lingZhuMap := additionsysManager.GetAdditionSysLingZhuMap(typ)
	for _, obj := range lingZhuMap {
		objLevel := obj.GetLevel()
		totalLevel += objLevel
	}
	skillTemp, beforeSkillTemp := additionsystemplate.GetAdditionSysTemplateService().GetLingZhuSkillTemplate(int32(id), totalLevel)
	if skillTemp == nil {
		return false
	}
	beforeSkillId := int32(0)
	if beforeSkillTemp != nil {
		beforeSkillId = beforeSkillTemp.SkillId
	}
	err := skilllogic.TempSkillChange(pl, beforeSkillId, skillTemp.SkillId)
	if err != nil {
		return false
	}
	return true
}

//根据系统类型更新属性
func UpdataAdditionSysPropertyByType(pl player.Player, typ additionsystypes.AdditionSysType) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	effectType, ok := typ.ConvertToPropertyEffectType()
	if !ok {
		return
	}
	propertyManager.UpdateBattleProperty(effectType.Mask())

	return
}

//根据类型判断套装
func GetAdditionSysTaoZhuangByType(pl player.Player, typ additionsystypes.AdditionSysType) *gametemplate.SystemTaozhuangTemplate {
	additionsysManager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	taoZhuangMap := make(map[additionsystypes.SlotPositionType]itemtypes.ItemQualityType)

	for pos := additionsystypes.MinPosition; pos <= additionsystypes.MaxPosition; pos++ {
		slot := additionsysManager.GetAdditionSysByArg(typ, pos)
		if slot == nil || slot.IsEmpty() {
			return nil
		}
		itemTemplate := item.GetItemService().GetItem(int(slot.ItemId))
		taoZhuangTemplate := additionsystemplate.GetAdditionSysTemplateService().GetTaoZhuangByArg(typ, itemTemplate.GetQualityType())
		if taoZhuangTemplate == nil {
			return nil
		}
		taoZhuangMap[pos] = taoZhuangTemplate.GetTaozhuangQuality()
	}

	qualityType := taoZhuangMap[additionsystypes.MinPosition]
	for _, num := range taoZhuangMap {
		if qualityType != num {
			return nil
		}
	}
	temp := additionsystemplate.GetAdditionSysTemplateService().GetTaoZhuangByArg(typ, qualityType)
	return temp
}

//系统升级判断
func AdditionSysLevelJudge(pl player.Player, curTimesNum int32, curBless int32, template additionsystemplate.SystemShengJiCommonTemplate) (sucess bool, pro, randBless, addTimes int32) {
	updateRate := template.GetUpdateWfb()
	blessMax := template.GetZhufuMax()
	addMin := template.GetAddMin()
	addMax := template.GetAddMax() + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	addTimes = int32(1)
	curTimesNum += addTimes

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, template.GetTimesMin(), template.GetTimesMax(), randBless, updateRate, blessMax)
	return
}

//系统新升级判断通用接口
func AdditionSysUpgradeCommonJudge(pl player.Player, curTimesNum int32, curBless int32, template additionsystemplate.SystemUpgradeCommonTemplate) (sucess bool, pro, randBless, addTimes int32) {
	updateRate := template.GetUpdateWfb()
	blessMax := template.GetZhufuMax()
	addMin := template.GetAddMin()
	addMax := template.GetAddMax() + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	addTimes = int32(1)
	curTimesNum += addTimes

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, template.GetTimesMin(), template.GetTimesMax(), randBless, updateRate, blessMax)
	return
}

//系统部位神铸判断
func AdditionSysShenZhuJudge(pl player.Player, curTimesNum int32, curBless int32, template *gametemplate.SystemShenZhuTemplate) (sucess bool, pro, randBless, addTimes int32) {
	updateRate := template.UpdateWfb
	blessMax := template.ZhufuMax
	addMin := template.AddMin
	addMax := template.AddMax + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	addTimes = int32(1)
	curTimesNum += addTimes

	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, template.TimesMin, template.TimesMax, randBless, updateRate, blessMax)
	return
}

//推送装备改变
func SnapInventoryAdditionSysEquipChangedByType(pl player.Player, typ additionsystypes.AdditionSysType) (err error) {
	additionsysManager := pl.GetPlayerDataManager(types.PlayerAdditionSysDataManagerType).(*playeradditionsys.PlayerAdditionSysDataManager)
	slotChangedList := additionsysManager.GetChangedEquipmentSlotAndReset(typ)
	if len(slotChangedList) <= 0 {
		return
	}
	equipmentChanged := pbutil.BuildSCAdditionSysSlotChanged(int32(typ), slotChangedList)
	pl.SendMsg(equipmentChanged)
	return
}
