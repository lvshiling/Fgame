package types

type ShenMoEventType string

const (
	EventTypeShenMoGongXunNumChanged ShenMoEventType = "ShenMoGongXunNumChanged" //神魔战场功勋值改变事件
)

const (
	//玩家进入神魔战场场景
	EventTypeShenMoPlayerEnter ShenMoEventType = "ShenMoPlayerEnter"
	//神魔战场场景完成
	EventTypeShenMoSceneFinish ShenMoEventType = "ShenMoSceneFinish"
	//玩家取消排队
	EventTypeShenMoCancleLineUp ShenMoEventType = "ShenMoCancleLineUp"
	//玩家退出神魔战场场景
	EventTypeShenMoPlayerExit ShenMoEventType = "ShenMoPlayerExit"
	//玩家排队完成
	EventTypeShenMoPlayerLineUpFinish ShenMoEventType = "ShenMoPlayerLineUpFinish"
)
