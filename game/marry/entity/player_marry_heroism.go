package entity

//玩家婚姻数据
type PlayerMarryHeroismEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Heroism    int32 `gorm:"column:heroism"`
	OutOfTime  int64 `gorm:"column:outOfTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (p *PlayerMarryHeroismEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerMarryHeroismEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerMarryHeroismEntity) TableName() string {
	return "t_player_marry_heroism"
}
