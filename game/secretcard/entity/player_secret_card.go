package entity

//玩家天机牌数据
type PlayerSecretCardEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	TotalNum   int64  `gorm:"column:totalNum"`
	Num        int32  `gorm:"column:num"`
	TotalStar  int32  `gorm:"column:totalStar"`
	OpenBoxs   string `gorm:"column:openBoxs"`
	CardId     int32  `gorm:"column:cardId"`
	Star       int32  `gorm:"column:star"`
	Cards      string `gorm:"column:cards"`
	UsedCards  string `gorm:"column:usedCards"`
	LastTime   int64  `gorm:"column:lastTime"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (p *PlayerSecretCardEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerSecretCardEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerSecretCardEntity) TableName() string {
	return "t_player_secret_card"
}
