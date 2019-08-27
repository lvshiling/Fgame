package entity

//玩家元神金装槽位数据
type PlayerGoldEquipSlotEntity struct {
	Id                int64  `gorm:"primary_key;column:id"`
	PlayerId          int64  `gorm:"column:playerId"`
	SlotId            int32  `gorm:"column:slotId"`
	Level             int32  `gorm:"column:level"`      
	NewStLevel        int32  `gorm:"column:newStLevel"` 
	ItemId            int32  `gorm:"column:itemId"`
	BindType          int32  `gorm:"column:bindType"`
	PropertyData      string `gorm:"column:porpertyData"`
	GemInfo           string `gorm:"column:gemInfo"`
	GemUnlockInfo     string `gorm:"column:gemUnlockInfo"`
	CastingSpiritInfo string `gorm:"column:castingSpiritInfo"`
	ForgeSoulInfo     string `gorm:"column:forgeSoulInfo"`
	UpdateTime        int64  `gorm:"column:updateTime"`
	CreateTime        int64  `gorm:"column:createTime"`
	DeleteTime        int64  `gorm:"column:deleteTime"`
}

func (e *PlayerGoldEquipSlotEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerGoldEquipSlotEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerGoldEquipSlotEntity) TableName() string {
	return "t_player_gold_equip_slot"
}
