package entity

//玩家商城促销数据
type PlayerShopDiscountEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Typ        int32 `gorm:"column:typ"`
	StartTime  int64 `gorm:"column:startTime"`
	EndTime    int64 `gorm:"column:endTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerShopDiscountEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerShopDiscountEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerShopDiscountEntity) TableName() string {
	return "t_player_shop_discount"
}
