package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

//次数
type CountRank struct {
	countList       []*rankentity.PlayerPropertyData
	countTime       int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newCountRank(config *ranktypes.RankConfig) RankTypeData {
	d := &CountRank{
		countList:       make([]*rankentity.PlayerPropertyData, 0, 8),
		countTime:       0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func (r *CountRank) init(timestamp int64) (err error) {
	r.countList, err = dao.GetRankDao().GetRedisRankCountList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.countList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.countList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.countTime = timestamp
	r.buildRankingInfo()
	return
}

func (r *CountRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.countList) == 0 {
		return
	}
	return r.countList[0].PlayerId
}

func (r *CountRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *CountRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.countTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.countList, err = dao.GetRankDao().GetRankCountList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankCountList(timestamp, r.config, r.countList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.countList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.countTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)

	return
}

func (r *CountRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *CountRank) ResetRankTime() {
	r.countTime = 0
}

func (r *CountRank) GetListAndTime() ([]*rankentity.PlayerPropertyData, int64) {
	return r.countList, r.countTime
}

func (r *CountRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *CountRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *CountRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.countList))
	for index, obj := range r.countList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Num)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
