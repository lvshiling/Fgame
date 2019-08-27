package entity

//玩家活动充值
type PlayerOpenActivityChargeEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	GroupId    int32 `gorm:"column:groupId"`
	GoldNum    int32 `gorm:"column:goldNum"`
	StartTime  int64 `gorm:"column:startTime"`
	EndTime    int64 `gorm:"column:endTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerOpenActivityChargeEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerOpenActivityChargeEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerOpenActivityChargeEntity) TableName() string {
	return "t_player_open_activity_charge"
}
