package entity

//仙盟boss
type AllianceBossEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	ServerId   int32 `gorm:"column:serverId"`
	AllianceId int64 `gorm:"column:allianceId"`
	SummonTime int64 `gorm:"column:summonTime"`
	BossLevel  int32 `gorm:"column:bossLevel"`
	BossExp    int32 `gorm:"column:bossExp"`
	IsSummon   int32 `gorm:"column:isSummon"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *AllianceBossEntity) GetId() int64 {
	return e.Id
}

func (e *AllianceBossEntity) TableName() string {
	return "t_alliance_boss"
}
