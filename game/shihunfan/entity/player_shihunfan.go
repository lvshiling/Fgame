package entity

//玩家噬魂幡数据
type PlayerShiHunFanEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	AdvancedId int   `gorm:"column:advancedId"`
	TimesNum   int32 `gorm:"column:timesNum"`
	Bless      int32 `gorm:"column:bless"`
	BlessTime  int64 `gorm:"column:blessTime"`
	DanLevel   int32 `gorm:"column:danLevel"`
	DanNum     int32 `gorm:"column:danNum"`
	DanPro     int32 `gorm:"column:danPro"`
	ChargeVal  int32 `gorm:"column:chargeVal"`
	Power      int64 `gorm:"column:power"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerShiHunFanEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerShiHunFanEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerShiHunFanEntity) TableName() string {
	return "t_player_shihunfan"
}
