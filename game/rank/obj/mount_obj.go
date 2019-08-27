package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
	rankpb "fgame/fgame/rank/protocol/pb"
)

//坐骑
type MountRank struct {
	mountList       []*rankentity.PlayerOrderData
	mountTime       int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newMountRank(config *ranktypes.RankConfig) RankTypeData {
	d := &MountRank{
		mountList:       make([]*rankentity.PlayerOrderData, 0, 8),
		mountTime:       0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func convertFromOrderInfo(orderInfo *rankpb.RankOrderInfo) *rankentity.PlayerOrderData {
	orderData := &rankentity.PlayerOrderData{}
	orderData.ServerId = orderInfo.ServerId
	orderData.PlayerId = orderInfo.PlayerId
	orderData.PlayerName = orderInfo.PlayerName
	orderData.Order = orderInfo.Order
	orderData.Power = orderInfo.Power
	return orderData
}

func (r *MountRank) ConvertFromMountInfo(rankMount *rankpb.AreaRankMount) {
	r.mountList = make([]*rankentity.PlayerOrderData, 0, 8)
	r.mountTime = rankMount.RankTime
	r.rankingMap = make(map[int64]int32)
	for index, mountInfo := range rankMount.MountList {
		playerData := convertFromOrderInfo(mountInfo)
		r.mountList = append(r.mountList, playerData)
		r.rankingMap[mountInfo.PlayerId] = int32(index + 1)
	}
}

func (r *MountRank) init(timestamp int64) (err error) {
	r.mountList, err = dao.GetRankDao().GetRedisRankMountList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.mountList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.mountList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.mountTime = timestamp
	r.buildRankingInfo()
	return
}

func (r *MountRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.mountList) == 0 {
		return
	}
	return r.mountList[0].PlayerId
}

func (r *MountRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *MountRank) updateRankList(timestamp int64) (err error) {
	// 根据配置时间更新
	if r.config.StartTime != 0 || r.config.EndTime != 0 {
		if timestamp < r.config.StartTime || timestamp > r.config.EndTime {
			return
		}
	}

	diffTime := timestamp - r.mountTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldRankInfoList := r.rankingInfoList
	r.mountList, err = dao.GetRankDao().GetRankMountList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankMountList(timestamp, r.config, r.mountList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.mountList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.mountTime = timestamp
	r.rankingMap = rankingMap
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)
	return
}

func (r *MountRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *MountRank) ResetRankTime() {
	r.mountTime = 0
}

func (r *MountRank) GetListAndTime() ([]*rankentity.PlayerOrderData, int64) {
	return r.mountList, r.mountTime
}

func (r *MountRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *MountRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *MountRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.mountList))
	for index, obj := range r.mountList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Order)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
