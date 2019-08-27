package model

type GamePlayer struct {
	Id                       int64  `gorm:"primary_key;column:id"`
	UserId                   int64  `gorm:"column:userId"`
	ServerId                 int64  `gorm:"column:serverId"`
	Name                     string `gorm:"column:name"`
	Role                     int    `gorm:"column:role"`
	Sex                      int    `gorm:"column:sex"`
	LastLoginTime            int64  `gorm:"column:lastLoginTime"`
	LastLogoutTime           int64  `gorm:"column:lastLogoutTime"`
	OnlineTime               int64  `gorm:"column:onlineTime"`
	OfflineTime              int64  `gorm:"column:offlineTime"`
	TotalOnlineTime          int64  `gorm:"column:totalOnlineTime"`
	TodayOnlineTime          int64  `gorm:"column:todayOnlineTime"`
	UpdateTime               int64  `gorm:"column:updateTime"`
	CreateTime               int64  `gorm:"column:createTime"`
	DeleteTime               int64  `gorm:"column:deleteTime"`
	Forbid                   int    `gorm:"column:forbid"`
	ForbidText               string `gorm:"column:forbidText"`
	ForbidTime               int64  `gorm:"column:forbidTime"`
	ForbidEndTime            int64  `gorm:"column:forbidEndTime"`
	ForbidName               string `gorm:"column:forbidName"`
	ForbidChat               int    `gorm:"column:forbidChat"`
	ForbidChatText           string `gorm:"column:forbidChatText"`
	ForbidChatTime           int64  `gorm:"column:forbidChatTime"`
	ForbidChatEndTime        int64  `gorm:"column:forbidChatEndTime"`
	ForbidChatName           string `gorm:"column:forbidChatName"`
	IgnoreChat               int    `gorm:"column:ignoreChat"`
	IgnoreChatText           string `gorm:"column:ignoreChatText"`
	IgnoreChatTime           int64  `gorm:"column:ignoreChatTime"`
	IgnoreChatEndTime        int64  `gorm:"column:ignoreChatEndTime"`
	IgnoreChatName           string `gorm:"column:ignoreChatName"`
	TotalChargeMoney         int    `gorm:"column:totalChargeMoney"`
	TotalChargeGold          int    `gorm:"column:totalChargeGold"`
	TotalPrivilegeChargeGold int    `gorm:"column:totalPrivilegeChargeGold"`
	PrivilegeType            int    `gorm:"column:privilegeType"`
	Ip                       string `gorm:"column:ip"`
}

func (m *GamePlayer) TableName() string {
	return "t_player"
}
