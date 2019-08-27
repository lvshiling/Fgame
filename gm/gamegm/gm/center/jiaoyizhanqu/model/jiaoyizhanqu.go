package model

type JiaoYiZhanQuEntity struct {
	Id         int32  `gorm:"primary_key;column:id"`
	PlatformId int32  `gorm:"column:platformId"`
	ServerId   int32  `gorm:"column:serverId"`
	JiaoYiName string `gorm:"column:jiaoYiName"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (m *JiaoYiZhanQuEntity) TableName() string {
	return "t_jiaoyi_zhanqu"
}
