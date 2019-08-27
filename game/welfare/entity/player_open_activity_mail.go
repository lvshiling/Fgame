package entity

//活动开启邮件记录
type PlayerActivityOpenMailEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Group      int32 `gorm:"column:group"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerActivityOpenMailEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerActivityOpenMailEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerActivityOpenMailEntity) TableName() string {
	return "t_player_open_activity_mail"
}
