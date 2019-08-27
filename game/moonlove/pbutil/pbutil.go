package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/core/types"
	"fgame/fgame/game/friend/friend"
	moonlovetypes "fgame/fgame/game/moonlove/types"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/scene/scene"
)

const (
	rankLimit = 10
)

func BuildSCMoonloveCharmChanged(playerId int64, charmNum, charmRank int32) *uipb.SCMoonloveCharmChanged {
	scMoonloveCharmChanged := &uipb.SCMoonloveCharmChanged{}

	scMoonloveCharmChanged.PlayerId = &playerId
	scMoonloveCharmChanged.CharmNum = &charmNum
	scMoonloveCharmChanged.CharmRank = &charmRank

	return scMoonloveCharmChanged
}

func BuildSCMoonloveGenerousChanged(playerId int64, generousNum, generousRank int32) *uipb.SCMoonloveGenerousChanged {
	scMoonloveGenerousChanged := &uipb.SCMoonloveGenerousChanged{}

	scMoonloveGenerousChanged.PlayerId = &playerId
	scMoonloveGenerousChanged.GenerousNum = &generousNum
	scMoonloveGenerousChanged.GenerousRank = &generousRank

	return scMoonloveGenerousChanged
}

func BuildMoonloveSceneResult(playerId int64, exp int64) *uipb.SCMoonloveSceneResult {
	scMoonloveSceneResult := &uipb.SCMoonloveSceneResult{}

	scMoonloveSceneResult.PlayerId = &playerId
	scMoonloveSceneResult.RewExp = &exp

	return scMoonloveSceneResult
}

func BuildMoonloveSceneInfo(charmArr []*moonlovetypes.RankData, generousArr []*moonlovetypes.RankData, expCount, startTime int64, charmNum, generousNum int32, playerId int64) *uipb.SCMoonloveSceneInfo {
	scMoonloveSceneInfo := &uipb.SCMoonloveSceneInfo{}

	charmRank := int32(0)
	charmRankArr := make([]*uipb.MoonloveRank, 0, rankLimit)
	for index, charm := range charmArr {
		charmRankArr = append(charmRankArr, buildMoonloveCharmRank(charm))
		if charm.PlayerId == playerId {
			charmRank = int32(index) + 1
		}
	}

	generousRank := int32(0)
	generousRankArr := make([]*uipb.MoonloveRank, 0, rankLimit)
	for index, generous := range generousArr {
		generousRankArr = append(generousRankArr, buildMoonloveGenerousRank(generous))
		if generous.PlayerId == playerId {
			generousRank = int32(index) + 1
		}
	}

	moonloveInfo := &uipb.PlayerMoonloveInfo{}
	moonloveInfo.CharmNum = &charmNum
	moonloveInfo.GenerousNum = &generousNum
	moonloveInfo.CharmRank = &charmRank
	moonloveInfo.GenerousRank = &generousRank

	scMoonloveSceneInfo.CharmRank = charmRankArr
	scMoonloveSceneInfo.GenerousRank = generousRankArr
	scMoonloveSceneInfo.MoonloveInfo = moonloveInfo
	scMoonloveSceneInfo.RewExp = &expCount
	scMoonloveSceneInfo.CreateTime = &startTime

	return scMoonloveSceneInfo

}

func buildMoonloveCharmRank(charmData *moonlovetypes.RankData) *uipb.MoonloveRank {
	charmRank := &uipb.MoonloveRank{}

	charmRank.PlayerName = &charmData.Name
	charmRank.Number = &charmData.Number

	return charmRank
}

func buildMoonloveGenerousRank(generousData *moonlovetypes.RankData) *uipb.MoonloveRank {
	generousRank := &uipb.MoonloveRank{}

	generousRank.PlayerName = &generousData.Name
	generousRank.Number = &generousData.Number

	return generousRank
}

func BuildMoonloveRankRewards(rd *propertytypes.RewData) *uipb.SCMoonloveRankRewards {

	rewProperty := &uipb.RewProperty{}
	rewExp := rd.GetRewExp()
	rewExpPoint := rd.GetRewExpPoint()
	rewGold := rd.GetRewGold()
	rewBindGold := rd.GetRewBindGold()
	rewSilver := rd.GetRewSilver()

	rewProperty.Exp = &rewExp
	rewProperty.ExpPoint = &rewExpPoint
	rewProperty.Silver = &rewSilver
	rewProperty.Gold = &rewGold
	rewProperty.BindGold = &rewBindGold

	sc := &uipb.SCMoonloveRankRewards{}
	sc.RewProperty = rewProperty

	return sc
}

