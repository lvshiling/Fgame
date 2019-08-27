package entity

//补偿数据
type CompensateEntity struct {
	Id             int64  `gorm:"primary_key;column:id"`
	ServerId       int32  `gorm:"column:serverId"`
	Titlte         string `gorm:"column:titlte"`
	Content        string `gorm:"column:content"`
	Attachment     string `gorm:"column:attachment"`
	RoleLevel      int32  `gorm:"column:roleLevel"`
	RoleCreateTime int64  `gorm:"column:roleCreateTime"`
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
}

func (pwe *CompensateEntity) GetId() int64 {
	return pwe.Id
}

func (pwe *CompensateEntity) TableName() string {
	return "t_compensate"
}
