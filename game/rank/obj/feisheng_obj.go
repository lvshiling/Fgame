package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

//飞升
type FeiShengRank struct {
	feiShengList    []*rankentity.PlayerPropertyData
	feiShengTime    int64
	rankingMap      map[int64]int32
	rankingInfoList []*ranktypes.RankingInfo
	config          *ranktypes.RankConfig
}

func newFeiShengRank(config *ranktypes.RankConfig) RankTypeData {
	d := &FeiShengRank{
		feiShengList:    make([]*rankentity.PlayerPropertyData, 0, 8),
		feiShengTime:    0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func (r *FeiShengRank) init(timestamp int64) (err error) {
	r.feiShengList, err = dao.GetRankDao().GetRedisRankFeiShengList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.feiShengList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.feiShengList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.buildRankingInfo()
	r.feiShengTime = timestamp
	return
}

func (r *FeiShengRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.feiShengList) == 0 {
		return
	}
	return r.feiShengList[0].PlayerId
}

func (r *FeiShengRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *FeiShengRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.feiShengTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.feiShengList, err = dao.GetRankDao().GetRankFeiShengList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankFeiShengList(timestamp, r.config, r.feiShengList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.feiShengList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.feiShengTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)

	return
}

func (r *FeiShengRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *FeiShengRank) ResetRankTime() {
	r.feiShengTime = 0
}

func (r *FeiShengRank) GetListAndTime() ([]*rankentity.PlayerPropertyData, int64) {
	return r.feiShengList, r.feiShengTime
}

func (r *FeiShengRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *FeiShengRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *FeiShengRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.feiShengList))
	for index, obj := range r.feiShengList {
		ranking := int32(index + 1)
		rankNum := obj.Num
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
