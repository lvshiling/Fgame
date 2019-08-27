package monitor

import (
	"context"
	"fgame/fgame/gm/gamegm/session"
	"fmt"
	"sync"
)

type Player interface {
	Id() int64
	Ip() string
	Auth(id int64, ip string) bool
	Session() session.Session
	Send(msg []byte)
	Ping() bool
	End()
	Close()
	IsInit() bool
	IsAuth() bool
}

const (
	playerKey = contextKey("qipai.qipai.player")
)

func PlayerInContext(ctx context.Context) Player {
	p, ok := ctx.Value(playerKey).(Player)
	if !ok {
		return nil
	}
	return p
}

func WithPlayer(ctx context.Context, pl Player) context.Context {
	return context.WithValue(ctx, playerKey, pl)
}

//玩家管理器
type PlayerManager struct {
	rwm     sync.RWMutex
	players map[int64]Player
}

func (pm *PlayerManager) AddPlayer(p Player) error {
	pm.rwm.Lock()
	defer pm.rwm.Unlock()
	_, exist := pm.players[p.Id()]
	if exist {
		return fmt.Errorf("玩家已经存在")
	}
	pm.players[p.Id()] = p
	return nil
}

func (pm *PlayerManager) RemovePlayer(p Player) {
	pm.rwm.Lock()
	defer pm.rwm.Unlock()
	delete(pm.players, p.Id())
}

func (pm *PlayerManager) Players() map[int64]Player {
	pm.rwm.RLock()
	defer pm.rwm.RUnlock()
	return pm.players
}

func (pm *PlayerManager) GetPlayerById(id int64) Player {
	pm.rwm.RLock()
	defer pm.rwm.RUnlock()
	pl, exist := pm.players[id]
	if !exist {
		return nil
	}
	return pl
}

func NewPlayerManager() *PlayerManager {
	pm := &PlayerManager{}
	pm.players = make(map[int64]Player)
	return pm
}

type OfflinePlayer struct {
	Id    int64
	Ip    string
	Image string
	Name  string
	Sex   int
	Gold  int64
	Robot int32
}
