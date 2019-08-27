package model

type PlayerMail struct {
	Id              int64  `gorm:"primary_key;column:id"`
	PlayerId        int64  `gorm:"column:playerId"`
	IsRead          int64  `gorm:"column:isRead"`
	IsGetAttachment int64  `gorm:"column:isGetAttachment"`
	Title           string `gorm:"column:title"`
	Content         string `gorm:"column:content"`
	AttachementInfo string `gorm:"column:attachementInfo"`
	UpdateTime      int64  `gorm:"column:updateTime"`
	CreateTime      int64  `gorm:"column:createTime"`
	DeleteTime      int64  `gorm:"column:deleteTime"`
}

func (m *PlayerMail) TableName() string {
	return "t_player_email"
}
