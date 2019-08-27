package entity

//玩家绝学数据
type PlayerJueXueEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Type       int32 `gorm:"column:typ"`
	Level      int32 `gorm:"column:level"`
	Insight    int32 `gorm:"column:insight"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pjxe *PlayerJueXueEntity) GetId() int64 {
	return pjxe.Id
}

func (pjxe *PlayerJueXueEntity) GetPlayerId() int64 {
	return pjxe.PlayerId
}

func (pjxe *PlayerJueXueEntity) TableName() string {
	return "t_player_juexue"
}
