package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

type BabyForceRank struct {
	babyForceList   []*rankentity.PlayerPropertyData
	babyForceTime   int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newBabyForceRank(config *ranktypes.RankConfig) RankTypeData {
	babyLevelRank := &BabyForceRank{
		babyForceList:   make([]*rankentity.PlayerPropertyData, 0, 8),
		babyForceTime:   0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return babyLevelRank
}

func (l *BabyForceRank) init(timestamp int64) (err error) {
	l.babyForceList, err = dao.GetRankDao().GetRedisRankBabyForceList(timestamp, l.config)
	if err != nil {
		return
	}
	if l.babyForceList == nil {
		err = l.updateRankList(timestamp)
		if err != nil {
			return
		}
	}

	for index, obj := range l.babyForceList {
		l.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	l.buildRankingInfo()
	l.babyForceTime = timestamp
	return
}

func (l *BabyForceRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(l.babyForceList) == 0 {
		return
	}
	return l.babyForceList[0].PlayerId
}

func (l *BabyForceRank) UpdateRankDataList(timestamp int64) (err error) {
	return l.updateRankList(timestamp)
}

func (l *BabyForceRank) updateRankList(timestamp int64) (err error) {
	if l.config.StartTime != 0 || l.config.EndTime != 0 {
		if timestamp < l.config.StartTime || timestamp > l.config.EndTime {
			return
		}
	}

	diffTime := timestamp - l.babyForceTime
	if diffTime < l.config.RefreshTime {
		return
	}

	oldRankInfoList := l.rankingInfoList
	l.babyForceList, err = dao.GetRankDao().GetRankBabyForceList(l.config)
	if err != nil {
		return
	}

	err = dao.GetRankDao().SetRedisRankBabyForceList(timestamp, l.config, l.babyForceList)
	if err != nil {
		return
	}

	rankingMap := make(map[int64]int32)
	for index, obj := range l.babyForceList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}

	l.babyForceTime = timestamp
	l.rankingMap = rankingMap
	l.buildRankingInfo()
	newRankInfoList := l.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, l.config.ClassType, l.config.RankType, l.config.GroupId)

	return
}

func (l *BabyForceRank) GetPos(playerId int64) (pos int32) {
	pos = l.rankingMap[playerId]
	return
}

func (l *BabyForceRank) ResetRankTime() {
	l.babyForceTime = 0
}

func (l *BabyForceRank) GetListAndTime() ([]*rankentity.PlayerPropertyData, int64) {
	return l.babyForceList, l.babyForceTime
}

func (l *BabyForceRank) GetRankingMap() map[int64]int32 {
	return l.rankingMap
}

func (l *BabyForceRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return l.rankingInfoList
}

func (l *BabyForceRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(l.babyForceList))
	for index, obj := range l.babyForceList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Num)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	l.rankingInfoList = newList
}
