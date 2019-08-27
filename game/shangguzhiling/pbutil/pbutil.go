package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droptemplate "fgame/fgame/game/drop/template"
	playershangguzhiling "fgame/fgame/game/shangguzhiling/player"
	shangguzhilingtypes "fgame/fgame/game/shangguzhiling/types"
)

func buildLingShouInfo(obj *playershangguzhiling.PlayerShangguzhilingObject, lingwenList []shangguzhilingtypes.LingwenType, posList []shangguzhilingtypes.LinglianPosType) *uipb.LingShouInfo {
	typ := int32(obj.GetLingShouType())
	level := obj.GetLevel()
	exp := obj.GetExperience()
	rank := obj.GetUprankLevel()
	bless := obj.GetUprankBless()
	times := obj.GetUprankTimes()
	lingwen := []*uipb.LingWenInfo{}
	linglian := []*uipb.LingLianInfo{}
	receiveTime := obj.GetLastReceiveTime()
	for _, subType := range lingwenList {
		info := obj.GetLingWenInfo(subType)
		lingwen = append(lingwen, buildLingWenInfo(subType, info))
	}
	for _, subType := range posList {
		info := obj.GetLingLianInfo(subType)
		linglian = append(linglian, buildLingLianInfo(subType, info))
	}
	linglianTimes := obj.GetLingLianTimes()

	scMsg := &uipb.LingShouInfo{
		Type:          &typ,
		Level:         &level,
		Exp:           &exp,
		Rank:          &rank,
		UprankBless:   &bless,
		UprankTimes:   &times,
		LingwenList:   lingwen,
		LinglianList:  linglian,
		LinglianTimes: &linglianTimes,
		ReceiveTime:   &receiveTime,
	}
	return scMsg
}

func buildLingWenInfo(subType shangguzhilingtypes.LingwenType, info *shangguzhilingtypes.LingwenInfo) *uipb.LingWenInfo {
	subTyp := int32(subType)
	level := info.Level
	exp := info.Experience

	scMsg := &uipb.LingWenInfo{
		SubType: &subTyp,
		Level:   &level,
		Exp:     &exp,
	}
	return scMsg
}

func buildLingLianInfo(subType shangguzhilingtypes.LinglianPosType, info *shangguzhilingtypes.LinglianInfo) *uipb.LingLianInfo {
	subTyp := int32(subType)
	mark := info.PoolMark
	// isLock := info.IsLock
	// lockTimes := info.LockTimes

	scMsg := &uipb.LingLianInfo{
		SubType:  &subTyp,
		PoolMark: &mark,
		// IsLock:    &isLock,
		// LockTimes: &lockTimes,
	}
	return scMsg
}

func BuildSCShangguzhilingInfo(objList []*playershangguzhiling.PlayerShangguzhilingObject, unlockLingwenMap map[shangguzhilingtypes.LingshouType][]shangguzhilingtypes.LingwenType, jiesuoLinglianPosMap map[shangguzhilingtypes.LingshouType][]shangguzhilingtypes.LinglianPosType) *uipb.SCShangguzhilingInfo {
	infoList := []*uipb.LingShouInfo{}
	for _, obj := range objList {
		infoList = append(infoList, buildLingShouInfo(obj, unlockLingwenMap[obj.GetLingShouType()], jiesuoLinglianPosMap[obj.GetLingShouType()]))
	}
	scMsg := &uipb.SCShangguzhilingInfo{
		LingshouList: infoList,
	}
	return scMsg
}

func BuildSCShangguzhilingUplevel(typ shangguzhilingtypes.LingshouType, level int32, exp int64) *uipb.SCShangguzhilingUplevel {
	typint := int32(typ)
	scMsg := &uipb.SCShangguzhilingUplevel{
		Type:  &typint,
		Level: &level,
		Exp:   &exp,
	}
	return scMsg
}

func BuildSCShangguzhilingLingWenUplevel(typ shangguzhilingtypes.LingshouType, subType shangguzhilingtypes.LingwenType, info *shangguzhilingtypes.LingwenInfo) *uipb.SCShangguzhilingLingWenUplevel {
	typint := int32(typ)
	scInfo := buildLingWenInfo(subType, info)
	scMsg := &uipb.SCShangguzhilingLingWenUplevel{
		Type:    &typint,
		Lingwen: scInfo,
	}
	return scMsg
}

func BuildSCShangguzhilingUpRank(typ shangguzhilingtypes.LingshouType, rank int32, bless int64, times int32) *uipb.SCShangguzhilingUpRank {
	typint := int32(typ)
	scMsg := &uipb.SCShangguzhilingUpRank{
		Type:        &typint,
		Rank:        &rank,
		UprankBless: &bless,
		UprankTimes: &times,
	}
	return scMsg
}

func BuildSCShangguzhilingLingLian(typ shangguzhilingtypes.LingshouType, linglianMap map[shangguzhilingtypes.LinglianPosType]*shangguzhilingtypes.LinglianInfo, linglianTimes int32) *uipb.SCShangguzhilingLingLian {
	typint := int32(typ)
	linglian := []*uipb.LingLianInfo{}
	for subType, info := range linglianMap {
		linglian = append(linglian, buildLingLianInfo(subType, info))
	}
	scMsg := &uipb.SCShangguzhilingLingLian{
		Type:          &typint,
		Linglian:      linglian,
		LinglianTimes: &linglianTimes,
	}
	return scMsg
}

func BuildSCShangguzhilingReceive(typ shangguzhilingtypes.LingshouType, receiveTime int64, itemData *droptemplate.DropItemData) *uipb.SCShangguzhilingReceive {
	typint := int32(typ)
	itemId := itemData.GetItemId()
	num := itemData.GetNum()
	level := itemData.GetLevel()
	scMsg := &uipb.SCShangguzhilingReceive{
		Type:        &typint,
		ReceiveTime: &receiveTime,
		DropInfo:    buildDropInfo(itemId, num, level),
	}
	return scMsg
}

func buildDropInfo(itemId, num, level int32) *uipb.DropInfo {
	dropInfo := &uipb.DropInfo{}
	dropInfo.ItemId = &itemId
	dropInfo.Num = &num
	dropInfo.Level = &level

	return dropInfo
}
