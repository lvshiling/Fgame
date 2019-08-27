package types

import (
	tulongequiptypes "fgame/fgame/game/tulongequip/types"
)

type TuLongEquipEventType string

const (
	EventTypeTuLongEquipPutOn         TuLongEquipEventType = "TuLongEquipPutOn"         // 屠龙装穿戴
	EventTypeTuLongEquipTakeOff                            = "TuLongEquipTakeOff"       // 屠龙装卸下
	EventTypeTuLongEquipStrengSuccess                      = "TuLongEquipStrengSuccess" // 屠龙装强化成功
	EventTypeTuLongEquipSkillUpgrade                       = "TuLongEquipSkillUpgrade"  // 屠龙技能升级
	EventTypeTuLongEquipUseItem                            = "TuLongEquipUseItem"       //屠龙装消耗物品
)

//
type PlayerTuLongEquipChangedEventData struct {
	suitType tulongequiptypes.TuLongSuitType
	itemId   int32
}

func CreatePlayerTuLongEquipChangedEventData(suitType tulongequiptypes.TuLongSuitType, itemId int32) *PlayerTuLongEquipChangedEventData {
	d := &PlayerTuLongEquipChangedEventData{
		suitType: suitType,
		itemId:   itemId,
	}
	return d
}

func (d *PlayerTuLongEquipChangedEventData) GetSuitType() tulongequiptypes.TuLongSuitType {
	return d.suitType
}

func (d *PlayerTuLongEquipChangedEventData) GetItemId() int32 {
	return d.itemId
}

type PlayerTuLongEquipUseItemEventData struct {
	itemId  int32
	itemNum int32
}

func CreatePlayerTuLongEquipUseItemEventData(itemId, itemNum int32) *PlayerTuLongEquipUseItemEventData {
	d := &PlayerTuLongEquipUseItemEventData{
		itemId:  itemId,
		itemNum: itemNum,
	}
	return d
}

func (data *PlayerTuLongEquipUseItemEventData) GetUseItemMap() map[int32]int32 {
	d := map[int32]int32{
		data.itemId: data.itemNum,
	}
	return d
}
