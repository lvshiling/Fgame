package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
	rankpb "fgame/fgame/rank/protocol/pb"
)

//帮派
type GangRank struct {
	gangList        []*rankentity.PlayerGangData
	gangTime        int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newGangRank(config *ranktypes.RankConfig) RankTypeData {
	d := &GangRank{
		gangList:        make([]*rankentity.PlayerGangData, 0, 8),
		gangTime:        0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func convertFromGangInfo(gangInfo *rankpb.RankGangInfo) *rankentity.PlayerGangData {
	gangData := &rankentity.PlayerGangData{}
	gangData.ServerId = gangInfo.ServerId
	gangData.GangId = gangInfo.GangId
	gangData.GangName = gangInfo.GangName
	gangData.LeadId = gangInfo.LeaderId
	gangData.LeadName = gangInfo.LeaderName
	gangData.Power = gangInfo.Power
	gangData.Role = gangInfo.Role
	gangData.Sex = gangInfo.Sex
	return gangData
}

func (r *GangRank) ConvertFromGangInfo(rankForce *rankpb.AreaRankGang) {
	r.gangList = make([]*rankentity.PlayerGangData, 0, 8)
	r.gangTime = rankForce.RankTime
	r.rankingMap = make(map[int64]int32)
	for index, gangInfo := range rankForce.GangList {
		playerGangData := convertFromGangInfo(gangInfo)
		r.gangList = append(r.gangList, playerGangData)
		r.rankingMap[gangInfo.GangId] = int32(index + 1)
	}
}

func (r *GangRank) init(timestamp int64) (err error) {
	r.gangList, err = dao.GetRankDao().GetRedisRankGangList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.gangList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.gangList {
		r.rankingMap[obj.GangId] = int32(index + 1)
	}
	r.gangTime = timestamp
	r.buildRankingInfo()
	return
}

func (r *GangRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.gangList) == 0 {
		return
	}
	return r.gangList[0].GangId
}

func (r *GangRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *GangRank) updateRankList(timestamp int64) (err error) {
	diffTime := timestamp - r.gangTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.gangList, err = dao.GetRankDao().GetRankGangList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankGangList(timestamp, r.config, r.gangList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.gangList {
		rankingMap[obj.GangId] = int32(index + 1)
	}
	r.gangTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)

	return
}

func (r *GangRank) GetPos(id int64) (pos int32) {
	pos = r.rankingMap[id]
	return
}

func (r *GangRank) ResetRankTime() {
	r.gangTime = 0
}

func (r *GangRank) GetListAndTime() ([]*rankentity.PlayerGangData, int64) {
	return r.gangList, r.gangTime
}

func (r *GangRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *GangRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *GangRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.gangList))
	for index, obj := range r.gangList {
		ranking := int32(index + 1)
		rankNum := obj.Power
		playerId := obj.LeadId
		playerName := obj.LeadName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
