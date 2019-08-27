package entity

//玩家战力记录数据
type PlayerPowerRecordEntity struct {
	Id             int64 `gorm:"primary_key;column:id"`
	PlayerId       int64 `gorm:"column:playerId"`
	TodayInitPower int64 `gorm:"column:todayInitPower"`
	HisMaxPower    int64 `gorm:"column:hisMaxPower"`
	UpdateTime     int64 `gorm:"column:updateTime"`
	CreateTime     int64 `gorm:"column:createTime"`
	DeleteTime     int64 `gorm:"column:deleteTime"`
}

func (psm *PlayerPowerRecordEntity) GetId() int64 {
	return psm.Id
}

func (psm *PlayerPowerRecordEntity) GetPlayerId() int64 {
	return psm.PlayerId
}

func (psm *PlayerPowerRecordEntity) TableName() string {
	return "t_player_power_record"
}
