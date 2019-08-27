package types

import (
	inventorytypes "fgame/fgame/game/inventory/types"
)

type GoldEquipPropertyData struct {
	*inventorytypes.ItemPropertyDataBase
	OpenLightLevel  int32   `json:"openLightLevel"`  //开光等级
	OpenTimes       int32   `json:"openMin"`         //开光次数
	UpstarLevel     int32   `json:"upstarLevel"`     //强化升星等级
	AttrList        []int32 `json:"attrList"`        //附加属性
	IsHadCountAttr  bool    `json:"isHadCountAttr"`  //是否生成过随机属性
	IsFix           bool    `json:"isFix"`           //是否未修正强化升星
	GodCastingTimes int32   `json:"godCastingTimes"` //神铸本次升阶已使用次数
}

func NewGoldEquipPropertyData() *GoldEquipPropertyData {
	d := &GoldEquipPropertyData{}
	d.IsFix = true
	return d
}

func (gd *GoldEquipPropertyData) InitBase() {
	if gd.ItemPropertyDataBase == nil {
		gd.ItemPropertyDataBase = inventorytypes.CreateDefaultItemPropertyDataBase()
	}
}

func (gd *GoldEquipPropertyData) FixUpstarLevel(maxLevel int32) {
	if gd.IsFix {
		return
	}
	fixLevel := gd.UpstarLevel * 3
	if fixLevel > maxLevel {
		fixLevel = maxLevel
	}
	gd.IsFix = true
	gd.UpstarLevel = fixLevel
}

func (gd *GoldEquipPropertyData) Copy() inventorytypes.ItemPropertyData {
	data := &GoldEquipPropertyData{}
	data.ItemPropertyDataBase = gd.ItemPropertyDataBase.CopyBase()
	data.OpenLightLevel = gd.OpenLightLevel
	data.OpenTimes = gd.OpenTimes
	data.UpstarLevel = gd.UpstarLevel
	data.GodCastingTimes = gd.GodCastingTimes
	data.AttrList = make([]int32, len(gd.AttrList))
	copy(data.AttrList, gd.AttrList)
	data.IsHadCountAttr = gd.IsHadCountAttr
	data.IsFix = gd.IsFix
	return data
}

func CreateGoldEquipPropertyData(base *inventorytypes.ItemPropertyDataBase) inventorytypes.ItemPropertyData {
	d := &GoldEquipPropertyData{}
	d.IsFix = true
	d.ItemPropertyDataBase = base
	return d
}
