package entity

//玩家阵旗仙火数据
type PlayerZhenQiXianHuoEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Type       int32 `gorm:"column:type"`
	Level      int32 `gorm:"column:level"`
	LuckyStar  int32 `gorm:"column:luckyStar"`
	LevelNum   int32 `gorm:"column:levelNum"`
	LevelPro   int32 `gorm:"column:levelPro"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pe *PlayerZhenQiXianHuoEntity) GetId() int64 {
	return pe.Id
}

func (pe *PlayerZhenQiXianHuoEntity) GetPlayerId() int64 {
	return pe.PlayerId
}

func (pe *PlayerZhenQiXianHuoEntity) TableName() string {
	return "t_player_zhenqi_xianhuo"
}
