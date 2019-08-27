package global

import (
	centertypes "fgame/fgame/center/types"
	"fgame/fgame/common/message"
	fgamedb "fgame/fgame/core/db"
	fgameredis "fgame/fgame/core/redis"
	coretime "fgame/fgame/core/time"
	"sync"
)

var (
	DEBUG    = false
	PRESSURE = false
)

type Game interface {
	GetDB() fgamedb.DBService
	GetRedisService() fgameredis.RedisService
	GetOperationService() OpeartionService
	GetGlobalRunner() *GlobalRunner
	GetMessageHandler() message.Handler
	GetTimeService() coretime.TimeService
	GetServerId() int32
	GetServerIndex() int32
	GetServerType() centertypes.GameServerType
	GetGlobalUpdater() *GlobalUpdater
	GetPlatform() int32
	GetServerIp() string
	GetServerPort() int32
	GetServerTime() int64
	Open() bool
	GMOpen() bool
	CrossDisable() bool
}

var (
	once sync.Once
	g    Game
)

func SetupGame(tg Game) {
	once.Do(func() {
		g = tg
	})
}

func GetGame() Game {
	return g
}
