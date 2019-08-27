package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
	rankpb "fgame/fgame/rank/protocol/pb"
)

//领域
type LingYuRank struct {
	lingYuList      []*rankentity.PlayerOrderData
	lingYuTime      int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newLingYuRank(config *ranktypes.RankConfig) RankTypeData {
	d := &LingYuRank{
		lingYuList:      make([]*rankentity.PlayerOrderData, 0, 8),
		lingYuTime:      0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func (r *LingYuRank) ConvertFromLingYuInfo(rankLingYu *rankpb.AreaRankLingYu) {
	r.lingYuList = make([]*rankentity.PlayerOrderData, 0, 8)
	r.lingYuTime = rankLingYu.RankTime
	r.rankingMap = make(map[int64]int32)
	for index, lingYuInfo := range rankLingYu.LingYuList {
		playerData := convertFromOrderInfo(lingYuInfo)
		r.lingYuList = append(r.lingYuList, playerData)
		r.rankingMap[lingYuInfo.PlayerId] = int32(index + 1)
	}
}

func (r *LingYuRank) init(timestamp int64) (err error) {
	r.lingYuList, err = dao.GetRankDao().GetRedisRankLingYuList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.lingYuList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.lingYuList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.lingYuTime = timestamp
	r.buildRankingInfo()
	return
}

func (r *LingYuRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.lingYuList) == 0 {
		return
	}
	return r.lingYuList[0].PlayerId
}

func (r *LingYuRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *LingYuRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.lingYuTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.lingYuList, err = dao.GetRankDao().GetRankLingYuList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankLingYuList(timestamp, r.config, r.lingYuList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.lingYuList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.lingYuTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)
	return
}

func (r *LingYuRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *LingYuRank) ResetRankTime() {
	r.lingYuTime = 0
}

func (r *LingYuRank) GetListAndTime() ([]*rankentity.PlayerOrderData, int64) {
	return r.lingYuList, r.lingYuTime
}

func (r *LingYuRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *LingYuRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *LingYuRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.lingYuList))
	for index, obj := range r.lingYuList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Order)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
