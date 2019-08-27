package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

type TuLongEquipForceRank struct {
	tuLongEquipForceList []*rankentity.PlayerPropertyData
	tuLongEquipForceTime int64
	rankingMap           map[int64]int32
	config               *ranktypes.RankConfig
	rankingInfoList      []*ranktypes.RankingInfo
}

func newTuLongEquipForceRank(config *ranktypes.RankConfig) RankTypeData {
	tuLongEquipLevelRank := &TuLongEquipForceRank{
		tuLongEquipForceList: make([]*rankentity.PlayerPropertyData, 0, 8),
		tuLongEquipForceTime: 0,
		rankingMap:           make(map[int64]int32),
		rankingInfoList:      make([]*ranktypes.RankingInfo, 0, 8),
		config:               config,
	}
	return tuLongEquipLevelRank
}

func (l *TuLongEquipForceRank) init(timestamp int64) (err error) {
	l.tuLongEquipForceList, err = dao.GetRankDao().GetRedisRankTuLongEquipForceList(timestamp, l.config)
	if err != nil {
		return
	}
	if l.tuLongEquipForceList == nil {
		err = l.updateRankList(timestamp)
		if err != nil {
			return
		}
	}

	for index, obj := range l.tuLongEquipForceList {
		l.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	l.buildRankingInfo()
	l.tuLongEquipForceTime = timestamp
	return
}

func (l *TuLongEquipForceRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(l.tuLongEquipForceList) == 0 {
		return
	}
	return l.tuLongEquipForceList[0].PlayerId
}

func (l *TuLongEquipForceRank) UpdateRankDataList(timestamp int64) (err error) {
	return l.updateRankList(timestamp)
}

func (l *TuLongEquipForceRank) updateRankList(timestamp int64) (err error) {
	if l.config.StartTime != 0 || l.config.EndTime != 0 {
		if timestamp < l.config.StartTime || timestamp > l.config.EndTime {
			return
		}
	}

	diffTime := timestamp - l.tuLongEquipForceTime
	if diffTime < l.config.RefreshTime {
		return
	}

	oldRankInfoList := l.rankingInfoList
	l.tuLongEquipForceList, err = dao.GetRankDao().GetRankTuLongEquipForceList(l.config)
	if err != nil {
		return
	}

	err = dao.GetRankDao().SetRedisRankTuLongEquipForceList(timestamp, l.config, l.tuLongEquipForceList)
	if err != nil {
		return
	}

	rankingMap := make(map[int64]int32)
	for index, obj := range l.tuLongEquipForceList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}

	l.tuLongEquipForceTime = timestamp
	l.rankingMap = rankingMap
	l.buildRankingInfo()
	newRankInfoList := l.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, l.config.ClassType, l.config.RankType, l.config.GroupId)

	return
}

func (l *TuLongEquipForceRank) GetPos(playerId int64) (pos int32) {
	pos = l.rankingMap[playerId]
	return
}

func (l *TuLongEquipForceRank) ResetRankTime() {
	l.tuLongEquipForceTime = 0
}

func (l *TuLongEquipForceRank) GetListAndTime() ([]*rankentity.PlayerPropertyData, int64) {
	return l.tuLongEquipForceList, l.tuLongEquipForceTime
}

func (l *TuLongEquipForceRank) GetRankingMap() map[int64]int32 {
	return l.rankingMap
}

func (l *TuLongEquipForceRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return l.rankingInfoList
}

func (l *TuLongEquipForceRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(l.tuLongEquipForceList))
	for index, obj := range l.tuLongEquipForceList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Num)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	l.rankingInfoList = newList
}
