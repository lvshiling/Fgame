package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

type ZhenFaForceRank struct {
	zhenFaForceList []*rankentity.PlayerPropertyData
	zhenFaForceTime int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newZhenFaForceRank(config *ranktypes.RankConfig) RankTypeData {
	zhenFaLevelRank := &ZhenFaForceRank{
		zhenFaForceList: make([]*rankentity.PlayerPropertyData, 0, 8),
		zhenFaForceTime: 0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return zhenFaLevelRank
}

func (l *ZhenFaForceRank) init(timestamp int64) (err error) {
	l.zhenFaForceList, err = dao.GetRankDao().GetRedisRankZhenFaForceList(timestamp, l.config)
	if err != nil {
		return
	}
	if l.zhenFaForceList == nil {
		err = l.updateRankList(timestamp)
		if err != nil {
			return
		}
	}

	for index, obj := range l.zhenFaForceList {
		l.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	l.buildRankingInfo()
	l.zhenFaForceTime = timestamp
	return
}

func (l *ZhenFaForceRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(l.zhenFaForceList) == 0 {
		return
	}
	return l.zhenFaForceList[0].PlayerId
}

func (l *ZhenFaForceRank) UpdateRankDataList(timestamp int64) (err error) {
	return l.updateRankList(timestamp)
}

func (l *ZhenFaForceRank) updateRankList(timestamp int64) (err error) {
	if l.config.StartTime != 0 || l.config.EndTime != 0 {
		if timestamp < l.config.StartTime || timestamp > l.config.EndTime {
			return
		}
	}

	diffTime := timestamp - l.zhenFaForceTime
	if diffTime < l.config.RefreshTime {
		return
	}

	oldRankInfoList := l.rankingInfoList
	l.zhenFaForceList, err = dao.GetRankDao().GetRankZhenFaForceList(l.config)
	if err != nil {
		return
	}

	err = dao.GetRankDao().SetRedisRankZhenFaForceList(timestamp, l.config, l.zhenFaForceList)
	if err != nil {
		return
	}

	rankingMap := make(map[int64]int32)
	for index, obj := range l.zhenFaForceList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}

	l.zhenFaForceTime = timestamp
	l.rankingMap = rankingMap
	l.buildRankingInfo()
	newRankInfoList := l.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, l.config.ClassType, l.config.RankType, l.config.GroupId)

	return
}

func (l *ZhenFaForceRank) GetPos(playerId int64) (pos int32) {
	pos = l.rankingMap[playerId]
	return
}

func (l *ZhenFaForceRank) ResetRankTime() {
	l.zhenFaForceTime = 0
}

func (l *ZhenFaForceRank) GetListAndTime() ([]*rankentity.PlayerPropertyData, int64) {
	return l.zhenFaForceList, l.zhenFaForceTime
}

func (l *ZhenFaForceRank) GetRankingMap() map[int64]int32 {
	return l.rankingMap
}

func (l *ZhenFaForceRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return l.rankingInfoList
}

func (l *ZhenFaForceRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(l.zhenFaForceList))
	for index, obj := range l.zhenFaForceList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Num)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	l.rankingInfoList = newList
}
