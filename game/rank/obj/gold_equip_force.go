package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

type GoldEquipForceRank struct {
	goldEquipForceList []*rankentity.PlayerPropertyData
	goldEquipForceTime int64
	rankingMap         map[int64]int32
	config             *ranktypes.RankConfig
	rankingInfoList    []*ranktypes.RankingInfo
}

func newGoldEquipForceRank(config *ranktypes.RankConfig) RankTypeData {
	goldEquipLevelRank := &GoldEquipForceRank{
		goldEquipForceList: make([]*rankentity.PlayerPropertyData, 0, 8),
		goldEquipForceTime: 0,
		rankingMap:         make(map[int64]int32),
		rankingInfoList:    make([]*ranktypes.RankingInfo, 0, 8),
		config:             config,
	}
	return goldEquipLevelRank
}

func (l *GoldEquipForceRank) init(timestamp int64) (err error) {
	l.goldEquipForceList, err = dao.GetRankDao().GetRedisRankGoldEquipForceList(timestamp, l.config)
	if err != nil {
		return
	}
	if l.goldEquipForceList == nil {
		err = l.updateRankList(timestamp)
		if err != nil {
			return
		}
	}

	for index, obj := range l.goldEquipForceList {
		l.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	l.buildRankingInfo()
	l.goldEquipForceTime = timestamp
	return
}

func (l *GoldEquipForceRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(l.goldEquipForceList) == 0 {
		return
	}
	return l.goldEquipForceList[0].PlayerId
}

func (l *GoldEquipForceRank) UpdateRankDataList(timestamp int64) (err error) {
	return l.updateRankList(timestamp)
}

func (l *GoldEquipForceRank) updateRankList(timestamp int64) (err error) {
	if l.config.StartTime != 0 || l.config.EndTime != 0 {
		if timestamp < l.config.StartTime || timestamp > l.config.EndTime {
			return
		}
	}

	diffTime := timestamp - l.goldEquipForceTime
	if diffTime < l.config.RefreshTime {
		return
	}

	oldRankInfoList := l.rankingInfoList
	l.goldEquipForceList, err = dao.GetRankDao().GetRankGoldEquipForceList(l.config)
	if err != nil {
		return
	}

	err = dao.GetRankDao().SetRedisRankGoldEquipForceList(timestamp, l.config, l.goldEquipForceList)
	if err != nil {
		return
	}

	rankingMap := make(map[int64]int32)
	for index, obj := range l.goldEquipForceList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}

	l.goldEquipForceTime = timestamp
	l.rankingMap = rankingMap
	l.buildRankingInfo()
	newRankInfoList := l.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, l.config.ClassType, l.config.RankType, l.config.GroupId)

	return
}

func (l *GoldEquipForceRank) GetPos(playerId int64) (pos int32) {
	pos = l.rankingMap[playerId]
	return
}

func (l *GoldEquipForceRank) ResetRankTime() {
	l.goldEquipForceTime = 0
}

func (l *GoldEquipForceRank) GetListAndTime() ([]*rankentity.PlayerPropertyData, int64) {
	return l.goldEquipForceList, l.goldEquipForceTime
}

func (l *GoldEquipForceRank) GetRankingMap() map[int64]int32 {
	return l.rankingMap
}

func (l *GoldEquipForceRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return l.rankingInfoList
}

func (l *GoldEquipForceRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(l.goldEquipForceList))
	for index, obj := range l.goldEquipForceList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Num)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	l.rankingInfoList = newList
}
