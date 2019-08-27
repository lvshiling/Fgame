package entity

type PlayerTransportationEntity struct {
	Id                     int64  `gorm:"primary_key;column:id"`
	PlayerId               int64  `gorm:"column:playerId"`
	RobList                string `gorm:"column:robList"`
	PersonalTransportTimes int32  `gorm:"column:personalTransportTimes"`
	UpdateTime             int64  `gorm:"column:updateTime"`
	CreateTime             int64  `gorm:"column:createTime"`
	DeleteTime             int64  `gorm:"column:deleteTime"`
}

func (e *PlayerTransportationEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerTransportationEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerTransportationEntity) TableName() string {
	return "t_player_biaoche"
}
