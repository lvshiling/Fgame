package entity

//玩家时装使用数据
type PlayerFashionTrialEntity struct {
	Id             int64 `gorm:"primary_key;column:id"`
	PlayerId       int64 `gorm:"column:playerId"`
	TrialFashionId int32 `gorm:"column:trialFashionId"`
	ExpireTime     int64 `gorm:"column:expireTime"`
	UpdateTime     int64 `gorm:"column:updateTime"`
	CreateTime     int64 `gorm:"column:createTime"`
	DeleteTime     int64 `gorm:"column:deleteTime"`
}

func (e *PlayerFashionTrialEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerFashionTrialEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerFashionTrialEntity) TableName() string {
	return "t_player_fashion_trial"
}
