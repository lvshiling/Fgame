package entity

//玩家周卡数据
type PlayerWeekEntity struct {
	Id                   int64 `gorm:"primary_key;column:id"`
	PlayerId             int64 `gorm:"column:playerId"`
	SeniorExpireTime     int64 `gorm:"column:seniorExpireTime"`
	SeniorLastDayRewTime int64 `gorm:"column:seniorLastDayRewTime"`
	SeniorCycDay         int32 `gorm:"column:seniorCycDay"`
	JuniorExpireTime     int64 `gorm:"column:juniorExpireTime"`
	JuniorLastDayRewTime int64 `gorm:"column:juniorLastDayRewTime"`
	JuniorCycDay         int32 `gorm:"column:juniorCycDay"`
	UpdateTime           int64 `gorm:"column:updateTime"`
	CreateTime           int64 `gorm:"column:createTime"`
	DeleteTime           int64 `gorm:"column:deleteTime"`
}

func (e *PlayerWeekEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerWeekEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerWeekEntity) TableName() string {
	return "t_player_week"
}
