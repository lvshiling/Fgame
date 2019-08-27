package player

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/funcopen/funcopen"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/pkg/mathutils"
	"sync"

	"github.com/golang/protobuf/proto"
)

//在线玩家管理器
type OnlinePlayerManager interface {
	GetPlayerByUserId(id int64) Player
	GetPlayerById(playerId int64) Player
	GetPlayerByName(name string) Player
	AddPlayer(p Player) bool
	PlayerEnterServer(p Player) bool
	PlayerLeaveServer(p Player)
	RemovePlayer(p Player)
	BroadcastMsg(msg proto.Message)
	BroadcastMsgExclude(excludeIdList []int64, msg proto.Message)
	BroadcastMsgRelated(messageCall message.ScheduleMessageCallBack, result interface{})
	RecommentPlayersExclude(map[int64]struct{}) (pList []Player)
	RecommentSpouses(pl Player, excludePlayers map[int64]struct{}) (pList []Player)
	Count() int32
	GetAllPlayers() []Player
	CountNeigua() int32
	OfflineAllPlayers()
}

type onlinePlayerManager struct {
	rwm sync.RWMutex
	//用户map
	userMap map[int64]Player
	//角色map
	playerMap map[int64]Player
	//名字map
	namePlayerMap map[string]Player
}

const (
	recommentNum       = 6
	RecommentSpouseNum = 3
)

func (ops *onlinePlayerManager) GetAllPlayers() []Player {
	ops.rwm.RLock()
	defer ops.rwm.RUnlock()
	pList := make([]Player, 0, len(ops.playerMap))
	for _, p := range ops.playerMap {
		pList = append(pList, p)
	}
	return pList
}

func (ops *onlinePlayerManager) RecommentPlayersExclude(playerIdMap map[int64]struct{}) (pList []Player) {
	ops.rwm.RLock()
	defer ops.rwm.RUnlock()
	tempList := make([]Player, 0, 8)
	tempWeightList := make([]int64, 0, 8)
	for _, p := range ops.playerMap {
		_, ok := playerIdMap[p.GetId()]
		if ok {
			continue
		}
		tempList = append(tempList, p)
		tempWeightList = append(tempWeightList, int64(1))
	}
	maxRecommentNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAddFriendInviteBatchLimit)
	if len(tempList) <= int(maxRecommentNum) {
		pList = tempList
		return
	}
	randomList := mathutils.RandomListFromWeights(tempWeightList, maxRecommentNum)
	for _, randomIndex := range randomList {
		pList = append(pList, tempList[randomIndex])
	}
	return
}

func (ops *onlinePlayerManager) RecommentSpouses(pl Player, excludePlayers map[int64]struct{}) (pList []Player) {
	ops.rwm.RLock()
	defer ops.rwm.RUnlock()

	tempList := make([]Player, 0, 8)
	tempWeightList := make([]int64, 0, 8)
	for _, p := range ops.playerMap {
		_, ok := excludePlayers[p.GetId()]
		if ok {
			continue
		}
		if p.GetSex() == pl.GetSex() {
			continue
		}
		funcOpenTemplate := funcopen.GetFuncOpenService().GetFuncOpenTemplate(funcopentypes.FuncOpenTypeMarry)
		if p.GetLevel() < funcOpenTemplate.OpenedLevel {
			continue
		}
		spouseId := p.GetSpouseId()
		if spouseId != 0 {
			continue
		}

		tempList = append(tempList, p)
		tempWeightList = append(tempWeightList, int64(1))
	}
	if len(tempList) <= RecommentSpouseNum {
		pList = tempList
		return
	}
	randomList := mathutils.RandomListFromWeights(tempWeightList, RecommentSpouseNum)
	for _, randomIndex := range randomList {
		pList = append(pList, tempList[randomIndex])
	}
	return
}

