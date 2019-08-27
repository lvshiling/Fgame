package entity

//合服标志
type MergeEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	ServerId   int32 `gorm:"column:serverId"`
	Merge      int32 `gorm:"column:merge"`
	MergeTime  int64 `gorm:"column:mergeTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *MergeEntity) GetId() int64 {
	return e.Id
}

func (e *MergeEntity) TableName() string {
	return "t_merge"
}
