package entity

//玩家物品技能数据
type PlayerItemSkillEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	Typ        int32 `gorm:"column:typ"`
	Level      int32 `gorm:"column:level"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerItemSkillEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerItemSkillEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerItemSkillEntity) TableName() string {
	return "t_player_item_skill"
}
