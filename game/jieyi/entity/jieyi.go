package entity

// 结义数据
type JieYiEntity struct {
	Id             int64  `gorm:"primary_key;column:id"`
	ServerId       int32  `gorm:"column:serverId"`
	OriginServerId int32  `gorm:"column:originServerId"`
	Name           string `gorm:"column:name"`
	MemberNum      int32  `gorm:"column:memberNum"` //(弃用)
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
}

func (e *JieYiEntity) GetId() int64 {
	return e.Id
}

func (e *JieYiEntity) TableName() string {
	return "t_jieyi"
}
