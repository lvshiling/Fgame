package merge

type MergeRecordEntity struct {
	Id            int64 `gorm:"primary_key;column:id"`
	Platform      int32 `gorm:"column:platform"`
	FromServerId  int32 `gorm:"column:fromServerId"`
	ToServerId    int32 `gorm:"column:toServerId"`
	FinalServerId int32 `gorm:"column:finalServerId"`
	MergeTime     int64 `gorm:"column:mergeTime"`
	UpdateTime    int64 `gorm:"column:updateTime"`
	CreateTime    int64 `gorm:"column:createTime"`
	DeleteTime    int64 `gorm:"column:deleteTime"`
}

func (e *MergeRecordEntity) TableName() string {
	return "t_merge_record"
}
