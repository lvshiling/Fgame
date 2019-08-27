package types

type YuXiReborType int32

const (
	YuXiReborTypeInitNone        YuXiReborType = iota //默认
	YuXiReborTypePlayerDead                           //玩家死亡
	YuXiReborTypePlayerEnterSafe                      //进入安全区
	YuXiReborTypePlayerExitScene                      //玩家退出场景
)
