package entity

//玩家查看过婚帖
type PlayerViewWedCardEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	CardId     int64 `gorm:"column:cardId"`
	ViewTime   int64 `gorm:"column:viewTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (p *PlayerViewWedCardEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerViewWedCardEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerViewWedCardEntity) TableName() string {
	return "t_player_view_wedcard"
}
