package entity

//玩家原石数据
type PlayerMiningEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Level      int32 `gorm:"column:level"`
	Storage    int32 `gorm:"column:storage"`
	Stone      int64 `gorm:"column:stone"`
	LastTime   int64 `gorm:"column:lastTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pme *PlayerMiningEntity) GetId() int64 {
	return pme.Id
}

func (pme *PlayerMiningEntity) GetPlayerId() int64 {
	return pme.PlayerId
}

func (pme *PlayerMiningEntity) TableName() string {
	return "t_player_mining"
}
