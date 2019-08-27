package entity

//玩家称号数据
type PlayerTitleEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	TitleId    int32 `gorm:"column:titleId"`
	ActiveFlag int32 `gorm:"column:activeFlag"`
	ActiveTime int64 `gorm:"column:activeTime"`
	ValidTime  int64 `gorm:"column:validTime"`
	StarLev    int32 `gorm:"column:starLev"`
	StarNum    int32 `gorm:"column:starNum"`
	StarBless  int32 `gorm:"column:starBless"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pte *PlayerTitleEntity) GetId() int64 {
	return pte.Id
}

func (pte *PlayerTitleEntity) GetPlayerId() int64 {
	return pte.PlayerId
}

func (pte *PlayerTitleEntity) TableName() string {
	return "t_player_title"
}
