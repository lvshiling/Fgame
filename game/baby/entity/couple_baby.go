package entity

//全局宝宝数据
type CoupleBabyEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	ServerId   int32  `gorm:"column:serverId"`
	PlayerId   int64  `gorm:"column:playerId"`
	BabyList   string `gorm:"column:babyList"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (e *CoupleBabyEntity) GetId() int64 {
	return e.Id
}

func (e *CoupleBabyEntity) TableName() string {
	return "t_couple_baby"
}
