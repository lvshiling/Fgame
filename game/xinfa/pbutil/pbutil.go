package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playerxinfa "fgame/fgame/game/xinfa/player"
	xinfatypes "fgame/fgame/game/xinfa/types"
)

func BuildSCXinFaGet(xinfaMap map[xinfatypes.XinFaType]*playerxinfa.PlayerXinFaObject) *uipb.SCXinFaGet {
	xinfaGet := &uipb.SCXinFaGet{}
	for _, jx := range xinfaMap {
		xinfaGet.XinFaList = append(xinfaGet.XinFaList, buildXinFa(jx))
	}
	return xinfaGet
}

func BuildSCXinFaActive(xfId int32) *uipb.SCXinFaActive {
	xinfaActive := &uipb.SCXinFaActive{}
	xinfaActive.Id = &xfId
	return xinfaActive
}

func BuildSCXinFaUpgrade(xfId int32) *uipb.SCXinFaUpgrade {
	xinfaUpgrade := &uipb.SCXinFaUpgrade{}
	xinfaUpgrade.Id = &xfId
	return xinfaUpgrade
}

func buildXinFa(xinfa *playerxinfa.PlayerXinFaObject) *uipb.XinFaInfo {
	xinfaInfo := &uipb.XinFaInfo{}
	typ := int32(xinfa.Type)
	level := xinfa.Level
	xinfaInfo.Typ = &typ
	xinfaInfo.Level = &level
	return xinfaInfo
}
