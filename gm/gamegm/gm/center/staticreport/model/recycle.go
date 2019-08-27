package model

type RecycleStatic struct {
	Id               *RecycleStaticGroupKey `json:"_id"`
	TotalRecycleGold int                    `json:"recycleGold"`
}

type RecycleStaticGroupKey struct {
	Date        int `json:"datestr"`
	MinuteIndex int `json:"minuteindex"`
	ServerId    int `json:"serverid"`
	Platform    int `json:"platform"`
}
