package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playertitle "fgame/fgame/game/title/player"
)

func BuildSCTitleGet(pl player.Player, titleWear int32, titleIdMap map[int32]*playertitle.TitleData) *uipb.SCTitleGet {
	titleGet := &uipb.SCTitleGet{}
	titleGet.TitleWear = &titleWear
	titleGet.TitleList = buildTitleList(pl, titleIdMap)
	return titleGet
}

func BuildSCTitleActive(title int32, activeTime, validTime int64) *uipb.SCTitleActive {
	titleActive := &uipb.SCTitleActive{}
	titleActive.TitleId = &title
	titleActive.ActiveTime = &activeTime
	titleActive.ValidTime = &validTime
	return titleActive
}

func BuildSCTitleAddValidTime(titleId int32, validTime int64) *uipb.SCTitleAddValidTime {
	scMsg := &uipb.SCTitleAddValidTime{}
	scMsg.TitleId = &titleId
	scMsg.ValidTime = &validTime
	return scMsg
}

func BuildSCTitleWear(title int32) *uipb.SCTitleWear {
	titleWear := &uipb.SCTitleWear{}
	titleWear.TitleId = &title
	return titleWear
}

func BuildSCTitleUnload(titleWear int32) *uipb.SCTitleUnLoad {
	titleUnload := &uipb.SCTitleUnLoad{}
	titleUnload.TitleWear = &titleWear
	return titleUnload
}

func BuildSCTitleAdd(titleId int32) *uipb.SCTitleAdd {
	titleAdd := &uipb.SCTitleAdd{}
	titleAdd.TitleId = &titleId
	return titleAdd
}

func BuildSCTitleRemove(titleId int32) *uipb.SCTitleRemove {
	titleRemove := &uipb.SCTitleRemove{}
	titleRemove.TitleId = &titleId
	return titleRemove
}

func BuildSCTitleUpstar(titleId int32, starLev int32, bless int32) *uipb.SCTitleUpStar {
	scMsg := &uipb.SCTitleUpStar{}
	scMsg.TitleId = &titleId
	scMsg.StarLev = &starLev
	scMsg.Bless = &bless
	return scMsg
}

func buildTitleList(pl player.Player, titleMap map[int32]*playertitle.TitleData) (infoList []*uipb.TitleInfo) {
	for titleId, data := range titleMap {
		activeTime := data.GetActiveTime()
		validTime := data.GetValidTime()

		// 获得称号星级
		starLev := int32(0)
		manager := pl.GetPlayerDataManager(types.PlayerTitleDataManagerType).(*playertitle.PlayerTitleDataManager)
		obj := manager.GetTitleObjectById(titleId)
		if obj != nil {
			starLev = obj.GetStarLev()
		}

		infoList = append(infoList, buildTitle(titleId, activeTime, validTime, starLev))
	}
	return infoList
}

func buildTitle(titleId int32, activeTime, validTime int64, starLev int32) *uipb.TitleInfo {
	titleInfo := &uipb.TitleInfo{}
	titleInfo.TitleId = &titleId
	titleInfo.AcitveTime = &activeTime
	titleInfo.ValidTime = &validTime
	titleInfo.StarLev = &starLev
	return titleInfo
}
