package types

type CrossPlayerEventType string

const (
	EventTypeCrossPlayerAfterLoad             CrossPlayerEventType = "CrossPlayerAfterLoad"
	EventTypeCrossPlayerLogout                                     = "CrossPlayerLogout"
	EventTypeCrossPlayerBeforeLogout                               = "CrossPlayerBeforeLogout"
	EventTypeCrossPlayerExitSceneBeforeLogout                      = "CrossPlayerExitSceneBeforeLogout"
)
