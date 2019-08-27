package entity

//玩家屠龙套装数据
type PlayerTuLongEquipEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Power      int64 `gorm:"column:power"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerTuLongEquipEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerTuLongEquipEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerTuLongEquipEntity) TableName() string {
	return "t_player_tulong_equip"
}
