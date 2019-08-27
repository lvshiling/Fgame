package types

type TuLongEquipSlotInfo struct {
	SlotId       int32                    `json:"slotId"`
	Level        int32                    `json:"level"`
	ItemId       int32                    `json:"itemId"`
	PropertyData *TuLongEquipPropertyData `json:"propertyData"`
	Gems         map[int32]int32          `json:"gems"`
}
