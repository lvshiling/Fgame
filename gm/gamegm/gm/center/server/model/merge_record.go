package model

type ServerMergeRecord struct {
	ID            int64 `gorm:"primary_key;gorm:"column:id"`
	Platform      int64 `gorm:"column:platform"`
	FromServerID  int   `gorm:"column:fromServerId"`
	ToServerID    int   `gorm:"column:toServerId"`
	FinalServerID int   `gorm:"column:finalServerId"`
	MergeTime     int64 `gorm:"column:mergeTime"`
	UpdateTime    int64 `gorm:"column:updateTime"`
	CreateTime    int64 `gorm:"column:createTime"`
	DeleteTime    int64 `gorm:"column:deleteTime"`
}

func (m *ServerMergeRecord) TableName() string {
	return "t_merge_record"
}
