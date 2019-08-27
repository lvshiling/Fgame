package entity

//玩家赞赏数据
type PlayerFriendFeedbackEntity struct {
	Id           int64  `gorm:"primary_key;column:id"`
	PlayerId     int64  `gorm:"column:playerId"`
	FriendId     int64  `gorm:"column:friendId"`
	FriendName   string `gorm:"column:friendName"`
	NoticeType   int32  `gorm:"column:noticeType"`
	FeedbackType int32  `gorm:"column:feedbackType"`
	Condition    int32  `gorm:"column:condition"`
	UpdateTime   int64  `gorm:"column:updateTime"`
	CreateTime   int64  `gorm:"column:createTime"`
	DeleteTime   int64  `gorm:"column:deleteTime"`
}

func (e *PlayerFriendFeedbackEntity) GetId() int64 {
	return e.Id
}

func (e *PlayerFriendFeedbackEntity) GetPlayerId() int64 {
	return e.PlayerId
}

func (e *PlayerFriendFeedbackEntity) TableName() string {
	return "t_player_friend_feedback"
}
