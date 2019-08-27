package entity

//3v3排行榜刷新时间排行榜
type ArenaRankTimeEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	ServerId   int32 `gorm:"column:serverId"`
	LastTime   int64 `gorm:"column:lastTime"`
	ThisTime   int64 `gorm:"column:thisTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (se *ArenaRankTimeEntity) GetId() int64 {
	return se.Id
}

func (se *ArenaRankTimeEntity) TableName() string {
	return "t_arena_rank_time"
}
