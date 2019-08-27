package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
	rankpb "fgame/fgame/rank/protocol/pb"
)

//身法
type ShenFaRank struct {
	shenFaList      []*rankentity.PlayerOrderData
	shenFaTime      int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newShenFaRank(config *ranktypes.RankConfig) RankTypeData {
	d := &ShenFaRank{
		shenFaList:      make([]*rankentity.PlayerOrderData, 0, 8),
		shenFaTime:      0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func (r *ShenFaRank) ConvertFromShenFaInfo(rankShenFa *rankpb.AreaRankShenFa) {
	r.shenFaList = make([]*rankentity.PlayerOrderData, 0, 8)
	r.shenFaTime = rankShenFa.RankTime
	r.rankingMap = make(map[int64]int32)
	for index, shenFaInfo := range rankShenFa.ShenFaList {
		playerData := convertFromOrderInfo(shenFaInfo)
		r.shenFaList = append(r.shenFaList, playerData)
		r.rankingMap[shenFaInfo.PlayerId] = int32(index + 1)
	}
}

func (r *ShenFaRank) init(timestamp int64) (err error) {
	r.shenFaList, err = dao.GetRankDao().GetRedisRankShenFaList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.shenFaList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.shenFaList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.shenFaTime = timestamp
	r.buildRankingInfo()
	return
}

func (r *ShenFaRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.shenFaList) == 0 {
		return
	}
	return r.shenFaList[0].PlayerId
}

func (r *ShenFaRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *ShenFaRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.shenFaTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.shenFaList, err = dao.GetRankDao().GetRankShenFaList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankShenFaList(timestamp, r.config, r.shenFaList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.shenFaList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.shenFaTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)
	return
}

func (r *ShenFaRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *ShenFaRank) ResetRankTime() {
	r.shenFaTime = 0
}

func (r *ShenFaRank) GetListAndTime() ([]*rankentity.PlayerOrderData, int64) {
	return r.shenFaList, r.shenFaTime
}

func (r *ShenFaRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *ShenFaRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *ShenFaRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.shenFaList))
	for index, obj := range r.shenFaList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Order)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
