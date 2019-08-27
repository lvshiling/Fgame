package model

type CenterServer struct {
	Id                   int64  `gorm:"primary_key;gorm:"column:id"`
	ServerType           int    `gorm:"column:serverType"`
	ServerId             int    `gorm:"column:serverId"`
	Platform             int64  `gorm:"column:platform"`
	ServerName           string `gorm:"column:name"`
	StartTime            int64  `gorm:"column:startTime"`
	UpdateTime           int64  `gorm:"column:updateTime"`
	CreateTime           int64  `gorm:"column:createTime"`
	DeleteTime           int64  `gorm:"column:deleteTime"`
	ServerIp             string `gorm:"column:serverIp"`
	ServerPort           string `gorm:"column:serverPort"`
	ServerRemoteIp       string `gorm:"column:serverRemoteIp"`
	ServerRemotePort     string `gorm:"column:serverRemotePort"`
	ServerDbIp           string `gorm:"column:serverDBIp"`
	ServerDbPort         string `gorm:"column:serverDBPort"`
	ServerDBName         string `gorm:"column:serverDBName"`
	ServerDBUser         string `gorm:"column:serverDBUser"`
	ServerDBPassword     string `gorm:"column:serverDBPassword"`
	ServerTag            int    `gorm:"column:serverTag"`
	ServerStatus         int    `gorm:"column:serverStatus"`
	ParentServerId       int    `gorm:"column:parentServerId"`
	PreShow              int    `gorm:"column:preShow"`
	JiaoYiZhanQuServerId int32  `gorm:"column:jiaoYiZhanQuServerId"`
	PingTaiFuServerId    int32  `gorm:"column:pingTaiFuServerId"`
	ChengZhanServerId    int32  `gorm:"column:chengZhanServerId"`
}

func (m *CenterServer) TableName() string {
	return "t_server"
}
