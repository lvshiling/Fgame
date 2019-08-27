package entity

//玩家心法数据
type PlayerXinFaEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Type       int32 `gorm:"column:typ"`
	Level      int32 `gorm:"column:level"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pxfe *PlayerXinFaEntity) GetId() int64 {
	return pxfe.Id
}

func (pxfe *PlayerXinFaEntity) GetPlayerId() int64 {
	return pxfe.PlayerId
}

func (pxfe *PlayerXinFaEntity) TableName() string {
	return "t_player_xinfa"
}
