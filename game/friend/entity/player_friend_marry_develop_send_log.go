package entity

//我的表白记录数据
type PlayerFriendMarryDevelopSendLogEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	PlayerId   int64  `gorm:"column:playerId"`
	RecvId     int64  `gorm:"column:recvId"`
	RecvName   string `gorm:"column:recvName"`
	ItemId     int32  `gorm:"column:itemId"`
	ItemNum    int32  `gorm:"column:itemNum"`
	CharmNum   int32  `gorm:"column:charmNum"`
	DevelopExp int32  `gorm:"column:developExp"`
	ContextStr string `gorm:"column:contextStr"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (pbse *PlayerFriendMarryDevelopSendLogEntity) GetId() int64 {
	return pbse.Id
}

func (pbse *PlayerFriendMarryDevelopSendLogEntity) GetPlayerId() int64 {
	return pbse.PlayerId
}

func (pbse *PlayerFriendMarryDevelopSendLogEntity) TableName() string {
	return "t_player_friend_marry_develop_send_log"
}
