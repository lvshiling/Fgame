package entity

//仙盟仓库数据
type AllianceDepotEntity struct {
	Id           int64  `gorm:"primary_key;column:id"`
	AllianceId   int64  `gorm:"column:allianceId"`
	ItemId       int32  `gorm:"column:itemId"`
	Num          int32  `gorm:"column:num"`
	Index        int32  `grom:"column:index"`
	Level        int32  `grom:"column:level"`
	Used         int32  `gorm:"column:used"`
	BindType     int32  `gorm:"column:bindType"`
	PropertyData string `gorm:"column:porpertyData"`
	UpdateTime   int64  `gorm:"column:updateTime"`
	CreateTime   int64  `gorm:"column:createTime"`
	DeleteTime   int64  `gorm:"column:deleteTime"`
}

func (e *AllianceDepotEntity) GetId() int64 {
	return e.Id
}

func (e *AllianceDepotEntity) TableName() string {
	return "t_alliance_depot"
}
