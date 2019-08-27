package entity

//玩家互动数据
type PlayerFriendAddRewEntity struct {
	Id                   int64  `gorm:"primary_key;column:id"`
	PlayerId             int64  `gorm:"column:playerId"`
	FrDummyNum           int32  `gorm:"column:frDummyNum"`
	RewRecord            string `gorm:"column:rewRecord"`
	LastAddDummyTime     int64  `gorm:"column:lastAddDummyTime"`
	LastCongratulateTime int64  `gorm:"column:lastCongratulateTime"`
	CongratulateTimes    int32  `gorm:"column:congratulateTimes"`
	UpdateTime           int64  `gorm:"column:updateTime"`
	CreateTime           int64  `gorm:"column:createTime"`
	DeleteTime           int64  `gorm:"column:deleteTime"`
}

func (e *PlayerFriendAddRewEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerFriendAddRewEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerFriendAddRewEntity) TableName() string {
	return "t_player_friend_add_rew"
}
