package entity

//玩家当日宝库商店购买限购道具数据
type PlayerEquipBaoKuShopEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	ShopId     int32 `gorm:"column:shopId"`
	DayCount   int32 `gorm:"column:dayCount"`
	LastTime   int64 `gorm:"column:lastTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (psde *PlayerEquipBaoKuShopEntity) GetId() int64 {
	return psde.Id
}

func (psde *PlayerEquipBaoKuShopEntity) GetPlayerId() int64 {
	return psde.PlayerId
}

func (psde *PlayerEquipBaoKuShopEntity) TableName() string {
	return "t_player_equipbaoku_shop"
}
