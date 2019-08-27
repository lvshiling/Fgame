package entity

//玩家膜拜数据
type PlayerEmperorWorshipEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Num        int32 `gorm:"column:num"`
	LastTime   int64 `gorm:"column:lastTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pewe *PlayerEmperorWorshipEntity) GetId() int64 {
	return pewe.Id
}

func (pewe *PlayerEmperorWorshipEntity) GetPlayerId() int64 {
	return pewe.PlayerId
}

func (pewe *PlayerEmperorWorshipEntity) TableName() string {
	return "t_player_emperor_worship"
}
