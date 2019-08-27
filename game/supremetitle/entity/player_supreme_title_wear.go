package entity

//玩家穿戴至尊称号数据
type PlayerWearSupremeTitleEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	TitleWear  int32 `gorm:"column:titleWear"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pwte *PlayerWearSupremeTitleEntity) GetId() int64 {
	return pwte.Id
}

func (pwte *PlayerWearSupremeTitleEntity) GetPlayerId() int64 {
	return pwte.PlayerId
}

func (pwte *PlayerWearSupremeTitleEntity) TableName() string {
	return "t_player_supreme_title_wear"
}
