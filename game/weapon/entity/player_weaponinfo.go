package entity

//玩家兵魂信息数据
type PlayerWeaponInfoEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	WeaponWear int32 `gorm:"column:weaponWear"`
	Star       int32 `gorm:"column:star"`
	Power      int64 `gorm:"column:power`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pwwe *PlayerWeaponInfoEntity) GetId() int64 {
	return pwwe.Id
}

func (pwwe *PlayerWeaponInfoEntity) GetPlayerId() int64 {
	return pwwe.PlayerId
}

func (pwwe *PlayerWeaponInfoEntity) TableName() string {
	return "t_player_weapon_info"
}
