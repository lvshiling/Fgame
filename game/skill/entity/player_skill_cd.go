package entity

//玩家技能cd数据
type PlayerSkillCdEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	SkillId    int32 `gorm:"column:skillId"`
	LastTime   int64 `gorm:"column:lastTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pse *PlayerSkillCdEntity) GetId() int64 {
	return pse.Id
}

func (pse *PlayerSkillCdEntity) GetPlayerId() int64 {
	return pse.PlayerId
}

func (pse *PlayerSkillCdEntity) TableName() string {
	return "t_player_skill_cd"
}
