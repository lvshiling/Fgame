package entity

//玩家命盘数据
type PlayerMingGeBuchangEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Buchang    int32 `gorm:"column:buchang"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pmne *PlayerMingGeBuchangEntity) GetId() int64 {
	return pmne.Id
}

func (pmne *PlayerMingGeBuchangEntity) GetPlayerId() int64 {
	return pmne.PlayerId
}

func (pmoe *PlayerMingGeBuchangEntity) TableName() string {
	return "t_player_mingge_buchang"
}
