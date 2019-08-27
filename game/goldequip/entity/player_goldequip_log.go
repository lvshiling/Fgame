package entity

//玩家元神金装日志数据
type PlayerGoldEquipLogEntity struct {
	Id           int64  `gorm:"primary_key;column:id"`
	PlayerId     int64  `gorm:"column:playerId"`
	FenJieItemId string `gorm:"column:fenJieItemId"`
	RewItemStr   string `gorm:"column:rewItemStr"`
	UpdateTime   int64  `gorm:"column:updateTime"`
	CreateTime   int64  `gorm:"column:createTime"`
	DeleteTime   int64  `gorm:"column:deleteTime"`
}

func (e *PlayerGoldEquipLogEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerGoldEquipLogEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerGoldEquipLogEntity) TableName() string {
	return "t_player_goldequip_log"
}
