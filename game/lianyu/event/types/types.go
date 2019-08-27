package types

type LianYuEventType string

const (
	//无间炼狱玩家取消排队
	EventTypeLianYuPlayerCancleLineUp LianYuEventType = "LianYuPlayerCancleLineUp"
)

//无间炼狱活动 事件
type LianYuSceneEventType string

const (
	EventTypeLianYuBossStatusRefresh  LianYuSceneEventType = "LianYuBossStatusRefresh"  //Boss状态刷新
	EventTypeLianYuPlayerEnter                             = "LianYuPlayerEnter"        //玩家进入无间炼狱场景
	EventTypeLianYuSceneFinish                             = "LianYuSceneFinish"        //无间炼狱场景完成
	EventTypeLianYuCancleLineUp                            = "LianYuCancleLineUp"       //玩家取消排队
	EventTypeLianYuPlayerExit                              = "LianYuPlayerExit"         //玩家退出无间炼狱场景
	EventTypeLianYuPlayerLineUpFinish                      = "LianYuPlayerLineUpFinish" //玩家排队完成
	EventTypeLianYuShaQiRankChanged                        = "LianYuShaQiRankChanged"   //杀气排行榜刷新
	EventTypeLianYuPlayerShaQiChanged                      = "LianYuPlayerShaQiChanged" //玩家杀气变化
)
