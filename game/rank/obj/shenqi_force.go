package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

type ShenQiForceRank struct {
	shenQiForceList []*rankentity.PlayerPropertyData
	shenQiForceTime int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newShenQiForceRank(config *ranktypes.RankConfig) RankTypeData {
	shenQiLevelRank := &ShenQiForceRank{
		shenQiForceList: make([]*rankentity.PlayerPropertyData, 0, 8),
		shenQiForceTime: 0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return shenQiLevelRank
}

func (l *ShenQiForceRank) init(timestamp int64) (err error) {
	l.shenQiForceList, err = dao.GetRankDao().GetRedisRankShenQiForceList(timestamp, l.config)
	if err != nil {
		return
	}
	if l.shenQiForceList == nil {
		err = l.updateRankList(timestamp)
		if err != nil {
			return
		}
	}

	for index, obj := range l.shenQiForceList {
		l.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	l.buildRankingInfo()
	l.shenQiForceTime = timestamp
	return
}

func (l *ShenQiForceRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(l.shenQiForceList) == 0 {
		return
	}
	return l.shenQiForceList[0].PlayerId
}

func (l *ShenQiForceRank) UpdateRankDataList(timestamp int64) (err error) {
	return l.updateRankList(timestamp)
}

func (l *ShenQiForceRank) updateRankList(timestamp int64) (err error) {
	if l.config.StartTime != 0 || l.config.EndTime != 0 {
		if timestamp < l.config.StartTime || timestamp > l.config.EndTime {
			return
		}
	}

	diffTime := timestamp - l.shenQiForceTime
	if diffTime < l.config.RefreshTime {
		return
	}

	oldRankInfoList := l.rankingInfoList
	l.shenQiForceList, err = dao.GetRankDao().GetRankShenQiForceList(l.config)
	if err != nil {
		return
	}

	err = dao.GetRankDao().SetRedisRankShenQiForceList(timestamp, l.config, l.shenQiForceList)
	if err != nil {
		return
	}

	rankingMap := make(map[int64]int32)
	for index, obj := range l.shenQiForceList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}

	l.shenQiForceTime = timestamp
	l.rankingMap = rankingMap
	l.buildRankingInfo()
	newRankInfoList := l.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, l.config.ClassType, l.config.RankType, l.config.GroupId)

	return
}

func (l *ShenQiForceRank) GetPos(playerId int64) (pos int32) {
	pos = l.rankingMap[playerId]
	return
}

func (l *ShenQiForceRank) ResetRankTime() {
	l.shenQiForceTime = 0
}

func (l *ShenQiForceRank) GetListAndTime() ([]*rankentity.PlayerPropertyData, int64) {
	return l.shenQiForceList, l.shenQiForceTime
}

func (l *ShenQiForceRank) GetRankingMap() map[int64]int32 {
	return l.rankingMap
}

func (l *ShenQiForceRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return l.rankingInfoList
}

func (l *ShenQiForceRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(l.shenQiForceList))
	for index, obj := range l.shenQiForceList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Num)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	l.rankingInfoList = newList
}
