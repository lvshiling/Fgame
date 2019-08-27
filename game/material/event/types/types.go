package types

import (
	materialtypes "fgame/fgame/game/material/types"
)

type MaterialEventType string

const (
	//材料副本挑战事件
	EventTypeMaterialChallenge MaterialEventType = "MaterialChallenge"
	//材料副本挑战成功
	EventTypeMaterialFinish MaterialEventType = "MaterialFinish"
	//材料副本刷新怪
	EventTypeMaterialRefreshGroup MaterialEventType = "MaterialRefreshGroup"
)

type RefreshGroupEventData struct {
	group int32
	typ   materialtypes.MaterialType
}

func CreateRefreshGroupEventData(group int32, typ materialtypes.MaterialType) *RefreshGroupEventData {
	data := &RefreshGroupEventData{
		group: group,
		typ:   typ,
	}
	return data
}

func (d *RefreshGroupEventData) GetGroup() int32 {
	return d.group
}

func (d *RefreshGroupEventData) GetMaterialType() materialtypes.MaterialType {
	return d.typ
}

type MaterialChallengeEventData struct {
	typ materialtypes.MaterialType
	num int32
}

func (d *MaterialChallengeEventData) GetType() materialtypes.MaterialType {
	return d.typ
}

func (d *MaterialChallengeEventData) GetNum() int32 {
	return d.num
}

func CreateMaterialChallengeEventData(typ materialtypes.MaterialType, num int32) *MaterialChallengeEventData {
	return &MaterialChallengeEventData{
		typ: typ,
		num: num,
	}
}

type MaterialFinishEventData struct {
	typ materialtypes.MaterialType
	num int32
}

func (d *MaterialFinishEventData) GetType() materialtypes.MaterialType {
	return d.typ
}

func (d *MaterialFinishEventData) GetNum() int32 {
	return d.num
}

func CreateMaterialFinishEventData(typ materialtypes.MaterialType, num int32) *MaterialFinishEventData {
	return &MaterialFinishEventData{
		typ: typ,
		num: num,
	}
}
