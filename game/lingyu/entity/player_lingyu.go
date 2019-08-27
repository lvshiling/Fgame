package entity

//玩家领域数据
type PlayerLingyuEntity struct {
	Id          int64  `gorm:"primary_key;column:id"`
	PlayerId    int64  `gorm:"column:playerId"`
	AdvancedId  int    `gorm:"column:advancedId"`
	LingyuId    int32  `gorm:"column:lingyuId"`
	UnrealLevel int32  `gorm:"column:unrealLevel"`
	UnrealNum   int32  `gorm:"column:unrealNum"`
	UnrealPro   int32  `gorm:"column:unrealPro"`
	UnrealInfo  string `gorm:"column:unrealInfo"`
	TimesNum    int32  `gorm:"column:timesNum"`
	Bless       int32  `gorm:"column:bless"`
	BlessTime   int64  `gorm:"column:blessTime"`
	Hidden      int32  `gorm:"column:hidden"`
	Power       int64  `gorm:"column:power"`
	ChargeVal   int64  `gorm:"column:chargeVal"`
	UpdateTime  int64  `gorm:"column:updateTime"`
	CreateTime  int64  `gorm:"column:createTime"`
	DeleteTime  int64  `gorm:"column:deleteTime"`
}

func (e *PlayerLingyuEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerLingyuEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerLingyuEntity) TableName() string {
	return "t_player_lingyu"
}
