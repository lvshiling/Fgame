package entity

//玩家灵童时装试用数据
type PlayerLingTongFashionTrialEntity struct {
	Id             int64 `gorm:"primary_key;column:id"`
	PlayerId       int64 `gorm:"column:playerId"`
	TrialFashionId int32 `gorm:"column:trialFashionId"`
	ActivateTime   int64 `gorm:"column:activateTime"`
	DurationTime   int64 `gorm:"column:durationTime"`
	IsExpire       int32 `gorm:"column:isExpire"`
	UpdateTime     int64 `gorm:"column:updateTime"`
	CreateTime     int64 `gorm:"column:createTime"`
	DeleteTime     int64 `gorm:"column:deleteTime"`
}

func (e *PlayerLingTongFashionTrialEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerLingTongFashionTrialEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerLingTongFashionTrialEntity) TableName() string {
	return "t_player_lingtong_fashion_trial"
}
