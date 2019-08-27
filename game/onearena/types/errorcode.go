package types

type OneArenaRobCodeType int32

const (
	//进入抢夺
	OneArenaRobCodeTypeEnter OneArenaRobCodeType = 1 + iota
	//当前正在抢夺中
	OneArenaRobCodeTypeIsRobbing
)
