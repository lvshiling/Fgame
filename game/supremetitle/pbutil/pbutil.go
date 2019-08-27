package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playersupremetitle "fgame/fgame/game/supremetitle/player"
)

func BuildSCSupremeTitleGet(titleWear int32, titleMap map[int32]*playersupremetitle.PlayerSupremeTitleObject) *uipb.SCSupremeTitleGet {
	scSupremeTitleGet := &uipb.SCSupremeTitleGet{}
	scSupremeTitleGet.TitleWear = &titleWear
	for titleId, _ := range titleMap {
		scSupremeTitleGet.TitleList = append(scSupremeTitleGet.TitleList, titleId)
	}
	return scSupremeTitleGet
}

func BuildSCSupremeTitleActive(title int32) *uipb.SCSupremeTitleActive {
	titleActive := &uipb.SCSupremeTitleActive{}
	titleActive.TitleId = &title
	return titleActive
}

func BuildSCSupremeTitleWear(title int32) *uipb.SCSupremeTitleWear {
	titleWear := &uipb.SCSupremeTitleWear{}
	titleWear.TitleId = &title
	return titleWear
}

func BuildSCSupremeTitleUnload(titleWear int32) *uipb.SCSupremeTitleUnLoad {
	titleUnload := &uipb.SCSupremeTitleUnLoad{}
	titleUnload.TitleWear = &titleWear
	return titleUnload
}
