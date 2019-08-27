package model

type LoginNotice struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlatformId int    `gorm:"column:platformId"`
	Content    string `gorm:"column:content"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (m *LoginNotice) TableName() string {
	return "t_notice_login"
}
