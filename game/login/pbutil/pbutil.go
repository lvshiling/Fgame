package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	accountypes "fgame/fgame/login/types"
)

func BuildSCTestLogin(userId int64, state accountypes.RealNameState, now int64, openServerTime, mergeTime, activityOpenServerTime, activityMergeTime int64) *uipb.SCTestLogin {
	scTestLogin := &uipb.SCTestLogin{}

	scTestLogin.UserId = &userId
	stateInt := int32(state)
	scTestLogin.RealNameState = &stateInt
	scTestLogin.Now = &now
	scTestLogin.OpenServerTime = &openServerTime
	scTestLogin.MergeServerTime = &mergeTime
	scTestLogin.ActivityOpenServerTime = &activityOpenServerTime
	scTestLogin.ActivityMergeServerTime = &activityMergeTime

	return scTestLogin
}

func BuildSCLogin(userId int64, state accountypes.RealNameState, now int64, openServerTime, mergeTime, activityOpenServerTime, activityMergeTime int64, gm int32, iosVersion string, androidVersion string) *uipb.SCLogin {
	scLogin := &uipb.SCLogin{}

	scLogin.UserId = &userId
	scLogin.UserId = &userId
	stateInt := int32(state)
	scLogin.RealNameState = &stateInt
	scLogin.Now = &now
	scLogin.OpenServerTime = &openServerTime
	scLogin.MergeServerTime = &mergeTime
	scLogin.Gm = &gm
	scLogin.ActivityOpenServerTime = &activityOpenServerTime
	scLogin.ActivityMergeServerTime = &activityMergeTime
	scLogin.IosVersion = &iosVersion
	scLogin.AndroidVersion = &androidVersion

	return scLogin
}
