package model

type OnLineStatic struct {
	Id        *OnLineStaticGroupKey `json:"_id"`
	MaxPlayer int                   `json:"maxplayer"`
}

type OnLineStaticGroupKey struct {
	Date        int `json:"datestr"`
	MinuteIndex int `json:"minuteindex"`
	ServerId    int `json:"serverid"`
	Platform    int `json:"platform"`
}

//当日最大的在线人数
type OnLinePlayerStaticDaily struct {
	Id        *OnLinePlayerStaticDailyKey `json:"_id"`
	MaxPlayer int                         `json:"maxplayer"`
}

type OnLinePlayerStaticDailyKey struct {
	ServerId int `json:"serverid"`
	Platform int `json:"platform"`
}
