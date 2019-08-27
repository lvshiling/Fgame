package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

//灵童等级
type LingTongLevelRank struct {
	levelList       []*rankentity.PlayerPropertyData
	levelTime       int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newLingTongLevelRank(config *ranktypes.RankConfig) RankTypeData {
	d := &LingTongLevelRank{
		levelList:       make([]*rankentity.PlayerPropertyData, 0, 8),
		levelTime:       0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func (r *LingTongLevelRank) init(timestamp int64) (err error) {
	r.levelList, err = dao.GetRankDao().GetRedisRankLingTongLevelList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.levelList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.levelList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.levelTime = timestamp
	r.buildRankingInfo()
	return
}

func (r *LingTongLevelRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.levelList) == 0 {
		return
	}
	return r.levelList[0].PlayerId
}

func (r *LingTongLevelRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *LingTongLevelRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.levelTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.levelList, err = dao.GetRankDao().GetRankLingTongLevelList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankLingTongLevelList(timestamp, r.config, r.levelList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.levelList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.levelTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)

	return
}

func (r *LingTongLevelRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *LingTongLevelRank) ResetRankTime() {
	r.levelTime = 0
}

func (r *LingTongLevelRank) GetListAndTime() ([]*rankentity.PlayerPropertyData, int64) {
	return r.levelList, r.levelTime
}

func (r *LingTongLevelRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *LingTongLevelRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *LingTongLevelRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.levelList))
	for index, obj := range r.levelList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Num)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
