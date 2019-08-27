package entity

//玩家收到的好友邀请数据
type PlayerFriendInviteEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	InviteId   int64  `gorm:"column:inviteId"`
	Name       string `gorm:"column:name"`
	Role       int32  `gorm:"column:role"`
	Sex        int32  `gorm:"column:sex"`
	Force      int64  `gorm:"column:force"`
	Level      int32  `gorm:"column:level"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (e *PlayerFriendInviteEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerFriendInviteEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerFriendInviteEntity) TableName() string {
	return "t_player_friend_invite"
}
