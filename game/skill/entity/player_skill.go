package entity

//玩家技能数据
type PlayerSkillEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	SkillId    int32  `gorm:"column:skillId"`
	Level      int32  `gorm:"column:level"`
	TianFuInfo string `gorm:"column:tianFuInfo"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (pse *PlayerSkillEntity) GetId() int64 {
	return pse.Id
}

func (pse *PlayerSkillEntity) GetPlayerId() int64 {
	return pse.PlayerId
}

func (pse *PlayerSkillEntity) TableName() string {
	return "t_player_skill"
}
