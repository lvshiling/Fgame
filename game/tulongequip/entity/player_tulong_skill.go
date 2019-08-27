package entity

//玩家屠龙套装技能数据
type PlayerTuLongSuitSkillEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	SuitType   int32 `gorm:"column:suitType"`
	Level      int32 `gorm:"column:level"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerTuLongSuitSkillEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerTuLongSuitSkillEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerTuLongSuitSkillEntity) TableName() string {
	return "t_player_tulong_suit_skill"
}
