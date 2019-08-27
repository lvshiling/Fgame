package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	additionsyscommon "fgame/fgame/game/additionsys/common"
	playeradditionsys "fgame/fgame/game/additionsys/player"
	additionsystypes "fgame/fgame/game/additionsys/types"
)

func BuildSCAdditionSysLingZhuInfo(lingtongId int32, objList []*playeradditionsys.PlayerAdditionSysLingZhuObject) *uipb.SCAdditionSysLingTongLingZhuInfo {
	msg := &uipb.SCAdditionSysLingTongLingZhuInfo{}
	msg.LingtongId = &lingtongId
	for _, obj := range objList {
		msg.LingzhuInfo = append(msg.LingzhuInfo, buildLingzhuInfo(obj))
	}
	return msg
}

func buildLingzhuInfo(obj *playeradditionsys.PlayerAdditionSysLingZhuObject) *uipb.LingTongLingZhuInfo {
	info := &uipb.LingTongLingZhuInfo{}
	level := obj.GetLevel()
	typ := int32(obj.GetLingZhuType())
	times := obj.GetTimes()
	bless := obj.GetBless()
	info.Level = &level
	info.LingZhuType = &typ
	info.Times = &times
	info.Bless = &bless
	return info
}

func BuildSCAdditionSysLingZhuUplevel(lingtongId int32, obj *playeradditionsys.PlayerAdditionSysLingZhuObject) *uipb.SCAdditionSysLingTongLingZhuUplevel {
	msg := &uipb.SCAdditionSysLingTongLingZhuUplevel{}
	msg.LingtongId = &lingtongId
	msg.LingzhuInfo = buildLingzhuInfo(obj)
	return msg
}

func BuildSCUseAdditionSysEquip(sysType int32, index int32) *uipb.SCUseAdditionSysEquip {
	scUseAdditionSysEquip := &uipb.SCUseAdditionSysEquip{}
	scUseAdditionSysEquip.SysType = &sysType
	scUseAdditionSysEquip.Index = &index
	return scUseAdditionSysEquip
}

func BuildSCTakeOffAdditionSysEquip(sysType int32, slotInt int32) *uipb.SCTakeOffAdditionSysEquip {
	scTakeOffAdditionSysEquip := &uipb.SCTakeOffAdditionSysEquip{}
	scTakeOffAdditionSysEquip.SysType = &sysType
	scTakeOffAdditionSysEquip.SlotId = &slotInt
	return scTakeOffAdditionSysEquip
}

// func BuildSCTakeOffAdditionSysEquip(sysType int32, pos inventorytypes.BodyPositionType) *uipb.SCTakeOffAdditionSysEquip {
// 	scTakeOffAdditionSysEquip := &uipb.SCTakeOffAdditionSysEquip{}
// 	slotInt := int32(pos)
// 	scTakeOffAdditionSysEquip.SysType = &sysType
// 	scTakeOffAdditionSysEquip.SlotId = &slotInt
// 	return scTakeOffAdditionSysEquip
// }

func BuildSCAdditionSysStrengthenBody(sysType int32, slotId int32, result int32, auto bool, isProtect bool) *uipb.SCAdditionSysStrengthenBody {
	scAdditionSysStrengthenBody := &uipb.SCAdditionSysStrengthenBody{}
	scAdditionSysStrengthenBody.SysType = &sysType
	scAdditionSysStrengthenBody.SlotId = &slotId
	scAdditionSysStrengthenBody.Result = &result
	scAdditionSysStrengthenBody.Auto = &auto
	scAdditionSysStrengthenBody.IsProtect = &isProtect
	return scAdditionSysStrengthenBody
}

func BuildSCAdditionSysSlotChanged(sysType int32, slotList []*playeradditionsys.PlayerAdditionSysSlotObject) *uipb.SCAdditionSysSlotChanged {
	scAdditionSysSlotChanged := &uipb.SCAdditionSysSlotChanged{}
	scAdditionSysSlotChanged.SysType = &sysType
	scAdditionSysSlotChanged.SlotList = buildEquipmentSlotList(slotList)
	return scAdditionSysSlotChanged
}

func BuildSCAdditionSysAwakeEat(sysType int32, level int32) *uipb.SCAdditionSysAwakeEat {
	scAwake := &uipb.SCAdditionSysAwakeEat{}
	scAwake.SysType = &sysType
	scAwake.AwakeLevel = &level
	return scAwake
}

func buildEquipmentSlotList(slotList []*playeradditionsys.PlayerAdditionSysSlotObject) (slotItemList []*uipb.AdditionSysSlotInfo) {
	for _, slot := range slotList {
		slotItemList = append(slotItemList, buildEquipmentSlot(slot))
	}
	return slotItemList
}

