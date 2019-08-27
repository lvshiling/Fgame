package entity

//boss首杀记录
type OpenActivityBossKillEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	ServerId   int32  `gorm:"column:serverId"`
	GroupId    int32  `gorm:"column:groupId"`
	BossIdList string `gorm:"column:bossIdList"`
	StartTime  int64  `gorm:"column:startTime"`
	EndTime    int64  `gorm:"column:endTime"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (e *OpenActivityBossKillEntity) GetId() int64 {
	return e.Id
}

func (e *OpenActivityBossKillEntity) TableName() string {
	return "t_open_activity_boss_kill"
}
