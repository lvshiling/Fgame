package entity

//龙椅抢夺记录数据
type EmperorRecordsEntity struct {
	Id          int64  `gorm:"primary_key;column:id"`
	ServerId    int32  `gorm:"column:serverId"`
	Type        int32  `gorm:"column:type"`
	EmperorName string `gorm:"column:emperorName"`
	RobbedName  string `gorm:"column:robbedName"`
	RobTime     int64  `gorm:"column:robTime"`
	ItemInfo    string `gorm:"column:itemInfo"`
	UpdateTime  int64  `gorm:"column:updateTime"`
	CreateTime  int64  `gorm:"column:createTime"`
	DeleteTime  int64  `gorm:"column:deleteTime"`
}

func (ere *EmperorRecordsEntity) GetId() int64 {
	return ere.Id
}

func (ere *EmperorRecordsEntity) TableName() string {
	return "t_emperor_records"
}
