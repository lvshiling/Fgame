package entity

//玩家至尊称号数据
type PlayerSupremeTitleEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	TitleId    int32 `gorm:"column:titleId"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pte *PlayerSupremeTitleEntity) GetId() int64 {
	return pte.Id
}

func (pte *PlayerSupremeTitleEntity) GetPlayerId() int64 {
	return pte.PlayerId
}

func (pte *PlayerSupremeTitleEntity) TableName() string {
	return "t_player_supreme_title"
}
