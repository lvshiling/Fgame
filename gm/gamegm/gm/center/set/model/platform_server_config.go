package model

type PlatformServerConfigEntity struct {
	Id            int64  `gorm:"primary_key;column:id"`
	TradeServerIp string `gorm:"column:tradeServerIp"`
	UpdateTime    int64  `gorm:"column:updateTime"`
	CreateTime    int64  `gorm:"column:createTime"`
	DeleteTime    int64  `gorm:"column:deleteTime"`
}

func (m *PlatformServerConfigEntity) TableName() string {
	return "t_platform_server_config"
}
