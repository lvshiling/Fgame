package model

type ServerRegisterSettingLog struct {
	Id         int64 `gorm:"primary_key;column:id"`
	ServerId   int   `gorm:"column:serverId"`
	Open       int   `gorm:"column:open"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (m *ServerRegisterSettingLog) TableName() string {
	return "t_register_setting_log"
}
