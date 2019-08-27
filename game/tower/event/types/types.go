package types

type TowerEventType string

var (
	EventTypeTowerCrossDay    TowerEventType = "TowerCrossDay"    //打宝塔跨天刷新
	EventTypeTowerStartDaBao                 = "TowerStartDaBao"  //开始打宝
	EventTypeTowerEndDaBao                   = "TowerEndDaBao"    //结束打宝
	EventTypeTowerDaBaoNotice                = "TowerDaBaoNotice" //打宝时间推送
)
