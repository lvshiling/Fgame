package model

type IpForbidInfo struct {
	Id            int    `gorm:"primary_key;gorm:"column:id"`
	Ip            string `gorm:"column:ip"`
	UpdateTime    int64  `gorm:"column:updateTime"`
	CreateTime    int64  `gorm:"column:createTime"`
	DeleteTime    int64  `gorm:"column:deleteTime"`
	Forbid        int    `gorm:"column:forbid"`
	ForbidTime    int64  `gorm:"column:forbidTime"`
	ForbidEndTime int64  `gorm:"column:forbidEndTime"`
	ForbidName    string `gorm:"column:forbidName"`
	ForbidText    string `gorm:"column:forbidText"`
}

func (u *IpForbidInfo) TableName() string {
	return "t_ipforbid"
}
