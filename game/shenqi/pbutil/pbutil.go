package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playershenqi "fgame/fgame/game/shenqi/player"
	shenqitypes "fgame/fgame/game/shenqi/types"
)

func BuildSCShenQiInfoGet(qiLingMap map[shenqitypes.ShenQiType]map[shenqitypes.QiLingType]map[shenqitypes.QiLingSubType]*playershenqi.PlayerShenQiQiLingObject, debrisMap map[shenqitypes.ShenQiType]map[shenqitypes.DebrisType]*playershenqi.PlayerShenQiDebrisObject, smeltMap map[shenqitypes.ShenQiType]map[shenqitypes.SmeltType]*playershenqi.PlayerShenQiSmeltObject, num int64) *uipb.SCShenqiInfoGet {
	scMsg := &uipb.SCShenqiInfoGet{}
	for _, tempMM := range qiLingMap {
		for _, tempM := range tempMM {
			for _, slot := range tempM {
				scMsg.QiLingList = append(scMsg.QiLingList, buildShenQiQiLing(slot))
			}
		}
	}
	for _, tempM := range debrisMap {
		for _, slot := range tempM {
			scMsg.DebrisList = append(scMsg.DebrisList, buildShenQiDebris(slot))
		}
	}
	for _, tempM := range smeltMap {
		for _, slot := range tempM {
			scMsg.SmeltList = append(scMsg.SmeltList, buildShenQiSmelt(slot))
		}
	}
	scMsg.LingQiNum = &num
	return scMsg
}

func BuildSCShenQiKindInfoGet(qiLingMap map[shenqitypes.QiLingType]map[shenqitypes.QiLingSubType]*playershenqi.PlayerShenQiQiLingObject, debrisMap map[shenqitypes.DebrisType]*playershenqi.PlayerShenQiDebrisObject, smeltMap map[shenqitypes.SmeltType]*playershenqi.PlayerShenQiSmeltObject, typ shenqitypes.ShenQiType) *uipb.SCShenqiKindInfoGet {
	scMsg := &uipb.SCShenqiKindInfoGet{}
	shenQiType := int32(typ)
	scMsg.ShenQiType = &shenQiType
	for _, tempM := range qiLingMap {
		for _, slot := range tempM {
			scMsg.QiLingList = append(scMsg.QiLingList, buildShenQiQiLing(slot))
		}
	}
	for _, slot := range debrisMap {
		scMsg.DebrisList = append(scMsg.DebrisList, buildShenQiDebris(slot))
	}
	for _, slot := range smeltMap {
		scMsg.SmeltList = append(scMsg.SmeltList, buildShenQiSmelt(slot))
	}
	return scMsg
}

func BuildSCShenQiUseQiling(obj *playershenqi.PlayerShenQiQiLingObject) *uipb.SCShenqiUseQiling {
	scMsg := &uipb.SCShenqiUseQiling{}
	scMsg.QiLing = buildShenQiQiLing(obj)
	return scMsg
}

func BuildSCShenQiZhuling(obj *playershenqi.PlayerShenQiQiLingObject, num int64, auto bool) *uipb.SCShenqiZhuling {
	scMsg := &uipb.SCShenqiZhuling{}
	scMsg.QiLing = buildShenQiQiLing(obj)
	scMsg.LingQiNum = &num
	scMsg.Auto = &auto
	return scMsg
}

func BuildSCShenQiQilingResolve(num int64) *uipb.SCShenqiQilingResolve {
	scMsg := &uipb.SCShenqiQilingResolve{}
	scMsg.LingQiNum = &num
	return scMsg
}

func BuildSCShenQiDebrisUp(obj *playershenqi.PlayerShenQiDebrisObject, auto bool) *uipb.SCShenqiDebrisUp {
	scMsg := &uipb.SCShenqiDebrisUp{}
	scMsg.Debris = buildShenQiDebris(obj)
	scMsg.Auto = &auto
	return scMsg
}

func BuildSCShenQiLingQiNumChanged(lingQiNum int64) *uipb.SCShenQiLingQiNumChanged {
	scMsg := &uipb.SCShenQiLingQiNumChanged{}
	scMsg.LingQiNum = &lingQiNum
	return scMsg
}

func BuildSCShenQiSmeltUp(obj *playershenqi.PlayerShenQiSmeltObject, auto bool) *uipb.SCShenqiSmeltUp {
	scMsg := &uipb.SCShenqiSmeltUp{}
	scMsg.Smelt = buildShenQiSmelt(obj)
	scMsg.Auto = &auto
	return scMsg
}

func buildShenQiDebris(obj *playershenqi.PlayerShenQiDebrisObject) *uipb.ShenqiDebrisInfo {
	slot := &uipb.ShenqiDebrisInfo{}
	typ := int32(obj.ShenQiType)
	slotId := int32(obj.SlotId)
	level := obj.Level
	upPro := obj.UpPro
	slot.ShenQiType = &typ
	slot.SlotId = &slotId
	slot.Level = &level
	slot.UpPro = &upPro

	return slot
}

func buildShenQiSmelt(obj *playershenqi.PlayerShenQiSmeltObject) *uipb.ShenqiSmeltInfo {
	slot := &uipb.ShenqiSmeltInfo{}
	typ := int32(obj.ShenQiType)
	slotId := int32(obj.SlotId)
	level := obj.Level
	upPro := obj.UpPro
	slot.ShenQiType = &typ
	slot.SlotId = &slotId
	slot.Level = &level
	slot.UpPro = &upPro

	return slot
}

func buildShenQiQiLing(obj *playershenqi.PlayerShenQiQiLingObject) *uipb.ShenqiQilingInfo {
	slot := &uipb.ShenqiQilingInfo{}
	typ := int32(obj.ShenQiType)
	qiLingType := int32(obj.QiLingType)
	slotId := obj.SlotId.SubType()
	level := obj.Level
	upPro := obj.UpPro
	itemId := obj.ItemId
	bindInt := int32(obj.BindType)

	slot.ShenQiType = &typ
	slot.QiLingType = &qiLingType
	slot.SlotId = &slotId
	slot.Level = &level
	slot.UpPro = &upPro
	slot.ItemId = &itemId
	slot.BindType = &bindInt

	return slot
}
