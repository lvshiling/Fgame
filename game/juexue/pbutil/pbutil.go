package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playerjx "fgame/fgame/game/juexue/player"
	juexuetypes "fgame/fgame/game/juexue/types"
)

func BuildSCJueXueGet(useId int32, jueXueMap map[juexuetypes.JueXueType]*playerjx.PlayerJueXueObject) *uipb.SCJueXueGet {
	jxGet := &uipb.SCJueXueGet{}
	jxGet.Id = &useId
	for _, jx := range jueXueMap {
		jxGet.JueXueList = append(jxGet.JueXueList, buildJueXue(jx))
	}
	return jxGet
}

func BuildSCJueXueActive(jxId int32) *uipb.SCJueXueActive {
	jxActive := &uipb.SCJueXueActive{}
	jxActive.Id = &jxId
	return jxActive
}

func BuildSCJueXueUpgrade(jxId int32) *uipb.SCJueXueUpgrade {
	jxUpgrade := &uipb.SCJueXueUpgrade{}
	jxUpgrade.Id = &jxId
	return jxUpgrade
}

func BuildSCJueXueInsight(jxId int32) *uipb.SCJueXueInsight {
	jxInsight := &uipb.SCJueXueInsight{}
	jxInsight.Id = &jxId
	return jxInsight
}

func BuildSCJueXueUse(jxId int32) *uipb.SCJueXueUse {
	jxUse := &uipb.SCJueXueUse{}
	jxUse.Id = &jxId
	return jxUse
}

func BuildSCJueXueUnload(jxId int32) *uipb.SCJueXueUnLoad {
	jxUnload := &uipb.SCJueXueUnLoad{}
	jxUnload.Id = &jxId
	return jxUnload
}

func buildJueXue(jx *playerjx.PlayerJueXueObject) *uipb.JueXueInfo {
	jxInfo := &uipb.JueXueInfo{}
	typ := int32(jx.Type)
	level := jx.Level
	insight := int32(jx.Insight)
	jxInfo.Typ = &typ
	jxInfo.Level = &level
	jxInfo.Insight = &insight
	return jxInfo
}
