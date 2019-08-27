package player

import (
	playertypes "fgame/fgame/game/player/types"
)

type PlayerBasicManager struct {
	p        *Player
	playerId int64
	name     string
	role     playertypes.RoleType
	sex      playertypes.SexType
}

func (m *PlayerBasicManager) GetPlayer() *Player {
	return m.p
}

func (m *PlayerBasicManager) Load(playerId int64, name string, role playertypes.RoleType, sex playertypes.SexType) {
	m.playerId = playerId
	m.name = name
	m.role = role
	m.sex = sex
}

func (m *PlayerBasicManager) GetPlayerId() int64 {
	return m.playerId
}

func (m *PlayerBasicManager) GetName() string {
	return m.name
}

func (m *PlayerBasicManager) GetRole() playertypes.RoleType {
	return m.role
}

func (m *PlayerBasicManager) GetSex() playertypes.SexType {
	return m.sex
}

func CreatePlayerBasicManager(pl *Player) PlayerDataManager {
	m := &PlayerBasicManager{
		p: pl,
	}

	return m
}

func init() {
	RegisterPlayerDataManager(PlayerDataKeyBasic, PlayerDataManagerFactoryFunc(CreatePlayerBasicManager))
}
