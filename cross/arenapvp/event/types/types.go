package types

type ArenapvpEventType string

//竞技场
const (
	// //竞技场匹配完成
	// EventTypeArenapvpMatched ArenapvpEventType = "ArenapvpMatched"
	//机器人退出
	EventTypeArenapvpRobotTeamEnd = "ArenapvpRobotTeamEnd"
)

//海选竞技场场景
const (
	//竞技场场景完成
	EventTypeArenapvpElectionSceneFinish = "ArenapvpElectionSceneFinish"
	//竞技场场景关闭
	EventTypeArenapvpElectionSceneStop = "ArenapvpElectionSceneStop"
)

//对战竞技场场景
const (
	//竞技场场景完成
	EventTypeArenapvpBattleSceneFinish = "ArenapvpBattleSceneFinish"
	//竞技场场景关闭
	EventTypeArenapvpBattleSceneStop = "ArenapvpBattleSceneStop"
)
