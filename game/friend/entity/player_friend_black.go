package entity

//好友系统数据
type PlayerFriendBlackEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	FriendId   int64 `gorm:"column:friendId"`
	Black      int32 `gorm:"column:black"`
	RevBlack   int32 `gorm:"column:revBlack"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pbse *PlayerFriendBlackEntity) GetId() int64 {
	return pbse.Id
}

func (pbse *PlayerFriendBlackEntity) GetPlayerId() int64 {
	return pbse.PlayerId
}

func (pbse *PlayerFriendBlackEntity) TableName() string {
	return "t_player_friend_black"
}
