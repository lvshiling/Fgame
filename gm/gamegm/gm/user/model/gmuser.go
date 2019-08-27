package model

type DBGmUserInfo struct {
	UserId         int64  `gorm:"primary_key;column:id"`
	UserName       string `gorm:"column:userName"`
	Psd            string `gorm:"column:psd"`
	Avator         string `gorm:"column:avator"`
	PrivilegeLevel int    `gorm:"privilege_level"`
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
	ChannelID      int64  `gorm:"column:channelId"`
	PlatformId     int64  `gorm:"column:platformId"`
}

func (m *DBGmUserInfo) TableName() string {
	return "t_gmuser"
}
