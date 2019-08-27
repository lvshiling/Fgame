package model

type User struct {
	Id            int64  `gorm:"primary_key;column:id"`
	DeviceMac     string `gorm:"column:deviceMac"`
	Name          string `gorm:"column:name"`
	Password      string `gorm:"column:password"`
	IdCard        string `gorm:"column:idCard"`
	PhoneNum      string `gorm:"column:phoneNum"`
	RealName      string `gorm:"column:realName"`
	RealNameState int32  `gorm:"column:realNameState"`
	UpdateTime    int64  `gorm:"column:updateTime"`
	CreateTime    int64  `gorm:"column:createTime"`
	DeleteTime    int64  `gorm:"column:deleteTime"`
}

func (u *User) TableName() string {
	return "t_user"
}
