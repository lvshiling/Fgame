package entity

//玩家命理数据
type PlayerMingGeMingLiEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	Type       int32  `gorm:"column:type"`
	SubType    int32  `gorm:"column:subType"`
	MingLiList string `gorm:"column:mingLiList"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (pmne *PlayerMingGeMingLiEntity) GetId() int64 {
	return pmne.Id
}

func (pmne *PlayerMingGeMingLiEntity) GetPlayerId() int64 {
	return pmne.PlayerId
}

func (pmoe *PlayerMingGeMingLiEntity) TableName() string {
	return "t_player_mingge_mingli"
}
