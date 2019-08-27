package model

type ServerSupportPool struct {
	Id               int64 `gorm:"primary_key;column:id"`
	ServerId         int   `gorm:"column:serverId"`
	BeginGold        int   `gorm:"column:beginGold"`
	CurGold          int   `gorm:"column:curGold"`
	DelGold          int   `gorm:"column:delGold"`
	UpdateTime       int64 `gorm:"column:updateTime"`
	CreateTime       int64 `gorm:"column:createTime"`
	DeleteTime       int64 `gorm:"column:deleteTime"`
	SdkType          int   `gorm:"column:sdkType"`
	CenterPlatformId int64 `gorm:"column:centerPlatformId"`
	OrderGoldPer     int32 `gorm:"column:orderGoldPer"`
	OrderGold        int32 `gorm:"column:orderGold"`
	BeginOrderTime   int64 `gorm:"column:beginOrderTime"`
	CurOrderTime     int64 `gorm:"column:curOrderTime"`
}

func (m *ServerSupportPool) TableName() string {
	return "t_server_support_pool"
}
