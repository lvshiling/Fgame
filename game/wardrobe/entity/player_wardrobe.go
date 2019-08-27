package entity

//玩家衣橱数据
type PlayerWardrobeEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Type       int32 `gorm:"column:type"`
	SubType    int32 `gorm:"column:subType"`
	ActiveFlag int32 `gorm:"column:activeFlag"`
	Permanent  int32 `gorm:"column:permanent"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerWardrobeEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerWardrobeEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerWardrobeEntity) TableName() string {
	return "t_player_wardrobe"
}
