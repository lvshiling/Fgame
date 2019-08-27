package entity

//玩家四神遗迹数据
type PlayerFourGodEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	KeyNum     int32  `gorm:"column:keyNum"`
	Exp        int64  `gorm:"column:exp"`
	ItemInfo   string `gorm:"column:itemInfo"`
	EndTime    int64  `gorm:"column:endTime"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (p *PlayerFourGodEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerFourGodEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerFourGodEntity) TableName() string {
	return "t_player_four_god"
}
