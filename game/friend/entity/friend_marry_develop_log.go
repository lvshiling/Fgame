package entity

//所有表白记录数据
type FriendMarryDevelopLogEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	ServerId   int32  `gorm:"column:serverId"`
	SendId     int64  `gorm:"column:sendId"`
	RecvId     int64  `gorm:"column:recvId"`
	SendName   string `gorm:"column:sendName"`
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

func (pbse *FriendMarryDevelopLogEntity) GetId() int64 {
	return pbse.Id
}

func (pbse *FriendMarryDevelopLogEntity) TableName() string {
	return "t_friend_marry_develop_log"
}
