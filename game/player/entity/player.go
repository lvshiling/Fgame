package entity

//玩家基础数据
type PlayerEntity struct {
	Id                       int64  `gorm:"primary_key;column:id"`
	ServerId                 int32  `gorm:"column:serverId"`
	OriginServerId           int32  `gorm:"column:originServerId"`
	UserId                   int64  `gorm:"column:userId"`
	Name                     string `gorm:"column:name"`
	Role                     int32  `gorm:"column:role"`
	Sex                      int32  `gorm:"column:sex"`
	LastLoginTime            int64  `gorm:"column:lastLoginTime"`
	OnlineTime               int64  `gorm:"column:onlineTime"`
	LastLogoutTime           int64  `gorm:"column:lastLogoutTime"`
	OfflineTime              int64  `gorm:"column:offlineTime"`
	TotalOnlineTime          int64  `gorm:"column:totalOnlineTime"`
	TodayOnlineTime          int64  `gorm:"column:todayOnlineTime"`
	Forbid                   int32  `gorm:"column:forbid"`
	ForbidText               string `gorm:"column:forbidText"`
	ForbidTime               int64  `gorm:"column:forbidTime"`
	ForbidEndTime            int64  `gorm:"column:forbidEndTime"`
	ForbidName               string `gorm:"column:forbidName"`
	ForbidChat               int32  `gorm:"column:forbidChat"`
	ForbidChatText           string `gorm:"column:forbidChatText"`
	ForbidChatTime           int64  `gorm:"column:forbidChatTime"`
	ForbidChatEndTime        int64  `gorm:"column:forbidChatEndTime"`
	ForbidChatName           string `gorm:"column:forbidChatName"`
	IgnoreChat               int32  `gorm:"column:ignoreChat"`
	IgnoreChatText           string `gorm:"column:ignoreChatText"`
	IgnoreChatTime           int64  `gorm:"column:ignoreChatTime"`
	IgnoreChatEndTime        int64  `gorm:"column:ignoreChatEndTime"`
	IgnoreChatName           string `gorm:"column:ignoreChatName"`
	IsOpenVideo              int32  `gorm:"column:isOpenVideo"`
	PrivilegeType            int32  `gorm:"column:privilegeType"`
	TotalChargeMoney         int64  `gorm:"column:totalChargeMoney"`
	TotalChargeGold          int64  `gorm:"column:totalChargeGold"`
	TotalPrivilegeChargeGold int64  `gorm:"column:totalPrivilegeChargeGold"`
	Online                   int32  `gorm:"column:online"`
	GetNewReward             int32  `gorm:"column:getNewReward"`
	SystemCompensate         int32  `gorm:"column:systemCompensate"`
	SdkType                  int32  `gorm:"column:sdkType"`
	Ip                       string `gorm:"column:ip"`
	TodayChargeMoney         int64  `gorm:"column:todayChargeMoney"`
	YesterdayChargeMoney     int64  `gorm:"column:yesterdayChargeMoney"`
	ChargeTime               int64  `gorm:"column:chargeTime"`
	UpdateTime               int64  `gorm:"column:updateTime"`
	CreateTime               int64  `gorm:"column:createTime"`
	DeleteTime               int64  `gorm:"column:deleteTime"`
}

func (p *PlayerEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerEntity) GetPlayerId() int64 {
	return p.Id
}

func (p *PlayerEntity) TableName() string {
	return "t_player"
}
