package entity

//排行榜重置周期时间
type RankTimeEntity struct {
	Id            int64 `gorm:"primary_key;column:id"`
	ServerId      int32 `gorm:"column:serverId"`
	ClassRankType int32 `gorm:"column:classRankType"`
	RankType      int32 `gorm:"column:rankType"`
	ThisTime      int64 `gorm:"column:thisTime"`
	UpdateTime    int64 `gorm:"column:updateTime"`
	CreateTime    int64 `gorm:"column:createTime"`
	DeleteTime    int64 `gorm:"column:deleteTime"`
}

func (se *RankTimeEntity) GetId() int64 {
	return se.Id
}

func (se *RankTimeEntity) TableName() string {
	return "t_rank_time"
}
