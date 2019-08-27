package entity

//玩家好友日志数据
type PlayerFriendLogEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	PlayerId   int64 `gorm:"column:playerId"`
	FriendId   int64 `gorm:"column:friendId"`
	Type       int32 `gorm:"column:type"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pble *PlayerFriendLogEntity) GetId() int64 {
	return pble.Id
}

func (pble *PlayerFriendLogEntity) GetPlayerId() int64 {
	return pble.PlayerId
}

func (pble *PlayerFriendLogEntity) TableName() string {
	return "t_player_friend_log"
}
