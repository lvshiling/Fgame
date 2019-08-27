package entity

//玩家帝陵遗迹章节奖励
type PlayerSoulRuinsRewChapterEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Chapter    int32 `gorm:"column:chapter"`
	Type       int32 `gorm:"column:type"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (p *PlayerSoulRuinsRewChapterEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerSoulRuinsRewChapterEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerSoulRuinsRewChapterEntity) TableName() string {
	return "t_player_soulruins_rewchapter"
}
