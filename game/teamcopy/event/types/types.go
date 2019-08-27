package types

type TeamCopyEventType string

const (
	//参加组队副本
	EventTypeTeamCopyAttend TeamCopyEventType = "TeamCopyAttend"
	//通关组队副本
	EventTypeTeamCopyFinishSucess = "TeamCopyFinishSucess"
)
