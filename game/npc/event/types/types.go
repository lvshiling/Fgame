package types

type NPCEventType string

const (
	EventTypeNPCCampChanged NPCEventType = "NPCCampChanged"
	EventTypeNPCAutoRecover              = "NPCAutoRecover"
	EventTypeNPCHPChanged                = "NPCHPChanged"
)
