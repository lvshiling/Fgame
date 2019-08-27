package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

//噬魂幡
type ShiHunFanRank struct {
	shiHunFanList   []*rankentity.PlayerOrderData
	shiHunFanTime   int64
	rankingMap      map[int64]int32
	rankingInfoList []*ranktypes.RankingInfo
	config          *ranktypes.RankConfig
}

func newShiHunFanRank(config *ranktypes.RankConfig) RankTypeData {
	d := &ShiHunFanRank{
		shiHunFanList:   make([]*rankentity.PlayerOrderData, 0, 8),
		shiHunFanTime:   0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func (r *ShiHunFanRank) init(timestamp int64) (err error) {
	r.shiHunFanList, err = dao.GetRankDao().GetRedisRankShiHunFanList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.shiHunFanList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.shiHunFanList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.buildRankingInfo()
	r.shiHunFanTime = timestamp
	return
}

func (r *ShiHunFanRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.shiHunFanList) == 0 {
		return
	}
	return r.shiHunFanList[0].PlayerId
}

func (r *ShiHunFanRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *ShiHunFanRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.shiHunFanTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.shiHunFanList, err = dao.GetRankDao().GetRankShiHunFanList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankShiHunFanList(timestamp, r.config, r.shiHunFanList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.shiHunFanList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.shiHunFanTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)
	return
}

func (r *ShiHunFanRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *ShiHunFanRank) ResetRankTime() {
	r.shiHunFanTime = 0
}

func (r *ShiHunFanRank) GetListAndTime() ([]*rankentity.PlayerOrderData, int64) {
	return r.shiHunFanList, r.shiHunFanTime
}

func (r *ShiHunFanRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *ShiHunFanRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *ShiHunFanRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.shiHunFanList))
	for index, obj := range r.shiHunFanList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Order)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
