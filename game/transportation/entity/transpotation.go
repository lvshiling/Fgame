package entity

type TransportationEntity struct {
	Id                     int64  `gorm:"primary_key;column:id"`
	PlayerId               int64  `gorm:"column:playerId"`
	ServerId               int32  `gorm:"column:serverId"`
	AllianceId             int64  `gorm:"column:allianceId"`
	TransportMoveId        int32  `gorm:"column:transportMoveId"`
	TransportType          int32  `gorm:"column:transportType"`
	State                  int32  `gorm:"column:state"`
	OwerName               string `gorm:"column:owerName"`
	RobName                string `gorm:"column:robName"`
	LastDistressUpdateTime int64  `gorm:"column:lastDistressUpdateTime"`
	UpdateTime             int64  `gorm:"column:updateTime"`
	CreateTime             int64  `gorm:"column:createTime"`
	DeleteTime             int64  `gorm:"column:deleteTime"`
}

func (e *TransportationEntity) GetId() int64 {
	return e.Id
}

func (e *TransportationEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *TransportationEntity) TableName() string {
	return "t_biaoche"
}
