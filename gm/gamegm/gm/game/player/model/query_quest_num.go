package model

type QueryPlayerQuestStatic struct {
	QuestId     int32 `gorm:"column:questId"`
	PlayerCount int32 `gorm:"column:playerCount"`
}
