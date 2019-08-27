package model

type GameOrderInfo struct {
	Id          int64  `gorm:"primary_key;gorm:"column:id"`
	ServerId    int    `gorm:"column:serverId"`
	OrderId     string `gorm:"column:orderId"`
	OrderStatus int    `gorm:"column:orderStatus"`
	UserId      int64  `gorm:"column:userId"`
	PlayerId    int64  `gorm:"column:playerId"`
	ChargeId    int    `gorm:"column:chargeId"`
	Money       int    `gorm:"column:money"`
	UpdateTime  int64  `gorm:"column:updateTime"`
	CreateTime  int64  `gorm:"column:createTime"`
	DeleteTime  int64  `gorm:"column:deleteTime"`
	PlayerLevel int    `gorm:"column:playerLevel"`
	Gold        int    `gorm:"column:gold"`
	PlayerName  string `gorm:"column:name"`
}

func (m *GameOrderInfo) TableName() string {
	return "t_order"
}
