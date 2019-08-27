package entity

//对我表白记录数据
type PlayerFriendMarryDevelopRecvLogEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	SendId     int64  `gorm:"column:sendId"`
	SendName   string `gorm:"column:sendName"`
	ItemId     int32  `gorm:"column:itemId"`
	ItemNum    int32  `gorm:"column:itemNum"`
	CharmNum   int32  `gorm:"column:charmNum"`
	DevelopExp int32  `gorm:"column:developExp"`
	ContextStr string `gorm:"column:contextStr"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (pbse *PlayerFriendMarryDevelopRecvLogEntity) GetId() int64 {
	return pbse.Id
}

func (pbse *PlayerFriendMarryDevelopRecvLogEntity) GetPlayerId() int64 {
	return pbse.PlayerId
}

func (pbse *PlayerFriendMarryDevelopRecvLogEntity) TableName() string {
	return "t_player_friend_marry_develop_recv_log"
}
