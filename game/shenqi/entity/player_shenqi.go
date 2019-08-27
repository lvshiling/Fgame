package entity

//玩家神器数据
type PlayerShenQiEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	LingQiNum  int64 `gorm:"column:lingQiNum"`
	Power      int64 `gorm:"column:power"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pmoe *PlayerShenQiEntity) GetId() int64 {
	return pmoe.Id
}

func (pmoe *PlayerShenQiEntity) GetPlayerId() int64 {
	return pmoe.PlayerId
}

func (pmoe *PlayerShenQiEntity) TableName() string {
	return "t_player_shenqi"
}
