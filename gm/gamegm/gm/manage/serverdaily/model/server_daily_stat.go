package model

type ServerDailyStats struct {
	Id             int64 `gorm:"primary_key;column:id"`
	PlatformId     int32 `gorm:"column:platformId"`
	ServerType     int32 `gorm:"column:serverType"`
	ServerId       int32 `gorm:"column:serverId"`
	CurDate        int64 `gorm:"column:curDate"`
	MaxOnlineNum   int32 `gorm:"column:maxOnlineNum"`
	LoginNum       int32 `gorm:"column:loginNum"`
	OrderPlayerNum int   `gorm:"column:orderPlayerNum"`
	OrderNum       int   `gorm:"column:orderNum"`
	OrderMoney     int   `gorm:"column:orderMoney"`
	OrderGold      int   `gorm:"column:orderGold"`
	UpdateTime     int64 `gorm:"column:updateTime"`
	CreateTime     int64 `gorm:"column:createTime"`
	DeleteTime     int64 `gorm:"column:deleteTime"`
}

func (m *ServerDailyStats) TableName() string {
	return "t_server_daily_stats"
}
