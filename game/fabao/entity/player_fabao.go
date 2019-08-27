package entity

//玩家法宝数据
type PlayerFaBaoEntity struct {
	Id            int64  `gorm:"primary_key;column:id"`
	PlayerId      int64  `gorm:"column:playerId"`
	AdvancedId    int    `gorm:"column:advancedId"`
	FaBaoId       int32  `gorm:"column:faBaoId"`
	UnrealLevel   int32  `gorm:"column:unrealLevel"`
	UnrealNum     int32  `gorm:"column:unrealNum"`
	UnrealPro     int32  `gorm:"column:unrealPro"`
	UnrealInfo    string `gorm:"column:unrealInfo"`
	TimesNum      int32  `gorm:"column:timesNum"`
	Bless         int32  `gorm:"column:bless"`
	BlessTime     int64  `gorm:"column:blessTime"`
	TongLingLevel int32  `gorm:"column:tonglingLevel"`
	TongLingNum   int32  `gorm:"column:tonglingNum"`
	TongLingPro   int32  `gorm:"column:tonglingPro"`
	Hidden        int32  `gorm:"column:hidden"`
	Power         int64  `gorm:"column:power"`
	UpdateTime    int64  `gorm:"column:updateTime"`
	CreateTime    int64  `gorm:"column:createTime"`
	DeleteTime    int64  `gorm:"column:deleteTime"`
}

func (pwe *PlayerFaBaoEntity) GetId() int64 {
	return pwe.Id
}

func (pwe *PlayerFaBaoEntity) GetPlayerId() int64 {
	return pwe.PlayerId
}

func (pwe *PlayerFaBaoEntity) TableName() string {
	return "t_player_fabao"
}
