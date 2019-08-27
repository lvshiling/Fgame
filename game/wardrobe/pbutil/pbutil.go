package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
)

func BuildSCWardrobeGet(pl player.Player) *uipb.SCWardrobeGet {
	wardrobeGet := &uipb.SCWardrobeGet{}
	manager := pl.GetPlayerDataManager(types.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	wardrobeMap := manager.GetWardrobeMap()
	for typ, suitMap := range wardrobeMap {
		num := manager.GetWardrobePeiYangNum(typ)
		wardrobeGet.WardrobeList = append(wardrobeGet.WardrobeList, buildWardrobeInfo(int32(typ), num, suitMap))
	}
	return wardrobeGet
}

func buildWardrobeInfo(typ int32, num int32, suitMap map[int32]*playerwardrobe.PlayerWardrobeObject) *uipb.WardrobeInfo {
	wardrobeInfo := &uipb.WardrobeInfo{}
	wardrobeInfo.Type = &typ
	wardrobeInfo.PeiYangNum = &num

	for subType, wardrobeObj := range suitMap {
		if !wardrobeObj.GetIsActive() {
			continue
		}
		wardrobeInfo.SubTypeList = append(wardrobeInfo.SubTypeList, subType)
	}
	return wardrobeInfo
}

func BuildSCWardrobeActive(wardrobeObj *playerwardrobe.PlayerWardrobeObject) *uipb.SCWardrobeActive {
	wardrobeActive := &uipb.SCWardrobeActive{}
	typ := int32(wardrobeObj.GetType())
	subType := wardrobeObj.GetSubType()
	wardrobeActive.SuitInfo = buildWardrobe(typ, subType)
	return wardrobeActive
}

func BuildSCWardrobeRemove(typ int32, subType int32) *uipb.SCWardrobeRemove {
	wardrobeRemove := &uipb.SCWardrobeRemove{}
	wardrobeRemove.SuitInfo = buildWardrobe(typ, subType)
	return wardrobeRemove
}

func BuildSCWardrobePeiYang(typ int32, peiYangNum int32) *uipb.SCWardrobePeiYang {
	wardrobePeiYang := &uipb.SCWardrobePeiYang{}
	wardrobePeiYang.Type = &typ
	wardrobePeiYang.PeiYangNum = &peiYangNum
	return wardrobePeiYang
}

func buildWardrobe(typ int32, subType int32) *uipb.WardrobeSuitInfo {
	wardrobeSuitInfo := &uipb.WardrobeSuitInfo{}
	wardrobeSuitInfo.Type = &typ
	wardrobeSuitInfo.SubType = &subType
	return wardrobeSuitInfo
}
