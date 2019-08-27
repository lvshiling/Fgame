package entity

//珍稀boss
type PlayerZhenXiBossEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	ReliveTime int32 `gorm:"column:reliveTime"`
	EnterTimes int32 `gorm:"column:enterTimes"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pmne *PlayerZhenXiBossEntity) GetId() int64 {
	return pmne.Id
}

func (pmne *PlayerZhenXiBossEntity) GetPlayerId() int64 {
	return pmne.PlayerId
}

func (pmoe *PlayerZhenXiBossEntity) TableName() string {
	return "t_player_zhenxi_boss"
}
