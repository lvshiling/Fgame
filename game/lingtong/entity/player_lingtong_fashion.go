package entity

//玩家灵童时装数据
type PlayerLingTongFashionEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	LingTongId int32 `gorm:"column:lingTongId"`
	FashionId  int32 `gorm:"column:fashionId"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerLingTongFashionEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerLingTongFashionEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerLingTongFashionEntity) TableName() string {
	return "t_player_lingtong_fashion"
}
