package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playerfushi "fgame/fgame/game/fushi/player"
	fushitypes "fgame/fgame/game/fushi/types"
)

func BuildSCFuShiInfo(fushiMap map[fushitypes.FuShiType]*playerfushi.PlayerFuShiObject) *uipb.SCFushiInfo {
	scMsg := &uipb.SCFushiInfo{}
	for _, fushi := range fushiMap {
		typ := int32(fushi.GetType())
		level := fushi.GetFushiLevel()
		scMsg.FushiList = append(scMsg.FushiList, buildFushiInfo(typ, level))
	}

	return scMsg
}

func BuildSCFuShiActivite(typ int32) *uipb.SCFuShiActivite {
	scMsg := &uipb.SCFuShiActivite{}
	scMsg.Typ = &typ
	return scMsg
}

func BuildSCFuShiUpLevel(typ int32, level int32) *uipb.SCFuShiUplevel {
	scMsg := &uipb.SCFuShiUplevel{}
	scMsg.Typ = &typ
	scMsg.Level = &level
	return scMsg
}

func buildFushiInfo(typ int32, level int32) *uipb.FushiInfo {
	fushiInfo := &uipb.FushiInfo{}
	fushiInfo.Typ = &typ
	fushiInfo.Level = &level
	return fushiInfo
}