func (ops *onlinePlayerManager) GetPlayerByUserId(id int64) Player {
	ops.rwm.RLock()
	defer ops.rwm.RUnlock()
	p, ok := ops.userMap[id]
	if !ok {
		return nil
	}
	return p
}
func (ops *onlinePlayerManager) GetPlayerById(playerId int64) Player {
	ops.rwm.RLock()
	defer ops.rwm.RUnlock()
	p, ok := ops.playerMap[playerId]
	if !ok {
		return nil
	}
	return p
}
func (ops *onlinePlayerManager) GetPlayerByName(name string) Player {
	ops.rwm.RLock()
	defer ops.rwm.RUnlock()
	p, ok := ops.namePlayerMap[name]
	if !ok {
		return nil
	}
	return p
}

func (ops *onlinePlayerManager) AddPlayer(p Player) bool {
	ops.rwm.Lock()
	defer ops.rwm.Unlock()
	_, ok := ops.userMap[p.GetUserId()]
	if ok {
		return false
	}
	ops.userMap[p.GetUserId()] = p
	return true
}

func (ops *onlinePlayerManager) PlayerEnterServer(p Player) bool {
	ops.rwm.Lock()
	defer ops.rwm.Unlock()
	_, ok := ops.userMap[p.GetUserId()]
	if !ok {
		return false
	}
	_, ok = ops.playerMap[p.GetId()]
	if ok {
		return false
	}
	ops.playerMap[p.GetId()] = p
	ops.namePlayerMap[p.GetName()] = p
	return true
}

func (ops *onlinePlayerManager) PlayerLeaveServer(p Player) {
	ops.rwm.Lock()
	defer ops.rwm.Unlock()
	playerId := p.GetId()
	if playerId != 0 {
		delete(ops.playerMap, playerId)
		delete(ops.namePlayerMap, p.GetName())
	}
}

func (ops *onlinePlayerManager) BroadcastMsg(msg proto.Message) {
	ops.rwm.Lock()
	defer ops.rwm.Unlock()
	for _, p := range ops.playerMap {
		p.SendMsg(msg)
	}
}

func (ops *onlinePlayerManager) BroadcastMsgExclude(excludeIdList []int64, msg proto.Message) {
	ops.rwm.Lock()
	defer ops.rwm.Unlock()
	for playerId, p := range ops.playerMap {
		flag := utils.ContainInt64(excludeIdList, playerId)
		if flag {
			continue
		}
		p.SendMsg(msg)
	}
}

func (ops *onlinePlayerManager) BroadcastMsgRelated(messageCall message.ScheduleMessageCallBack, result interface{}) {
	ops.rwm.Lock()
	defer ops.rwm.Unlock()

	for _, p := range ops.playerMap {
		ctx := scene.WithPlayer(context.Background(), p)
		relatedMsg := message.NewScheduleMessage(messageCall, ctx, result, nil)
		p.Post(relatedMsg)
	}
}

func (ops *onlinePlayerManager) RemovePlayer(p Player) {
	ops.rwm.Lock()
	defer ops.rwm.Unlock()
	delete(ops.userMap, p.GetUserId())
}

func (ops *onlinePlayerManager) Count() int32 {
	ops.rwm.Lock()
	defer ops.rwm.Unlock()
	return int32(len(ops.playerMap))
}

func (ops *onlinePlayerManager) CountNeigua() int32 {
	ops.rwm.Lock()
	defer ops.rwm.Unlock()
	neiGua := int32(0)
	for _, p := range ops.playerMap {
		if p.IsGuaJiPlayer() {
			neiGua += 1
		}
	}
	return neiGua
}

func (ops *onlinePlayerManager) OfflineAllPlayers() {
	ops.rwm.Lock()
	defer ops.rwm.Unlock()
	for _, p := range ops.playerMap {
		p.Close(nil)
	}
}

var (
	opm = &onlinePlayerManager{
		userMap:       make(map[int64]Player),
		playerMap:     make(map[int64]Player),
		namePlayerMap: make(map[string]Player),
	}
)

func GetOnlinePlayerManager() OnlinePlayerManager {
	return opm
}
