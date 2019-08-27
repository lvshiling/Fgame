package entity

//玩家赞赏数据
type PlayerFriendAdmireEntity struct {
	Id          int64 `gorm:"primary_key;column:id"`
	PlayerId    int64 `gorm:"column:playerId"`
	FriId       int64 `gorm:"column:friId"`
	AdmireTimes int32 `gorm:"column:admireTimes"`
	UpdateTime  int64 `gorm:"column:updateTime"`
	CreateTime  int64 `gorm:"column:createTime"`
	DeleteTime  int64 `gorm:"column:deleteTime"`
}

func (e *PlayerFriendAdmireEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerFriendAdmireEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerFriendAdmireEntity) TableName() string {
	return "t_player_friend_admire"
}
