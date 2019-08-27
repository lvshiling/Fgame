package entity

//玩家下线灵池产生的鲲
type PlayerOneArenaKunEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	KunInfo    string `gorm:"column:kunInfo"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (pme *PlayerOneArenaKunEntity) GetId() int64 {
	return pme.Id
}

func (pme *PlayerOneArenaKunEntity) GetPlayerId() int64 {
	return pme.PlayerId
}

func (pme *PlayerOneArenaKunEntity) TableName() string {
	return "t_player_onearena_kun"
}
