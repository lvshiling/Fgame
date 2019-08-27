package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

type LingTongForceRank struct {
	lingTongForceList []*rankentity.PlayerPropertyData
	lingTongForceTime int64
	rankingMap        map[int64]int32
	config            *ranktypes.RankConfig
	rankingInfoList   []*ranktypes.RankingInfo
}

func newLingTongForceRank(config *ranktypes.RankConfig) RankTypeData {
	lingTongLevelRank := &LingTongForceRank{
		lingTongForceList: make([]*rankentity.PlayerPropertyData, 0, 8),
		lingTongForceTime: 0,
		rankingMap:        make(map[int64]int32),
		rankingInfoList:   make([]*ranktypes.RankingInfo, 0, 8),
		config:            config,
	}
	return lingTongLevelRank
}

func (l *LingTongForceRank) init(timestamp int64) (err error) {
	l.lingTongForceList, err = dao.GetRankDao().GetRedisRankLingTongForceList(timestamp, l.config)
	if err != nil {
		return
	}
	if l.lingTongForceList == nil {
		err = l.updateRankList(timestamp)
		if err != nil {
			return
		}
	}

	for index, obj := range l.lingTongForceList {
		l.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	l.buildRankingInfo()
	l.lingTongForceTime = timestamp
	return
}

func (l *LingTongForceRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(l.lingTongForceList) == 0 {
		return
	}
	return l.lingTongForceList[0].PlayerId
}

func (l *LingTongForceRank) UpdateRankDataList(timestamp int64) (err error) {
	return l.updateRankList(timestamp)
}

func (l *LingTongForceRank) updateRankList(timestamp int64) (err error) {
	if l.config.StartTime != 0 || l.config.EndTime != 0 {
		if timestamp < l.config.StartTime || timestamp > l.config.EndTime {
			return
		}
	}

	diffTime := timestamp - l.lingTongForceTime
	if diffTime < l.config.RefreshTime {
		return
	}

	oldRankInfoList := l.rankingInfoList
	l.lingTongForceList, err = dao.GetRankDao().GetRankLingTongForceList(l.config)
	if err != nil {
		return
	}

	err = dao.GetRankDao().SetRedisRankLingTongForceList(timestamp, l.config, l.lingTongForceList)
	if err != nil {
		return
	}

	rankingMap := make(map[int64]int32)
	for index, obj := range l.lingTongForceList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}

	l.lingTongForceTime = timestamp
	l.rankingMap = rankingMap
	l.buildRankingInfo()
	newRankInfoList := l.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, l.config.ClassType, l.config.RankType, l.config.GroupId)

	return
}

func (l *LingTongForceRank) GetPos(playerId int64) (pos int32) {
	pos = l.rankingMap[playerId]
	return
}

func (l *LingTongForceRank) ResetRankTime() {
	l.lingTongForceTime = 0
}

func (l *LingTongForceRank) GetListAndTime() ([]*rankentity.PlayerPropertyData, int64) {
	return l.lingTongForceList, l.lingTongForceTime
}

func (l *LingTongForceRank) GetRankingMap() map[int64]int32 {
	return l.rankingMap
}

func (l *LingTongForceRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return l.rankingInfoList
}

func (l *LingTongForceRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(l.lingTongForceList))
	for index, obj := range l.lingTongForceList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Num)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	l.rankingInfoList = newList
}
