package types

type TuLongEventType string

const (
	//大龙蛋状态刷新
	EventTypeTuLongBigEggStatusRefresh TuLongEventType = "TuLongBigEggStatusRefresh"
	//玩家进入屠龙场景
	EventTypeTuLongPlayerEnter TuLongEventType = "TuLongPlayerEnter"
	//采集龙蛋完成
	EventTypeTuLongCollectFinish TuLongEventType = "TuLongCollectFinish"
	//屠龙场景完成
	EventTypeTuLongSceneFinish TuLongEventType = "TuLongSceneFinish"
)
