package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
	rankpb "fgame/fgame/rank/protocol/pb"
)

//护体盾
type BodyShieldRank struct {
	bodyShieldList  []*rankentity.PlayerOrderData
	bodyShieldTime  int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newBodyShieldRank(config *ranktypes.RankConfig) RankTypeData {
	d := &BodyShieldRank{
		bodyShieldList:  make([]*rankentity.PlayerOrderData, 0, 8),
		bodyShieldTime:  0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func (r *BodyShieldRank) ConvertFromBodyShieldInfo(rankBodyShield *rankpb.AreaRankBodyShield) {
	r.bodyShieldList = make([]*rankentity.PlayerOrderData, 0, 8)
	r.bodyShieldTime = rankBodyShield.RankTime
	r.rankingMap = make(map[int64]int32)
	for index, bodyShieldInfo := range rankBodyShield.BodyShieldList {
		playerData := convertFromOrderInfo(bodyShieldInfo)
		r.bodyShieldList = append(r.bodyShieldList, playerData)
		r.rankingMap[bodyShieldInfo.PlayerId] = int32(index + 1)
	}
}

func (r *BodyShieldRank) init(timestamp int64) (err error) {
	r.bodyShieldList, err = dao.GetRankDao().GetRedisRankBodyShieldList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.bodyShieldList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.bodyShieldList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.buildRankingInfo()
	r.bodyShieldTime = timestamp
	return
}

func (r *BodyShieldRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.bodyShieldList) == 0 {
		return
	}
	return r.bodyShieldList[0].PlayerId
}

func (r *BodyShieldRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *BodyShieldRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.bodyShieldTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.bodyShieldList, err = dao.GetRankDao().GetRankBodyShieldList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankBodyShieldList(timestamp, r.config, r.bodyShieldList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.bodyShieldList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.bodyShieldTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)
	return
}

func (r *BodyShieldRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *BodyShieldRank) ResetRankTime() {
	r.bodyShieldTime = 0
}

func (r *BodyShieldRank) GetListAndTime() ([]*rankentity.PlayerOrderData, int64) {
	return r.bodyShieldList, r.bodyShieldTime
}

func (r *BodyShieldRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *BodyShieldRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *BodyShieldRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.bodyShieldList))
	for index, obj := range r.bodyShieldList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Order)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
