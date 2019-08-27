package entity

//玩家兵魂数据
type PlayerWeaponEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	WeaponId   int32 `gorm:"column:weaponId"`
	ActiveFlag int32 `gorm:"column:activeFlag"`
	Level      int32 `gorm:"column:level"`
	UpNum      int32 `gorm:"column:upNum"`
	UpPro      int32 `gorm:"column:upPro"`
	CulLevel   int32 `gorm:"column:culLevel"`
	CulNum     int32 `gorm:"column:culNum"`
	CulPro     int32 `gorm:"column:culPro"`
	State      int32 `gorm:"column:state"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pwe *PlayerWeaponEntity) GetId() int64 {
	return pwe.Id
}

func (pwe *PlayerWeaponEntity) GetPlayerId() int64 {
	return pwe.PlayerId
}

func (pwe *PlayerWeaponEntity) TableName() string {
	return "t_player_weapon"
}
