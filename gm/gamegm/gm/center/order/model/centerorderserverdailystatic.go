package model

type CenterOrderServerDailyStatic struct {
	SdkType        int `gorm:"column:sdkType"`
	ServerId       int `gorm:"column:serverId"`
	OrderPlayerNum int `gorm:"column:orderPlayerNum"`
	OrderNum       int `gorm:"column:orderNum"`
	OrderMoney     int `gorm:"column:orderMoney"`
	OrderGold      int `gorm:"column:orderGold"`
}
