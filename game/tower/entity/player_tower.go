package entity

//玩家打宝塔数据
type PlayerTowerEntity struct {
	Id            int64 `gorm:"primary_key;column:id"`
	PlayerId      int64 `gorm:"column:playerId"`
	UseTime       int64 `gorm:"column:useTime"`
	ExtralTime    int64 `gorm:"column:extraTime"`
	LastResetTime int64 `gorm:"column:lastResetTime"`
	UpdateTime    int64 `gorm:"column:updateTime"`
	CreateTime    int64 `gorm:"column:createTime"`
	DeleteTime    int64 `gorm:"column:deleteTime"`
}

func (e *PlayerTowerEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerTowerEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerTowerEntity) TableName() string {
	return "t_player_tower"
}
