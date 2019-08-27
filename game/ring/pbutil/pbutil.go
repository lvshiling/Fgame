package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droptemplate "fgame/fgame/game/drop/template"
	ringcommon "fgame/fgame/game/ring/common"
	playerring "fgame/fgame/game/ring/player"
	ringtypes "fgame/fgame/game/ring/types"
)

func BuildSCRingInfoList(ringInfoList []*ringcommon.RingInfo) (msgList []*uipb.RingInfo) {
	msgList = make([]*uipb.RingInfo, 0, 8)
	for _, info := range ringInfoList {
		msg := &uipb.RingInfo{}
		typ := int32(info.Typ)
		msg.Type = &typ
		msg.ItemId = &info.ItemId
		msg.Advance = &info.PropertyData.Advance
		msg.StrengthenLevel = &info.PropertyData.StrengthLevel
		msg.JingLingLevel = &info.PropertyData.JingLingLevel
		msgList = append(msgList, msg)
	}
	return
}

func BuildSCRingInfoGet(ringObjMap map[ringtypes.RingType]*playerring.PlayerRingObject) *uipb.SCRingInfoGet {
	scRingInfoGet := &uipb.SCRingInfoGet{}
	for _, obj := range ringObjMap {
		scRingInfoGet.RingInfo = append(scRingInfoGet.RingInfo, buildRingInfo(obj))
	}
	return scRingInfoGet
}

func buildRingInfo(obj *playerring.PlayerRingObject) *uipb.RingInfo {
	ringInfo := &uipb.RingInfo{}
	itemId := obj.GetItemId()
	typ := int32(obj.GetRingType())
	data := obj.GetPropertyData()
	ringData := data.(*ringtypes.RingPropertyData)
	ringInfo.ItemId = &itemId
	ringInfo.Type = &typ
	ringInfo.Advance = &ringData.Advance
	ringInfo.AdvancePro = &ringData.AdvancePro
	ringInfo.StrengthenLevel = &ringData.StrengthLevel
	ringInfo.JingLingLevel = &ringData.JingLingLevel
	return ringInfo
}

func BuildSCRingBaoKuInfo(obj *playerring.PlayerRingBaoKuObject) *uipb.SCRingBaoKuInfo {
	scRingBaoKuInfo := &uipb.SCRingBaoKuInfo{}
	scRingBaoKuInfo.RingBaoKuInfo = buildRingBaoKuInfo(obj)
	return scRingBaoKuInfo
}

func buildRingBaoKuInfo(obj *playerring.PlayerRingBaoKuObject) *uipb.RingBaoKuInfo {
	ringInfo := &uipb.RingBaoKuInfo{}
	luckyPoints := obj.GetLuckyPoints()
	attendPoints := obj.GetAttendPoints()
	ringInfo.LuckyPoints = &luckyPoints
	ringInfo.AttendPoints = &attendPoints
	return ringInfo
}

func BuildSCRingSlotChanged(obj *playerring.PlayerRingObject) *uipb.SCRingSlotChanged {
	scRingSlotChanged := &uipb.SCRingSlotChanged{}
	scRingSlotChanged.RingInfo = buildRingInfo(obj)
	return scRingSlotChanged
}

func BuildSCRingBaoKuAttend(autoFlag bool, attendType int32, obj *playerring.PlayerRingBaoKuObject, rewItemList []*droptemplate.DropItemData, extraRewList []*droptemplate.DropItemData) *uipb.SCRingBaoKuAttend {
	scRingBaoKuAttend := &uipb.SCRingBaoKuAttend{}
	typ := int32(obj.GetType())
	luckyPoints := obj.GetLuckyPoints()
	attendPoints := obj.GetAttendPoints()
	scRingBaoKuAttend.Type = &typ
	scRingBaoKuAttend.AttendType = &attendType
	scRingBaoKuAttend.AutoFlag = &autoFlag
	scRingBaoKuAttend.LuckyPoints = &luckyPoints
	scRingBaoKuAttend.AttendPoints = &attendPoints
	for i := int(0); i < len(rewItemList); i++ {
		itemId := rewItemList[i].GetItemId()
		num := rewItemList[i].GetNum()
		level := rewItemList[i].GetLevel()

		scRingBaoKuAttend.DropInfo = append(scRingBaoKuAttend.DropInfo, buildDropInfo(itemId, num, level))
	}
	for i := int(0); i < len(extraRewList); i++ {
		itemId := extraRewList[i].GetItemId()
		num := extraRewList[i].GetNum()
		level := extraRewList[i].GetLevel()

		scRingBaoKuAttend.RewDropInfo = append(scRingBaoKuAttend.RewDropInfo, buildDropInfo(itemId, num, level))
	}
	return scRingBaoKuAttend
}

func BuildSCRingLuckyPointsChange(typ int32, luckyPoints int32) *uipb.SCRingLuckyPointsChange {
	scRingLuckyPointsChange := &uipb.SCRingLuckyPointsChange{}
	scRingLuckyPointsChange.Type = &typ
	scRingLuckyPointsChange.LuckyPoints = &luckyPoints
	return scRingLuckyPointsChange
}

func buildDropInfo(itemId, num, level int32) *uipb.DropInfo {
	dropInfo := &uipb.DropInfo{}
	dropInfo.ItemId = &itemId
	dropInfo.Num = &num
	dropInfo.Level = &level

	return dropInfo
}

func BuildSCRingStrengthen(success bool, typ int32, level int32) *uipb.SCRingStrengthen {
	scRingStrengthen := &uipb.SCRingStrengthen{}
	scRingStrengthen.Type = &typ
	scRingStrengthen.StrengthenLevel = &level
	scRingStrengthen.Success = &success
	return scRingStrengthen
}

func BuildSCRingJingLing(success bool, typ int32, level int32) *uipb.SCRingJingLing {
	scRingJingLing := &uipb.SCRingJingLing{}
	scRingJingLing.Type = &typ
	scRingJingLing.JingLingLevel = &level
	scRingJingLing.Success = &success
	return scRingJingLing
}

func BuildSCRingAdvance(success bool, typ int32, advance int32, pro int32, bless int32) *uipb.SCRingAdvance {
	scRingAdvance := &uipb.SCRingAdvance{}
	scRingAdvance.Type = &typ
	scRingAdvance.Advance = &advance
	scRingAdvance.AdvancePro = &pro
	scRingAdvance.Bless = &bless
	scRingAdvance.Success = &success
	return scRingAdvance
}

func BuildSCRingFuse(success bool, isBag bool, typ int32, index int32, needIndex int32, createItemId int32, createItemIdNum int32) *uipb.SCRingFuse {
	scRingFuse := &uipb.SCRingFuse{}
	scRingFuse.IsBag = &isBag
	scRingFuse.Type = &typ
	scRingFuse.Index = &index
	scRingFuse.NeedIndex = &needIndex
	scRingFuse.CreateItemId = &createItemId
	scRingFuse.CreateItemIdNum = &createItemIdNum
	scRingFuse.Success = &success
	return scRingFuse
}

func BuildSCRingUnload(typ int32) *uipb.SCRingUnload {
	scRingUnload := &uipb.SCRingUnload{}
	scRingUnload.Type = &typ
	return scRingUnload
}

func BuildSCRingEquip(index int32) *uipb.SCRingEquip {
	scRingEquip := &uipb.SCRingEquip{}
	scRingEquip.Index = &index
	return scRingEquip
}
