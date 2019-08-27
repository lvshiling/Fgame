package model

type CenterOrderStatic struct {
	SdkType                int `gorm:"column:sdkType"`
	TodayAmount            int `gorm:"column:todayAmount"`
	TodayPerson            int `gorm:"column:todayPerson"`
	TodayRegisterPerson    int `gorm:"column:todayRegisterPerson"`
	TodayArppu             int `gorm:"column:todayArppu"`
	TodayArpu              int `gorm:"column:todayArpu"`
	YestodayAmount         int `gorm:"column:yestodayAmount"`
	YestodayPerson         int `gorm:"column:yestodayPerson"`
	YestodayRegisterPerson int `gorm:"column:yestodayRegisterPerson"`
	YestodayArppu          int `gorm:"column:yestodayArppu"`
	YestodayArpu           int `gorm:"column:yestodayArpu"`
	TotalAmount            int `gorm:"column:totalAmount"`
	TotalPerson            int `gorm:"column:totalPerson"`
	TotalRegisterPerson    int `gorm:"column:totalRegisterPerson"`
	TotalArppu             int `gorm:"column:totalArppu"`
	TotalArpu              int `gorm:"column:totalArpu"`
	MonthAmount            int `gorm:"monthAmount"`
	MonthPerson            int `gorm:"monthPerson"`
}
