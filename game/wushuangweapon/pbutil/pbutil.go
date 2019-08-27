package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playerwushuangweapon "fgame/fgame/game/wushuangweapon/player"
	wushuangweapontypes "fgame/fgame/game/wushuangweapon/types"
)

func BuildSCWushuangWeaponDevouring(bodyPos wushuangweapontypes.WushuangWeaponPart, level int32, experience int64) *uipb.SCWushuangWeaponDevouring {
	bp := uipb.WushuangWeaponBodyPos(bodyPos)
	scMsg := &uipb.SCWushuangWeaponDevouring{
		BodyPos:    &bp,
		Level:      &level,
		Experience: &experience,
	}
	return scMsg
}

func BuildSCWushuangWeaponBreakthrough(bodyPos wushuangweapontypes.WushuangWeaponPart, level int32) *uipb.SCWushuangWeaponBreakthrough {
	bp := uipb.WushuangWeaponBodyPos(bodyPos)
	scMsg := &uipb.SCWushuangWeaponBreakthrough{
		BodyPos: &bp,
		Level:   &level,
	}
	return scMsg
}

func BuildSCWushuangWeaponPutOn(bodyPos wushuangweapontypes.WushuangWeaponPart, level int32, ex int64) *uipb.SCWushuangWeaponPutOn {
	bp := uipb.WushuangWeaponBodyPos(bodyPos)
	scMsg := &uipb.SCWushuangWeaponPutOn{
		BodyPos:    &bp,
		Level:      &level,
		Experience: &ex,
	}
	return scMsg
}

func BuildSCWushuangWeaponTakeOff() *uipb.SCWushuangWeaponTakeOff {
	scMsg := &uipb.SCWushuangWeaponTakeOff{}
	return scMsg
}

// func BuildSCWushuangWeaponReddot(allBodyPos []wushuangweapontypes.WushuangWeaponPart) *uipb.SCWushuangWeaponReddot {
// 	all := make([]uipb.WushuangWeaponBodyPos, 0, len(allBodyPos))
// 	for _, a := range allBodyPos {
// 		ab := uipb.WushuangWeaponBodyPos(a)
// 		all = append(all, ab)
// 	}
// 	scMsg := &uipb.SCWushuangWeaponReddot{
// 		AllBodyPos: all,
// 	}
// 	return scMsg
// }

func BuildSCWushuangWeaponInfo(allBodyPosInfo []*playerwushuangweapon.PlayerWushuangWeaponSlotObject) *uipb.SCWushuangWeaponInfo {
	allBpInfo := make([]*uipb.WushuangWeaponBodyPosInfo, 0, len(allBodyPosInfo))
	for _, bodyPosInfo := range allBodyPosInfo {
		bpinfo := BuildWushuangWeaponBodyPosInfo(bodyPosInfo)
		allBpInfo = append(allBpInfo, bpinfo)
	}
	scMsg := &uipb.SCWushuangWeaponInfo{
		AllBodyPosInfo: allBpInfo,
	}
	return scMsg
}

func BuildWushuangWeaponBodyPosInfo(bpinfo *playerwushuangweapon.PlayerWushuangWeaponSlotObject) *uipb.WushuangWeaponBodyPosInfo {
	bp := uipb.WushuangWeaponBodyPos(bpinfo.GetBodyPart())
	itemId := bpinfo.GetItemId()
	level := bpinfo.GetLevel()
	ex := bpinfo.GetExperience()
	scMsg := &uipb.WushuangWeaponBodyPosInfo{
		BodyPos:    &bp,
		ItemId:     &itemId,
		Level:      &level,
		Experience: &ex,
	}
	return scMsg
}

func BuildWushuangBodyPosList(slotList []*wushuangweapontypes.WushuangInfo) (es []*uipb.WushuangWeaponBodyPosInfo) {
	for _, slot := range slotList {
		es = append(es, buildWushuangBodyPos(slot))
	}
	return
}

func buildWushuangBodyPos(slot *wushuangweapontypes.WushuangInfo) *uipb.WushuangWeaponBodyPosInfo {
	info := &uipb.WushuangWeaponBodyPosInfo{}
	bp := uipb.WushuangWeaponBodyPos(slot.BodyPos)
	lv := slot.Level
	it := slot.ItemId
	exp := slot.Exp
	info.BodyPos = &bp
	info.Level = &lv
	info.ItemId = &it
	info.Experience = &exp
	return info
}
