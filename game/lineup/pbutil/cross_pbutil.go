package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

func BuildSILineupCancle(crossType int32, sceneId int64) *crosspb.SILineupCancle {
	siMsg := &crosspb.SILineupCancle{}
	siMsg.CrossType = &crossType
	siMsg.SceneId = &sceneId
	return siMsg
}

func BuildSILineupAttend(crossType int32, sceneId int64) *crosspb.SILineupAttend {
	siMsg := &crosspb.SILineupAttend{}
	siMsg.CrossType = &crossType
	siMsg.SceneId = &sceneId
	return siMsg
}
