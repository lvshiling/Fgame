package player

import "fgame/fgame/game/player/types"

//玩家数据管理器
type PlayerDataManager interface {
	//玩家
	Player() Player
	//加载
	Load() error
	//加载后
	AfterLoad() error
	//心跳
	Heartbeat()
}

type PlayerDataManagerFactory interface {
	CreatePlayerDataManager(p Player) PlayerDataManager
}

type PlayerDataManagerFactoryFunc func(p Player) PlayerDataManager

func (pdmff PlayerDataManagerFactoryFunc) CreatePlayerDataManager(p Player) PlayerDataManager {
	return pdmff(p)
}

var (
	playerDataManagerMap map[types.PlayerDataManagerType]PlayerDataManagerFactory = make(map[types.PlayerDataManagerType]PlayerDataManagerFactory)
)

func RegisterPlayerDataManager(typ types.PlayerDataManagerType, pdmf PlayerDataManagerFactory) {
	_, exist := playerDataManagerMap[typ]
	if exist {
		panic("repeate register player data manager")
	}
	playerDataManagerMap[typ] = pdmf
	return
}

func GetPlayerDataManagerMap() map[types.PlayerDataManagerType]PlayerDataManagerFactory {
	return playerDataManagerMap
}
