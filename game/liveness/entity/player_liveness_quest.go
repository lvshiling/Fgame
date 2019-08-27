package entity

//玩家活跃度任务数据
type PlayerLivenessQuestEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	QuestId    int32 `gorm:"column:questId"`
	Num        int32 `gorm:"column:num"`
	LastTime   int64 `gorm:"column:lastTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerLivenessQuestEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerLivenessQuestEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerLivenessQuestEntity) TableName() string {
	return "t_player_liveness_quest"
}
