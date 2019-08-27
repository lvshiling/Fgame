package entity

//玩家衣橱套装培养数据
type PlayerWardrobePeiYangEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Type       int32 `gorm:"column:type"`
	PeiYangNum int32 `gorm:"column:peiYangNum"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerWardrobePeiYangEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerWardrobePeiYangEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerWardrobePeiYangEntity) TableName() string {
	return "t_player_wardrobe_peiyang"
}