func buildEquipmentSlot(slot *playeradditionsys.PlayerAdditionSysSlotObject) *uipb.AdditionSysSlotInfo {
	slotItem := &uipb.AdditionSysSlotInfo{}
	sysType := int32(slot.GetSysType())
	slotItem.SysType = &sysType
	slotId := int32(slot.GetSlotId())
	slotItem.SlotId = &slotId
	level := slot.GetLevel()
	slotItem.Level = &level
	shenZhuLev := slot.GetShenZhuLev()
	slotItem.ShenZhuLev = &shenZhuLev
	shenZhuPro := slot.GetShenZhuPro()
	slotItem.ShenZhuPro = &shenZhuPro
	itemId := slot.GetItemId()
	slotItem.ItemId = &itemId
	bindInt := int32(slot.GetBindType())
	slotItem.BindType = &bindInt

	return slotItem
}

func BuildSCAdditionSysSlotInfoList(bags map[additionsystypes.AdditionSysType]*playeradditionsys.BodyBag, levelInfoAll map[additionsystypes.AdditionSysType]*playeradditionsys.PlayerAdditionSysLevelObject, tongLingInfoAll map[additionsystypes.AdditionSysType]*playeradditionsys.PlayerAdditionSysTongLingObject, awakeInfoAll map[additionsystypes.AdditionSysType]*playeradditionsys.PlayerAdditionSysAwakeObject) *uipb.SCAdditionSysSlotInfo {
	scAdditionSysSlotInfo := &uipb.SCAdditionSysSlotInfo{}
	//部位信息
	for _, bag := range bags {
		for _, slot := range bag.GetAll() {
			slotPb := buildEquipmentSlot(slot)
			scAdditionSysSlotInfo.SlotList = append(scAdditionSysSlotInfo.SlotList, slotPb)
		}
	}
	//等级信息
	for sysType, levelInfo := range levelInfoAll {
		isAwake := awakeInfoAll[sysType].IsAwake
		levelPb := buildLevelInfo(levelInfo, isAwake)
		scAdditionSysSlotInfo.LevelList = append(scAdditionSysSlotInfo.LevelList, levelPb)
	}

	//通灵信息
	for _, tongLingInfo := range tongLingInfoAll {
		tongLingPb := buildAdditionSysTongLingInfo(tongLingInfo)
		scAdditionSysSlotInfo.TongLingList = append(scAdditionSysSlotInfo.TongLingList, tongLingPb)
	}

	return scAdditionSysSlotInfo
}

func buildLevelInfo(info *playeradditionsys.PlayerAdditionSysLevelObject, isAwake int32) *uipb.AdditionSysLevelInfo {
	levelInfo := &uipb.AdditionSysLevelInfo{}
	sysType := int32(info.SysType)
	levelInfo.SysType = &sysType
	level := info.Level
	levelInfo.Level = &level
	upPro := info.UpPro
	levelInfo.Progress = &upPro
	lingLevel := info.LingLevel
	levelInfo.LingLevel = &lingLevel
	lingPro := info.LingPro
	levelInfo.LingPro = &lingPro
	levelInfo.IsAwake = &isAwake
	return levelInfo
}

func buildAdditionSysTongLingInfo(info *playeradditionsys.PlayerAdditionSysTongLingObject) *uipb.AdditionSysTongLingInfo {
	tongLingInfo := &uipb.AdditionSysTongLingInfo{}
	sysType := int32(info.SysType)
	tongLingLev := info.TongLingLev
	tongLingPro := info.TongLingPro

	tongLingInfo.SysType = &sysType
	tongLingInfo.TongLingLev = &tongLingLev
	tongLingInfo.TongLingPro = &tongLingPro
	return tongLingInfo
}

func BuildSCAdditionSysShengJi(sysType int32, level int32, progress int32, result int32) *uipb.SCAdditionSysShengJi {
	scAdditionSysShengJi := &uipb.SCAdditionSysShengJi{}
	scAdditionSysShengJi.SysType = &sysType
	scAdditionSysShengJi.Level = &level
	scAdditionSysShengJi.Progress = &progress
	scAdditionSysShengJi.Result = &result
	return scAdditionSysShengJi
}

func BuildSCAdditionSysHuaLingEat(sysType int32, level int32, progress int32) *uipb.SCAdditionSysHualingEat {
	scAdditionSysHualingEat := &uipb.SCAdditionSysHualingEat{}
	scAdditionSysHualingEat.SysType = &sysType
	scAdditionSysHualingEat.Level = &level
	scAdditionSysHualingEat.Progress = &progress
	return scAdditionSysHualingEat
}

