package entity

type PlayerWushuangSettingsEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	ItemId     int32 `gorm:"column:itemId"`
	Level      int32 `gorm:"column:level"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerWushuangSettingsEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerWushuangSettingsEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerWushuangSettingsEntity) TableName() string {
	return "t_player_wushuang_settings"
}
