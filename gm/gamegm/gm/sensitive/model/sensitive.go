package model

type UserSensitive struct {
	Id         int64  `gorm:"primary_key;column:id"`
	UserId     int64  `gorm:"column:userId"`
	Content    string `gorm:"column:content"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (m *UserSensitive) TableName() string {
	return "t_user_sensitive"
}