func BuildMoonloveViewDouble(playerId, otherPlayerId int64, target types.Position) *uipb.SCMoonloveViewDouble {
	scMoonloveViewDouble := &uipb.SCMoonloveViewDouble{}
	targetPosition := &uipb.Position{}
	x := float32(target.X)
	y := float32(target.Y)
	z := float32(target.Z)
	targetPosition.PosX = &x
	targetPosition.PosY = &y
	targetPosition.PosZ = &z

	scMoonloveViewDouble.PlayerId = &playerId
	scMoonloveViewDouble.TargetPlayerId = &otherPlayerId
	scMoonloveViewDouble.TargetPosition = targetPosition

	return scMoonloveViewDouble
}

func BuildMoonloveViewDoubleState(targetPlayerId int64, isSuccess bool) *uipb.SCMoonloveViewDoubleState {
	scMoonloveViewDoubleState := &uipb.SCMoonloveViewDoubleState{}

	scMoonloveViewDoubleState.TargetPlayerId = &targetPlayerId
	scMoonloveViewDoubleState.IsSuccess = &isSuccess

	return scMoonloveViewDoubleState
}

func BuildMoonloveViewDoubleRelease(playerId int64) *uipb.SCMoonloveViewDoubleRelease {
	scMoonloveViewDoubleRelease := &uipb.SCMoonloveViewDoubleRelease{}
	scMoonloveViewDoubleRelease.PlayerId = &playerId

	return scMoonloveViewDoubleRelease
}

func BuildMoonlovePushCharmRank(rankArr []*moonlovetypes.RankData) *uipb.SCMoonlovePushCharmRank {
	scMoonlovePushCharmRank := &uipb.SCMoonlovePushCharmRank{}

	var charmRankArr []*uipb.MoonloveRank
	for _, charm := range rankArr {
		charmRankArr = append(charmRankArr, buildMoonloveCharmRank(charm))
	}
	scMoonlovePushCharmRank.CharmRank = charmRankArr

	return scMoonlovePushCharmRank
}

func BuildMoonlovePushGenerousRank(rankArr []*moonlovetypes.RankData) *uipb.SCMoonlovePushGenerousRank {
	scMoonlovePushGenerousRank := &uipb.SCMoonlovePushGenerousRank{}

	var generousRankArr []*uipb.MoonloveRank
	for _, generous := range rankArr {
		generousRankArr = append(generousRankArr, buildMoonloveGenerousRank(generous))
	}
	scMoonlovePushGenerousRank.GenerousRank = generousRankArr

	return scMoonlovePushGenerousRank
}

func BuildMoonlovePlayerList(playerId int64, playerMap map[int64]scene.Player, friendMap map[int64]*friend.FriendObject, coupleMap map[int64]*moonlovetypes.MoonloveDoubleData) *uipb.SCMoonlovePlayerList {
	scMoonlovePlayerList := &uipb.SCMoonlovePlayerList{}
	scMoonlovePlayerList.PlayerList = buildScenePlayerInfo(playerId, playerMap, friendMap, coupleMap)

	return scMoonlovePlayerList
}

func BuildSCMoonloveGiftNotice(playerName, targetName string, num int32) *uipb.SCMoonloveGiftNotice {
	scMoonloveGiftNotice := &uipb.SCMoonloveGiftNotice{}

	scMoonloveGiftNotice.Name = &playerName
	scMoonloveGiftNotice.TargetName = &targetName
	scMoonloveGiftNotice.Num = &num

	return scMoonloveGiftNotice
}

func BuildSCMoonloveExpCountNotice(expCount int64) *uipb.SCMoonloveExpCountNotice {
	scMoonloveExpCountNotice := &uipb.SCMoonloveExpCountNotice{}
	scMoonloveExpCountNotice.ExpCount = &expCount

	return scMoonloveExpCountNotice
}

func buildScenePlayerInfo(playerId int64, playerMap map[int64]scene.Player, friendMap map[int64]*friend.FriendObject, coupleMap map[int64]*moonlovetypes.MoonloveDoubleData) (infoList []*uipb.ScenePlayer) {
	for plId, spl := range playerMap {
		if plId == playerId {
			continue
		}

		splayerId := spl.GetId()
		name := spl.GetName()
		sex := int32(spl.GetSex())
		point := int32(0)

		// 亲密度
		friend, ok := friendMap[plId]
		if ok {
			point = friend.Point
		}

		// 状态
		couple := int32(0)
		_, isCouple := coupleMap[plId]
		if isCouple {
			couple = 1
		}

		info := &uipb.ScenePlayer{}
		info.PlayerId = &splayerId
		info.Name = &name
		info.Sex = &sex
		info.Point = &point
		info.Status = &couple

		infoList = append(infoList, info)
	}

	return infoList
}
