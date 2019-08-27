package entity

//玩家功能开启数据
type PlayerFuncOpenEntity struct {
	Id           int64  `gorm:"primary_key;column:id"`
	PlayerId     int64  `gorm:"column:playerId"`
	FuncOpenList string `gorm:"column:funcOpenList"`
	UpdateTime   int64  `gorm:"column:updateTime"`
	CreateTime   int64  `gorm:"column:createTime"`
	DeleteTime   int64  `gorm:"column:deleteTime"`
}

func (e *PlayerFuncOpenEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerFuncOpenEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerFuncOpenEntity) TableName() string {
	return "t_player_func_open"
}
