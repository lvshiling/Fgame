package model

type CenterOrderDatePlatformStatic struct {
	SdkType        int `gorm:"column:sdkType"`
	OrderPlayerNum int `gorm:"column:orderPlayerNum"`
	OrderNum       int `gorm:"column:orderNum"`
	OrderMoney     int `gorm:"column:orderMoney"`
}
