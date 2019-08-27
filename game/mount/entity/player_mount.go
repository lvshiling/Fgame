package entity

//玩家坐骑数据
type PlayerMountEntity struct {
	Id          int64  `gorm:"primary_key;column:id"`
	PlayerId    int64  `gorm:"column:playerId"`
	AdvancedId  int    `gorm:"column:advancedId"`
	MountId     int32  `gorm:"column:mountId"`
	UnrealLevel int32  `gorm:"column:unrealLevel"`
	UnrealNum   int32  `gorm:"column:unrealNum"`
	UnrealPro   int32  `gorm:"column:unrealPro"`
	CulLevel    int32  `gorm:"column:culLevel"`
	CulNum      int32  `gorm:"column:culNum"`
	CulPro      int32  `gorm:"column:culPro"`
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

func (pme *PlayerMountEntity) GetId() int64 {
	return pme.Id
}

func (pme *PlayerMountEntity) GetPlayerId() int64 {
	return pme.PlayerId
}

func (pme *PlayerMountEntity) TableName() string {
	return "t_player_mount"
}
