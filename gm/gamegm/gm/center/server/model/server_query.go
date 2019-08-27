package model

type CenterQueryServer struct {
	Id                   int64  `gorm:"primary_key;gorm:"column:id"`
	ServerType           int    `gorm:"column:serverType"`
	ServerId             int    `gorm:"column:serverId"`
	Platform             int64  `gorm:"column:platform"`
	ServerName           string `gorm:"column:name"`
	ParentServerId       int    `gorm:"column:parentServerId"`
	PreShow              int    `gorm:"column:preShow"`
	JiaoYiZhanQuServerId int32  `gorm:"column:jiaoYiZhanQuServerId"`
	PingTaiFuServerId    int32  `gorm:"column:pingTaiFuServerId"`
	FinnalServerId       int32  `gorm:"column:finnalServerId"`
	ChengZhanServerId    int32  `gorm:"column:chengZhanServerId"`
}
