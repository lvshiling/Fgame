package types

import (
	scenetypes "fgame/fgame/game/scene/types"
)

type BuffEventType string

const (
	//buff触发事件
	EventTypeBuffTouch BuffEventType = "BuffTouch"
	//buff添加
	EventTypeBuffAdd = "BuffAdd"
	//buff更新
	EventTypeBuffUpdate = "BuffUpdate"
	//buff移除
	EventTypeBuffRemove = "BuffRemove"
	//buff特殊效果
	EventTypeBuffEffectCost = "BuffEffectCost"
)

const (
	EventTypeBuffAddExp = "BuffAddExp"
)

type BuffEffectCostEventData struct {
	effectType scenetypes.BuffEffectType
	costNum    int64
}

func CraetBuffEffectCostEventData(effectType scenetypes.BuffEffectType, costNum int64) *BuffEffectCostEventData {
	d := &BuffEffectCostEventData{
		effectType: effectType,
		costNum:    costNum,
	}

	return d
}

func (d *BuffEffectCostEventData) GetEffectType() scenetypes.BuffEffectType {
	return d.effectType
}

func (d *BuffEffectCostEventData) GetCostNum() int64 {
	return d.costNum
}
