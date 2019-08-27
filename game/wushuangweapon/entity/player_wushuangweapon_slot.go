package entity

type PlayerWushuangWeaponSlotEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	SlotId     int32 `gorm:"column:slotId"` //身体部位
	ItemId     int32 `gorm:"column:itemId"`
	Level      int32 `gorm:"column:level"`
	Experience int64 `gorm:"column:experience"` //升级积攒的经验值
	Bind       int32 `gorm:"column:bind"`
	IsActive   int32 `gorm:"column:isActive"` //是否激活
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerWushuangWeaponSlotEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerWushuangWeaponSlotEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerWushuangWeaponSlotEntity) TableName() string {
	return "t_player_wushuangweapon_slot"
}
