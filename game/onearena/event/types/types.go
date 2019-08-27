package types

import (
	onearenatypes "fgame/fgame/game/onearena/types"
)

type OneArenaEventType string

const (
	//玩家抢夺结束
	EventTypeOneArenaRobEnd OneArenaEventType = "oneArenaRobEnd"
	//玩家灵池争夺成功
	EventTypePlayerOneArenaSucess OneArenaEventType = "playerOneArenaSucess"
	//玩家灵池争夺失败
	EventTypePlayerOneArenaFail OneArenaEventType = "playerOneArenaFail"
	//灵池产出鲲
	EventTypeOneArenaOutputKun OneArenaEventType = "oneArenaOutputKun"
	//灵池争夺合服
	EventTypeOneArenaMergeServer OneArenaEventType = "oneArenaMergeServer"
	//灵池占领时间
	EventTypeOneArenaOccupyTime OneArenaEventType = "oneArenaOccupyTime"
)

type OneArenaData struct {
	playerId int64
	level    onearenatypes.OneArenaLevelType
	pos      int32
}

func (o *OneArenaData) GetPlayerId() int64 {
	return o.playerId
}

func (o *OneArenaData) GetLevel() onearenatypes.OneArenaLevelType {
	return o.level
}

func (o *OneArenaData) GetPos() int32 {
	return o.pos
}

func CreateOneArenaData(playerId int64, level onearenatypes.OneArenaLevelType, pos int32) *OneArenaData {
	d := &OneArenaData{
		playerId: playerId,
		level:    level,
		pos:      pos,
	}
	return d
}

type OneArenaRobFailEventData struct {
	playerId         int64
	peerOneArenaData *OneArenaData
}

func (o *OneArenaRobFailEventData) GetPlayerId() int64 {
	return o.playerId
}

func (o *OneArenaRobFailEventData) GetPeerOneArenaData() *OneArenaData {
	return o.peerOneArenaData
}

func CreateOneArenaRobFailEventData(playerId int64, peerOneArenaData *OneArenaData) *OneArenaRobFailEventData {
	d := &OneArenaRobFailEventData{
		playerId:         playerId,
		peerOneArenaData: peerOneArenaData,
	}
	return d
}

type OneArenaRobSucessEventData struct {
	oneArenaData     *OneArenaData
	peerOneArenaData *OneArenaData
}

func (o *OneArenaRobSucessEventData) GetOneArenaData() *OneArenaData {
	return o.oneArenaData
}

func (o *OneArenaRobSucessEventData) GetPeerOneArenaData() *OneArenaData {
	return o.peerOneArenaData
}

func CreateOneArenaRobSucessEventData(oneArenaData *OneArenaData, peerOneArenaData *OneArenaData) *OneArenaRobSucessEventData {
	d := &OneArenaRobSucessEventData{
		oneArenaData:     oneArenaData,
		peerOneArenaData: peerOneArenaData,
	}
	return d
}

type OneArenaRobEndEventData struct {
	sucess    bool
	level     onearenatypes.OneArenaLevelType
	pos       int32
	ownerName string
}

func (o *OneArenaRobEndEventData) GetSucess() bool {
	return o.sucess
}

func (o *OneArenaRobEndEventData) GetLevel() onearenatypes.OneArenaLevelType {
	return o.level
}

func (o *OneArenaRobEndEventData) GetPos() int32 {
	return o.pos
}

func (o *OneArenaRobEndEventData) GetOwnerName() string {
	return o.ownerName
}

func CreateOneArenaRobEndEventData(sucess bool, level onearenatypes.OneArenaLevelType, pos int32, ownerName string) *OneArenaRobEndEventData {
	d := &OneArenaRobEndEventData{
		sucess:    sucess,
		level:     level,
		pos:       pos,
		ownerName: ownerName,
	}
	return d
}
