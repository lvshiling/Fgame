package model

type CenterOrderDateStatic struct {
	OrderDate        string `gorm:"column:orderDate"`
	OrderPlayerNum   int    `gorm:"column:orderPlayerNum"`
	OrderNum         int    `gorm:"column:orderNum"`
	OrderMoney       int    `gorm:"column:orderMoney"`
	OrderGold        int    `gorm:"column:orderGold"`
	OrderNewPlayer   int    `gorm:"column:orderNewPlayer"`
	OrderNewMoney    int    `gorm:"column:orderNewMoney"`
	OrderFirstPlayer int    `gorm:"column:orderFirstPlayer"`
	OrderFirstMoney  int    `gorm:"column:orderFirstMoney"`
}
