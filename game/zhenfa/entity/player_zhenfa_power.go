package entity

//玩家阵法战力数据
type PlayerZhenFaPowerEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Power      int64 `gorm:"column:power"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pmne *PlayerZhenFaPowerEntity) GetId() int64 {
	return pmne.Id
}

func (pmne *PlayerZhenFaPowerEntity) GetPlayerId() int64 {
	return pmne.PlayerId
}

func (pmoe *PlayerZhenFaPowerEntity) TableName() string {
	return "t_player_zhenfa_power"
}
