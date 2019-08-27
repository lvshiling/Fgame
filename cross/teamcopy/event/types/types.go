package types

type TeamCopyEventType string

const (
	//组队副本玩家进入场景
	EvnetTypeTeamCopySceneEnter TeamCopyEventType = "TeamCopySceneEnter"
	//组队副本场景完成
	EventTypeTeamCopySceneFinish TeamCopyEventType = "TeamCopySceneFinish"
	//组队副本场景创建完成
	EventTypeTeamCopySceneCreateFinish TeamCopyEventType = "TeamCopySceneCreateFinish"
	//组队副本场景伤害改变
	EventTypeTeamCopySceneDamageChanged TeamCopyEventType = "TeamCopySceneDamageChanged"
)

//玩家队伍变化
const (
	EventTypeTeamCopyMemberOnline  = "TeamCopyMemberOnline"
	EventTypeTeamCopyMemberOffline = "TeamCopyMemberOffline"
	EventTypeTeamCopyMemberExit    = "TeamCopyMemberExit"
	EventTypeTeamCopyMemberGiveUp  = "TeamCopyMemberGiveUp"
)
