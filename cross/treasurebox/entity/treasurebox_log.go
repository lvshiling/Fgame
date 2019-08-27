package entity

//跨服宝箱日志
type TreasureBoxLogEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	AreaId     int32  `gorm:"column:areaId"`
	ServerId   int32  `gorm:"column:serverId"`
	PlayerName string `gorm:"column:playerName"`
	ItemInfo   string `gorm:"column:itemInfo"`
	LastTime   int64  `gorm:"column:lastTime"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (tlre *TreasureBoxLogEntity) GetId() int64 {
	return tlre.Id
}

func (tlre *TreasureBoxLogEntity) TableName() string {
	return "t_treasurebox_log"
}
