package entity

//玩家穿戴称号数据
type PlayerWearTitleEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	TitleWear  int32 `gorm:"column:titleWear"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pwte *PlayerWearTitleEntity) GetId() int64 {
	return pwte.Id
}

func (pwte *PlayerWearTitleEntity) GetPlayerId() int64 {
	return pwte.PlayerId
}

func (pwte *PlayerWearTitleEntity) TableName() string {
	return "t_player_title_wear"
}
