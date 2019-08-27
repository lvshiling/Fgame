package types

type ChristmasEventType string

var (
	EventTypeChristmasRefreshCollect ChristmasEventType = "ChristmasRefreshCollect" //刷新采集物
	EventTypeChristmasStopCollect                       = "ChristmasStopCollect"    //结束采集物
)
