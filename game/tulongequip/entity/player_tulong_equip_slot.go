package entity

//玩家屠龙装备槽位数据
type PlayerTuLongEquipSlotEntity struct {
	Id           int64  `gorm:"primary_key;column:id"`
	PlayerId     int64  `gorm:"column:playerId"`
	SuitType     int32  `gorm:"column:suitType"`
	SlotId       int32  `gorm:"column:slotId"`
	Level        int32  `gorm:"column:level"`
	ItemId       int32  `gorm:"column:itemId"`
	BindType     int32  `gorm:"column:bindType"`
	PropertyData string `gorm:"column:porpertyData"`
	GemInfo      string `gorm:"column:gemInfo"`
	UpdateTime   int64  `gorm:"column:updateTime"`
	CreateTime   int64  `gorm:"column:createTime"`
	DeleteTime   int64  `gorm:"column:deleteTime"`
}

func (e *PlayerTuLongEquipSlotEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerTuLongEquipSlotEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerTuLongEquipSlotEntity) TableName() string {
	return "t_player_tulong_equip_slot"
}
