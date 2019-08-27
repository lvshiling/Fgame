package types

type CrossEventType string

const (
	EventTypePlayerCrossEnter     CrossEventType = "CrossEnter"
	EventTypePlayerCrossExit                     = "CrossExit"
	EventTypePlayerCrossHeartbeat                = "CrossHeartbeat"
)
