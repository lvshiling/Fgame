package types

type RobotType int32

const (
	RobotTypeTest RobotType = iota
	RobotTypeArena
	RobotTypeTeamCopy
	RobotTypeQuest
	RobotTypeModel
	RobotTypeArenapvp
)

var (
	robotTypeMap = map[RobotType]string{
		RobotTypeTest:     "测试机器人",
		RobotTypeArena:    "竞技机器人",
		RobotTypeTeamCopy: "组队副本",
		RobotTypeQuest:    "任务机器人",
		RobotTypeModel:    "站街机器人",
		RobotTypeArenapvp: "比武大會机器人",
	}
)

func (t RobotType) String() string {
	return robotTypeMap[t]
}
