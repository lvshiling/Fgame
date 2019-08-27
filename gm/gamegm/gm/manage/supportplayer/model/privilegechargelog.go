package model

type PrivilegeChargeLog struct {
	Id               int64  `gorm:"primary_key;column:id"`
	PlayerId         int64  `gorm:"column:playerId"`
	PlayerName       string `gorm:"column:playerName"`
	ChannelId        int    `gorm:"column:channelId"`
	PlatformId       int    `gorm:"column:platformId"`
	CenterPlatformId int    `gorm:"column:centerPlatformId"`
	ServerId         int    `gorm:"column:serverId"`
	ServerName       string `gorm:"column:serverName"`
	Gold             int    `gorm:"column:gold"`
	ChargeTime       int64  `gorm:"column:chargeTime"`
	UserName         string `gorm:"column:userName"`
	Reason           string `gorm:"column:reason"`
	UpdateTime       int64  `gorm:"column:updateTime"`
	CreateTime       int64  `gorm:"column:createTime"`
	DeleteTime       int64  `gorm:"column:deleteTime"`
}

func (m *PrivilegeChargeLog) TableName() string {
	return "t_charge_log"
}
