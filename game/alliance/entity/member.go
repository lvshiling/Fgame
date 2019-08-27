package entity

//仙盟成员数据
type AllianceMemberEntity struct {
	Id             int64  `gorm:"primary_key;column:id"`
	AllianceId     int64  `gorm:"column:allianceId"`
	MemberId       int64  `gorm:"column:memberId"`
	LingyuId       int32  `gorm:"column:lingyuId"`
	Role           int32  `gorm:"column:role"`
	Sex            int32  `gorm:"column:sex"`
	Position       int32  `gorm:"column:position"`
	Name           string `gorm:"column:name"`
	Level          int32  `gorm:"column:level"`
	Force          int64  `gorm:"column:force"`
	Vip            int32  `gorm:"column:vip"`
	ZhuanSheng     int32  `gorm:"column:zhuanSheng"`
	GongXian       int64  `gorm:"column:gongXian"`
	JoinTime       int64  `gorm:"column:joinTime"`
	LastLogoutTime int64  `gorm:"column:lastLogoutTime"`
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
}

func (e *AllianceMemberEntity) GetId() int64 {
	return e.Id
}

func (e *AllianceMemberEntity) TableName() string {
	return "t_alliance_member"
}
