package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
)

func BuildSCLineupNotice(beforeNum, crossType int32) *uipb.SCLineupNotice {
	scMsg := &uipb.SCLineupNotice{}
	scMsg.BeforeNum = &beforeNum
	scMsg.CrossType = &crossType
	return scMsg
}

func BuildSCLineupCancel(crossType int32) *uipb.SCLineupCancel {
	scMsg := &uipb.SCLineupCancel{}
	scMsg.CrossType = &crossType
	return scMsg
}

func BuildSCLineupSuccess(crossType int32) *uipb.SCLineupSuccess {
	scMsg := &uipb.SCLineupSuccess{}
	scMsg.CrossType = &crossType
	return scMsg
}
