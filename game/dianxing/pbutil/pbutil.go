package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	commontypes "fgame/fgame/game/common/types"
	dianxingcommon "fgame/fgame/game/dianxing/common"
	playerdianxing "fgame/fgame/game/dianxing/player"
)

func BuildSCDianXingGet(info *playerdianxing.PlayerDianXingObject) *uipb.SCDianxingGet {
	scDianXingGet := &uipb.SCDianxingGet{}
	xingChenNum := info.XingChenNum
	scDianXingGet.DianXingInfo = buildDianXingInfo(info)
	scDianXingGet.JieFengInfo = buildJieFengInfo(info)
	scDianXingGet.XingChenNum = &xingChenNum
	return scDianXingGet
}

func BuildSCDianXingAdavancedFinshed(info *playerdianxing.PlayerDianXingObject, typ commontypes.AdvancedType) *uipb.SCDianxingAdvanced {
	dianxingAdvanced := &uipb.SCDianxingAdvanced{}
	xingChenNum := info.XingChenNum
	dianxingAdvanced.DianXingInfo = buildDianXingInfo(info)
	dianxingAdvanced.XingChenNum = &xingChenNum
	typInt := int32(typ)
	dianxingAdvanced.AdvancedType = &typInt
	return dianxingAdvanced
}

func BuildSCDianXingAdavanced(info *playerdianxing.PlayerDianXingObject, typ commontypes.AdvancedType, isDouble bool, isAutoBuy bool, isFu bool) *uipb.SCDianxingAdvanced {
	dianxingAdvanced := &uipb.SCDianxingAdvanced{}
	xingChenNum := info.XingChenNum
	dianxingAdvanced.DianXingInfo = buildDianXingInfo(info)
	dianxingAdvanced.XingChenNum = &xingChenNum
	typInt := int32(typ)
	dianxingAdvanced.AdvancedType = &typInt
	dianxingAdvanced.IsDouble = &isDouble
	dianxingAdvanced.BuyFlag = &isAutoBuy
	dianxingAdvanced.FuFlag = &isFu
	return dianxingAdvanced
}

func BuildSCDianXingJieFengAdavancedFinshed(info *playerdianxing.PlayerDianXingObject, typ commontypes.AdvancedType) *uipb.SCDianxingJiefengAdvanced {
	dianxingJiefengAdvanced := &uipb.SCDianxingJiefengAdvanced{}
	dianxingJiefengAdvanced.JieFengInfo = buildJieFengInfo(info)
	typInt := int32(typ)
	dianxingJiefengAdvanced.AdvancedType = &typInt
	return dianxingJiefengAdvanced
}

func BuildSCDianXingJieFengAdavanced(info *playerdianxing.PlayerDianXingObject, typ commontypes.AdvancedType, isDouble bool, isAutoBuy bool) *uipb.SCDianxingJiefengAdvanced {
	dianxingJiefengAdvanced := &uipb.SCDianxingJiefengAdvanced{}
	dianxingJiefengAdvanced.JieFengInfo = buildJieFengInfo(info)
	typInt := int32(typ)
	dianxingJiefengAdvanced.AdvancedType = &typInt
	dianxingJiefengAdvanced.IsDouble = &isDouble
	dianxingJiefengAdvanced.BuyFlag = &isAutoBuy
	return dianxingJiefengAdvanced
}

func BuildSCDianxingXingchenVary(num int64) *uipb.SCDianxingXingchenVary {
	scMsg := &uipb.SCDianxingXingchenVary{}
	scMsg.XingChenNum = &num

	return scMsg
}

func buildDianXingInfo(info *playerdianxing.PlayerDianXingObject) *uipb.DianXingInfo {
	dianXingInfo := &uipb.DianXingInfo{}

	currType := info.CurrType
	currLevel := info.CurrLevel
	dianXingBless := info.DianXingBless
	dianXingBlessTime := info.DianXingBlessTime

	dianXingInfo.XingPu = &currType
	dianXingInfo.Level = &currLevel
	dianXingInfo.Bless = &dianXingBless
	dianXingInfo.BlessTime = &dianXingBlessTime
	return dianXingInfo
}

func buildJieFengInfo(info *playerdianxing.PlayerDianXingObject) *uipb.JieFengInfo {
	jieFengInfo := &uipb.JieFengInfo{}

	jieFengLev := info.JieFengLev
	jieFengBless := info.JieFengBless

	jieFengInfo.Level = &jieFengLev
	jieFengInfo.Bless = &jieFengBless
	return jieFengInfo
}

func BuildDianXingCacheInfo(info *dianxingcommon.DianXingInfo) *uipb.DianXingCacheInfo {
	dianXingCacheInfo := &uipb.DianXingCacheInfo{}
	xingPu := info.CurrType
	level := info.CurrLevel
	jieFengLev := info.JieFengLev
	jieFengBless := info.JieFengBless

	dianXingCacheInfo.XingPu = &xingPu
	dianXingCacheInfo.Level = &level
	dianXingCacheInfo.JieFengLev = &jieFengLev
	dianXingCacheInfo.JieFengBless = &jieFengBless

	return dianXingCacheInfo
}
