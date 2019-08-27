package model

type ChatMsg struct {
	PlayerId         int64  `json:"playerId"`
	PlayerName       string `json:"playerName"`
	VipLevel         int32  `json:"vipLevel"`
	GameLevel        int32  `json:"gameLevel"`
	ChatType         int32  `json:"chatType"`
	ChatMethod       int32  `json:"chatMethod"`
	ChatMsg          string `json:"chatMsg"`
	ChatTime         int64  `json:"chatTime"`
	ToPlayerId       int64  `json:"toPlayerId"`
	ToPlayerName     string `json:"toPlayerName"`
	Ip               string `json:"ip"`
	CenterPlatformId int32  `json:"centerPlatformId"`
	CenterServerId   int32  `json:"centerServerId"`
	UserId           int64  `json:"userId"`
}
