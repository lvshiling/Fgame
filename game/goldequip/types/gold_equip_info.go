package types

type GoldEquipSlotInfo struct {
	SlotId            int32                             `json:"slotId"`
	Level             int32                             `json:"level"`
	NewStLevel        int32                             `json:"newStLevel"`
	ItemId            int32                             `json:"itemId"`
	PropertyData      *GoldEquipPropertyData            `json:"propertyData"`
	Gems              map[int32]int32                   `json:"gems"`
	GemUnlockInfo     map[int32]int32                   `json:"gemUnlockInfo"`
	CastingSpiritInfo map[SpiritType]*CastingSpiritInfo `json:"castingSpiritInfo"`
	ForgeSoulInfo     map[ForgeSoulType]*ForgeSoulInfo  `json:"forgeSoulInfo"`
}

//铸灵
type CastingSpiritInfo struct {
	Level int32 `json:"level"` //等级
	Times int32 `json:"times"` //次数
	Bless int32 `json:"bless"` //祝福值
}

//锻魂
type ForgeSoulInfo struct {
	Level int32 `json:"level"` //等级
	Times int32 `json:"times"` //次数
}
