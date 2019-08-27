package types

import (
	fushitypes "fgame/fgame/game/fushi/types"
)

// 八卦符石事件类型
type FushiEventType string

const (
	FushiEventTypeLevelChanged FushiEventType = "FuShiLevelChanged" // 八卦符石等级改变
)

// 八卦符石数据
type FuShiLevelChangedData struct {
	typ   fushitypes.FuShiType
	level int32
}

func (data *FuShiLevelChangedData) GetType() fushitypes.FuShiType {
	return data.typ
}

func (data *FuShiLevelChangedData) GetFushiLevel() int32 {
	return data.level
}

func CreateFuShiLevelChangedData(typ fushitypes.FuShiType, level int32) *FuShiLevelChangedData {
	data := &FuShiLevelChangedData{
		typ:   typ,
		level: level,
	}
	return data
}
