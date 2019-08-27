package battle

import (
	"fgame/fgame/core/fsm"
	"fgame/fgame/core/heartbeat"

	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"

	log "github.com/Sirupsen/logrus"
)

const (
	//空闲
	PlayerStateIdle fsm.State = iota
	//追踪
	PlayerStateTrace
	//攻击
	PlayerStateAttack
	//被攻击
	PlayerStateAttacked
	//逃跑
	PlayerStateRun
	//死亡
	PlayerStateDead
)

const (
	EventPlayerIdle     fsm.Event = "idle"
	EventPlayerTrace              = "trace"
	EventPlayerAttack             = "attack"
	EventPlayerAttacked           = "attacked"
	EventPlayerRun                = "run"
	EventPlayerDead               = "dead"
)

var (
	guaJiStateMachine *fsm.StateMachine
)

var (
	transitions = []*fsm.Trasition{
		//初始化->追击
		&fsm.Trasition{
			From:  PlayerStateIdle,
			To:    PlayerStateTrace,
			Event: EventPlayerTrace,
		},
		//初始化->攻击
		&fsm.Trasition{
			From:  PlayerStateIdle,
			To:    PlayerStateAttack,
			Event: EventPlayerAttack,
		},
		//初始化->被攻击
		&fsm.Trasition{
			From:  PlayerStateIdle,
			To:    PlayerStateAttacked,
			Event: EventPlayerAttacked,
		},
		//初始化->逃跑
		&fsm.Trasition{
			From:  PlayerStateIdle,
			To:    PlayerStateRun,
			Event: EventPlayerRun,
		},
		//初始化->死亡
		&fsm.Trasition{
			From:  PlayerStateIdle,
			To:    PlayerStateDead,
			Event: EventPlayerDead,
		},

		// 追击->初始化
		&fsm.Trasition{
			From:  PlayerStateTrace,
			To:    PlayerStateIdle,
			Event: EventPlayerIdle,
		},
		// 追击->攻击
		&fsm.Trasition{
			From:  PlayerStateTrace,
			To:    PlayerStateAttack,
			Event: EventPlayerAttack,
		},
		// 追击->被攻击
		&fsm.Trasition{
			From:  PlayerStateTrace,
			To:    PlayerStateAttacked,
			Event: EventPlayerAttacked,
		},
		//追击->逃跑
		&fsm.Trasition{
			From:  PlayerStateTrace,
			To:    PlayerStateRun,
			Event: EventPlayerRun,
		},
		//追击->死亡
		&fsm.Trasition{
			From:  PlayerStateTrace,
			To:    PlayerStateDead,
			Event: EventPlayerDead,
		},
		// 攻击->初始化
		&fsm.Trasition{
			From:  PlayerStateAttack,
			To:    PlayerStateIdle,
			Event: EventPlayerIdle,
		},
		// 攻击->追击
		&fsm.Trasition{
			From:  PlayerStateAttack,
			To:    PlayerStateTrace,
			Event: EventPlayerTrace,
		},
		// 攻击->被攻击
		&fsm.Trasition{
			From:  PlayerStateAttack,
			To:    PlayerStateAttacked,
			Event: EventPlayerAttacked,
		},
		//攻击->逃跑
		&fsm.Trasition{
			From:  PlayerStateAttack,
			To:    PlayerStateRun,
			Event: EventPlayerRun,
		},
		//攻击->死亡
		&fsm.Trasition{
			From:  PlayerStateAttack,
			To:    PlayerStateDead,
			Event: EventPlayerDead,
		},

		// 被攻击->初始化
		&fsm.Trasition{
			From:  PlayerStateAttacked,
			To:    PlayerStateIdle,
			Event: EventPlayerIdle,
		},
		// 被攻击->追击
		&fsm.Trasition{
			From:  PlayerStateAttacked,
			To:    PlayerStateTrace,
			Event: EventPlayerTrace,
		},
		// 被攻击->攻击
		&fsm.Trasition{
			From:  PlayerStateAttacked,
			To:    PlayerStateAttack,
			Event: EventPlayerAttack,
		},
		//被攻击->逃跑
		&fsm.Trasition{
			From:  PlayerStateAttacked,
			To:    PlayerStateRun,
			Event: EventPlayerRun,
		},
		//被攻击->死亡
		&fsm.Trasition{
			From:  PlayerStateAttacked,
			To:    PlayerStateDead,
			Event: EventPlayerDead,
		},

		// 逃跑->初始化
		&fsm.Trasition{
			From:  PlayerStateRun,
			To:    PlayerStateIdle,
			Event: EventPlayerIdle,
		},
		// 逃跑->追击
		&fsm.Trasition{
			From:  PlayerStateRun,
			To:    PlayerStateTrace,
			Event: EventPlayerTrace,
		},
		// 逃跑->攻击
		&fsm.Trasition{
			From:  PlayerStateRun,
			To:    PlayerStateAttack,
			Event: EventPlayerAttack,
		},
		//逃跑->被攻击
		&fsm.Trasition{
			From:  PlayerStateRun,
			To:    PlayerStateAttacked,
			Event: EventPlayerAttacked,
		},
		//逃跑->死亡
		&fsm.Trasition{
			From:  PlayerStateRun,
			To:    PlayerStateDead,
			Event: EventPlayerDead,
		},
		//死亡->初始化
		&fsm.Trasition{
			From:  PlayerStateDead,
			To:    PlayerStateIdle,
			Event: EventPlayerIdle,
		},
	}
)
var (
	stateMachine *fsm.StateMachine
)

func init() {
	stateMachine = fsm.NewStateMachine(transitions)
}

