package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
	rankpb "fgame/fgame/rank/protocol/pb"
)

//战翼
type WingRank struct {
	wingList        []*rankentity.PlayerOrderData
	wingTime        int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newWingRank(config *ranktypes.RankConfig) RankTypeData {
	d := &WingRank{
		wingList:        make([]*rankentity.PlayerOrderData, 0, 8),
		wingTime:        0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func (r *WingRank) ConvertFromWingInfo(rankWing *rankpb.AreaRankWing) {
	r.wingList = make([]*rankentity.PlayerOrderData, 0, 8)
	r.wingTime = rankWing.RankTime
	r.rankingMap = make(map[int64]int32)
	for index, wingInfo := range rankWing.WingList {
		playerData := convertFromOrderInfo(wingInfo)
		r.wingList = append(r.wingList, playerData)
		r.rankingMap[wingInfo.PlayerId] = int32(index + 1)
	}
}

func (r *WingRank) init(timestamp int64) (err error) {
	r.wingList, err = dao.GetRankDao().GetRedisRankWingList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.wingList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.wingList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.wingTime = timestamp
	r.buildRankingInfo()
	return
}

func (r *WingRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.wingList) == 0 {
		return
	}
	return r.wingList[0].PlayerId
}

func (r *WingRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *WingRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.wingTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.wingList, err = dao.GetRankDao().GetRankWingList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankWingList(timestamp, r.config, r.wingList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.wingList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.wingTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)
	return
}

func (r *WingRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *WingRank) ResetRankTime() {
	r.wingTime = 0
}

func (r *WingRank) GetListAndTime() ([]*rankentity.PlayerOrderData, int64) {
	return r.wingList, r.wingTime
}

func (r *WingRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *WingRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *WingRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.wingList))
	for index, obj := range r.wingList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Order)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
