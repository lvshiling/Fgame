package entity

//玩家时装数据
type PlayerFashionEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	FashionId  int32 `gorm:"column:fashionId"`
	Star       int32 `gorm:"column:star"`
	UpStarNum  int32 `gorm:"column:upStarNum"`
	UpStarPro  int32 `gorm:"column:upStarPro"`
	IsExpire   int32 `gorm:"column:isExpire"`
	ActiveTime int64 `gorm:"column:activeTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerFashionEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerFashionEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerFashionEntity) TableName() string {
	return "t_player_fashion"
}
