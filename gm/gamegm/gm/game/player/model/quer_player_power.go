package model

type QueryPlayerPower struct {
	PlayerId int64 `gorm:"column:playerId"`
	Power    int32 `gorm:"column:power"`
}
