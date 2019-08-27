package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/game/arena/arena"
	playerarena "fgame/fgame/game/arena/player"
	arenatypes "fgame/fgame/game/arena/types"
)

var (
	csArenaStopMatch      = &uipb.CSArenaStopMatch{}
	scArenaMatch          = &uipb.SCArenaMatch{}
	scArenaMatchBroadcast = &uipb.SCArenaMatchBroadcast{}
)

func BuildSCArenaMatch() *uipb.SCArenaMatch {
	return scArenaMatch
}

func BuildSCArenaMatchBroadcast() *uipb.SCArenaMatchBroadcast {
	return scArenaMatchBroadcast
}

var (
	scArenaStopMatch          = &uipb.SCArenaStopMatch{}
	scArenaStopMatchBroadcast = &uipb.SCArenaStopMatchBroadcast{}
)

func BuildSCArenaStopMatch() *uipb.SCArenaStopMatch {
	return scArenaStopMatch
}

func BuildSCArenaStopMatchBroadcast() *uipb.SCArenaStopMatchBroadcast {
	return scArenaStopMatchBroadcast
}

func BuildSCArenaMatchResult(result bool) *uipb.SCArenaMatchResult {
	scArenaMatchResult := &uipb.SCArenaMatchResult{}
	scArenaMatchResult.Result = &result
	return scArenaMatchResult
}

func BuildSCArenaSelectFourGod(fourGodType arenatypes.FourGodType) *uipb.SCArenaSelectFourGod {
	fourGodTypeInt := int32(fourGodType)
	scArenaSelectFourGod := &uipb.SCArenaSelectFourGod{}
	scArenaSelectFourGod.FourGodType = &fourGodTypeInt
	return scArenaSelectFourGod
}

//gm模拟使用
func BuildCSArenaSelectFourGod(fourGodType arenatypes.FourGodType) *uipb.CSArenaSelectFourGod {
	fourGodTypeInt := int32(fourGodType)
	csArenaSelectFourGod := &uipb.CSArenaSelectFourGod{}
	csArenaSelectFourGod.FourGodType = &fourGodTypeInt
	return csArenaSelectFourGod
}

var (
	csArenaNextMatch = &uipb.CSArenaNextMatch{}
)

//gm模拟使用
func BuildCSArenaNextMatch() *uipb.CSArenaNextMatch {
	return csArenaNextMatch
}

var (
	scArenaInvite = &uipb.SCArenaInvite{}
)

func BuildSCArenaInvite() *uipb.SCArenaInvite {
	return scArenaInvite
}

func BuildSCPlayerArenaInfo(arenaObj *playerarena.PlayerArenaObject) *uipb.SCPlayerArenaInfo {
	rewardTime := arenaObj.GetCulRewardTime()
	curPoint := arenaObj.GetJiFenCount()
	dayPoint := arenaObj.GetJiFenDay()
	winCount := arenaObj.GetWinCount()
	failCount := arenaObj.GetFailCount()
	dayMaxWinCount := arenaObj.GetDayMaxWinCount()
	dayWinCount := arenaObj.GetDayWinCount()
	rankRewTime := arenaObj.GetRankRewTime()

	scMsg := &uipb.SCPlayerArenaInfo{}
	scMsg.RewardTime = &rewardTime
	scMsg.CurPoint = &curPoint
	scMsg.DayPoint = &dayPoint
	scMsg.WinCount = &winCount
	scMsg.FailCount = &failCount
	scMsg.DayMaxWinCount = &dayMaxWinCount
	scMsg.DayWinCount = &dayWinCount
	scMsg.RankRewTime = &rankRewTime
	return scMsg
}

func BuildSCArenaGetReward(rankTime int64) *uipb.SCArenaGetReward {
	scArenaGetReward := &uipb.SCArenaGetReward{}
	scArenaGetReward.RankTime = &rankTime
	return scArenaGetReward
}

func BuildSCArenaMyRank(timeType, pos, winCount int32, rankTime int64) *uipb.SCArenaMyRank {
	scArenaMyRank := &uipb.SCArenaMyRank{}
	scArenaMyRank.TimeType = &timeType
	scArenaMyRank.Pos = &pos
	scArenaMyRank.WinCount = &winCount
	scArenaMyRank.RankTime = &rankTime
	return scArenaMyRank
}

func BuildSCArenaRankGet(timeType arenatypes.RankTimeType, page int32, rankTime int64, dataList []*arena.ArenaRankObject) *uipb.SCArenaRankGet {
	timeTypeInt := int32(timeType)
	scArenaRankGet := &uipb.SCArenaRankGet{}
	scArenaRankGet.TimeType = &timeTypeInt
	scArenaRankGet.Page = &page
	scArenaRankGet.RankTime = &rankTime
	scArenaRankGet.RankList = buildArenaRankList(timeType, dataList)
	return scArenaRankGet
}

func buildArenaRankList(timeType arenatypes.RankTimeType, dataList []*arena.ArenaRankObject) (infoList []*uipb.ArenaRank) {
	for _, data := range dataList {
		serverId := data.GetServerId()
		playerId := data.GetPlayerId()
		playerName := data.GetPlayerName()
		winCount := int32(0)
		if timeType == arenatypes.RankTimeTypeThis {
			winCount = data.GetWinCount()
		} else {
			winCount = data.GetLastWinCount()
		}

		arenaRank := &uipb.ArenaRank{}
		arenaRank.ServerId = &serverId
		arenaRank.PlayerId = &playerId
		arenaRank.PlayerName = &playerName
		arenaRank.WinCount = &winCount

		infoList = append(infoList, arenaRank)
	}

	return infoList
}
