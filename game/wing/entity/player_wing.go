package entity

//玩家战翼数据
type PlayerWingEntity struct {
	Id          int64  `gorm:"primary_key;column:id"`
	PlayerId    int64  `gorm:"column:playerId"`
	AdvancedId  int    `gorm:"column:advancedId"`
	WingId      int32  `gorm:"column:wingId"`
	UnrealLevel int32  `gorm:"column:unrealLevel"`
	UnrealNum   int32  `gorm:"column:unrealNum"`
	UnrealPro   int32  `gorm:"column:unrealPro"`
	UnrealInfo  string `gorm:"column:unrealInfo"`
	TimesNum    int32  `gorm:"column:timesNum"`
	Bless       int32  `gorm:"column:bless"`
	BlessTime   int64  `gorm:"column:blessTime"`
	FeatherId   int32  `gorm:"column:featherId"`
	FeatherNum  int32  `gorm:"column:featherNum"`
	FeatherPro  int32  `gorm:"column:featherPro"`
	Hidden      int32  `gorm:"column:hidden"`
	Power       int64  `gorm:"column:power"`
	FPower      int64  `gorm:"column:fpower"`
	UpdateTime  int64  `gorm:"column:updateTime"`
	CreateTime  int64  `gorm:"column:createTime"`
	DeleteTime  int64  `gorm:"column:deleteTime"`
}

func (pwe *PlayerWingEntity) GetId() int64 {
	return pwe.Id
}

func (pwe *PlayerWingEntity) GetPlayerId() int64 {
	return pwe.PlayerId
}

func (pwe *PlayerWingEntity) TableName() string {
	return "t_player_wing"
}
