package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

//充值
type ChargeRank struct {
	chargeList      []*rankentity.PlayerPropertyData
	chargeTime      int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newChargeRank(config *ranktypes.RankConfig) RankTypeData {
	d := &ChargeRank{
		chargeList:      make([]*rankentity.PlayerPropertyData, 0, 8),
		chargeTime:      0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func (r *ChargeRank) init(timestamp int64) (err error) {
	r.chargeList, err = dao.GetRankDao().GetRedisRankChargeList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.chargeList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.chargeList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.buildRankingInfo()
	r.chargeTime = timestamp

	return
}

func (r *ChargeRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.chargeList) == 0 {
		return
	}
	return r.chargeList[0].PlayerId
}

func (r *ChargeRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *ChargeRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.chargeTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.chargeList, err = dao.GetRankDao().GetRankChargeList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankChargeList(timestamp, r.config, r.chargeList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.chargeList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.chargeTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)

	return
}

func (r *ChargeRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *ChargeRank) ResetRankTime() {
	r.chargeTime = 0
}

func (r *ChargeRank) GetListAndTime() ([]*rankentity.PlayerPropertyData, int64) {
	return r.chargeList, r.chargeTime
}

func (r *ChargeRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *ChargeRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *ChargeRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.chargeList))
	for index, obj := range r.chargeList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Num)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
