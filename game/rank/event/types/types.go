package types

import (
	ranktypes "fgame/fgame/game/rank/types"
)

type RankEventType string

const (
	//排行榜第一名变化
	RankEventTypeFirst RankEventType = "RankFirstChanged"
	//排行榜排名变化
	RankEventTypeRankChanged = "RankChanged"
	//排行榜重置
	RankEventTypeRankReset = "RankReset"
)

func CreateRankEventData(playerId int64, oldPlayerId int64, rankClassType ranktypes.RankClassType, rankType ranktypes.RankType) *RankEventData {
	tted := &RankEventData{
		playerId:      playerId,
		oldPlayerId:   oldPlayerId,
		rankClassType: rankClassType,
		rankType:      rankType,
	}
	return tted
}

type RankEventData struct {
	playerId      int64
	oldPlayerId   int64
	rankClassType ranktypes.RankClassType
	rankType      ranktypes.RankType
}

func (red *RankEventData) GetPlayerId() int64 {
	return red.playerId
}

func (red *RankEventData) GetOldPlayerId() int64 {
	return red.oldPlayerId
}

func (red *RankEventData) GetRankClassType() ranktypes.RankClassType {
	return red.rankClassType
}

func (red *RankEventData) GetRankType() ranktypes.RankType {
	return red.rankType
}

//
type RankChangedEventData struct {
	oldRankList   []*ranktypes.RankingInfo
	newRankList   []*ranktypes.RankingInfo
	rankClassType ranktypes.RankClassType
	rankType      ranktypes.RankType
	groupId       int32
}

func CreateRankChangedEventData(oldRankList, newRankList []*ranktypes.RankingInfo, rankClassType ranktypes.RankClassType, rankType ranktypes.RankType, groupId int32) *RankChangedEventData {
	d := &RankChangedEventData{
		oldRankList:   oldRankList,
		newRankList:   newRankList,
		rankClassType: rankClassType,
		rankType:      rankType,
		groupId:       groupId,
	}
	return d
}

func (red *RankChangedEventData) GetOldRankList() []*ranktypes.RankingInfo {
	return red.oldRankList
}

func (red *RankChangedEventData) GetNewRankList() []*ranktypes.RankingInfo {
	return red.newRankList
}

func (red *RankChangedEventData) GetRankClassType() ranktypes.RankClassType {
	return red.rankClassType
}

func (red *RankChangedEventData) GetRankType() ranktypes.RankType {
	return red.rankType
}
func (red *RankChangedEventData) GetGroupId() int32 {
	return red.groupId
}
