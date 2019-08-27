package entity

//玩家血池数据
type PlayerXueChiEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	BloodLine  int32 `gorm:"column:bloodLine"`
	Blood      int64 `gorm:"column:blood"`
	LastTime   int64 `gorm:"column:lastTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pxce *PlayerXueChiEntity) GetId() int64 {
	return pxce.Id
}

func (pxce *PlayerXueChiEntity) GetPlayerId() int64 {
	return pxce.PlayerId
}

func (pxfe *PlayerXueChiEntity) TableName() string {
	return "t_player_xuechi"
}
