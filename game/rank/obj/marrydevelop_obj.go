package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

//表白排行
type MarryDevelopRank struct {
	marryDevelopList []*rankentity.PlayerPropertyData
	marryDevelopTime int64
	rankingMap       map[int64]int32
	config           *ranktypes.RankConfig
	rankingInfoList  []*ranktypes.RankingInfo
}

func newMarryDevelopRank(config *ranktypes.RankConfig) RankTypeData {
	d := &MarryDevelopRank{
		marryDevelopList: make([]*rankentity.PlayerPropertyData, 0, 8),
		marryDevelopTime: 0,
		rankingMap:       make(map[int64]int32),
		rankingInfoList:  make([]*ranktypes.RankingInfo, 0, 8),
		config:           config,
	}
	return d
}

func (r *MarryDevelopRank) init(timestamp int64) (err error) {
	r.marryDevelopList, err = dao.GetRankDao().GetRedisRankMarryDevelopList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.marryDevelopList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.marryDevelopList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.marryDevelopTime = timestamp
	r.buildRankingInfo()
	return
}

func (r *MarryDevelopRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.marryDevelopList) == 0 {
		return
	}
	return r.marryDevelopList[0].PlayerId
}

func (r *MarryDevelopRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *MarryDevelopRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.marryDevelopTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.marryDevelopList, err = dao.GetRankDao().GetRankMarryDevelopList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankMarryDevelopList(timestamp, r.config, r.marryDevelopList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.marryDevelopList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.marryDevelopTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)
	return
}

func (r *MarryDevelopRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *MarryDevelopRank) ResetRankTime() {
	r.marryDevelopTime = 0
}

func (r *MarryDevelopRank) GetListAndTime() ([]*rankentity.PlayerPropertyData, int64) {
	return r.marryDevelopList, r.marryDevelopTime
}

func (r *MarryDevelopRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *MarryDevelopRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *MarryDevelopRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.marryDevelopList))
	for index, obj := range r.marryDevelopList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Num)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
