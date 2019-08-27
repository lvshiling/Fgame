package common

import (
	pktypes "fgame/fgame/game/pk/types"
)

type PlayerPkObject interface {
	GetPkState() pktypes.PkState
	GetPkCamp() pktypes.PkCamp
	GetPkValue() int32
	GetPkRedState() pktypes.PkRedState
	GetKillNum() int32
	GetOnlineTime() int64
	GetLoginTime() int64
	GetLastKillTime() int64
}

type playerPkObject struct {
	//玩家pk状态
	pkState pktypes.PkState
	//阵营id
	pkCamp pktypes.PkCamp
	//pk值
	pkValue int32

	killNum      int32
	onlineTime   int64
	loginTime    int64
	lastKillTime int64
}

func (o *playerPkObject) GetPkState() pktypes.PkState {
	return o.pkState
}

func (o *playerPkObject) GetPkCamp() pktypes.PkCamp {
	return o.pkCamp
}

func (o *playerPkObject) GetPkValue() int32 {
	return o.pkValue
}

func (o *playerPkObject) GetPkRedState() pktypes.PkRedState {
	return pktypes.PkRedStateFromValue(o.pkValue)
}

func (o *playerPkObject) GetKillNum() int32 {
	return o.killNum
}

func (o *playerPkObject) GetOnlineTime() int64 {
	return o.onlineTime
}

func (o *playerPkObject) GetLastKillTime() int64 {
	return o.lastKillTime
}

func (o *playerPkObject) GetLoginTime() int64 {
	return o.loginTime
}

func NewPlayerPkObject(pkState pktypes.PkState, pkCamp pktypes.PkCamp, pkValue int32) PlayerPkObject {
	o := &playerPkObject{
		pkState: pkState,
		pkCamp:  pkCamp,
		pkValue: pkValue,
	}
	return o
}
