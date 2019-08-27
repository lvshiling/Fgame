package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
	rankpb "fgame/fgame/rank/protocol/pb"
)

//护体仙羽
type FeatherRank struct {
	featherList     []*rankentity.PlayerOrderData
	featherTime     int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newFeatherRank(config *ranktypes.RankConfig) RankTypeData {
	d := &FeatherRank{
		featherList:     make([]*rankentity.PlayerOrderData, 0, 8),
		featherTime:     0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func (r *FeatherRank) ConvertFromFeatherInfo(rankFeather *rankpb.AreaRankFeather) {
	r.featherList = make([]*rankentity.PlayerOrderData, 0, 8)
	r.featherTime = rankFeather.RankTime
	r.rankingMap = make(map[int64]int32)
	for index, featherInfo := range rankFeather.FeatherList {
		playerData := convertFromOrderInfo(featherInfo)
		r.featherList = append(r.featherList, playerData)
		r.rankingMap[featherInfo.PlayerId] = int32(index + 1)
	}
}

func (r *FeatherRank) init(timestamp int64) (err error) {
	r.featherList, err = dao.GetRankDao().GetRedisRankFeatherList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.featherList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.featherList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.featherTime = timestamp
	r.buildRankingInfo()
	return
}

func (r *FeatherRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.featherList) == 0 {
		return
	}
	return r.featherList[0].PlayerId
}

func (r *FeatherRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *FeatherRank) updateRankList(timestamp int64) (err error) {
	diffTime := timestamp - r.featherTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.featherList, err = dao.GetRankDao().GetRankFeatherList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankFeatherList(timestamp, r.config, r.featherList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.featherList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.featherTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)

	return
}

func (r *FeatherRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *FeatherRank) ResetRankTime() {
	r.featherTime = 0
}

func (r *FeatherRank) GetListAndTime() ([]*rankentity.PlayerOrderData, int64) {
	return r.featherList, r.featherTime
}

func (r *FeatherRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *FeatherRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *FeatherRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.featherList))
	for index, obj := range r.featherList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Order)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
