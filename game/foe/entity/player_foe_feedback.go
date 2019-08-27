package entity

//仇人反馈
type PlayerFoeFeedbackEntity struct {
	Id           int64  `gorm:"primary_key;column:id"`
	PlayerId     int64  `gorm:"column:playerId"`
	IsProtected  int32  `gorm:"column:isProtected"`
	FeedbackName string `gorm:"column:feedbackName"`
	UpdateTime   int64  `gorm:"column:updateTime"`
	CreateTime   int64  `gorm:"column:createTime"`
	DeleteTime   int64  `gorm:"column:deleteTime"`
}

func (e *PlayerFoeFeedbackEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerFoeFeedbackEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerFoeFeedbackEntity) TableName() string {
	return "t_player_foe_feedback"
}
