package model

type QueryPlayerLevelStatic struct {
	Level       int32 `gorm:"column:level"`
	PlayerCount int32 `gorm:"column:playerCount"`
}
