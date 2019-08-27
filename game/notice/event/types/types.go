package types

type NoticeEventType string

const (
	EventTypeBroadcastNotice NoticeEventType = "BroadcastNotice" //GM跑马灯公告
	EventTypeBroadcastSystem                 = "BroadcastSystem" //GM系统公告
)
