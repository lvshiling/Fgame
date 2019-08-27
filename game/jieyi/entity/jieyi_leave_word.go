package entity

// 结义留言数据
type JieYiLeaveWordEntity struct {
	Id           int64  `gorm:"primary_key;column:id"`
	ServerId     int32  `gorm:"column:serverId"`
	PlayerId     int64  `gorm:"column:playerId"`
	Name         string `gorm:"column:name"`
	Level        int32  `gorm:"column:level"`
	Role         int32  `gorm:"column:role"`
	Sex          int32  `gorm:"column:sex"`
	Force        int64  `gorm:"column:force"`
	LeaveWord    string `gorm:"column:leaveWord"`
	LastPostTime int64  `gorm:"column:lastPostTime"`
	UpdateTime   int64  `gorm:"column:updateTime"`
	CreateTime   int64  `gorm:"column:createTime"`
	DeleteTime   int64  `gorm:"column:deleteTime"`
}

func (e *JieYiLeaveWordEntity) GetId() int64 {
	return e.Id
}

func (e *JieYiLeaveWordEntity) TableName() string {
	return "t_jieyi_leave_word"
}
