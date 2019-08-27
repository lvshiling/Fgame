package entity

//玩家天书数据
type PlayerTianShuEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Typ        int32 `gorm:"column:typ"`
	Level      int32 `gorm:"column:level"`
	IsReceive  int32 `gorm:"column:isReceive"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerTianShuEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerTianShuEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerTianShuEntity) TableName() string {
	return "t_player_tianshu"
}
