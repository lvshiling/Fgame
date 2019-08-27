package entity

//玩家天魔体数据
type PlayerTianMoEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	AdvancedId int32 `gorm:"column:advancedId"`
	DanLevel   int32 `gorm:"column:danLevel"`
	DanNum     int32 `gorm:"column:danNum"`
	DanPro     int32 `gorm:"column:danPro"`
	TimesNum   int32 `gorm:"column:timesNum"`
	Bless      int32 `gorm:"column:bless"`
	BlessTime  int64 `gorm:"column:blessTime"`
	ChargeVal  int64 `gorm:"column:chargeVal"`
	Power      int64 `gorm:"column:power"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerTianMoEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerTianMoEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerTianMoEntity) TableName() string {
	return "t_player_tianmo"
}
