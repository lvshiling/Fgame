package model

type DBConfigInfo struct {
	PlatformId int64  `json:"platformId"`
	ServerId   int    `json:"serverId"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	UserName   string `json:"userName"`
	PassWord   string `json:"password"`
	DBName     string `json:"dbName"`
}

func (info *DBConfigInfo) Equal(info2 *DBConfigInfo) bool {
	return info.Host == info2.Host && info.Port == info2.Port && info.DBName == info2.DBName
}
