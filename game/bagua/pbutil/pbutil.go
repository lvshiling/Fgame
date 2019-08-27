package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
)

func BuildSCBaGuaLevel(level int32) *uipb.SCBaGuaLevel {
	baguaLevel := &uipb.SCBaGuaLevel{}
	baguaLevel.Level = &level
	return baguaLevel
}

func BuildSCBaGuaToKill(level int32) *uipb.SCBaGuaToKill {
	baguaTokill := &uipb.SCBaGuaToKill{}
	baguaTokill.Level = &level
	return baguaTokill
}

func BuildSCBaGuaToKillResult(state bool, level int32) *uipb.SCBaGuaToKillResult {
	baguaToKillResult := &uipb.SCBaGuaToKillResult{}
	baguaToKillResult.State = &state
	baguaToKillResult.Level = &level
	return baguaToKillResult
}

func BuildSCBaGuaScene(starTime int64, level int32, ownerId int64, spouseId int64) *uipb.SCBaGuaSceneInfo {
	baguaScene := &uipb.SCBaGuaSceneInfo{}
	baguaScene.CreateTime = &starTime
	baguaScene.Level = &level
	baguaScene.OwnerId = &ownerId
	baguaScene.SpouseId = &spouseId
	return baguaScene
}

func BuildSCBaGuaPair(inviteTime int64) *uipb.SCBaGuaPair {
	baguaPair := &uipb.SCBaGuaPair{}
	baguaPair.InviteTime = &inviteTime
	return baguaPair
}

func BuildSCBaGuaPairPushSpouse(playerId int64, level int32) *uipb.SCBaGuaPairPushSpouse {
	baguaPairPushSpouse := &uipb.SCBaGuaPairPushSpouse{}
	baguaPairPushSpouse.PlayerId = &playerId
	baguaPairPushSpouse.Level = &level
	return baguaPairPushSpouse
}

func BuildSCBaGuaSpouseRefused(name string) *uipb.SCBaGuaSpouseRefused {
	baguaSpouseRefused := &uipb.SCBaGuaSpouseRefused{}
	baguaSpouseRefused.Name = &name
	return baguaSpouseRefused
}

func BuildSCBaGuaInviteOffonline(name string) *uipb.SCBaGuaInviteOffonline {
	baguaInviteOffonline := &uipb.SCBaGuaInviteOffonline{}
	baguaInviteOffonline.InviteName = &name
	return baguaInviteOffonline
}

func BuildSCBaGuaPairDeal(result int32) *uipb.SCBaGuaPairDeal {
	baguaPairDeal := &uipb.SCBaGuaPairDeal{}
	baguaPairDeal.Result = &result
	return baguaPairDeal
}

func BuildSCBaGuaPairCancle(result int32) *uipb.SCBaGuaPairCancle {
	baguaPairCancle := &uipb.SCBaGuaPairCancle{}
	baguaPairCancle.Result = &result
	return baguaPairCancle
}

func BuildSCBaGuaPairPushCancle(name string) *uipb.SCBaGuaPairPushCancle {
	baguaPairPushCancle := &uipb.SCBaGuaPairPushCancle{}
	baguaPairPushCancle.Name = &name
	return baguaPairPushCancle
}

func BuildSCBaGuaPairResult(identity bool, state bool, level int32) *uipb.SCBaGuaPairResult {
	baguaPairResult := &uipb.SCBaGuaPairResult{}
	baguaPairResult.Identity = &identity
	baguaPairResult.State = &state
	baguaPairResult.Level = &level
	return baguaPairResult
}

func BuildSCBaGuaNext(level int32) *uipb.SCBaGuaNext {
	baguaNext := &uipb.SCBaGuaNext{}
	baguaNext.Level = &level
	return baguaNext
}
