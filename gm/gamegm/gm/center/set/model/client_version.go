package model

type ClientVersionEntity struct {
	Id             int64  `gorm:"primary_key;column:id"`
	IosVersion     string `gorm:"column:iosVersion"`
	AndroidVersion string `gorm:"column:androidVersion"`
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
}

func (m *ClientVersionEntity) TableName() string {
	return "t_client_version"
}
