package entity

//充值订单
type OrderEntity struct {
	Id          int64  `gorm:"primary_key;column:id"`
	ServerId    int32  `gorm:"column:serverId"`
	OrderId     string `gorm:"column:orderId"`
	OrderStatus int32  `gorm:"column:orderStatus"`
	PlayerId    int64  `gorm:"column:playerId"`
	PlayerLevel int32  `gorm:"column:playerLevel"`
	ChargeId    int32  `gorm:"column:chargeId"`
	Money       int32  `gorm:"column:money"`
	Gold        int32  `gorm:"column:gold"`
	UpdateTime  int64  `gorm:"column:updateTime"`
	CreateTime  int64  `gorm:"column:createTime"`
	DeleteTime  int64  `gorm:"column:deleteTime"`
}

func (e *OrderEntity) GetId() int64 {
	return e.Id
}

func (e *OrderEntity) TableName() string {
	return "t_order"
}
