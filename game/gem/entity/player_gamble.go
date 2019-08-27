package entity

//玩家赌石数据
type PlayerGambleEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Type       int32 `gorm:"column:type"`
	Num        int32 `gorm:"column:num"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pge *PlayerGambleEntity) GetId() int64 {
	return pge.Id
}

func (pge *PlayerGambleEntity) GetPlayerId() int64 {
	return pge.PlayerId
}

func (pge *PlayerGambleEntity) TableName() string {
	return "t_player_gamble"
}
