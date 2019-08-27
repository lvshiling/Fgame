package entity

//仙盟日志
type AllianceLogEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	Content    string `gorm:"column:content"`
	AllianceId int64  `gorm:"column:allianceId"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (e *AllianceLogEntity) GetId() int64 {
	return e.Id
}

func (e *AllianceLogEntity) TableName() string {
	return "t_alliance_log"
}
