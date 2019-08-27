package entity

//玩家仙体数据
type PlayerXianTiEntity struct {
	Id          int64  `gorm:"primary_key;column:id"`
	PlayerId    int64  `gorm:"column:playerId"`
	AdvancedId  int    `gorm:"column:advancedId"`
	XianTiId    int32  `gorm:"column:xianTiId"`
	UnrealLevel int32  `gorm:"column:unrealLevel"`
	UnrealNum   int32  `gorm:"column:unrealNum"`
	UnrealPro   int32  `gorm:"column:unrealPro"`
	UnrealInfo  string `gorm:"column:unrealInfo"`
	TimesNum    int32  `gorm:"column:timesNum"`
	Bless       int32  `gorm:"column:bless"`
	BlessTime   int64  `gorm:"column:blessTime"`
	Hidden      int32  `gorm:"column:hidden"`
	Power       int64  `gorm:"column:power"`
	UpdateTime  int64  `gorm:"column:updateTime"`
	CreateTime  int64  `gorm:"column:createTime"`
	DeleteTime  int64  `gorm:"column:deleteTime"`
}

func (pme *PlayerXianTiEntity) GetId() int64 {
	return pme.Id
}

func (pme *PlayerXianTiEntity) GetPlayerId() int64 {
	return pme.PlayerId
}

func (pme *PlayerXianTiEntity) TableName() string {
	return "t_player_xianti"
}
