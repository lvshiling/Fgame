package entity

//玩家VIP数据
type PlayerVipEntity struct {
	Id           int64  `gorm:"primary_key;column:id"`
	PlayerId     int64  `gorm:"column:playerId"`
	VipLevel     int32  `gorm:"column:vipLevel"`
	VipStar      int32  `gorm:"column:vipStar"`
	ConsumeLevel int32  `gorm:"column:consumeLevel"`
	ChargeNum    int64  `gorm:"column:chargeNum"`
	FreeGiftMap  string `gorm:"column:freeGiftMap"`
	DiscountMap  string `gorm:"column:discountMap"`
	UpdateTime   int64  `gorm:"column:updateTime"`
	CreateTime   int64  `gorm:"column:createTime"`
	DeleteTime   int64  `gorm:"column:deleteTime"`
}

func (e *PlayerVipEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerVipEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerVipEntity) TableName() string {
	return "t_player_vip"
}
