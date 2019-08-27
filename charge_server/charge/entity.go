package charge

//充值订单
type OrderEntity struct {
	Id             int64  `gorm:"primary_key;column:id"`
	OrderId        string `gorm:"column:orderId"`
	SdkOrderId     string `gorm:"column:sdkOrderId"`
	Status         int32  `gorm:"column:status"`
	SdkType        int32  `gorm:"column:sdkType"`
	DevicePlatform int32  `gorm:"column:devicePlatform"`
	PlatformUserId string `gorm:"column:platformUserId"`
	ServerId       int32  `gorm:"column:serverId"`
	UserId         int64  `gorm:"column:userId"`
	PlayerId       int64  `gorm:"column:playerId"`
	PlayerLevel    int32  `gorm:"column:playerLevel"`
	PlayerName     string `gorm:"column:playerName"`
	ChargeId       int32  `gorm:"column:chargeId"`
	Money          int32  `gorm:"column:money"`
	Gold           int32  `gorm:"column:gold"`
	ReceivePayTime int64  `gorm:"column:receivePayTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
	UpdateTime     int64  `gorm:"column:updateTime"`
}

func (e *OrderEntity) TableName() string {
	return "t_order"
}
