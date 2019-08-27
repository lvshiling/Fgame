package entity

//好友系统数据
type FriendEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	ServerId   int32 `gorm:"column:serverId"`
	PlayerId   int64 `gorm:"column:playerId"`
	FriendId   int64 `gorm:"column:friendId"`
	Point      int32 `gorm:"column:point"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (pbse *FriendEntity) GetId() int64 {
	return pbse.Id
}

func (pbse *FriendEntity) TableName() string {
	return "t_friend"
}
