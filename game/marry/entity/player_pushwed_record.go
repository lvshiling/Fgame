package entity

//玩家查看过的喜帖
type PlayerPushWedRecordEntity struct {
	Id          int64 `gorm:"primary_key;column:id"`
	PlayerId    int64 `gorm:"column:playerId"`
	WedId       int64 `gorm:"column:wedId"`
	HunCheTime  int64 `gorm:"column:hunCheTime"`
	BanquetTime int64 `gorm:"column:banquetTime"`
	UpdateTime  int64 `gorm:"column:updateTime"`
	CreateTime  int64 `gorm:"column:createTime"`
	DeleteTime  int64 `gorm:"column:deleteTime"`
}

func (p *PlayerPushWedRecordEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerPushWedRecordEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerPushWedRecordEntity) TableName() string {
	return "t_player_pushwed_record"
}
