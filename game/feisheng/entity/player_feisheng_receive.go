package entity

//玩家飞升数据
type PlayerFeiShengReceiveEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Num        int32 `gorm:"column:num"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerFeiShengReceiveEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerFeiShengReceiveEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerFeiShengReceiveEntity) TableName() string {
	return "t_player_feisheng_receive"
}
