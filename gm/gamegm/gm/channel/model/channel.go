package model

type ChannelInfo struct {
	ChannelID   int64  `gorm:"primary_key;column:channelId"`
	ChannelName string `gorm:"column:channelName"`
	UpdateTime  int64  `gorm:"column:updateTime"`
	CreateTime  int64  `gorm:"column:createTime"`
	DeleteTime  int64  `gorm:"column:deleteTime"`
}

func (m *ChannelInfo) TableName() string {
	return "t_channel"
}
