package robot

import (
	"fgame/fgame/core/fsm"
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/game/scene/scene"
)

const (
	RobotPlayerStateIdle fsm.State = iota
	RobotPlayerStateTrace
	RobotPlayerStateAttack
	RobotPlayerStateAttacked
	RobotPlayerStateRun
	RobotPlayerStateDead
)

const (
	EventRobotPlayerIdle     fsm.Event = "idle"
	EventRobotPlayerTrace              = "trace"
	EventRobotPlayerAttack             = "attack"
	EventRobotPlayerAttacked           = "attacked"
	EventRobotPlayerRun                = "run"
	EventRobotPlayerDead               = "dead"
)

var (
	robotStateMachine *fsm.StateMachine
)

var (
	transitions = []*fsm.Trasition{
		//初始化->追击
		&fsm.Trasition{
			From:  RobotPlayerStateIdle,
			To:    RobotPlayerStateTrace,
			Event: EventRobotPlayerTrace,
		},
		//初始化->攻击
		&fsm.Trasition{
			From:  RobotPlayerStateIdle,
			To:    RobotPlayerStateAttack,
			Event: EventRobotPlayerAttack,
		},
		//初始化->被攻击
		&fsm.Trasition{
			From:  RobotPlayerStateIdle,
			To:    RobotPlayerStateAttacked,
			Event: EventRobotPlayerAttacked,
		},
		//初始化->逃跑
		&fsm.Trasition{
			From:  RobotPlayerStateIdle,
			To:    RobotPlayerStateRun,
			Event: EventRobotPlayerRun,
		},
		//初始化->死亡
		&fsm.Trasition{
			From:  RobotPlayerStateIdle,
			To:    RobotPlayerStateDead,
			Event: EventRobotPlayerDead,
		},

		// 追击->初始化
		&fsm.Trasition{
			From:  RobotPlayerStateTrace,
			To:    RobotPlayerStateIdle,
			Event: EventRobotPlayerIdle,
		},
		// 追击->攻击
		&fsm.Trasition{
			From:  RobotPlayerStateTrace,
			To:    RobotPlayerStateAttack,
			Event: EventRobotPlayerAttack,
		},
		// 追击->被攻击
		&fsm.Trasition{
			From:  RobotPlayerStateTrace,
			To:    RobotPlayerStateAttacked,
			Event: EventRobotPlayerAttacked,
		},
		//追击->逃跑
		&fsm.Trasition{
			From:  RobotPlayerStateTrace,
			To:    RobotPlayerStateRun,
			Event: EventRobotPlayerRun,
		},
		//追击->死亡
		&fsm.Trasition{
			From:  RobotPlayerStateTrace,
			To:    RobotPlayerStateDead,
			Event: EventRobotPlayerDead,
		},
		// 攻击->初始化
		&fsm.Trasition{
			From:  RobotPlayerStateAttack,
			To:    RobotPlayerStateIdle,
			Event: EventRobotPlayerIdle,
		},
		// 攻击->追击
		&fsm.Trasition{
			From:  RobotPlayerStateAttack,
			To:    RobotPlayerStateTrace,
			Event: EventRobotPlayerTrace,
		},
		// 攻击->被攻击
		&fsm.Trasition{
			From:  RobotPlayerStateAttack,
			To:    RobotPlayerStateAttacked,
			Event: EventRobotPlayerAttacked,
		},
		//攻击->逃跑
		&fsm.Trasition{
			From:  RobotPlayerStateAttack,
			To:    RobotPlayerStateRun,
			Event: EventRobotPlayerRun,
		},
		//攻击->死亡
		&fsm.Trasition{
			From:  RobotPlayerStateAttack,
			To:    RobotPlayerStateDead,
			Event: EventRobotPlayerDead,
		},

		// 被攻击->初始化
		&fsm.Trasition{
			From:  RobotPlayerStateAttacked,
			To:    RobotPlayerStateIdle,
			Event: EventRobotPlayerIdle,
		},
		// 被攻击->追击
		&fsm.Trasition{
			From:  RobotPlayerStateAttacked,
			To:    RobotPlayerStateTrace,
			Event: EventRobotPlayerTrace,
		},
		// 被攻击->攻击
		&fsm.Trasition{
			From:  RobotPlayerStateAttacked,
			To:    RobotPlayerStateAttack,
			Event: EventRobotPlayerAttack,
		},
		//被攻击->逃跑
		&fsm.Trasition{
			From:  RobotPlayerStateAttacked,
			To:    RobotPlayerStateRun,
			Event: EventRobotPlayerRun,
		},
		//被攻击->死亡
		&fsm.Trasition{
			From:  RobotPlayerStateAttacked,
			To:    RobotPlayerStateDead,
			Event: EventRobotPlayerDead,
		},

		// 逃跑->初始化
		&fsm.Trasition{
			From:  RobotPlayerStateRun,
			To:    RobotPlayerStateIdle,
			Event: EventRobotPlayerIdle,
		},
		// 逃跑->追击
		&fsm.Trasition{
			From:  RobotPlayerStateRun,
			To:    RobotPlayerStateTrace,
			Event: EventRobotPlayerTrace,
		},
		// 逃跑->攻击
		&fsm.Trasition{
			From:  RobotPlayerStateRun,
			To:    RobotPlayerStateAttack,
			Event: EventRobotPlayerAttack,
		},
		//逃跑->被攻击
		&fsm.Trasition{
			From:  RobotPlayerStateRun,
			To:    RobotPlayerStateAttacked,
			Event: EventRobotPlayerAttacked,
		},
		//逃跑->死亡
		&fsm.Trasition{
			From:  RobotPlayerStateRun,
			To:    RobotPlayerStateDead,
			Event: EventRobotPlayerDead,
		},
		//死亡->初始化
		&fsm.Trasition{
			From:  RobotPlayerStateDead,
			To:    RobotPlayerStateIdle,
			Event: EventRobotPlayerIdle,
		},
	}
)
var (
	stateMachine *fsm.StateMachine
)

