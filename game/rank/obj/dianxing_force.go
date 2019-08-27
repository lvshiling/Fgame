package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

type DianXingForceRank struct {
	dianXingForceList []*rankentity.PlayerPropertyData
	dianXingForceTime int64
	rankingMap        map[int64]int32
	config            *ranktypes.RankConfig
	rankingInfoList   []*ranktypes.RankingInfo
}

func newDianXingForceRank(config *ranktypes.RankConfig) RankTypeData {
	dianXingLevelRank := &DianXingForceRank{
		dianXingForceList: make([]*rankentity.PlayerPropertyData, 0, 8),
		dianXingForceTime: 0,
		rankingMap:        make(map[int64]int32),
		rankingInfoList:   make([]*ranktypes.RankingInfo, 0, 8),
		config:            config,
	}
	return dianXingLevelRank
}

func (l *DianXingForceRank) init(timestamp int64) (err error) {
	l.dianXingForceList, err = dao.GetRankDao().GetRedisRankDianXingForceList(timestamp, l.config)
	if err != nil {
		return
	}
	if l.dianXingForceList == nil {
		err = l.updateRankList(timestamp)
		if err != nil {
			return
		}
	}

	for index, obj := range l.dianXingForceList {
		l.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	l.buildRankingInfo()
	l.dianXingForceTime = timestamp
	return
}

func (l *DianXingForceRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(l.dianXingForceList) == 0 {
		return
	}
	return l.dianXingForceList[0].PlayerId
}

func (l *DianXingForceRank) UpdateRankDataList(timestamp int64) (err error) {
	return l.updateRankList(timestamp)
}

func (l *DianXingForceRank) updateRankList(timestamp int64) (err error) {
	if l.config.StartTime != 0 || l.config.EndTime != 0 {
		if timestamp < l.config.StartTime || timestamp > l.config.EndTime {
			return
		}
	}

	diffTime := timestamp - l.dianXingForceTime
	if diffTime < l.config.RefreshTime {
		return
	}

	oldRankInfoList := l.rankingInfoList
	l.dianXingForceList, err = dao.GetRankDao().GetRankDianXingForceList(l.config)
	if err != nil {
		return
	}

	err = dao.GetRankDao().SetRedisRankDianXingForceList(timestamp, l.config, l.dianXingForceList)
	if err != nil {
		return
	}

	rankingMap := make(map[int64]int32)
	for index, obj := range l.dianXingForceList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}

	l.dianXingForceTime = timestamp
	l.rankingMap = rankingMap
	l.buildRankingInfo()
	newRankInfoList := l.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, l.config.ClassType, l.config.RankType, l.config.GroupId)

	return
}

func (l *DianXingForceRank) GetPos(playerId int64) (pos int32) {
	pos = l.rankingMap[playerId]
	return
}

func (l *DianXingForceRank) ResetRankTime() {
	l.dianXingForceTime = 0
}

func (l *DianXingForceRank) GetListAndTime() ([]*rankentity.PlayerPropertyData, int64) {
	return l.dianXingForceList, l.dianXingForceTime
}

func (l *DianXingForceRank) GetRankingMap() map[int64]int32 {
	return l.rankingMap
}

func (l *DianXingForceRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return l.rankingInfoList
}

func (l *DianXingForceRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(l.dianXingForceList))
	for index, obj := range l.dianXingForceList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Num)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	l.rankingInfoList = newList
}
