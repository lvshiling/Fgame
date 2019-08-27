package player

import "fgame/fgame/client/player/player"
import coretypes "fgame/fgame/core/types"

//玩家场景管理器
type PlayerSceneDataManager struct {
	p     *player.Player
	mapId int32
	pos   coretypes.Position
}

func (m *PlayerSceneDataManager) GetPlayer() *player.Player {
	return m.p
}

func (m *PlayerSceneDataManager) GetMapId() int32 {
	return m.mapId
}

func (m *PlayerSceneDataManager) GetPos() coretypes.Position {
	return m.pos
}

func (m *PlayerSceneDataManager) EnterScene(mapId int32, pos coretypes.Position) {
	m.mapId = mapId
	m.pos = pos
}

func (m *PlayerSceneDataManager) Move(pos coretypes.Position) {
	m.pos = pos
}

func CreatePlayerSceneDataManager(pl *player.Player) player.PlayerDataManager {
	m := &PlayerSceneDataManager{
		p: pl,
	}

	return m
}

func init() {
	player.RegisterPlayerDataManager(player.PlayerDataKeyScene, player.PlayerDataManagerFactoryFunc(CreatePlayerSceneDataManager))
}
