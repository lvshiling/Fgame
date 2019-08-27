package model

//当日最大的在线人数
type ServerPlayerLoginStatic struct {
	Id          *ServerPlayerLoginStaticKey `json:"_id"`
	TotalPlayer int                         `json:"totalplayer"`
}

type ServerPlayerLoginStaticKey struct {
	ServerId int `json:"serverid"`
	Platform int `json:"platform"`
}
