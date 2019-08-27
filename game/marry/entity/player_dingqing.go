package entity

//玩家定情套装
type PlayerMarryDingQingEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	Suit       string `gorm:"column:suit"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (p *PlayerMarryDingQingEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerMarryDingQingEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerMarryDingQingEntity) TableName() string {
	return "t_player_marry_dingqing"
}
