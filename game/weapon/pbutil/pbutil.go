package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playerweapon "fgame/fgame/game/weapon/player"
	weapontypes "fgame/fgame/game/weapon/types"
)

func BuildSCWeaponGet(weaponWear int32, weaponList map[int32]*playerweapon.PlayerWeaponObject) *uipb.SCWeaponGet {
	weaponGet := &uipb.SCWeaponGet{}
	weaponGet.WeaponWear = &weaponWear
	weaponGet.WeaponList = buildWeaponList(weaponList)
	return weaponGet
}

func BuildSCWeaponActive(weaponId int32) *uipb.SCWeaponActive {
	weaponActive := &uipb.SCWeaponActive{}
	weaponActive.WeaponId = &weaponId
	return weaponActive
}

func BuildSCWeaponEatDan(weaponId int32, culLevel int32, culPro int32) *uipb.SCWeaponEatDan {
	eatDan := &uipb.SCWeaponEatDan{}
	eatDan.WeaponId = &weaponId
	eatDan.CulLevel = &culLevel
	eatDan.CulPro = &culPro
	return eatDan
}

func BuildSCWeaponUpstar(weaponId int32, level int32, pro int32) *uipb.SCWeaponUpstar {
	upStar := &uipb.SCWeaponUpstar{}
	upStar.WeaponId = &weaponId
	upStar.Level = &level
	upStar.UpPro = &pro
	return upStar
}

func BuildSCWeaponAwaken(weaponId int32, sucess bool) *uipb.SCWeaponAwaken {
	awaken := &uipb.SCWeaponAwaken{}
	result := int32(0)
	if sucess {
		result = 1
	}
	awaken.Result = &result
	awaken.WeaponId = &weaponId
	return awaken
}

func BuildSCWeaponWear(weaponId int32) *uipb.SCWeaponWear {
	wear := &uipb.SCWeaponWear{}
	wear.WeaponWear = &weaponId
	return wear
}

func BuildSCWeaponUnload(weaponId int32) *uipb.SCWeaponUnLoad {
	unLoad := &uipb.SCWeaponUnLoad{}
	unLoad.WeaponWear = &weaponId
	return unLoad
}

func buildWeaponList(weapons map[int32]*playerweapon.PlayerWeaponObject) (weaponList []*uipb.WeaponInfo) {

	for _, weapon := range weapons {
		if weapon.ActiveFlag == 0 {
			continue
		}
		weaponList = append(weaponList, buildWeapon(weapon))
	}
	return weaponList
}

func buildWeapon(weapon *playerweapon.PlayerWeaponObject) *uipb.WeaponInfo {
	weaponInfo := &uipb.WeaponInfo{}
	weaponId := weapon.WeaponId
	level := weapon.Level
	culLevel := weapon.CulLevel
	culPro := weapon.CulPro
	state := int32(weapon.State)

	weaponInfo.WeaponId = &weaponId
	weaponInfo.Level = &level
	weaponInfo.CulLevel = &culLevel
	weaponInfo.CulPro = &culPro
	weaponInfo.State = &state

	return weaponInfo
}

func BuildAllWeaponInfo(info *weapontypes.AllWeaponInfo) *uipb.AllWeaponInfo {
	allWeaponInfo := &uipb.AllWeaponInfo{}
	weaponWear := info.Wear
	allWeaponInfo.WeaponWear = &weaponWear
	for _, tempInfo := range info.WeaponList {
		allWeaponInfo.WeaponList = append(allWeaponInfo.WeaponList, BuildWeaponInfo(tempInfo))
	}
	return allWeaponInfo
}

func BuildWeaponInfo(info *weapontypes.WeaponInfo) *uipb.WeaponInfo {
	weaponInfo := &uipb.WeaponInfo{}
	weaponId := info.WeaponId
	level := info.Level
	culLevel := info.CulLevel
	culPro := info.CulPro
	state := info.State

	weaponInfo.WeaponId = &weaponId
	weaponInfo.Level = &level
	weaponInfo.CulLevel = &culLevel
	weaponInfo.CulPro = &culPro
	weaponInfo.State = &state

	return weaponInfo
}
