package common

import (
	alliancetypes "fgame/fgame/game/alliance/types"
)

//仙盟信息
type PlayerAllianceObject interface {
	GetAllianceId() int64
	GetAllianceName() string
	GetMengZhuId() int64
	GetMemPos() alliancetypes.AlliancePosition
}

type playerAllianceObject struct {
	allianceId   int64
	allianceName string
	mengZhuId    int64
	memPos       alliancetypes.AlliancePosition
}

func (o *playerAllianceObject) GetAllianceId() int64 {
	return o.allianceId
}

func (o *playerAllianceObject) GetAllianceName() string {
	return o.allianceName
}

func (o *playerAllianceObject) GetMengZhuId() int64 {
	return o.mengZhuId
}

func (o *playerAllianceObject) GetMemPos() alliancetypes.AlliancePosition {
	return o.memPos
}

func CreatePlayerAllianceObject(allianceId int64, allianceName string, mengZhuId int64, pos alliancetypes.AlliancePosition) PlayerAllianceObject {
	obj := &playerAllianceObject{}
	obj.allianceId = allianceId
	obj.allianceName = allianceName
	obj.mengZhuId = mengZhuId
	obj.memPos = pos
	return obj
}
