package battle

import (
	alliancecommon "fgame/fgame/game/alliance/common"
	alliancetypes "fgame/fgame/game/alliance/types"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

type PlayerAllianceManager struct {
	p             scene.Player
	allianceId    int64
	allianceName  string
	allianceLevel int32
	mengZhuId     int64
	memPos        alliancetypes.AlliancePosition
}

func (m *PlayerAllianceManager) GetAllianceId() int64 {
	return m.allianceId
}

func (m *PlayerAllianceManager) GetAllianceName() string {
	return m.allianceName
}

func (m *PlayerAllianceManager) GetMengZhuId() int64 {
	return m.mengZhuId
}

func (m *PlayerAllianceManager) GetMemPos() alliancetypes.AlliancePosition {
	return m.memPos
}

func (m *PlayerAllianceManager) SyncAlliance(allianceId int64, allianceName string, mengZhuId int64, pos alliancetypes.AlliancePosition) {
	m.allianceId = allianceId
	m.allianceName = allianceName
	m.mengZhuId = mengZhuId
	m.memPos = pos
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerAllianceChanged, m.p, nil)
}

func CreatePlayerAllianceManagerWithObject(p scene.Player, playerAllianceObj alliancecommon.PlayerAllianceObject) *PlayerAllianceManager {
	m := &PlayerAllianceManager{
		p:            p,
		allianceId:   playerAllianceObj.GetAllianceId(),
		allianceName: playerAllianceObj.GetAllianceName(),
		mengZhuId:    playerAllianceObj.GetMengZhuId(),
		memPos:       playerAllianceObj.GetMemPos(),
	}
	return m
}

func CreatePlayerAllianceManager(p scene.Player) *PlayerAllianceManager {
	m := &PlayerAllianceManager{
		p:            p,
		allianceId:   0,
		allianceName: "",
		mengZhuId:    0,
	}
	return m
}
