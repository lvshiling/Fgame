package entity

//玩家绝学使用数据
type PlayerJueXueUseEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Type       int32 `gorm:"column:typ"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pjxue *PlayerJueXueUseEntity) GetId() int64 {
	return pjxue.Id
}

func (pjxue *PlayerJueXueUseEntity) GetPlayerId() int64 {
	return pjxue.PlayerId
}

func (pjxue *PlayerJueXueUseEntity) TableName() string {
	return "t_player_juexue_use"
}
