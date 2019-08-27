package entity

//神魔战场排行榜
type ShenMoRankEntity struct {
	Id           int64  `gorm:"primary_key;column:id"`
	ServerId     int32  `gorm:"column:serverId"`
	AllianceId   int64  `gorm:"column:allianceId"`
	AllianceName string `gorm:"column:allianceName"`
	JiFenNum     int32  `gorm:"column:jiFenNum"`
	LastJiFenNum int32  `gorm:"column:lastJiFenNum"`
	LastTime     int64  `gorm:"column:lastTime"`
	UpdateTime   int64  `gorm:"column:updateTime"`
	CreateTime   int64  `gorm:"column:createTime"`
	DeleteTime   int64  `gorm:"column:deleteTime"`
}

func (se *ShenMoRankEntity) GetId() int64 {
	return se.Id
}

func (se *ShenMoRankEntity) TableName() string {
	return "t_shenmo_rank"
}
