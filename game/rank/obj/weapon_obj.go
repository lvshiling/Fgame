package obj

import (
	"fgame/fgame/game/rank/dao"
	rankentity "fgame/fgame/game/rank/entity"
	ranklogic "fgame/fgame/game/rank/logic"
	ranktypes "fgame/fgame/game/rank/types"
	rankpb "fgame/fgame/rank/protocol/pb"
)

//兵魂
type WeaponRank struct {
	weaponList      []*rankentity.PlayerWeaponData
	weaponTime      int64
	rankingMap      map[int64]int32
	config          *ranktypes.RankConfig
	rankingInfoList []*ranktypes.RankingInfo
}

func newWeaponRank(config *ranktypes.RankConfig) RankTypeData {
	d := &WeaponRank{
		weaponList:      make([]*rankentity.PlayerWeaponData, 0, 8),
		weaponTime:      0,
		rankingMap:      make(map[int64]int32),
		rankingInfoList: make([]*ranktypes.RankingInfo, 0, 8),
		config:          config,
	}
	return d
}

func convertFromWeaponInfo(weaponInfo *rankpb.RankWeaponInfo) *rankentity.PlayerWeaponData {
	weaponData := &rankentity.PlayerWeaponData{}
	weaponData.ServerId = weaponData.ServerId
	weaponData.PlayerId = weaponData.PlayerId
	weaponData.PlayerName = weaponData.PlayerName
	weaponData.Star = weaponData.Star
	weaponData.WearId = weaponData.WearId
	weaponData.Role = weaponData.Role
	weaponData.Sex = weaponData.Sex
	return weaponData
}

func (r *WeaponRank) ConvertFromWeaponInfo(rankWeapon *rankpb.AreaRankWeapon) {
	r.weaponList = make([]*rankentity.PlayerWeaponData, 0, 8)
	r.weaponTime = rankWeapon.RankTime
	r.rankingMap = make(map[int64]int32)
	for index, weaponInfo := range rankWeapon.WeaponList {
		playerData := convertFromWeaponInfo(weaponInfo)
		r.weaponList = append(r.weaponList, playerData)
		r.rankingMap[weaponInfo.PlayerId] = int32(index + 1)
	}
}

func (r *WeaponRank) init(timestamp int64) (err error) {
	r.weaponList, err = dao.GetRankDao().GetRedisRankWeaponList(timestamp, r.config)
	if err != nil {
		return
	}
	if r.weaponList == nil {
		err = r.updateRankList(timestamp)
		if err != nil {
			return
		}
	}
	for index, obj := range r.weaponList {
		r.rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.weaponTime = timestamp
	r.buildRankingInfo()
	return
}

func (r *WeaponRank) GetFirstId() (fristId int64) {
	fristId = 0
	if len(r.weaponList) == 0 {
		return
	}
	return r.weaponList[0].PlayerId
}

func (r *WeaponRank) UpdateRankDataList(timestamp int64) (err error) {
	return r.updateRankList(timestamp)
}

func (r *WeaponRank) updateRankList(timestamp int64) (err error) {
	diffTime := timestamp - r.weaponTime
	if diffTime < r.config.RefreshTime {
		return
	}

	oldFirstId := r.GetFirstId()
	oldRankInfoList := r.rankingInfoList
	r.weaponList, err = dao.GetRankDao().GetRankWeaponList(r.config)
	if err != nil {
		return
	}
	err = dao.GetRankDao().SetRedisRankWeaponList(timestamp, r.config, r.weaponList)
	if err != nil {
		return
	}
	rankingMap := make(map[int64]int32)
	for index, obj := range r.weaponList {
		rankingMap[obj.PlayerId] = int32(index + 1)
	}
	r.weaponTime = timestamp
	r.rankingMap = rankingMap
	firstId := r.GetFirstId()
	//第一名有变
	if oldFirstId != firstId {
		err = ranklogic.RankFirstChange(firstId, oldFirstId, r.config.ClassType, ranktypes.RankTypeWeapon)
	}
	r.buildRankingInfo()
	newRankInfoList := r.rankingInfoList

	ranklogic.RankChanged(oldRankInfoList, newRankInfoList, r.config.ClassType, r.config.RankType, r.config.GroupId)
	return
}

func (r *WeaponRank) GetPos(playerId int64) (pos int32) {
	pos = r.rankingMap[playerId]
	return
}

func (r *WeaponRank) ResetRankTime() {
	r.weaponTime = 0
}

func (r *WeaponRank) GetListAndTime() ([]*rankentity.PlayerWeaponData, int64) {
	return r.weaponList, r.weaponTime
}

func (r *WeaponRank) GetRankingMap() map[int64]int32 {
	return r.rankingMap
}

func (r *WeaponRank) GetRankingInfoList() []*ranktypes.RankingInfo {
	return r.rankingInfoList
}

func (r *WeaponRank) buildRankingInfo() {
	newList := make([]*ranktypes.RankingInfo, 0, len(r.weaponList))
	for index, obj := range r.weaponList {
		ranking := int32(index + 1)
		rankNum := int64(obj.Star)
		playerId := obj.PlayerId
		playerName := obj.PlayerName
		info := ranktypes.CreateRankingInfo(playerId, playerName, ranking, rankNum)
		newList = append(newList, info)
	}
	r.rankingInfoList = newList
}
