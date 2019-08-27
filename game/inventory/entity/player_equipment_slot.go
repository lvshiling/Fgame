package entity

//玩家槽位数据
type PlayerEquipmentSlotEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	SlotId     int32  `gorm:"column:slotId"`
	Star       int32  `gorm:"column:star"`
	Level      int32  `gorm:"column:level"`
	ItemId     int32  `gorm:"column:itemId"`
	GemInfo    string `gorm:"column:gemInfo"`
	BindType   int32  `gorm:"column:bindType"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (e *PlayerEquipmentSlotEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerEquipmentSlotEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerEquipmentSlotEntity) TableName() string {
	return "t_player_equipment_slot"
}
