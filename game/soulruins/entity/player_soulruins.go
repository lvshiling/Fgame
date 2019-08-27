package entity

//玩家帝陵遗迹数据
type PlayerSoulRuinsEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Chapter    int32 `gorm:"column:chapter"`
	Type       int32 `gorm:"column:type"`
	Level      int32 `gorm:"column:level"`
	Star       int32 `gorm:"column:star"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (p *PlayerSoulRuinsEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerSoulRuinsEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerSoulRuinsEntity) TableName() string {
	return "t_player_soulruins"
}
