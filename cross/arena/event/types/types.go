package types

type ArenaEventType string

//竞技场
const (
	EventTypeArenaTeamCreate ArenaEventType = "ArenaTeamCreate"
	//竞技场匹配到
	EventTypeArenaMatched ArenaEventType = "ArenaMatched"
	//竞技场开始下一场匹配
	EventTypeArenaNextMatch = "ArenaNextMatch"
	//竞技场
	EventTypeArenaFourGod = "ArenaFourGod"
	//机器人退出
	EventTypeArenaRobotTeamEnd = "ArenaRobotTeamEnd"
)

//竞技场场景
const (
	//竞技场场景 开始
	EventTypeArenaSceneStart = "ArenaSceneStart"
	//玩家进入竞技场
	EventTypeArenaScenePlayerEnter = "ArenaScenePlayerEnter"
	//玩家退出竞技
	EventTypeArenaScenePlayerExit = "ArenaScenePlayerExit"
	//竞技场场景 结束
	EventTypeArenaSceneEnd = "ArenaSceneEnd"
)

//四神
const (
	//四神排队中
	EventTypeArenaFourGodSceneTeamQueue = "ArenaFourGodSceneTeamQueue"
	//四神排队变化
	EventTypeArenaFourGodSceneTeamQueueChanged = "ArenaFourGodSceneTeamQueueChanged"
	//四神排队取消
	EventTypeArenaFourGodSceneTeamQueueCancel = "ArenaFourGodSceneTeamQueueCancel"
	//进入四神
	EventTypeArenaFourGodSceneTeamJoin = "ArenaFourGodSceneTeamJoin"
	//离开四神
	EventTypeArenaFourGodSceneTeamLeave = "ArenaFourGodSceneTeamLeave"
	//四神场景结束
	EventTypeArenaFourGodSceneFinish = "ArenaFourGodSceneFinish"
	//玩家进入
	EventTypeArenaFourGodScenePlayerEnter = "ArenaFourGodScenePlayerEnter"
	//玩家退出
	EventTypeArenaFourGodScenePlayerExit = "ArenaFourGodScenePlayerExit"
	//经验树采集中
	EventTypeArenaFourGodSceneCollecting = "ArenaFourGodSceneCollecting"
	//经验树采集完成
	EventTypeArenaFourGodSceneCollect = "ArenaFourGodSceneCollect"
	//经验树采集中断
	EventTypeArenaFourGodSceneCollectStop = "ArenaFourGodSceneCollectStop"
)

//玩家队伍变化
const (
	EventTypeArenaTeamMemberOnline  = "ArenaTeamMemberOnline"
	EventTypeArenaTeamMemberOffline = "ArenaTeamMemberOffline"
	EventTypeArenaTeamMemberExit    = "ArenaTeamMemberExit"
	EventTypeArenaTeamMemberGiveUp  = "ArenaTeamMemberGiveUp"
)
