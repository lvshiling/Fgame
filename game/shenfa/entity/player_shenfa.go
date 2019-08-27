package entity

//玩家身法数据
type PlayerShenfaEntity struct {
	Id          int64  `gorm:"primary_key;column:id"`
	PlayerId    int64  `gorm:"column:playerId"`
	AdvancedId  int    `gorm:"column:advancedId"`
	ShenfaId    int32  `gorm:"column:shenfaId"`
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

func (e *PlayerShenfaEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerShenfaEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerShenfaEntity) TableName() string {
	return "t_player_shenfa"
}
