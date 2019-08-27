package entity

//玩家补偿数据
type PlayerCompensateEntity struct {
	Id           int64 `gorm:"primary_key;column:id"`
	PlayerId     int64 `gorm:"column:playerId"`
	CompensateId int64 `gorm:"column:compensateId"`
	State        int32 `gorm:"column:state"`
	UpdateTime   int64 `gorm:"column:updateTime"`
	CreateTime   int64 `gorm:"column:createTime"`
	DeleteTime   int64 `gorm:"column:deleteTime"`
}

func (pwe *PlayerCompensateEntity) GetId() int64 {
	return pwe.Id
}

func (pwe *PlayerCompensateEntity) GetPlayerId() int64 {
	return pwe.PlayerId
}

func (pwe *PlayerCompensateEntity) TableName() string {
	return "t_player_compensate"
}
