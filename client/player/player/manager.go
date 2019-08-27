package player

import (
	"fmt"
)

//玩家管理器
type PlayerDataManager interface {
	GetPlayer() *Player
}

type PlayerDataManagerFactory interface {
	Create(p *Player) PlayerDataManager
}
type PlayerDataManagerFactoryFunc func(p *Player) PlayerDataManager

func (pdmff PlayerDataManagerFactoryFunc) Create(p *Player) PlayerDataManager {
	return pdmff(p)
}

type PlayerDataKey int32

const (
	PlayerDataKeyBasic PlayerDataKey = iota
	PlayerDataKeyScene
	PlayerDataKeyInventory
	PlayerDataKeyProperty
)

var (
	playerDataFactoryMap = make(map[PlayerDataKey]PlayerDataManagerFactory)
)

func RegisterPlayerDataManager(pdk PlayerDataKey, pdm PlayerDataManagerFactory) {
	_, exist := playerDataFactoryMap[pdk]
	if exist {
		panic(fmt.Errorf("repeat register player data manager factory [%d]", pdk))
	}
	playerDataFactoryMap[pdk] = pdm
}
