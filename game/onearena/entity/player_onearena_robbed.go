package entity

//玩家灵池被抢记录数据
type PlayerOneArenaRobbedEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	RobName    string `gorm:"column:robName"`
	RobTime    int64  `gorm:"column:robTime"`
	Status     int32  `gorm:"column:status"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (pme *PlayerOneArenaRobbedEntity) GetId() int64 {
	return pme.Id
}

func (pme *PlayerOneArenaRobbedEntity) GetPlayerId() int64 {
	return pme.PlayerId
}

func (pme *PlayerOneArenaRobbedEntity) TableName() string {
	return "t_player_onearena_robbed"
}
