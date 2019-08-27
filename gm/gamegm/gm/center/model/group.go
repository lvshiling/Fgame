package model

type CenterGroupInfo struct {
	Id         int64 `gorm:"primary_key;gorm:"column:id"`
	GroupId    int64 `gorm:"column:groupId"`
	Platform   int64 `gorm:"column:platform"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (m *CenterGroupInfo) TableName() string {
	return "t_group"
}
