package player

import (
	"sync"

	"github.com/golang/protobuf/proto"
)

//在线玩家管理器
type OnlinePlayerManager interface {
	GetPlayerById(playerId int64) *Player
	BroadcastMsg(msg proto.Message)
	PlayerEnterServer(p *Player) bool
	PlayerLeaveServer(p *Player)
}

type onlinePlayerManager struct {
	rwm sync.RWMutex
	//角色map
	playerMap map[int64]*Player
}

func (ops *onlinePlayerManager) GetPlayerById(playerId int64) *Player {
	ops.rwm.RLock()
	defer ops.rwm.RUnlock()
	p, ok := ops.playerMap[playerId]
	if !ok {
		return nil
	}
	return p
}

func (ops *onlinePlayerManager) BroadcastMsg(msg proto.Message) {
	ops.rwm.Lock()
	defer ops.rwm.Unlock()
	for _, p := range ops.playerMap {
		p.SendMsg(msg)
	}
}

func (ops *onlinePlayerManager) PlayerEnterServer(p *Player) bool {
	ops.rwm.Lock()
	defer ops.rwm.Unlock()

	_, ok := ops.playerMap[p.GetId()]
	if ok {
		return false
	}

	ops.playerMap[p.GetId()] = p

	return true
}

func (ops *onlinePlayerManager) PlayerLeaveServer(p *Player) {
	ops.rwm.Lock()
	defer ops.rwm.Unlock()
	playerId := p.GetId()
	if playerId != 0 {
		delete(ops.playerMap, playerId)
	}
}

var (
	opm = &onlinePlayerManager{
		playerMap: make(map[int64]*Player),
	}
)

func GetOnlinePlayerManager() OnlinePlayerManager {
	return opm
}
