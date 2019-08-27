package entity

//玩家竞技场boss数据
type ArenaBossEntity struct {
	Id           int64 `gorm:"primary_key;column:id"`
	Platform     int32 `gorm:"column:platform"`
	ServerId     int32 `gorm:"column:serverId"`
	MapId        int32 `gorm:"column:mapId"`
	BossId       int32 `gorm:"column:bossId"`
	LastKillTime int64 `gorm:"column:lastKillTime"`
	UpdateTime   int64 `gorm:"column:updateTime"`
	CreateTime   int64 `gorm:"column:createTime"`
	DeleteTime   int64 `gorm:"column:deleteTime"`
}

func (e *ArenaBossEntity) GetId() int64 {
	return e.Id
}

func (e *ArenaBossEntity) TableName() string {
	return "t_arena_boss"
}
