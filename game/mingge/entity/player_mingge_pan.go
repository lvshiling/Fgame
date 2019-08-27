package entity

//玩家命盘数据
type PlayerMingGePanEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	Type       int32  `gorm:"column:type"`
	SubType    int32  `gorm:"column:subType"`
	ItemList   string `gorm:"column:itemList"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (pmne *PlayerMingGePanEntity) GetId() int64 {
	return pmne.Id
}

func (pmne *PlayerMingGePanEntity) GetPlayerId() int64 {
	return pmne.PlayerId
}

func (pmoe *PlayerMingGePanEntity) TableName() string {
	return "t_player_mingge_pan"
}
