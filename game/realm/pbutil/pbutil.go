package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	rankentity "fgame/fgame/game/rank/entity"
	realmtypes "fgame/fgame/game/realm/types"
)

func BuildSCRealmLevel(level int32) *uipb.SCRealmLevel {
	realmLevel := &uipb.SCRealmLevel{}
	realmLevel.Level = &level
	return realmLevel
}

func BuildSCRealmToKill(level int32) *uipb.SCRealmToKill {
	realmTokill := &uipb.SCRealmToKill{}
	realmTokill.Level = &level
	return realmTokill
}

func BuildSCRealmToKillResult(state bool, level int32) *uipb.SCRealmToKillResult {
	realmToKillResult := &uipb.SCRealmToKillResult{}
	realmToKillResult.State = &state
	realmToKillResult.Level = &level
	return realmToKillResult
}

func BuildSCRealmScene(starTime int64, level int32, ownerId int64, spouseId int64) *uipb.SCRealmSceneInfo {
	realmScene := &uipb.SCRealmSceneInfo{}
	realmScene.CreateTime = &starTime
	realmScene.Level = &level
	realmScene.OwnerId = &ownerId
	realmScene.SpouseId = &spouseId
	return realmScene
}

func BuildSCRealmRankGet(dataList []*rankentity.RankCommonData, pos int32) *uipb.SCRealmRankGet {
	realmRankGet := &uipb.SCRealmRankGet{}
	realmRankGet.Pos = &pos
	for _, data := range dataList {
		level, _ := realmtypes.Resolve(data.Value)
		realmRankGet.RankList = append(realmRankGet.RankList, buildRank(data.Name, level))
	}
	return realmRankGet
}

func BuildSCRealmPair(inviteTime int64) *uipb.SCRealmPair {
	realmPair := &uipb.SCRealmPair{}
	realmPair.InviteTime = &inviteTime
	return realmPair
}

func BuildSCRealmPairPushSpouse(playerId int64, level int32) *uipb.SCRealmPairPushSpouse {
	realmPairPushSpouse := &uipb.SCRealmPairPushSpouse{}
	realmPairPushSpouse.PlayerId = &playerId
	realmPairPushSpouse.Level = &level
	return realmPairPushSpouse
}

func BuildSCRealmSpouseRefused(name string) *uipb.SCRealmSpouseRefused {
	realmSpouseRefused := &uipb.SCRealmSpouseRefused{}
	realmSpouseRefused.Name = &name
	return realmSpouseRefused
}

func BuildSCRealmInviteOffonline(name string) *uipb.SCRealmInviteOffonline {
	realmInviteOffonline := &uipb.SCRealmInviteOffonline{}
	realmInviteOffonline.InviteName = &name
	return realmInviteOffonline
}

func BuildSCRealmPairDeal(result int32) *uipb.SCRealmPairDeal {
	realmPairDeal := &uipb.SCRealmPairDeal{}
	realmPairDeal.Result = &result
	return realmPairDeal
}

func BuildSCRealmPairCancle(result int32) *uipb.SCRealmPairCancle {
	realmPairCancle := &uipb.SCRealmPairCancle{}
	realmPairCancle.Result = &result
	return realmPairCancle
}

func BuildSCRealmPairPushCancle(name string) *uipb.SCRealmPairPushCancle {
	realmPairPushCancle := &uipb.SCRealmPairPushCancle{}
	realmPairPushCancle.Name = &name
	return realmPairPushCancle
}

func BuildSCRealmPairResult(identity bool, state bool, level int32) *uipb.SCRealmPairResult {
	realmPairResult := &uipb.SCRealmPairResult{}
	realmPairResult.Identity = &identity
	realmPairResult.State = &state
	realmPairResult.Level = &level
	return realmPairResult
}

func BuildSCRealmNext(level int32) *uipb.SCRealmNext {
	realmNext := &uipb.SCRealmNext{}
	realmNext.Level = &level
	return realmNext
}

func buildRank(name string, level int32) *uipb.RealmRank {
	realmRank := &uipb.RealmRank{}
	realmRank.Name = &name
	realmRank.Level = &level
	return realmRank
}
