package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

//消费
type CostRank struct {
	costList        []*rankentity.PlayerPropertyData
	costTime        int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newCostRank(config *ranktypes.RankConfig) RankTypeData {
	d := &CostRank{
		costList:        make([]*rankentity.PlayerPropertyData, 0, 8),
		costTime:        0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func (r *CostRank) init(timestamp int64) (err error) {
	r.costList, err = dao.GetRankDao().GetRedisRankCostList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.costList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.costList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.costTime = timestamp
	r.buildRankingInfo()
	return
}

func (r *CostRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.costList) == 0 {
		return
	}
	return r.costList[0].PlayerId
}

func (r *CostRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *CostRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.costTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.costList, err = dao.GetRankDao().GetRankCostList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankCostList(timestamp, r.config, r.costList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.costList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.costTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)

	return
}

func (r *CostRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *CostRank) ResetRankTime() {
	r.costTime = 0
}

func (r *CostRank) GetListAndTime() ([]*rankentity.PlayerPropertyData, int64) {
	return r.costList, r.costTime
}

func (r *CostRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *CostRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *CostRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.costList))
	for index, obj := range r.costList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Num)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
