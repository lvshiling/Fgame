package entity

//红包数据
type HongBaoEntity struct {
	Id          int64  `gorm:"primary_key;column:id"`
	ServerId    int32  `gorm:"column:serverId"`
	HongBaoType int32  `gorm:"column:hongBaoType"`
	SendId      int64  `gorm:"column:sendId"`
	AwardList   string `gorm:"column:awardList"`
	SnatchLog   string `gorm:"column:snatchLog"`
	UpdateTime  int64  `gorm:"column:updateTime"`
	CreateTime  int64  `gorm:"column:createTime"`
	DeleteTime  int64  `gorm:"column:deleteTime"`
}

func (e *HongBaoEntity) GetId() int64 {
	return e.Id
}

func (e *HongBaoEntity) TableName() string {
	return "t_hongbao"
}
