package entity

//注册设置
type RegisterSettingLogEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	ServerId   int32 `gorm:"column:serverId"`
	Open       int32 `gorm:"column:open"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *RegisterSettingLogEntity) GetId() int64 {
	return e.Id
}

func (e *RegisterSettingLogEntity) TableName() string {
	return "t_register_setting_log"
}
