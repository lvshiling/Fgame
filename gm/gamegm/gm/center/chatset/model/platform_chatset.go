package model

type PlatformChatSetInfo struct {
	Id               int    `gorm:"primary_key;column:id"`
	PlatformId       int    `gorm:"column:platformId"`
	MinVip           int    `gorm:"column:minVip"`
	MinPlayerlevel   int    `gorm:"column:minPlayerlevel"`
	StartTime        string `gorm:"column:startTime"`
	EndTime          string `gorm:"column:endTime"`
	UpdateTime       int64  `gorm:"column:updateTime"`
	CreateTime       int64  `gorm:"column:createTime"`
	DeleteTime       int64  `gorm:"column:deleteTime"`
	WorldVip         int    `gorm:"column:worldVip"`
	WorldPlayerLevel int    `gorm:"column:worldPlayerLevel"`
	PChatVip         int    `gorm:"column:pChatVip"`
	PChatPlayerLevel int    `gorm:"column:pChatPlayerLevel"`
	GuildVip         int    `gorm:"column:guildVip"`
	GuildPlayerLevel int    `gorm:"column:guildPlayerLevel"`
	TeamVip          int    `gorm:"column:teamVip"`
	TeamPlayerLevel  int    `gorm:"column:teamPlayerLevel"`
}

func (m *PlatformChatSetInfo) TableName() string {
	return "t_platform_chatset"
}
