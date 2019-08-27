package entity

//注册设置
type RegisterSettingEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	ServerId   int32 `gorm:"column:serverId"`
	Open       int32 `gorm:"column:open"`
	Auto       int32 `gorm:"column:auto"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *RegisterSettingEntity) GetId() int64 {
	return e.Id
}

func (e *RegisterSettingEntity) TableName() string {
	return "t_register_setting"
}
