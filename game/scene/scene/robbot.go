package scene

import (
	"fgame/fgame/core/fsm"
	robottypes "fgame/fgame/game/robot/types"
)

// 机器人接口
type RobotPlayer interface {
	Player
	RobotStateManagerInterface

	//机器人类型
	GetRobotType() robottypes.RobotType
	//获取复活次数
	GetCanReliveTime() int32
	//获取起始任务id
	GetQuestBeginId() int32
	//获取结束任务id
	GetQuestEndId() int32
}

type RobotStateManagerInterface interface {
	Idle() bool
	Trace() bool
	Attack() bool
	Run() bool
	RobotDead() bool
	GetCurrentAction() RobotAction
	GetState() fsm.State
}

//机器人行为
type RobotAction interface {
	OnEnter()
	Action(p RobotPlayer)
	OnExit()
}
