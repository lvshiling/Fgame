package types

type ActivityJoinEventType string

const (
	EventTypeActivityJoin        ActivityJoinEventType = "ActivityJoin"
	EventTypeActivityNoticeStart                      = "ActivityNoticeStart" //活动开始提示
	EventTypeActivityNoticeEnd                        = "ActivityNoticeEnd"   //活动结束提示
)
