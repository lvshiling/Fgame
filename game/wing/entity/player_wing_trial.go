package entity

//玩家战翼试用阶数
type PlayerWingTrialEntity struct {
	Id           int64 `gorm:"primary_key;column:id"`
	PlayerId     int64 `gorm:"column:playerId"`
	TrialOrderId int32 `gorm:"column:trialOrderId"`
	ActiveTime   int64 `gorm:"column:activeTime"`
	UpdateTime   int64 `gorm:"column:updateTime"`
	CreateTime   int64 `gorm:"column:createTime"`
	DeleteTime   int64 `gorm:"column:deleteTime"`
}

func (pwce *PlayerWingTrialEntity) GetId() int64 {
	return pwce.Id
}

func (pwce *PlayerWingTrialEntity) GetPlayerId() int64 {
	return pwce.PlayerId
}

func (pwce *PlayerWingTrialEntity) TableName() string {
	return "t_player_wing_trial"
}
