package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

//魅力
type CharmRank struct {
	charmList       []*rankentity.PlayerPropertyData
	charmTime       int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newCharmRank(config *ranktypes.RankConfig) RankTypeData {
	d := &CharmRank{
		charmList:       make([]*rankentity.PlayerPropertyData, 0, 8),
		charmTime:       0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func (r *CharmRank) init(timestamp int64) (err error) {
	r.charmList, err = dao.GetRankDao().GetRedisRankCharmList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.charmList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.charmList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.charmTime = timestamp
	r.buildRankingInfo()
	return
}

func (r *CharmRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.charmList) == 0 {
		return
	}
	return r.charmList[0].PlayerId
}

func (r *CharmRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *CharmRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.charmTime
	if diffTime < r.config.RefreshTime {
		return
	}
	oldRankInfoList := r.rankingInfoList
	r.charmList, err = dao.GetRankDao().GetRankCharmList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankCharmList(timestamp, r.config, r.charmList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.charmList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.charmTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)
	return
}

func (r *CharmRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *CharmRank) ResetRankTime() {
	r.charmTime = 0
}

func (r *CharmRank) GetListAndTime() ([]*rankentity.PlayerPropertyData, int64) {
	return r.charmList, r.charmTime
}

func (r *CharmRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *CharmRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *CharmRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.charmList))
	for index, obj := range r.charmList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Num)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
