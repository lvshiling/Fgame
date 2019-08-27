package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

//法宝
type FaBaoRank struct {
	faBaoList       []*rankentity.PlayerOrderData
	faBaoTime       int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newFaBaoRank(config *ranktypes.RankConfig) RankTypeData {
	d := &FaBaoRank{
		faBaoList:       make([]*rankentity.PlayerOrderData, 0, 8),
		faBaoTime:       0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

// func (r *FaBaoRank) ConvertFromFaBaoInfo(rankFaBao *rankpb.AreaRankFaBao) {
// 	r.faBaoList = make([]*rankentity.PlayerOrderData, 0, 8)
// 	r.faBaoTime = rankFaBao.RankTime
// 	r.rankingMap = make(map[int64]int32)
// 	for index, faBaoInfo := range rankFaBao.FaBaoList {
// 		playerData := convertFromOrderInfo(faBaoInfo)
// 		r.faBaoList = append(r.faBaoList, playerData)
// 		r.rankingMap[faBaoInfo.PlayerId] = int32(index + 1)
// 	}
// }

func (r *FaBaoRank) init(timestamp int64) (err error) {
	r.faBaoList, err = dao.GetRankDao().GetRedisRankFaBaoList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.faBaoList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.faBaoList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.faBaoTime = timestamp
	r.buildRankingInfo()
	return
}

func (r *FaBaoRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.faBaoList) == 0 {
		return
	}
	return r.faBaoList[0].PlayerId
}

func (r *FaBaoRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *FaBaoRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.faBaoTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.faBaoList, err = dao.GetRankDao().GetRankFaBaoList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankFaBaoList(timestamp, r.config, r.faBaoList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.faBaoList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.faBaoTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)

	return
}

func (r *FaBaoRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *FaBaoRank) ResetRankTime() {
	r.faBaoTime = 0
}

func (r *FaBaoRank) GetListAndTime() ([]*rankentity.PlayerOrderData, int64) {
	return r.faBaoList, r.faBaoTime
}

func (r *FaBaoRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *FaBaoRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *FaBaoRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.faBaoList))
	for index, obj := range r.faBaoList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Order)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
