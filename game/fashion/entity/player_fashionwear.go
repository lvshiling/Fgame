package entity

//玩家穿戴时装数据
type PlayerWearFashionEntity struct {
	Id          int64 `gorm:"primary_key;column:id"`
	PlayerId    int64 `gorm:"column:playerId"`
	FashionWear int32 `gorm:"column:fashionWear"`
	UpdateTime  int64 `gorm:"column:updateTime"`
	CreateTime  int64 `gorm:"column:createTime"`
	DeleteTime  int64 `gorm:"column:deleteTime"`
}

func (e *PlayerWearFashionEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerWearFashionEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerWearFashionEntity) TableName() string {
	return "t_player_fashion_wear"
}
