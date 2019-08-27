package player

import (
	"fgame/fgame/core/fsm"
	"fmt"
)

const (
	EventPlayerAuth              fsm.Event = "auth"
	EventPlayerLoadingRoleList             = "loadingRoleList"
	EventPlayerWaitingSelectRole           = "waitingSelectRole"
	EventPlayerCreatingRole                = "creatingRole"
	EventPlayerSelectRole                  = "selectRole"
	EventPlayerLoading                     = "loading"
	EventPlayerLoaded                      = "loaded"
	EventPlayerEnterScene                  = "enterScene"
	EventPlayerGaming                      = "gaming"
	EventPlayerLeaveScene                  = "leaveScene"
	EventPlayerEnterCross                  = "enterCross"
	EventPlayerCrossing                    = "crossing"
	EventPlayerLeaveCross                  = "leaveCross"

	EventPlayerLogout = "logout"
)

const (
	//初始化 0
	PlayerStateInit fsm.State = iota
	//认证 1
	PlayerStateAuth
	//加载角色 2
	PlayerStateLoadingRoleList
	//等候选择角色 3
	PlayerStateWaitingSelectRole
	//创建角色 4
	PlayerStateCreatingRole
	//选择角色 5
	PlayerStateSelectRole
	//加载中 6
	PlayerStateLoading
	//加载完成 7
	PlayerStateLoaded
	//进入场景 8
	PlayerStateEnterScene
	//游戏中 9
	PlayerStateGaming
	//离开场景 10
	PlayerStateLeaveScene
	//进入跨服中11
	PlayerStateEnterCross
	//跨服中12
	PlayerStateCrossing
	//退出跨服13
	PlayerStateLeaveCross
	//登出中 14
	PlayerStateLogouting
	//登出 15
	PlayerStateLogouted
)

var (
	stateMachine *fsm.StateMachine
)

var (
	transitions = []*fsm.Trasition{
		//初始化->认证
		&fsm.Trasition{
			From:  PlayerStateInit,
			To:    PlayerStateAuth,
			Event: EventPlayerAuth,
		},
		//初始化->退出
		&fsm.Trasition{
			From:  PlayerStateInit,
			To:    PlayerStateLogouting,
			Event: EventPlayerLogout,
		},
		// 认证->加载角色列表
		&fsm.Trasition{
			From:  PlayerStateAuth,
			To:    PlayerStateLoadingRoleList,
			Event: EventPlayerLoadingRoleList,
		},
		// 认证->退出
		&fsm.Trasition{
			From:  PlayerStateAuth,
			To:    PlayerStateLogouting,
			Event: EventPlayerLogout,
		},
		// 加载角色列表->等候选择角色
		&fsm.Trasition{
			From:  PlayerStateLoadingRoleList,
			To:    PlayerStateWaitingSelectRole,
			Event: EventPlayerWaitingSelectRole,
		},
		// 加载角色列表->加载数据
		&fsm.Trasition{
			From:  PlayerStateLoadingRoleList,
			To:    PlayerStateSelectRole,
			Event: EventPlayerSelectRole,
		},
		// 加载角色列表->退出
		&fsm.Trasition{
			From:  PlayerStateLoadingRoleList,
			To:    PlayerStateLogouting,
			Event: EventPlayerLogout,
		},

		// 等候选择角色 -> 创建角色
		&fsm.Trasition{
			From:  PlayerStateWaitingSelectRole,
			To:    PlayerStateCreatingRole,
			Event: EventPlayerCreatingRole,
		},

		// 等候选择角色->退出
		&fsm.Trasition{
			From:  PlayerStateWaitingSelectRole,
			To:    PlayerStateLogouting,
			Event: EventPlayerLogout,
		},

		// 创建角色->加载数据
		&fsm.Trasition{
			From:  PlayerStateCreatingRole,
			To:    PlayerStateSelectRole,
			Event: EventPlayerSelectRole,
		},
		// 创建角色->退出
		&fsm.Trasition{
			From:  PlayerStateCreatingRole,
			To:    PlayerStateLogouting,
			Event: EventPlayerLogout,
		},
		// 选择角色->加载数据
		&fsm.Trasition{
			From:  PlayerStateSelectRole,
			To:    PlayerStateLoading,
			Event: EventPlayerLoading,
		},
		// 选择角色->退出
		&fsm.Trasition{
			From:  PlayerStateSelectRole,
			To:    PlayerStateLogouting,
			Event: EventPlayerLogout,
		},
		// 加载数据 ->加载完数据
		&fsm.Trasition{
			From:  PlayerStateLoading,
			To:    PlayerStateLoaded,
			Event: EventPlayerLoaded,
		},
		// 加载数据->退出
		&fsm.Trasition{
			From:  PlayerStateLoading,
			To:    PlayerStateLogouting,
			Event: EventPlayerLogout,
		},

		//  加载完数据->进入场景
		&fsm.Trasition{
			From:  PlayerStateLoaded,
			To:    PlayerStateEnterScene,
			Event: EventPlayerEnterScene,
		},
		//  加载完数据->进入跨服
		&fsm.Trasition{
			From:  PlayerStateLoaded,
			To:    PlayerStateEnterCross,
			Event: EventPlayerEnterCross,
		},
		//  加载完数据->退出
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
		//退出场景->进入跨服
		&fsm.Trasition{
			From:  PlayerStateLeaveScene,
			To:    PlayerStateEnterCross,
			Event: EventPlayerEnterCross,
		},
		//退出场景 -> 退出
		&fsm.Trasition{
			From:  PlayerStateLeaveScene,
			To:    PlayerStateLogouting,
			Event: EventPlayerLogout,
		},
		//跨服->跨服中
		&fsm.Trasition{
			From:  PlayerStateEnterCross,
			To:    PlayerStateCrossing,
			Event: EventPlayerCrossing,
		},
		//跨服->退出跨服
		&fsm.Trasition{
			From:  PlayerStateEnterCross,
			To:    PlayerStateLeaveCross,
			Event: EventPlayerLeaveCross,
		},
		//跨服->退出跨服
		&fsm.Trasition{
			From:  PlayerStateEnterCross,
			To:    PlayerStateLogouting,
			Event: EventPlayerLogout,
		},
		//跨服中 ->退出跨服
		&fsm.Trasition{
			From:  PlayerStateCrossing,
			To:    PlayerStateLeaveCross,
			Event: EventPlayerLeaveCross,
		},
		//退出跨服 ->进入场景
		&fsm.Trasition{
			From:  PlayerStateLeaveCross,
			To:    PlayerStateEnterScene,
			Event: EventPlayerEnterScene,
		},
		//退出跨服 ->进入跨服
		&fsm.Trasition{
			From:  PlayerStateLeaveCross,
			To:    PlayerStateEnterCross,
			Event: EventPlayerEnterCross,
		},
		//退出跨服 ->退出
		&fsm.Trasition{
			From:  PlayerStateLeaveCross,
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
