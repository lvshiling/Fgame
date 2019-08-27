package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

type MingGeForceRank struct {
	mingGeForceList []*rankentity.PlayerPropertyData
	mingGeForceTime int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newMingGeForceRank(config *ranktypes.RankConfig) RankTypeData {
	mingGeLevelRank := &MingGeForceRank{
		mingGeForceList: make([]*rankentity.PlayerPropertyData, 0, 8),
		mingGeForceTime: 0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return mingGeLevelRank
}

func (l *MingGeForceRank) init(timestamp int64) (err error) {
	l.mingGeForceList, err = dao.GetRankDao().GetRedisRankMingGeForceList(timestamp, l.config)
	if err != nil {
		return
	}
	if l.mingGeForceList == nil {
		err = l.updateRankList(timestamp)
		if err != nil {
			return
		}
	}

	for index, obj := range l.mingGeForceList {
		l.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	l.buildRankingInfo()
	l.mingGeForceTime = timestamp
	return
}

func (l *MingGeForceRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(l.mingGeForceList) == 0 {
		return
	}
	return l.mingGeForceList[0].PlayerId
}

func (l *MingGeForceRank) UpdateRankDataList(timestamp int64) (err error) {
	return l.updateRankList(timestamp)
}

func (l *MingGeForceRank) updateRankList(timestamp int64) (err error) {
	if l.config.StartTime != 0 || l.config.EndTime != 0 {
		if timestamp < l.config.StartTime || timestamp > l.config.EndTime {
			return
		}
	}

	diffTime := timestamp - l.mingGeForceTime
	if diffTime < l.config.RefreshTime {
		return
	}

	oldRankInfoList := l.rankingInfoList
	l.mingGeForceList, err = dao.GetRankDao().GetRankMingGeForceList(l.config)
	if err != nil {
		return
	}

	err = dao.GetRankDao().SetRedisRankMingGeForceList(timestamp, l.config, l.mingGeForceList)
	if err != nil {
		return
	}

	rankingMap := make(map[int64]int32)
	for index, obj := range l.mingGeForceList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}

	l.mingGeForceTime = timestamp
	l.rankingMap = rankingMap
	l.buildRankingInfo()
	newRankInfoList := l.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, l.config.ClassType, l.config.RankType, l.config.GroupId)

	return
}

func (l *MingGeForceRank) GetPos(playerId int64) (pos int32) {
	pos = l.rankingMap[playerId]
	return
}

func (l *MingGeForceRank) ResetRankTime() {
	l.mingGeForceTime = 0
}

func (l *MingGeForceRank) GetListAndTime() ([]*rankentity.PlayerPropertyData, int64) {
	return l.mingGeForceList, l.mingGeForceTime
}

func (l *MingGeForceRank) GetRankingMap() map[int64]int32 {
	return l.rankingMap
}

func (l *MingGeForceRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return l.rankingInfoList
}

func (l *MingGeForceRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(l.mingGeForceList))
	for index, obj := range l.mingGeForceList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Num)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	l.rankingInfoList = newList
}
