package logic

import (
	gameevent "fgame/fgame/game/event"
	rankeventtypes "fgame/fgame/game/rank/event/types"
	ranktypes "fgame/fgame/game/rank/types"
)

//排行榜第一名变化
func RankFirstChange(playerId int64, oldPlayerId int64, classType ranktypes.RankClassType, rankType ranktypes.RankType) (err error) {
	eventData := rankeventtypes.CreateRankEventData(playerId, oldPlayerId, classType, rankType)
	err = gameevent.Emit(rankeventtypes.RankEventTypeFirst, nil, eventData)
	return
}

//排行榜变化
func RankChanged(oldRankList, newRankList []*ranktypes.RankingInfo, classType ranktypes.RankClassType, rankType ranktypes.RankType, groupId int32) (err error) {
	eventData := rankeventtypes.CreateRankChangedEventData(oldRankList, newRankList, classType, rankType, groupId)
	err = gameevent.Emit(rankeventtypes.RankEventTypeRankChanged, nil, eventData)
	return
}
