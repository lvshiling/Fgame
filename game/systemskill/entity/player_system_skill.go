package entity

//玩家系统技能数据
type PlayerSystemSkillEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Type       int32 `gorm:"column:typ"`
	SubType    int32 `gorm:"column:subType"`
	Level      int32 `gorm:"column:level"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerSystemSkillEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerSystemSkillEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerSystemSkillEntity) TableName() string {
	return "t_player_system_skill"
}
