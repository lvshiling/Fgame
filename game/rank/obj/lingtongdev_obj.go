package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

//灵童养成类
type LingTongDevRank struct {
	lingTongDevList []*rankentity.PlayerOrderData
	lingTongDevTime int64
	rankingMap      map[int64]int32
	rankingInfoList []*ranktypes.RankingInfo
	config          *ranktypes.RankConfig
}

func newLingTongDevRank(config *ranktypes.RankConfig) RankTypeData {
	d := &LingTongDevRank{
		lingTongDevList: make([]*rankentity.PlayerOrderData, 0, 8),
		lingTongDevTime: 0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func (r *LingTongDevRank) init(timestamp int64) (err error) {
	r.lingTongDevList, err = dao.GetRankDao().GetRedisRankLingTongDevList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.lingTongDevList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.lingTongDevList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.buildRankingInfo()
	r.lingTongDevTime = timestamp
	return
}

func (r *LingTongDevRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.lingTongDevList) == 0 {
		return
	}
	return r.lingTongDevList[0].PlayerId
}

func (r *LingTongDevRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *LingTongDevRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.lingTongDevTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.lingTongDevList, err = dao.GetRankDao().GetRankLingTongDevList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankLingTongDevList(timestamp, r.config, r.lingTongDevList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.lingTongDevList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.lingTongDevTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)
	return
}

func (r *LingTongDevRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *LingTongDevRank) ResetRankTime() {
	r.lingTongDevTime = 0
}

func (r *LingTongDevRank) GetListAndTime() ([]*rankentity.PlayerOrderData, int64) {
	return r.lingTongDevList, r.lingTongDevTime
}

func (r *LingTongDevRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *LingTongDevRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *LingTongDevRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.lingTongDevList))
	for index, obj := range r.lingTongDevList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Order)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
