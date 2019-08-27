package model

type LastOnLineStatic struct {
	Id        *LastOnLineStaticGroupKey `json:"_id"`
	MaxPlayer int                       `json:"lastplayer"`
}

type LastOnLineStaticGroupKey struct {
	ServerId int `json:"serverid"`
	Platform int `json:"platform"`
}

type OnLinePlayerStatic struct {
	Count int `json:"count"`
}

type OnLinePlayerStaticSdk struct {
	Id    *OnLinePlayerStaticSdkKey `json:"_id"`
	Count int                       `json:"sdkcount"`
}

type OnLinePlayerStaticSdkKey struct {
	SdkType int `json:"platform"`
}

type OnLinePlayerStaticServer struct {
	Id           *OnLinePlayerStaticServerKey `json:"_id"`
	MaxLoginTime int64                        `json:"maxlogintime"`
}

type OnLinePlayerStaticServerKey struct {
	ServerId int   `json:"serverid"`
	Platform int   `json:"platform"`
	PlayerId int64 `json:"playerid"`
}

type PlayerGoldChange struct {
	Id          *PlayerGoldChangeKey `json:"_id"`
	PlayerCount int                  `json:"playercount"`
	ChangedNum  int64                `json:"changednum"`
}

type PlayerGoldChangeKey struct {
	PlatformId int `json:"platform"`
	ServerId   int `json:"serverid"`
	//变更原因编号
	Reason int `json:"reason"`
}
