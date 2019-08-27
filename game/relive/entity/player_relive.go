package entity

//玩家复活数据
type PlayerReliveEntity struct {
	Id             int64 `gorm:"primary_key;column:id"`
	PlayerId       int64 `gorm:"column:playerId"`
	CulTime        int32 `gorm:"column:culTime"`
	LastReliveTime int64 `gorm:"column:lastReliveTime"`
	UpdateTime     int64 `gorm:"column:updateTime"`
	CreateTime     int64 `gorm:"column:createTime"`
	DeleteTime     int64 `gorm:"column:deleteTime"`
}

func (e *PlayerReliveEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerReliveEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (ptjt *PlayerReliveEntity) TableName() string {
	return "t_player_relive"
}
