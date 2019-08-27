package entity

//玩家元神金装
type PlayerGoldEquipEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Power      int64 `gorm:"column:power"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerGoldEquipEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerGoldEquipEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerGoldEquipEntity) TableName() string {
	return "t_player_goldequip"
}
