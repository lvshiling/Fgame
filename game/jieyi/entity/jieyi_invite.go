package entity

// 结义邀请数据
type JieYiInviteEntity struct {
	Id             int64  `gorm:"primary_key;column:id"`
	ServerId       int32  `gorm:"column:serverId"`
	State          int32  `gorm:"column:state"`
	DaoJu          int32  `gorm:"column:daoJu"`
	InviteDaoJu    int32  `gorm:"column:inviteDaoJu"`
	InviteToken    int32  `gorm:"column:inviteToken"`
	InviteTokenLev int32  `gorm:"column:inviteTokenLev"`
	NameLev        int32  `gorm:"column:nameLev"`
	InviteId       int64  `gorm:"column:inviteId"`
	InviteeId      int64  `gorm:"column:inviteeId"`
	Name           string `gorm:"column:name"`
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
}

func (e *JieYiInviteEntity) GetId() int64 {
	return e.Id
}

func (e *JieYiInviteEntity) TableName() string {
	return "t_jieyi_invite"
}
