package entity

//玩家帝魂镶嵌数据
type PlayerSoulEmbedEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	EmbedInfo  string `gorm:"column:embedInfo"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (psee *PlayerSoulEmbedEntity) GetId() int64 {
	return psee.Id
}

func (psee *PlayerSoulEmbedEntity) GetPlayerId() int64 {
	return psee.PlayerId
}

func (psee *PlayerSoulEmbedEntity) TableName() string {
	return "t_player_soul_embed"
}
