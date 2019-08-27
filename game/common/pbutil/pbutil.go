package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/exception"
	"fgame/fgame/core/types"
)

func BuildSCGetTime(now int64) *uipb.SCGetTime {
	scGetTime := &uipb.SCGetTime{}
	scGetTime.Now = &now
	return scGetTime
}

func BuildSCOpenServerTime(openTime, activityOpenTime int64) *uipb.SCOpenServerTime {
	scOpenServerTime := &uipb.SCOpenServerTime{}
	scOpenServerTime.OpenServerTime = &openTime
	scOpenServerTime.ActivityOpenServerTime = &activityOpenTime
	return scOpenServerTime
}

func BuildSCMergeServerTime(mergeTime, activityMergeTime int64) *uipb.SCMergeServerTime {
	scMergeServerTime := &uipb.SCMergeServerTime{}
	scMergeServerTime.MergeServerTime = &mergeTime
	scMergeServerTime.ActivityMergeServerTime = &activityMergeTime
	return scMergeServerTime
}

var (
	scHeartBeat = &uipb.SCHeartBeat{}
)

func BuildSCHeartBeat() *uipb.SCHeartBeat {

	return scHeartBeat
}

func BuildSCSystemMessage(content string, args ...string) *uipb.SCSystemMessage {
	scSystemMessage := &uipb.SCSystemMessage{}
	scSystemMessage.Content = &content
	scSystemMessage.Args = args
	return scSystemMessage
}

func BuildSCException(content string, code exception.ExceptionCode) *uipb.SCException {
	scException := &uipb.SCException{}
	codeInt := int32(code)
	scException.Code = &codeInt
	scException.Content = &content
	return scException
}

func BuildPos(pos types.Position) *uipb.Position {
	targetPosition := &uipb.Position{}
	x := float32(pos.X)
	y := float32(pos.Y)
	z := float32(pos.Z)
	targetPosition.PosX = &x
	targetPosition.PosY = &y
	targetPosition.PosZ = &z

	return targetPosition
}

// 通用取消排队
func BuildSILineupCancles() *crosspb.SILineupCancle {
	siMsg := &crosspb.SILineupCancle{}
	return siMsg
}
