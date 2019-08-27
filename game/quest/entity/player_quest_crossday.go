package entity

//玩家任务模块过12点数据
type PlayerQuestCrossDayEntity struct {
	Id           int64 `gorm:"primary_key;column:id"`
	PlayerId     int64 `gorm:"column:playerId"`
	CrossDayTime int64 `gorm:"column:crossDayTime"`
	UpdateTime   int64 `gorm:"column:updateTime"`
	CreateTime   int64 `gorm:"column:createTime"`
	DeleteTime   int64 `gorm:"column:deleteTime"`
}

func (p *PlayerQuestCrossDayEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerQuestCrossDayEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerQuestCrossDayEntity) TableName() string {
	return "t_player_quest_crossday"
}
