package entity

//玩家仙盟玉玺数据
type PlayerAlliancYuXiEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	IsReceive  int32 `gorm:"column:isReceive"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerAlliancYuXiEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerAlliancYuXiEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerAlliancYuXiEntity) TableName() string {
	return "t_player_alliance_yuxi"
}
