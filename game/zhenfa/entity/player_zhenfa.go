package entity

//玩家阵法数据
type PlayerZhenFaEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Type       int32 `gorm:"column:type"`
	Level      int32 `gorm:"column:level"`
	LevelNum   int32 `gorm:"column:levelNum"`
	LevelPro   int32 `gorm:"column:levelPro"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pmne *PlayerZhenFaEntity) GetId() int64 {
	return pmne.Id
}

func (pmne *PlayerZhenFaEntity) GetPlayerId() int64 {
	return pmne.PlayerId
}

func (pmoe *PlayerZhenFaEntity) TableName() string {
	return "t_player_zhenfa"
}
