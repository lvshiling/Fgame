package model

type QueryPlayer struct {
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
	Level                    int    `gorm:"column:level"`
	ZhuanSheng               int    `gorm:"column:zhuanSheng"`
	Silver                   int    `gorm:"column:silver"`
	Gold                     int    `gorm:"column:gold"`
	BindGold                 int    `gorm:"column:bindGold"`
	Yuanshi                  int    `gorm:"column:yuanshi"`
	AllianceName             string `gorm:"column:allianceName"`
	SpouseName               string `gorm:"column:spouseName"`
	Charm                    int    `gorm:"column:charm"`
	Power                    int    `gorm:"column:power"`
	TotalChargeMoney         int    `gorm:"column:totalChargeMoney"`
	TotalChargeGold          int    `gorm:"column:totalChargeGold"`
	TotalPrivilegeChargeGold int    `gorm:"column:totalPrivilegeChargeGold"`
	PrivilegeType            int    `gorm:"column:privilegeType"`
	Ip                       string `gorm:"column:ip"`
	OriginServerId           int32  `gorm:"column:originServerId"`
	TodayChargeMoney         int64  `gorm:"column:todayChargeMoney"`
	YesterdayChargeMoney     int64  `gorm:"column:yesterdayChargeMoney"`
	SdkType                  int32  `gorm:"column:sdkType"`
}
