package model

type ServerRegisterSetting struct {
	Id         int64 `gorm:"primary_key;column:id"`
	ServerId   int   `gorm:"column:serverId"`
	Open       int   `gorm:"column:open"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (m *ServerRegisterSetting) TableName() string {
	return "t_register_setting"
}
