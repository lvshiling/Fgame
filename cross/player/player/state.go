package player

import (
	"fgame/fgame/core/fsm"
	"fmt"
)

const (
	//认证
	EventPlayerAuth fsm.Event = "auth"
	//加载
	EventPlayerLoaded = "loaded"
	//进入场景
	EventPlayerEnterScene = "enterScene"
	//场景
	EventPlayerGaming = "gaming"
	//退出场景
	EventPlayerLeaveScene = "leaveScene"
	//退出
	EventPlayerLogout = "logout"
)

const (
	//认证 0
	PlayerStateAuth = iota
	//加载完成 1
	PlayerStateLoaded
	//进入场景 2
	PlayerStateEnterScene
	//游戏中 3
	PlayerStateGaming
	//离开场景 4
	PlayerStateLeaveScene
	//登出中 5
	PlayerStateLogouting
	//登出 6
	PlayerStateLogouted
)

var (
	stateMachine *fsm.StateMachine
)

var (
	transitions = []*fsm.Trasition{
		// 认证->加载数据
		&fsm.Trasition{
			From:  PlayerStateAuth,
			To:    PlayerStateLoaded,
			Event: EventPlayerLoaded,
		},
		// 认证->登出
		&fsm.Trasition{
			From:  PlayerStateAuth,
			To:    PlayerStateLogouting,
			Event: EventPlayerLogout,
		},
		//  加载完数据->进入场景
		&fsm.Trasition{
			From:  PlayerStateLoaded,
			To:    PlayerStateEnterScene,
			Event: EventPlayerEnterScene,
		},
		// 加载数据->退出
		&fsm.Trasition{
			From:  PlayerStateLoaded,
			To:    PlayerStateLogouting,
			Event: EventPlayerLogout,
		},

		//进入场景 ->场景中
		&fsm.Trasition{
			From:  PlayerStateEnterScene,
			To:    PlayerStateGaming,
			Event: EventPlayerGaming,
		},
		//进入场景 -> 退出
		&fsm.Trasition{
			From:  PlayerStateEnterScene,
			To:    PlayerStateLogouting,
			Event: EventPlayerLogout,
		},
		//场景中->退出场景
		&fsm.Trasition{
			From:  PlayerStateGaming,
			To:    PlayerStateLeaveScene,
			Event: EventPlayerLeaveScene,
		},
		//退出场景->进入场景
		&fsm.Trasition{
			From:  PlayerStateLeaveScene,
			To:    PlayerStateEnterScene,
			Event: EventPlayerEnterScene,
		},
		//退出场景 -> 退出
		&fsm.Trasition{
			From:  PlayerStateLeaveScene,
			To:    PlayerStateLogouting,
			Event: EventPlayerLogout,
		},
	}
)

func init() {
	stateMachine = fsm.NewStateMachine(transitions)
}

func GetPlayerStateMachine() *fsm.StateMachine {
	return stateMachine
}

type stateError struct {
	state fsm.State
	event fsm.Event
}

func (se *stateError) Error() string {
	return fmt.Sprintf("状态[%d]触发事件[%s],失败", se.state, se.event)
}

func NewStateError(state fsm.State, event fsm.Event) error {
	return &stateError{
		state: state,
		event: event,
	}
}
