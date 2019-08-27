package entity

//玩家聊天
type PlayerChatEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	ChatCount  int32 `gorm:"column:chatCount"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *PlayerChatEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerChatEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerChatEntity) TableName() string {
	return "t_player_chat"
}
