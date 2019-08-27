package entity

//玩家宝宝玩具槽位数据
type PlayerBabyToySlotEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	SuitType   int32 `gorm:"column:suitType"`
	ItemId     int32 `gorm:"column:itemId"`
	SlotId     int32 `gorm:"column:slotId"`
	Level      int32 `gorm:"column:level"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerBabyToySlotEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerBabyToySlotEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerBabyToySlotEntity) TableName() string {
	return "t_player_baby_toy_slot"
}
