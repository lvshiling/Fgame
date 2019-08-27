package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
	rankpb "fgame/fgame/rank/protocol/pb"
)

//暗器
type AnQiRank struct {
	anQiList        []*rankentity.PlayerOrderData
	anQiTime        int64
	rankingMap      map[int64]int32
	rankingInfoList []*ranktypes.RankingInfo
	config          *ranktypes.RankConfig
}

func newAnQiRank(config *ranktypes.RankConfig) RankTypeData {
	d := &AnQiRank{
		anQiList:        make([]*rankentity.PlayerOrderData, 0, 8),
		anQiTime:        0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func (r *AnQiRank) ConvertFromShieldInfo(rankAnQi *rankpb.AreaRankAnQi) {
	r.anQiList = make([]*rankentity.PlayerOrderData, 0, 8)
	r.anQiTime = rankAnQi.RankTime
	r.rankingMap = make(map[int64]int32)
	for index, anQiInfo := range rankAnQi.AnQiList {
		playerData := convertFromOrderInfo(anQiInfo)
		r.anQiList = append(r.anQiList, playerData)
		r.rankingMap[anQiInfo.PlayerId] = int32(index + 1)
	}
}

func (r *AnQiRank) init(timestamp int64) (err error) {
	r.anQiList, err = dao.GetRankDao().GetRedisRankAnQiList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.anQiList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.anQiList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.buildRankingInfo()
	r.anQiTime = timestamp
	return
}

func (r *AnQiRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.anQiList) == 0 {
		return
	}
	return r.anQiList[0].PlayerId
}

func (r *AnQiRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *AnQiRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.anQiTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.anQiList, err = dao.GetRankDao().GetRankAnQiList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankAnQiList(timestamp, r.config, r.anQiList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.anQiList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.anQiTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)
	return
}

func (r *AnQiRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *AnQiRank) ResetRankTime() {
	r.anQiTime = 0
}

func (r *AnQiRank) GetListAndTime() ([]*rankentity.PlayerOrderData, int64) {
	return r.anQiList, r.anQiTime
}

func (r *AnQiRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *AnQiRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *AnQiRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.anQiList))
	for index, obj := range r.anQiList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Order)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
