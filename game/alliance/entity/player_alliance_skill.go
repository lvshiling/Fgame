package entity

//玩家仙盟仙术数据
type PlayerAllianceSkillEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	SkillType  int32 `gorm:"column:skillType"`
	Level      int32 `gorm:"column:level"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerAllianceSkillEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerAllianceSkillEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerAllianceSkillEntity) TableName() string {
	return "t_player_alliance_skill"
}
