package entity

//黑货商店次数限制数据
type OpenActivityDiscountLimitEntity struct {
	Id          int64  `gorm:"primary_key;column:id"`
	ServerId    int32  `gorm:"column:serverId"`
	GroupId     int32  `gorm:"column:groupId"`
	DiscountDay int32  `gorm:"column:discountDay"`
	TimesMap    string `gorm:"column:timesMap"`
	StartTime   int64  `gorm:"column:startTime"`
	EndTime     int64  `gorm:"column:endTime"`
	UpdateTime  int64  `gorm:"column:updateTime"`
	CreateTime  int64  `gorm:"column:createTime"`
	DeleteTime  int64  `gorm:"column:deleteTime"`
}

func (e *OpenActivityDiscountLimitEntity) GetId() int64 {
	return e.Id
}

func (e *OpenActivityDiscountLimitEntity) TableName() string {
	return "t_open_activity_discount_limit"
}
