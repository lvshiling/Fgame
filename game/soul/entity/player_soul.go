package entity

//玩家帝魂数据
type PlayerSoulEntity struct {
	Id              int64 `gorm:"primary_key;column:id"`
	PlayerId        int64 `gorm:"column:playerId"`
	SoulTag         int32 `gorm:"column:soulTag"`
	Level           int32 `gorm:"column:level"`
	Experience      int32 `gorm:"column:experience"`
	IsAwaken        int32 `gorm:"column:isAwaken"`
	AwakenOrder     int32 `gorm:"column:awakenOrder"`
	StrengthenLevel int32 `gorm:"column:strengthenLevel"`
	StrengthenNum   int32 `gorm:"column:strengthenNum"`
	StrengthenPro   int32 `gorm:"column:strengthenPro"`
	UpdateTime      int64 `gorm:"column:updateTime"`
	CreateTime      int64 `gorm:"column:createTime"`
	DeleteTime      int64 `gorm:"column:deleteTime"`
}

func (pse *PlayerSoulEntity) GetId() int64 {
	return pse.Id
}

func (pse *PlayerSoulEntity) GetPlayerId() int64 {
	return pse.PlayerId
}

func (pse *PlayerSoulEntity) TableName() string {
	return "t_player_soul"
}
