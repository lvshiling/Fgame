package entity

type RegisterStaticPlayer struct {
	TotalPlayerCount   int32 `gorm:"column:totalPlayerCount"`
	TodayRegisterCount int32 `gorm:"column:todayRegisterCount"`
}
