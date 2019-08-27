package onlinestat

type ServerOnLineStatic struct {
	Id          int64 `gorm:"primary_key;column:id"`
	PlatformId  int   `gorm:"column:platformId"`
	ServerId    int   `gorm:"column:serverId"`
	PlayerId    int64 `gorm:"column:playerId"`
	OnLineIndex int   `gorm:"column:onLineIndex"`
	OnLineTime  int64 `gorm:"column:onLineTime"`
	OnLineDate  int64 `gorm:"column:onLineDate"`
	UpdateTime  int64 `gorm:"column:updateTime"`
	CreateTime  int64 `gorm:"column:createTime"`
	DeleteTime  int64 `gorm:"column:deleteTime"`
}

func (m *ServerOnLineStatic) TableName() string {
	return "t_server_online"
}
