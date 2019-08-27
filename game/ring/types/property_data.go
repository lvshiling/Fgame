package types

import (
	inventorytypes "fgame/fgame/game/inventory/types"
)

type RingPropertyData struct {
	*inventorytypes.ItemPropertyDataBase
	Advance       int32 `json:"advance"`       // 进阶等级
	AdvancePro    int32 `json:"advancedPro"`   // 进阶进度值
	AdvanceNum    int32 `json:"advanceNum"`    // 进阶次数
	StrengthLevel int32 `json:"strengthLevel"` // 强化等级
	StrengthNum   int32 `json:"strengthNum"`   // 强化次数
	JingLingLevel int32 `json:"jingLingLevel"` // 净灵等级
	JingLingNum   int32 `json:"jingLingNum"`   // 净灵次数
}

func NewRingPropertyData() *RingPropertyData {
	d := &RingPropertyData{}
	return d
}

func (gd *RingPropertyData) InitBase() {
	if gd.ItemPropertyDataBase == nil {
		gd.ItemPropertyDataBase = inventorytypes.CreateDefaultItemPropertyDataBase()
	}
}

func (gd *RingPropertyData) Copy() inventorytypes.ItemPropertyData {
	data := &RingPropertyData{}
	data.ItemPropertyDataBase = gd.ItemPropertyDataBase.CopyBase()
	data.Advance = gd.Advance
	data.AdvancePro = gd.AdvancePro
	data.AdvanceNum = gd.AdvanceNum
	data.StrengthLevel = gd.StrengthLevel
	data.StrengthNum = gd.StrengthNum
	data.JingLingLevel = gd.JingLingLevel
	data.JingLingNum = gd.JingLingNum
	return data
}

func CreateRingPropertyData(base *inventorytypes.ItemPropertyDataBase) inventorytypes.ItemPropertyData {
	d := &RingPropertyData{}
	d.ItemPropertyDataBase = base
	return d
}