func init() {
	stateMachine = fsm.NewStateMachine(transitions)
}

//状态管理器
type RobotStateManager struct {
	p scene.RobotPlayer
	//状态
	state fsm.State
	//定时器
	hbRunner heartbeat.HeartbeatTaskRunner
	//当前行为
	action scene.RobotAction
	//行为
	actionMap map[fsm.State]scene.RobotAction
}

func NewRobotStateManager(p scene.RobotPlayer) *RobotStateManager {
	m := &RobotStateManager{}
	m.p = p
	m.state = RobotPlayerStateIdle
	m.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	m.actionMap = make(map[fsm.State]scene.RobotAction)
	m.action = m.getAction()
	m.hbRunner.AddTask(CreateRobotTask(m.p))
	return m
}

func (m *RobotStateManager) Heartbeat() {
	m.hbRunner.Heartbeat()
}

//当前状态
func (m *RobotStateManager) CurrentState() fsm.State {
	return m.state
}

func (m *RobotStateManager) OnEnter(state fsm.State) {
	m.p.PauseMove()
	m.state = state
	//设置行为
	//通过脚本类型和状态获取行为
	m.action = m.getAction()
	m.action.OnEnter()

}

func (m *RobotStateManager) getAction() scene.RobotAction {
	action, ok := m.actionMap[m.state]
	if !ok {
		action = GetAction(m.p.GetRobotType(), m.state)
		m.actionMap[m.state] = action
	}
	return action
}

func (m *RobotStateManager) OnExit(state fsm.State) {
	m.action.OnExit()
	//清除行为
	m.action = nil
}

func (m *RobotStateManager) RobotDead() bool {
	flag := stateMachine.Trigger(m, EventRobotPlayerDead)
	if !flag {
		return false
	}

	return true
}

func (m *RobotStateManager) Trace() bool {
	flag := stateMachine.Trigger(m, EventRobotPlayerTrace)
	if !flag {
		return false
	}
	return true
}
func (m *RobotStateManager) Attack() bool {
	flag := stateMachine.Trigger(m, EventRobotPlayerAttack)
	if !flag {
		return false
	}
	return true
}

func (m *RobotStateManager) Idle() bool {
	flag := stateMachine.Trigger(m, EventRobotPlayerIdle)
	if !flag {
		return false
	}
	return true
}

func (m *RobotStateManager) Run() bool {
	flag := stateMachine.Trigger(m, EventRobotPlayerRun)
	if !flag {
		return false
	}
	return true
}

//获取主人
func (m *RobotStateManager) GetCurrentAction() scene.RobotAction {
	return m.action
}

func (m *RobotStateManager) GetState() fsm.State {
	return m.state
}
