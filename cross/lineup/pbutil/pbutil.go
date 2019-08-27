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

func BuildSCLineupSceneFinishToCancel(crossType int32) *uipb.SCLineupSceneFinishToCancel {
	scMsg := &uipb.SCLineupSceneFinishToCancel{}
	scMsg.CrossType = &crossType
	return scMsg
}
 