func BuildSCAdditionSysShenZhuBody(sysType additionsystypes.AdditionSysType, pos additionsystypes.SlotPositionType, level int32, progress int32) *uipb.SCAdditionSysShenZhuBody {
	scAdditionSysShenZhuBody := &uipb.SCAdditionSysShenZhuBody{}
	sysTypeInt := int32(sysType)
	posInt := int32(pos)

	scAdditionSysShenZhuBody.SysType = &sysTypeInt
	scAdditionSysShenZhuBody.SlotId = &posInt
	scAdditionSysShenZhuBody.Level = &level
	scAdditionSysShenZhuBody.Progress = &progress
	return scAdditionSysShenZhuBody
}

func BuildSCAdditionSysTongLingUpgrade(sysType additionsystypes.AdditionSysType, level int32, progress int32) *uipb.SCAdditionSysTongLingUpgrade {
	scAdditionSysTongLingUpgrade := &uipb.SCAdditionSysTongLingUpgrade{}
	sysTypeInt := int32(sysType)

	scAdditionSysTongLingUpgrade.SysType = &sysTypeInt
	scAdditionSysTongLingUpgrade.Level = &level
	scAdditionSysTongLingUpgrade.Progress = &progress
	return scAdditionSysTongLingUpgrade
}

func BuildSCAdditionSysUpgrade(sysType additionsystypes.AdditionSysType, posType additionsystypes.SlotPositionType, result int32) *uipb.SCAdditionSysUpgrade {
	scAdditionSysUpgrade := &uipb.SCAdditionSysUpgrade{}
	sysTypeInt := int32(sysType)
	slotInt := int32(posType)
	scAdditionSysUpgrade.SysType = &sysTypeInt
	scAdditionSysUpgrade.SlotId = &slotInt
	scAdditionSysUpgrade.Result = &result
	return scAdditionSysUpgrade
}

func BuildAllAdditionSysInfo(info *additionsyscommon.AllAdditionSysInfo) *uipb.AllAdditionSysInfo {
	allAdditionSysInfo := &uipb.AllAdditionSysInfo{}
	for _, additionSys := range info.AdditionSysList {
		allAdditionSysInfo.AdditionSysList = append(allAdditionSysInfo.AdditionSysList, buildAdditionSysCacheInfo(additionSys))
	}
	return allAdditionSysInfo
}

func buildAdditionSysCacheInfo(info *additionsyscommon.AdditionSysInfo) *uipb.AdditionSysCacheInfo {
	additionSysCacheInfo := &uipb.AdditionSysCacheInfo{}
	typ := info.SysType
	lev := info.Level
	pro := info.UpPro
	lingLev := info.LingLevel
	lingPro := info.LingPro
	isAwake := info.IsAwake

	additionSysCacheInfo.SysType = &typ
	additionSysCacheInfo.Level = &lev
	additionSysCacheInfo.UpPro = &pro
	additionSysCacheInfo.LingLevel = &lingLev
	additionSysCacheInfo.LingPro = &lingPro
	additionSysCacheInfo.IsAwake = &isAwake
	if info.TongLingInfo != nil {
		additionSysCacheInfo.TongLingInfo = buildAdditionSysTongLingCacheInfo(info.TongLingInfo)
	}
	for _, sysTypeSlot := range info.SysTypeSlotList {
		additionSysCacheInfo.SysTypeSlotList = append(additionSysCacheInfo.SysTypeSlotList, buildAdditionSysSlotCacheInfo(sysTypeSlot))
	}
	return additionSysCacheInfo
}

func buildAdditionSysSlotCacheInfo(info *additionsyscommon.AdditionSysSlotInfo) *uipb.AdditionSysSlotCacheInfo {
	additionSysSlotCacheInfo := &uipb.AdditionSysSlotCacheInfo{}
	slotId := info.SlotId
	level := info.Level
	itemId := info.ItemId
	shenZhuLev := info.ShenZhuLev
	shenZhuPro := info.ShenZhuPro

	additionSysSlotCacheInfo.SlotId = &slotId
	additionSysSlotCacheInfo.Level = &level
	additionSysSlotCacheInfo.ShenZhuLev = &shenZhuLev
	additionSysSlotCacheInfo.ShenZhuPro = &shenZhuPro
	additionSysSlotCacheInfo.ItemId = &itemId
	return additionSysSlotCacheInfo
}

func buildAdditionSysTongLingCacheInfo(info *additionsyscommon.AdditionSysTongLingInfo) *uipb.AdditionSysTongLingCacheInfo {
	additionSysTongLingCacheInfo := &uipb.AdditionSysTongLingCacheInfo{}
	tongLingLev := info.TongLingLev
	tongLingPro := info.TongLingPro

	additionSysTongLingCacheInfo.TongLingLev = &tongLingLev
	additionSysTongLingCacheInfo.TongLingPro = &tongLingPro
	return additionSysTongLingCacheInfo
}
