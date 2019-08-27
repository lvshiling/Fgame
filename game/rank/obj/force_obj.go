package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
	rankpb "fgame/fgame/rank/protocol/pb"
)

//战力
type ForceRank struct {
	forceList       []*rankentity.PlayerForceData
	forceTime       int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newForceRank(config *ranktypes.RankConfig) RankTypeData {
	d := &ForceRank{
		forceList:       make([]*rankentity.PlayerForceData, 0, 8),
		forceTime:       0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func convertFromForceInfo(forceInfo *rankpb.RankForceInfo) *rankentity.PlayerForceData {
	forceData := &rankentity.PlayerForceData{}
	forceData.ServerId = forceInfo.ServerId
	forceData.PlayerId = forceInfo.PlayerId
	forceData.PlayerName = forceInfo.PlayerName
	forceData.GangName = forceInfo.GangName
	forceData.Force = forceInfo.Power
	forceData.Role = forceInfo.Role
	forceData.Sex = forceInfo.Sex
	return forceData
}

func (r *ForceRank) ConvertFromForceInfo(rankForce *rankpb.AreaRankForce) {
	r.forceList = make([]*rankentity.PlayerForceData, 0, 8)
	r.forceTime = rankForce.RankTime
	r.rankingMap = make(map[int64]int32)
	for index, forceInfo := range rankForce.ForceList {
		playerForceData := convertFromForceInfo(forceInfo)
		r.forceList = append(r.forceList, playerForceData)
		r.rankingMap[forceInfo.PlayerId] = int32(index + 1)
	}
}

func (r *ForceRank) init(timestamp int64) (err error) {
	r.forceList, err = dao.GetRankDao().GetRedisRankForceList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.forceList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.forceList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.forceTime = timestamp
	r.buildRankingInfo()
	return
}

func (r *ForceRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.forceList) == 0 {
		return
	}
	return r.forceList[0].PlayerId
}

func (r *ForceRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *ForceRank) updateRankList(timestamp int64) (err error) {
	diffTime := timestamp - r.forceTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldFirstId := r.GetFirstId()
	oldRankInfoList := r.rankingInfoList
	r.forceList, err = dao.GetRankDao().GetRankForceList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankForceList(timestamp, r.config, r.forceList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.forceList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.forceTime = timestamp
	r.rankingMap = rankingMap
	firstId := r.GetFirstId()
	//第一名有变
	if oldFirstId != firstId {
		err = ranklogic.RankFirstChange(firstId, oldFirstId, r.config.ClassType, ranktypes.RankTypeForce)
	}
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)

	return
}

func (r *ForceRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *ForceRank) ResetRankTime() {
	r.forceTime = 0
}

func (r *ForceRank) GetListAndTime() ([]*rankentity.PlayerForceData, int64) {
	return r.forceList, r.forceTime
}

func (r *ForceRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *ForceRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *ForceRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.forceList))
	for index, obj := range r.forceList {
		ranking := int32(index + 1)
		rankNum := obj.Force
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
