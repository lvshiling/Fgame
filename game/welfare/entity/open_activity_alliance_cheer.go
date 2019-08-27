package entity

//城战助威
type OpenActivityAllianceCheerEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	ServerId   int32 `gorm:"column:serverId"`
	GroupId    int32 `gorm:"column:groupId"`
	AllianceId int64 `gorm:"column:allianceId"`
	StartTime  int64 `gorm:"column:startTime"`
	EndTime    int64 `gorm:"column:endTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *OpenActivityAllianceCheerEntity) GetId() int64 {
	return e.Id
}

func (e *OpenActivityAllianceCheerEntity) TableName() string {
	return "t_open_activity_alliance_cheer"
}
