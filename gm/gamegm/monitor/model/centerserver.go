package model

type CenterServer struct {
	Id         int64  `gorm:"primary_key;gorm:"column:id"`
	ServerType int    `gorm:"column:serverType"`
	ServerId   int    `gorm:"column:serverId"`
	Platform   int64  `gorm:"column:platform"`
	ServerName string `gorm:"column:name"`
	StartTime  int64  `gorm:"column:startTime"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (m *CenterServer) TableName() string {
	return "t_server"
}
