package entity

//仙盟邀请数据
type AllianceInvitationEntity struct {
	Id           int64  `gorm:"primary_key;column:id"`
	AllianceId   int64  `gorm:"column:allianceId"`
	InvitationId int64  `gorm:"column:invitationId"`
	Name         string `gorm:"column:name"`
	Role         int32  `gorm:"column:role"`
	Sex          int32  `gorm:"column:sex"`
	Force        int64  `gorm:"column:force"`
	Level        int32  `gorm:"column:level"`
	UpdateTime   int64  `gorm:"column:updateTime"`
	CreateTime   int64  `gorm:"column:createTime"`
	DeleteTime   int64  `gorm:"column:deleteTime"`
}

func (e *AllianceInvitationEntity) GetId() int64 {
	return e.Id
}

func (e *AllianceInvitationEntity) TableName() string {
	return "t_alliance_invitation"
}
