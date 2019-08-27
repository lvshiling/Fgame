package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
)

func BuildISLineupAttend(lineUpPos, crossType int32, sceneId int64) *crosspb.ISLineupAttend {
	isMsg := &crosspb.ISLineupAttend{}
	isMsg.BeforeNum = &lineUpPos
	isMsg.CrossType = &crossType
	isMsg.SceneId = &sceneId
	return isMsg
}

func BuildISLineupCancle(crossType int32) *crosspb.ISLineupCancle {
	isMsg := &crosspb.ISLineupCancle{}
	isMsg.CrossType = &crossType
	return isMsg
}

func BuildISLineupSuccess() *crosspb.ISLineupSuccess {
	isMsg := &crosspb.ISLineupSuccess{}
	return isMsg
}

func BuildISLineupSceneFinishToCancel(crossType int32) *crosspb.ISLineupSceneFinishToCancel {
	isMsg := &crosspb.ISLineupSceneFinishToCancel{}
	isMsg.CrossType = &crossType
	return isMsg
}
