package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
)

func BuildSCPlayerZhenXiBossInfo(enterTimes int32) *uipb.SCPlayerZhenXiBossInfo {
	scPlayerZhenXiBossInfo := &uipb.SCPlayerZhenXiBossInfo{}
	scPlayerZhenXiBossInfo.EnterTimes = &enterTimes
	return scPlayerZhenXiBossInfo
}
