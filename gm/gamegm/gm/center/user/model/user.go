package model

type UserInfo struct {
	Id             int    `gorm:"primary_key;gorm:"column:id"`
	Platform       int    `gorm:"column:platform"`
	PlatformUserId string `gorm:"column:platformUserId"`
	Name           string `gorm:"column:name"`
	Password       string `gorm:"column:password"`
	PhoneNum       string `gorm:"column:phoneNum"`
	IdCard         string `gorm:"column:idCard"`
	RealName       string `gorm:"column:realName"`
	RealNameState  int    `gorm:"column:realNameState"`
	UpdateTime     int64  `gorm:"column:updateTime"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
	Gm             int    `gorm:"column:gm"`
	Forbid         int    `gorm:"column:forbid"`
	ForbidTime     int64  `gorm:"column:forbidTime"`
	ForbidEndTime  int64  `gorm:"column:forbidEndTime"`
	ForbidName     string `gorm:"column:forbidName"`
	ForbidText     string `gorm:"column:forbidText"`
}

func (u *UserInfo) TableName() string {
	return "t_user"
}
