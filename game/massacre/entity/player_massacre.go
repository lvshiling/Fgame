package entity

//玩家戮仙刃数据
type PlayerMassacreEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	AdvancedId int   `gorm:"column:advancedId"`
	CurrLevel  int32 `gorm:"column:currLevel"`
	CurrStar   int32 `gorm:"column:currStar"`
	LastTime   int64 `gorm:"column:lastTime"`
	ShaQiNum   int64 `gorm:"column:shaQiNum"`
	TimesNum   int32 `gorm:"column:timesNum"`
	Power      int64 `gorm:"column:power"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerMassacreEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerMassacreEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerMassacreEntity) TableName() string {
	return "t_player_massacre"
}
