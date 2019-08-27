package entity

//龙椅数据
type EmperorEntity struct {
	Id                int64  `gorm:"primary_key;column:id"`
	ServerId          int32  `gorm:"column:serverId"`
	EmperorId         int64  `gorm:"column:emperorId"`
	Name              string `gorm:"column:name"`
	Sex               int32  `gorm:"column:sex"`
	SpouseName        string `gorm:"column:spouseName"`
	RobNum            int64  `gorm:"column:robNum"`
	Storage           int64  `gorm:"column:storage"`
	RobTime           int64  `gorm:"column:robTime"`
	LastTime          int64  `gorm:"column:lastTime"`
	BoxNum            int64  `gorm:"column:boxNum"`
	BoxOutNum         int64  `gorm:"column:boxOutNum"`
	SpecialBoxLeftNum int32  `gorm:"column:specialBoxLeftNum"`
	BoxLastTime       int64  `gorm:"column:boxLastTime"`
	UpdateTime        int64  `gorm:"column:updateTime"`
	CreateTime        int64  `gorm:"column:createTime"`
	DeleteTime        int64  `gorm:"column:deleteTime"`
}

func (ee *EmperorEntity) GetId() int64 {
	return ee.Id
}

func (ee *EmperorEntity) TableName() string {
	return "t_emperor"
}
