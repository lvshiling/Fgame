package entity

//仙盟成员数据
type AllianceHegemonEntity struct {
	Id                int64 `gorm:"primary_key;column:id"`
	ServerId          int32 `gorm:"column:serverId"`
	AllianceId        int64 `gorm:"column:allianceId"`
	WinNum            int32 `gorm:"column:winNum"`
	DefenceAllianceId int64 `gorm:"column:defenceAllianceId"`
	UpdateTime        int64 `gorm:"column:updateTime"`
	CreateTime        int64 `gorm:"column:createTime"`
	DeleteTime        int64 `gorm:"column:deleteTime"`
}

func (e *AllianceHegemonEntity) GetId() int64 {
	return e.Id
}

func (e *AllianceHegemonEntity) TableName() string {
	return "t_alliance_hegemon"
}
