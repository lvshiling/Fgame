package types

type SupremeTitleEventType string

const (
	//至尊称号改变事件
	EventTypeSupremeTitleChanged SupremeTitleEventType = "SupremeTitleChanged"
	//至尊称号激活
	EventTypeSupremeTitleActivate = "TitleActivate"
)
