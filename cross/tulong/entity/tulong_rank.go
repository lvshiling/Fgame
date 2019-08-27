package entity

//跨服屠龙排行榜
type TuLongRankEntity struct {
	Id           int64  `gorm:"primary_key;column:id"`
	Platform     int32  `gorm:"column:platform"`
	AreaId       int32  `gorm:"column:areaId"`
	ServerId     int32  `gorm:"column:serverId"`
	AllianceId   int64  `gorm:"column:allianceId"`
	AllianceName string `gorm:"column:allianceName"`
	KillNum      int32  `gorm:"column:killNum"`
	LastTime     int64  `gorm:"column:lastTime"`
	UpdateTime   int64  `gorm:"column:updateTime"`
	CreateTime   int64  `gorm:"column:createTime"`
	DeleteTime   int64  `gorm:"column:deleteTime"`
}

func (tlre *TuLongRankEntity) GetId() int64 {
	return tlre.Id
}

func (tlre *TuLongRankEntity) TableName() string {
	return "t_tulong_rank"
}
