package player

import "fgame/fgame/client/player/player"

import propertycommon "fgame/fgame/game/property/common"
import propertytypes "fgame/fgame/game/property/types"

//玩家属性管理器
type PlayerPropertyDataManager struct {
	p             *player.Player
	finalProperty *propertycommon.BattlePropertySegment
}

func (m *PlayerPropertyDataManager) UpdateSystemBattleProperty(properties map[int32]int64) {
	for k, v := range properties {
		pt := propertytypes.BattlePropertyType(k)
		if !pt.IsValid() {
			continue
		}
		m.finalProperty.Set(pt, v)
	}
}

func (m *PlayerPropertyDataManager) GetPlayer() *player.Player {
	return m.p
}

func (m *PlayerPropertyDataManager) GetBattleProperty(typ propertytypes.BattlePropertyType) int64 {
	return m.finalProperty.Get(typ)
}

func CreatePlayerPropertyDataManager(pl *player.Player) player.PlayerDataManager {
	m := &PlayerPropertyDataManager{
		p: pl,
	}
	m.finalProperty = propertycommon.NewBattlePropertySegment()
	return m
}

func init() {
	player.RegisterPlayerDataManager(player.PlayerDataKeyProperty, player.PlayerDataManagerFactoryFunc(CreatePlayerPropertyDataManager))
}
