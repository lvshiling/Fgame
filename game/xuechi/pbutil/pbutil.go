package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
)

func BuildSCXueChiGet(bloodLine int32, blood int64) *uipb.SCXueChiGet {
	xueChiGet := &uipb.SCXueChiGet{}
	xueChiGet.BloodLine = &bloodLine
	xueChiGet.Blood = &blood
	return xueChiGet
}

func BuildSCXueChiBloodLine(bloodLine int32) *uipb.SCXueChiBloodLine {
	xueChiBloodLine := &uipb.SCXueChiBloodLine{}
	xueChiBloodLine.BloodLine = &bloodLine
	return xueChiBloodLine
}

func BuildSCXueChiAutoBuy(itemId, itemNum int32, addBlood int64) *uipb.SCXueChiAutoBuy {
	scMsg := &uipb.SCXueChiAutoBuy{}
	scMsg.ItemId = &itemId
	scMsg.ItemNum = &itemNum
	scMsg.AddBlood = &addBlood
	return scMsg
}

func BuildSCXueChiBlood(blood int64) *uipb.SCXueChiBlood {
	xueChiBlood := &uipb.SCXueChiBlood{}
	xueChiBlood.Blood = &blood
	return xueChiBlood
}
