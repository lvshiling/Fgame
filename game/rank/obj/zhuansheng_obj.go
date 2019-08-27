package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
)

//次数
type ZhuanShengRank struct {
	zhuanShengList  []*rankentity.PlayerPropertyData
	zhuanShengTime  int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newZhuanShengRank(config *ranktypes.RankConfig) RankTypeData {
	d := &ZhuanShengRank{
		zhuanShengList:  make([]*rankentity.PlayerPropertyData, 0, 8),
		zhuanShengTime:  0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func (r *ZhuanShengRank) init(timestamp int64) (err error) {
	r.zhuanShengList, err = dao.GetRankDao().GetRedisRankZhuanShengList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.zhuanShengList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.zhuanShengList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.zhuanShengTime = timestamp
	r.buildRankingInfo()
	return
}

func (r *ZhuanShengRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.zhuanShengList) == 0 {
		return
	}
	return r.zhuanShengList[0].PlayerId
}

func (r *ZhuanShengRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *ZhuanShengRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.zhuanShengTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.zhuanShengList, err = dao.GetRankDao().GetRankZhuanShengList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankZhuanShengList(timestamp, r.config, r.zhuanShengList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.zhuanShengList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.zhuanShengTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)

	return
}

func (r *ZhuanShengRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *ZhuanShengRank) ResetRankTime() {
	r.zhuanShengTime = 0
}

func (r *ZhuanShengRank) GetListAndTime() ([]*rankentity.PlayerPropertyData, int64) {
	return r.zhuanShengList, r.zhuanShengTime
}

func (r *ZhuanShengRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *ZhuanShengRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *ZhuanShengRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.zhuanShengList))
	for index, obj := range r.zhuanShengList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Num)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
