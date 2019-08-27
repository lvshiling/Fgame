package entity

//玩家任务数据
type PlayerQuestEntity struct {
	Id              int64  `gorm:"primary_key;column:id"`
	PlayerId        int64  `gorm:"column:playerId"`
	QuestId         int32  `gorm:"column:questId"`
	QuestData       string `gorm:"column:questData"`
	CollectItemData string `gorm:"column:collectItemData"`
	QuestState      int32  `gorm:"column:questState"`
	UpdateTime      int64  `gorm:"column:updateTime"`
	CreateTime      int64  `gorm:"column:createTime"`
	DeleteTime      int64  `gorm:"column:deleteTime"`
}

func (p *PlayerQuestEntity) GetId() int64 {
	return p.Id
}

func (p *PlayerQuestEntity) GetPlayerId() int64 {
	return p.PlayerId
}

func (p *PlayerQuestEntity) TableName() string {
	return "t_player_quest"
}
