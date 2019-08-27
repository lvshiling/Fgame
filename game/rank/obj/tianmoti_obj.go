package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

//天魔体
type TianMoTiRank struct {
	tianMoTiList    []*rankentity.PlayerOrderData
	tianMoTiTime    int64
	rankingMap      map[int64]int32
	rankingInfoList []*ranktypes.RankingInfo
	config          *ranktypes.RankConfig
}

func newTianMoTiRank(config *ranktypes.RankConfig) RankTypeData {
	d := &TianMoTiRank{
		tianMoTiList:    make([]*rankentity.PlayerOrderData, 0, 8),
		tianMoTiTime:    0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func (r *TianMoTiRank) init(timestamp int64) (err error) {
	r.tianMoTiList, err = dao.GetRankDao().GetRedisRankTianMoTiList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.tianMoTiList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.tianMoTiList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.buildRankingInfo()
	r.tianMoTiTime = timestamp
	return
}

func (r *TianMoTiRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.tianMoTiList) == 0 {
		return
	}
	return r.tianMoTiList[0].PlayerId
}

func (r *TianMoTiRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *TianMoTiRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.tianMoTiTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.tianMoTiList, err = dao.GetRankDao().GetRankTianMoTiList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankTianMoTiList(timestamp, r.config, r.tianMoTiList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.tianMoTiList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.tianMoTiTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)
	return
}

func (r *TianMoTiRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *TianMoTiRank) ResetRankTime() {
	r.tianMoTiTime = 0
}

func (r *TianMoTiRank) GetListAndTime() ([]*rankentity.PlayerOrderData, int64) {
	return r.tianMoTiList, r.tianMoTiTime
}

func (r *TianMoTiRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *TianMoTiRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *TianMoTiRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.tianMoTiList))
	for index, obj := range r.tianMoTiList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Order)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
