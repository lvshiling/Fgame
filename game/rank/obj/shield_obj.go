package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
	rankpb "fgame/fgame/rank/protocol/pb"
)

//神盾尖刺
type ShieldRank struct {
	shieldList      []*rankentity.PlayerOrderData
	shieldTime      int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newShieldRank(config *ranktypes.RankConfig) RankTypeData {
	d := &ShieldRank{
		shieldList:      make([]*rankentity.PlayerOrderData, 0, 8),
		shieldTime:      0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func (r *ShieldRank) ConvertFromShieldInfo(rankShield *rankpb.AreaRankShield) {
	r.shieldList = make([]*rankentity.PlayerOrderData, 0, 8)
	r.shieldTime = rankShield.RankTime
	r.rankingMap = make(map[int64]int32)
	for index, shieldInfo := range rankShield.ShieldList {
		playerData := convertFromOrderInfo(shieldInfo)
		r.shieldList = append(r.shieldList, playerData)
		r.rankingMap[shieldInfo.PlayerId] = int32(index + 1)
	}
}

func (r *ShieldRank) init(timestamp int64) (err error) {
	r.shieldList, err = dao.GetRankDao().GetRedisRankShieldList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.shieldList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.shieldList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.shieldTime = timestamp
	r.buildRankingInfo()
	return
}

func (r *ShieldRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.shieldList) == 0 {
		return
	}
	return r.shieldList[0].PlayerId
}

func (r *ShieldRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *ShieldRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.shieldTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.shieldList, err = dao.GetRankDao().GetRankShieldList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankShieldList(timestamp, r.config, r.shieldList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.shieldList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.shieldTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)
	return
}

func (r *ShieldRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *ShieldRank) ResetRankTime() {
	r.shieldTime = 0
}

func (r *ShieldRank) GetListAndTime() ([]*rankentity.PlayerOrderData, int64) {
	return r.shieldList, r.shieldTime
}

func (r *ShieldRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *ShieldRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *ShieldRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.shieldList))
	for index, obj := range r.shieldList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Order)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
