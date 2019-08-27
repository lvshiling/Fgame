package model

type SystemLogMsg struct {
	LogTime    int64 `json:"logTime"`
	Platform   int32 `json:"platform"`
	ServerType int32 `json:"serverType"`
	ServerId   int32 `json:"serverId"`
}

type PlayerLogMsg struct {
	LogTime    int64  `json:"logTime"`
	Platform   int32  `json:"platform"`
	ServerType int32  `json:"serverType"`
	SdkType    int32  `json:"sdkType"`
	DeviceType int32  `json:"deviceType"`
	SdkUserId  string `json:"sdkUserId"`
	ServerId   int32  `json:"serverId"`
	UserId     int64  `json:"userId"`
	PlayerId   int64  `json:"playerId"`
	Ip         string `json:"ip"`
	Name       string `json:"name"`
	Role       int32  `json:"role"`
	Sex        int32  `json:"sex"`
	Level      int32  `json:"level"`
	Vip        int32  `json:"vip"`
}

type AllianceLogMsg struct {
	LogTime    int64  `json:"logTime"`
	Platform   int32  `json:"platform"`
	ServerType int32  `json:"serverType"`
	ServerId   int32  `json:"serverId"`
	AllianceId int64  `json:"allianceId"`
	Name       string `json:"name"`
	Level      int32  `json:"level"`
	JianShe    int64  `json:"jianShe"`
}

type PlayerTradeLogMsg struct {
	LogTime    int64 `json:"logTime"`
	Platform   int32 `json:"platform"`
	ServerType int32 `json:"serverType"`
	ServerId   int32 `json:"serverId"`
	PlayerId   int64 `json:"playerId"`
}

type JieYiLogMsg struct {
	LogTime    int64  `json:"logTime"`
	Platform   int32  `json:"platform"`
	ServerType int32  `json:"serverType"`
	ServerId   int32  `json:"serverId"`
	JieYiId    int64  `json:"jieYiId"`
	Name       string `json:"name"`
}
