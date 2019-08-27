package entity

//仙盟成员申请数据
type AllianceJoinApplyEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	AllianceId int64  `gorm:"column:allianceId"`
	JoinId     int64  `gorm:"column:joinId"`
	Name       string `gorm:"column:name"`
	Role       int32  `gorm:"column:role"`
	Sex        int32  `gorm:"column:sex"`
	Force      int64  `gorm:"column:force"`
	Level      int32  `gorm:"column:level"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (e *AllianceJoinApplyEntity) GetId() int64 {
	return e.Id
}

func (e *AllianceJoinApplyEntity) TableName() string {
	return "t_alliance_join_apply"
}
