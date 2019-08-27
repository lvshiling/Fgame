package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

type ShengHenForceRank struct {
	shengHenForceList []*rankentity.PlayerPropertyData
	shengHenForceTime int64
	rankingMap        map[int64]int32
	config            *ranktypes.RankConfig
	rankingInfoList   []*ranktypes.RankingInfo
}

func newShengHenForceRank(config *ranktypes.RankConfig) RankTypeData {
	shengHenLevelRank := &ShengHenForceRank{
		shengHenForceList: make([]*rankentity.PlayerPropertyData, 0, 8),
		shengHenForceTime: 0,
		rankingMap:        make(map[int64]int32),
		rankingInfoList:   make([]*ranktypes.RankingInfo, 0, 8),
		config:            config,
	}
	return shengHenLevelRank
}

func (l *ShengHenForceRank) init(timestamp int64) (err error) {
	l.shengHenForceList, err = dao.GetRankDao().GetRedisRankShengHenForceList(timestamp, l.config)
	if err != nil {
		return
	}
	if l.shengHenForceList == nil {
		err = l.updateRankList(timestamp)
		if err != nil {
			return
		}
	}

	for index, obj := range l.shengHenForceList {
		l.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	l.buildRankingInfo()
	l.shengHenForceTime = timestamp
	return
}

func (l *ShengHenForceRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(l.shengHenForceList) == 0 {
		return
	}
	return l.shengHenForceList[0].PlayerId
}

func (l *ShengHenForceRank) UpdateRankDataList(timestamp int64) (err error) {
	return l.updateRankList(timestamp)
}

func (l *ShengHenForceRank) updateRankList(timestamp int64) (err error) {
	if l.config.StartTime != 0 || l.config.EndTime != 0 {
		if timestamp < l.config.StartTime || timestamp > l.config.EndTime {
			return
		}
	}

	diffTime := timestamp - l.shengHenForceTime
	if diffTime < l.config.RefreshTime {
		return
	}

	oldRankInfoList := l.rankingInfoList
	l.shengHenForceList, err = dao.GetRankDao().GetRankShengHenForceList(l.config)
	if err != nil {
		return
	}

	err = dao.GetRankDao().SetRedisRankShengHenForceList(timestamp, l.config, l.shengHenForceList)
	if err != nil {
		return
	}

	rankingMap := make(map[int64]int32)
	for index, obj := range l.shengHenForceList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}

	l.shengHenForceTime = timestamp
	l.rankingMap = rankingMap
	l.buildRankingInfo()
	newRankInfoList := l.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, l.config.ClassType, l.config.RankType, l.config.GroupId)

	return
}

func (l *ShengHenForceRank) GetPos(playerId int64) (pos int32) {
	pos = l.rankingMap[playerId]
	return
}

func (l *ShengHenForceRank) ResetRankTime() {
	l.shengHenForceTime = 0
}

func (l *ShengHenForceRank) GetListAndTime() ([]*rankentity.PlayerPropertyData, int64) {
	return l.shengHenForceList, l.shengHenForceTime
}

func (l *ShengHenForceRank) GetRankingMap() map[int64]int32 {
	return l.rankingMap
}

func (l *ShengHenForceRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return l.rankingInfoList
}

func (l *ShengHenForceRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(l.shengHenForceList))
	for index, obj := range l.shengHenForceList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Num)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	l.rankingInfoList = newList
}
