package entity

type ChuangShiShenWangSignUpEntity struct {
	Id         int64 `gorm:"column:id"`
	ServerId   int32 `gorm:"column:serverId"`
	PlayerId   int64 `gorm:"column:playerId"`
	Status     int32 `gorm:"column:status"` //报名状态
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *ChuangShiShenWangSignUpEntity) GetId() int64 {
	return e.Id
}

func (e *ChuangShiShenWangSignUpEntity) TableName() string {
	return "t_chuangshi_shenwang_signup"
}
