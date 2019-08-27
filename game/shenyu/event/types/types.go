package types

type ShenYuEventType string

const (
	EventTypeShenYuKeyChange ShenYuEventType = "ShenYuKeyChange" //玩家钥匙改变
	EventTypeShenYuFinish                    = "ShenYuFinish"    //神域之战完成
	EventTypeShenYuStop                      = "ShenYuStop"      //神域之战关闭
	EventTypeShenYuLuckRew                   = "ShenYuLuckRew"   //神域之战幸运奖
)
