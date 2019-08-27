package scene

import (
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/types"
	"fmt"
	"sort"
)

const (
	MAX_RANKING_TOP = 10
	MAX_RANKING     = 100
)

//场景排行玩家信息
type SceneRankPlayerInfo struct {
	playerId   int64
	playerName string
	value      int64
}

func (info *SceneRankPlayerInfo) GetPlayerId() int64 {
	return info.playerId
}

func (info *SceneRankPlayerInfo) GetPlayerName() string {
	return info.playerName
}

func (info *SceneRankPlayerInfo) GetValue() int64 {
	return info.value
}

type SceneRankPlayerInfoList []*SceneRankPlayerInfo

func (l SceneRankPlayerInfoList) Len() int {
	return len(l)
}

func (l SceneRankPlayerInfoList) Less(i, j int) bool {
	return l[i].value < l[j].value
}

func (l SceneRankPlayerInfoList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func createSceneRankPlayerInfo(playerId int64, playerName string, val int64) *SceneRankPlayerInfo {
	info := &SceneRankPlayerInfo{}
	info.playerId = playerId
	info.playerName = playerName
	info.value = val
	return info
}

type SceneRankManager interface {
	GetScene() Scene
	GetAllRanks() map[types.SceneRankType]*SceneRank
	//获取排行
	GetPlayerRank(rankType types.SceneRankType, playerId int64) (rank int32, val int64)
	//获取排行榜
	GetRankList(rankType types.SceneRankType) []*SceneRankPlayerInfo
	//获取排行榜
	GetAllRankList(rankType types.SceneRankType) []*SceneRankPlayerInfo

	//移除玩家
	RemovePlayer(rankType types.SceneRankType, playerId int64)
	//更新值
	UpdatePlayer(rankType types.SceneRankType, playerId int64, playerName string, val int64)
	//排序
	Sort()
}

type sceneRankManager struct {
	s       Scene
	rankMap map[types.SceneRankType]*SceneRank
}

func (m *sceneRankManager) GetScene() Scene {
	return m.s
}

func (m *sceneRankManager) GetAllRanks() map[types.SceneRankType]*SceneRank {
	return m.rankMap
}

func (m *sceneRankManager) GetPlayerRank(rankType types.SceneRankType, playerId int64) (rank int32, val int64) {
	r, ok := m.rankMap[rankType]
	if !ok {
		return -1, 0
	}
	return r.GetPlayerRank(playerId)
}

func (m *sceneRankManager) GetRankList(rankType types.SceneRankType) []*SceneRankPlayerInfo {
	r, ok := m.rankMap[rankType]
	if !ok {
		return nil
	}
	return r.GetRankList()
}

func (m *sceneRankManager) GetAllRankList(rankType types.SceneRankType) []*SceneRankPlayerInfo {
	r, ok := m.rankMap[rankType]
	if !ok {
		return nil
	}
	return r.GetAllRankList()
}

func (m *sceneRankManager) UpdatePlayer(rankType types.SceneRankType, playerId int64, playerName string, val int64) {
	r, ok := m.rankMap[rankType]
	if !ok {
		r = newSceneRank(rankType, MAX_RANKING, MAX_RANKING_TOP)
		m.rankMap[rankType] = r
	}

	r.UpdatePlayer(playerId, playerName, val)
	return
}

func (m *sceneRankManager) RemovePlayer(rankType types.SceneRankType, playerId int64) {
	r, ok := m.rankMap[rankType]
	if !ok {
		return
	}

	r.RemovePlayer(playerId)
	return
}

func (m *sceneRankManager) Sort() {
	for _, r := range m.rankMap {
		flag := r.Sort()
		if flag {
			//发送排行变化
			gameevent.Emit(sceneeventtypes.EventTypeSceneRankChanged, m.s, r)
		}
	}
	return
}

func newSceneRankManager(s Scene) SceneRankManager {
	m := &sceneRankManager{}
	m.s = s
	m.rankMap = make(map[types.SceneRankType]*SceneRank)
	return m
}

type SceneRank struct {
	rankMap            map[int64]int32
	rankPlayerInfoList []*SceneRankPlayerInfo
	maxRank            int32
	topRank            int32
	rankType           types.SceneRankType
	dirty              bool
}

func (r *SceneRank) GetRankType() types.SceneRankType {
	return r.rankType
}

func (r *SceneRank) GetPlayerRank(playerId int64) (rank int32, val int64) {
	index, ok := r.rankMap[playerId]
	if !ok {
		return -1, 0
	}
	playerInfo := r.rankPlayerInfoList[index]
	return index, playerInfo.value
}

func (r *SceneRank) getPlayerInfo(playerId int64) *SceneRankPlayerInfo {
	index, ok := r.rankMap[playerId]
	if !ok {
		return nil
	}

	return r.rankPlayerInfoList[index]
}

func (r *SceneRank) UpdatePlayer(playerId int64, playerName string, val int64) {

	found := false
	//TODO: 优化
	for _, playerInfo := range r.rankPlayerInfoList {
		if playerInfo.GetPlayerId() == playerId {
			playerInfo.value = val
			found = true
			break
		}
	}
	if !found {
		//TODO: zrc循环使用
		info := createSceneRankPlayerInfo(playerId, playerName, val)
		r.rankPlayerInfoList = append(r.rankPlayerInfoList, info)
	}
	r.dirty = true
}

func (r *SceneRank) RemovePlayer(playerId int64) {
	delIndex := -1
	for index, playerInfo := range r.rankPlayerInfoList {
		if playerInfo.GetPlayerId() != playerId {
			continue
		}
		delIndex = index
	}

	if delIndex != -1 {
		r.rankPlayerInfoList = append(r.rankPlayerInfoList[:delIndex], r.rankPlayerInfoList[delIndex+1:]...)
		r.dirty = true
	}
}

func (r *SceneRank) GetRankList() []*SceneRankPlayerInfo {
	if len(r.rankPlayerInfoList) > int(r.topRank) {
		return r.rankPlayerInfoList[:r.topRank]
	}
	return r.rankPlayerInfoList
}

func (r *SceneRank) GetAllRankList() []*SceneRankPlayerInfo {

	return r.rankPlayerInfoList
}

func (r *SceneRank) Sort() bool {
	if !r.dirty {
		return false
	}
	sort.Sort(sort.Reverse(SceneRankPlayerInfoList(r.rankPlayerInfoList)))
	if len(r.rankPlayerInfoList) > int(r.maxRank) {
		r.rankPlayerInfoList = r.rankPlayerInfoList[:r.maxRank]
		//TODO zrc:回收
	}
	r.rankMap = make(map[int64]int32)
	for i, playerInfo := range r.rankPlayerInfoList {
		r.rankMap[playerInfo.GetPlayerId()] = int32(i)
	}
	r.dirty = false
	return true
}

func newSceneRank(rankType types.SceneRankType, maxRank int32, topRank int32) *SceneRank {
	if maxRank >= MAX_RANKING {
		maxRank = MAX_RANKING
	}
	if topRank >= MAX_RANKING_TOP {
		topRank = MAX_RANKING_TOP
	}
	if maxRank <= 0 || topRank <= 0 {
		panic(fmt.Errorf("排名不能为0"))
	}
	if topRank >= maxRank {
		topRank = maxRank
	}
	r := &SceneRank{}
	r.maxRank = maxRank
	r.topRank = topRank
	r.rankType = rankType
	r.rankPlayerInfoList = make([]*SceneRankPlayerInfo, 0, r.maxRank)
	r.rankMap = make(map[int64]int32)
	return r
}
