package entity

//玩家泣血枪数据
type PlayerQiXueEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	CurrLevel  int32 `gorm:"column:currLevel"`
	CurrStar   int32 `gorm:"column:currStar"`
	LastTime   int64 `gorm:"column:lastTime"`
	ShaLuNum   int64 `gorm:"column:shaLuNum"`
	TimesNum   int32 `gorm:"column:timesNum"`
	Power      int64 `gorm:"column:power"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerQiXueEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerQiXueEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerQiXueEntity) TableName() string {
	return "t_player_qixue"
}
