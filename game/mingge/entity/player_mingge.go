package entity

//玩家命格数据
type PlayerMingGeEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Power      int64 `gorm:"column:power"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pmne *PlayerMingGeEntity) GetId() int64 {
	return pmne.Id
}

func (pmne *PlayerMingGeEntity) GetPlayerId() int64 {
	return pmne.PlayerId
}

func (pmoe *PlayerMingGeEntity) TableName() string {
	return "t_player_mingge"
}