//玩家挂机管理器
type PlayerGuaJiManager struct {
	p scene.Player
	// guaJi bool
	//当前挂机类型
	currentGuaJiTypeList []scenetypes.GuaJiType
	//状态
	state fsm.State
	//定时器
	hbRunner      heartbeat.HeartbeatTaskRunner
	actionMapList []map[fsm.State]scene.GuaJiAction
	action        scene.GuaJiAction

	deadTimes int32
}

func CreatePlayerGuaJiManager(p scene.Player) *PlayerGuaJiManager {
	m := &PlayerGuaJiManager{}
	m.p = p
	m.state = PlayerStateIdle
	m.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	m.hbRunner.AddTask(CreateGuaJiTask(m.p))
	m.hbRunner.AddTask(CreateGuaJiGetItemTask(m.p))
	return m
}

func (m *PlayerGuaJiManager) IsGuaJi() bool {
	return len(m.currentGuaJiTypeList) != 0
}

func (m *PlayerGuaJiManager) StartGuaJi(guaJiType scenetypes.GuaJiType) bool {
	if len(m.currentGuaJiTypeList) != 0 {
		return false
	}
	m.EnterGuaJi(guaJiType)
	return true
}

func (m *PlayerGuaJiManager) StopGuaJi() {
	m.currentGuaJiTypeList = nil
}

func (m *PlayerGuaJiManager) EnterGuaJi(guaJiType scenetypes.GuaJiType) {
	log.Infof("玩家进入挂机[%s]", guaJiType.String())
	m.currentGuaJiTypeList = append(m.currentGuaJiTypeList, guaJiType)
	actionMap := make(map[fsm.State]scene.GuaJiAction)
	m.actionMapList = append(m.actionMapList, actionMap)
	m.state = PlayerStateIdle
	m.action = m.getAction()
	//清空死亡次数
	m.deadTimes = 0
	return
}
func (m *PlayerGuaJiManager) GetCurrentGuaJiType() (t scenetypes.GuaJiType) {
	if len(m.currentGuaJiTypeList) == 0 {
		return
	}
	t = m.currentGuaJiTypeList[len(m.currentGuaJiTypeList)-1]
	return
}

func (m *PlayerGuaJiManager) GetLastGuaJiType() (t scenetypes.GuaJiType, flag bool) {
	if len(m.currentGuaJiTypeList) < 2 {
		return
	}
	t = m.currentGuaJiTypeList[len(m.currentGuaJiTypeList)-2]
	flag = true
	return
}

func (m *PlayerGuaJiManager) ExitGuaJi() {
	if len(m.currentGuaJiTypeList) == 0 {
		return
	}
	log.Infof("玩家退出挂机[%s]", m.currentGuaJiTypeList[len(m.currentGuaJiTypeList)-1].String())
	m.currentGuaJiTypeList = m.currentGuaJiTypeList[0 : len(m.currentGuaJiTypeList)-1]
	if len(m.currentGuaJiTypeList) == 0 {
		return
	}
	m.actionMapList = m.actionMapList[:len(m.actionMapList)-1]
	m.state = PlayerStateIdle
	m.action = m.getAction()
}

func (m *PlayerGuaJiManager) Heartbeat() {
	if len(m.currentGuaJiTypeList) == 0 {
		return
	}
	m.hbRunner.Heartbeat()
}

func (m *PlayerGuaJiManager) onEnter(state fsm.State) {
	m.p.PauseMove()
	m.state = state
	if len(m.currentGuaJiTypeList) == 0 {
		return
	}
	//设置行为
	//通过脚本类型和状态获取行为
	m.action = m.getAction()
	m.action.OnEnter()
}

func (m *PlayerGuaJiManager) getAction() scene.GuaJiAction {
	actionMap := m.actionMapList[len(m.actionMapList)-1]
	action, ok := actionMap[m.state]
	if !ok {
		currentGuaJiType := m.currentGuaJiTypeList[len(m.currentGuaJiTypeList)-1]
		action = scene.GetGuaJiAction(currentGuaJiType, m.state)
		actionMap[m.state] = action
	}
	return action
}

func (m *PlayerGuaJiManager) onExit(state fsm.State) {
	if len(m.currentGuaJiTypeList) == 0 {
		return
	}

	m.action.OnExit()
	//清除行为
	m.action = nil

}

func (m *PlayerGuaJiManager) GuaJiDead() bool {
	// flag := stateMachine.Trigger(m, EventPlayerDead)
	// if !flag {
	// 	return false
	// }
	if m.state == PlayerStateDead {
		return false
	}
	m.onExit(m.state)
	m.onEnter(PlayerStateDead)
	m.deadTimes += 1
	return true
}

func (m *PlayerGuaJiManager) GuaJiTrace() bool {
	if m.state == PlayerStateTrace {
		return false
	}
	m.onExit(m.state)
	m.onEnter(PlayerStateTrace)
	return true
}

func (m *PlayerGuaJiManager) GuaJiAttack() bool {
	if m.state == PlayerStateAttack {
		return false
	}
	m.onExit(m.state)
	m.onEnter(PlayerStateAttack)
	return true
}

func (m *PlayerGuaJiManager) GuaJiIdle() bool {
	if m.state == PlayerStateIdle {
		return false
	}

	m.onExit(m.state)
	m.onEnter(PlayerStateIdle)

	return true
}

func (m *PlayerGuaJiManager) GuaJiRun() bool {
	if m.state == PlayerStateRun {
		return false
	}
	m.onExit(m.state)
	m.onEnter(PlayerStateRun)
	return true
}

func (m *PlayerGuaJiManager) GetGuaJiDeadTimes() int32 {
	return m.deadTimes
}

//获取主人
func (m *PlayerGuaJiManager) GetCurrentGuaJiAction() scene.GuaJiAction {
	return m.action
}

func (m *PlayerGuaJiManager) GetGuaJiState() fsm.State {
	return m.state
}
