package entity

//玩家元神金装设置
type PlayerGoldEquipSettingEntity struct {
	Id             int64 `gorm:"primary_key;column:id"`
	PlayerId       int64 `gorm:"column:playerId"`
	FenJieIsAuto   int32 `gorm:"column:fenJieIsAuto"`
	FenJieQuality  int32 `gorm:"column:fenJieQuality"`
	FenJieZhuanShu int32 `gorm:"column:fenJieZhuanShu"`
	IsCheckOldSt   int32 `gorm:"column:isCheckOldSt"`
	UpdateTime     int64 `gorm:"column:updateTime"`
	CreateTime     int64 `gorm:"column:createTime"`
	DeleteTime     int64 `gorm:"column:deleteTime"`
}

func (e *PlayerGoldEquipSettingEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerGoldEquipSettingEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerGoldEquipSettingEntity) TableName() string {
	return "t_player_goldequip_setting"
}
