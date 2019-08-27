package types

import (
	inventorytypes "fgame/fgame/game/inventory/types"
)

type TuLongEquipPropertyData struct {
	*inventorytypes.ItemPropertyDataBase
}

func NewTuLongEquipPropertyData() *TuLongEquipPropertyData {
	d := &TuLongEquipPropertyData{}
	return d
}

func (gd *TuLongEquipPropertyData) InitBase() {
	if gd.ItemPropertyDataBase == nil {
		gd.ItemPropertyDataBase = inventorytypes.CreateDefaultItemPropertyDataBase()
	}
}

func CreateTuLongEquipPropertyData(base *inventorytypes.ItemPropertyDataBase) inventorytypes.ItemPropertyData {
	d := &TuLongEquipPropertyData{}
	d.ItemPropertyDataBase = base
	return d
}
