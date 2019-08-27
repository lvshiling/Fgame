package battle

import (
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	jieyicommon "fgame/fgame/game/jieyi/common"
	jieyitypes "fgame/fgame/game/jieyi/types"
	"fgame/fgame/game/scene/scene"
	"fmt"
)

type PlayerJieYiManager struct {
	p         scene.Player
	jieYiName string
	jieYiRank int32
	jieYiId   int64
}

func (m *PlayerJieYiManager) GetSceneJieYiName() string {
	if m.jieYiId == 0 {
		return ""
	}

	rankName := jieyitypes.JieYiRank(m.jieYiRank).GetRankString()
	return fmt.Sprintf("%s·%s", m.jieYiName, rankName)
}

func (m *PlayerJieYiManager) GetJieYiName() string {
	return m.jieYiName
}

func (m *PlayerJieYiManager) GetJieYiId() int64 {
	return m.jieYiId
}

func (m *PlayerJieYiManager) GetJieYiRank() int32 {
	return m.jieYiRank
}

func (m *PlayerJieYiManager) SyncJieYi(jieYiId int64, jieYiName string, jieYiRank int32) {
	m.jieYiId = jieYiId
	m.jieYiName = jieYiName
	m.jieYiRank = jieYiRank
	// 发事件
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerJieYiChanged, m.p, nil)
}

func CreatePlayerJieYiManagerWithObject(p scene.Player, playerJiYiObj jieyicommon.PlayerJieYiObject) *PlayerJieYiManager {
	m := &PlayerJieYiManager{
		p:         p,
		jieYiName: playerJiYiObj.GetJieYiName(),
		jieYiRank: playerJiYiObj.GetJieYiRank(),
	}
	return m
}

func CreatePlayerJieYiManager(p scene.Player) *PlayerJieYiManager {
	m := &PlayerJieYiManager{
		p:         p,
		jieYiId:   0,
		jieYiName: "",
		jieYiRank: 0,
	}
	return m
}
