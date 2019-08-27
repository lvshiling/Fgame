package entity

//玩家点星数据
type DingShiBossEntity struct {
	Id           int64 `gorm:"primary_key;column:id"`
	ServerId     int32 `gorm:"column:serverId"`
	MapId        int32 `gorm:"column:mapId"`
	BossId       int32 `gorm:"column:bossId"`
	LastKillTime int64 `gorm:"column:lastKillTime"`
	UpdateTime   int64 `gorm:"column:updateTime"`
	CreateTime   int64 `gorm:"column:createTime"`
	DeleteTime   int64 `gorm:"column:deleteTime"`
}

func (e *DingShiBossEntity) GetId() int64 {
	return e.Id
}

func (e *DingShiBossEntity) TableName() string {
	return "t_dingshi_boss"
}
