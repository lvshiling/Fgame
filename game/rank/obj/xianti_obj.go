package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

//仙体
type XianTiRank struct {
	xianTiList      []*rankentity.PlayerOrderData
	xianTiTime      int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newXianTiRank(config *ranktypes.RankConfig) RankTypeData {
	d := &XianTiRank{
		xianTiList:      make([]*rankentity.PlayerOrderData, 0, 8),
		xianTiTime:      0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func (r *XianTiRank) init(timestamp int64) (err error) {
	r.xianTiList, err = dao.GetRankDao().GetRedisRankXianTiList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.xianTiList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.xianTiList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.xianTiTime = timestamp
	r.buildRankingInfo()
	return
}

func (r *XianTiRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.xianTiList) == 0 {
		return
	}
	return r.xianTiList[0].PlayerId
}

func (r *XianTiRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *XianTiRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.xianTiTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.xianTiList, err = dao.GetRankDao().GetRankXianTiList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankXianTiList(timestamp, r.config, r.xianTiList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.xianTiList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.xianTiTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)
	return
}

func (r *XianTiRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *XianTiRank) ResetRankTime() {
	r.xianTiTime = 0
}

func (r *XianTiRank) GetListAndTime() ([]*rankentity.PlayerOrderData, int64) {
	return r.xianTiList, r.xianTiTime
}

func (r *XianTiRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *XianTiRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *XianTiRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.xianTiList))
	for index, obj := range r.xianTiList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Order)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
