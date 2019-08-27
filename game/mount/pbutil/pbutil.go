package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	commontypes "fgame/fgame/game/common/types"
	mountcommon "fgame/fgame/game/mount/common"
	playermount "fgame/fgame/game/mount/player"
	mounttypes "fgame/fgame/game/mount/types"
)

func BuildSCMountGet(mountInfo *playermount.PlayerMountObject, mountOtherMap map[mounttypes.MountType]map[int32]*playermount.PlayerMountOtherObject) *uipb.SCMountGet {
	mountGet := &uipb.SCMountGet{}
	advancedId := int32(mountInfo.AdvanceId)
	mountGet.AdvancedId = &advancedId
	mountId := mountInfo.MountId
	mountGet.MountId = &mountId
	unrealLevel := mountInfo.UnrealLevel
	unrealPro := mountInfo.UnrealPro
	mountGet.UnrealLevel = &unrealLevel
	mountGet.UnrealPro = &unrealPro
	culLevel := mountInfo.CulLevel
	culPro := mountInfo.CulPro
	mountGet.CulLevel = &culLevel
	mountGet.CulPro = &culPro
	bless := mountInfo.Bless
	mountGet.Bless = &bless
	blessTime := mountInfo.BlessTime
	mountGet.BlessTime = &blessTime
	hidden := false
	if mountInfo.Hidden == 1 {
		hidden = true
	}
	mountGet.Hidden = &hidden
	for _, unrealId := range mountInfo.UnrealList {
		mountGet.UnrealList = append(mountGet.UnrealList, int32(unrealId))
	}

	for _, mountTypeOtherMap := range mountOtherMap {
		for _, mountOtherObj := range mountTypeOtherMap {
			mountGet.MountSkinList = append(mountGet.MountSkinList, buildMountOther(mountOtherObj))
		}
	}
	return mountGet
}

func buildMountOther(mountOtherObj *playermount.PlayerMountOtherObject) *uipb.MountSkinInfo {
	mountSkinInfo := &uipb.MountSkinInfo{}
	mountId := mountOtherObj.MountId
	level := mountOtherObj.Level
	pro := mountOtherObj.UpPro

	mountSkinInfo.MountId = &mountId
	mountSkinInfo.Level = &level
	mountSkinInfo.Pro = &pro
	return mountSkinInfo
}

func BuildSCMountCulDan(level int32, progress int32) *uipb.SCMountCulDan {
	mountCulDan := &uipb.SCMountCulDan{}
	mountCulDan.Level = &level
	mountCulDan.Progress = &progress
	return mountCulDan
}

func BuildSCMountUnrealDan(level int32, progress int32) *uipb.SCMountUnrealDan {
	mountUnrealDan := &uipb.SCMountUnrealDan{}
	mountUnrealDan.Level = &level
	mountUnrealDan.Progress = &progress
	return mountUnrealDan
}

func BuildSCMountUnreal(mountId int32) *uipb.SCMountUnreal {
	mountUnreal := &uipb.SCMountUnreal{}
	mountUnreal.MountId = &mountId
	return mountUnreal
}

func BuildSCMountUnload(mountId int32) *uipb.SCMountUnload {
	mountUnload := &uipb.SCMountUnload{}
	mountUnload.MountId = &mountId
	return mountUnload
}

func BuildSCMountHidden(hiddenFlag bool) *uipb.SCMountHidden {
	mountHidden := &uipb.SCMountHidden{}
	mountHidden.Hidden = &hiddenFlag
	return mountHidden
}

func BuildSCMountAdavancedFinshed(advancedId int32, mountId int32, typ commontypes.AdvancedType) *uipb.SCMountAdvanced {
	mountAdvanced := &uipb.SCMountAdvanced{}
	mountAdvanced.AdvancedId = &advancedId
	mountAdvanced.MountId = &mountId
	typeInt := int32(typ)
	mountAdvanced.AdvancedType = &typeInt
	return mountAdvanced
}

func BuildSCMountAdavanced(advancedId int32, mountId int32, bless int32, totalBless int32, blessTime int64, typ commontypes.AdvancedType, isDouble bool) *uipb.SCMountAdvanced {
	mountAdvanced := &uipb.SCMountAdvanced{}
	mountAdvanced.AdvancedId = &advancedId
	mountAdvanced.MountId = &mountId
	mountAdvanced.Bless = &bless
	mountAdvanced.BlessTime = &blessTime
	typeInt := int32(typ)
	mountAdvanced.AdvancedType = &typeInt
	mountAdvanced.IsDouble = &isDouble
	mountAdvanced.TotalBless = &totalBless
	return mountAdvanced
}

func BuildMountInfo(info *mountcommon.MountInfo) *uipb.MountInfo {
	mountInfo := &uipb.MountInfo{}
	advanceId := int32(info.AdvanceId)
	mountInfo.AdvancedId = &advanceId
	mountId := info.MountId
	mountInfo.MountId = &mountId
	unrealLevel := info.UnrealLevel
	unrealPro := info.UnrealPro
	mountInfo.UnrealLevel = &unrealLevel
	mountInfo.UnrealPro = &unrealPro
	culLevel := info.CulLevel
	culPro := info.CulPro
	mountInfo.CulLevel = &culLevel
	mountInfo.CulPro = &culPro
	for _, skinInfo := range info.SkinList {
		temp := buildMountSkinInfo(skinInfo)
		mountInfo.SkinList = append(mountInfo.SkinList, temp)
	}

	return mountInfo
}

func buildMountSkinInfo(info *mountcommon.MountSkinInfo) *uipb.MountSkinInfo {
	mountSkinInfo := &uipb.MountSkinInfo{}
	mountId := info.MountId
	level := info.Level
	pro := info.UpPro

	mountSkinInfo.MountId = &mountId
	mountSkinInfo.Level = &level
	mountSkinInfo.Pro = &pro
	return mountSkinInfo
}

func BuildSCMountUpstar(mountId int32, level int32, pro int32) *uipb.SCMountUpstar {
	mountUpstar := &uipb.SCMountUpstar{}
	mountUpstar.MountId = &mountId
	mountUpstar.Level = &level
	mountUpstar.UpPro = &pro
	return mountUpstar
}
