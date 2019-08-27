package player

import "fgame/fgame/core/fsm"

const (
	PlayerStateInit fsm.State = iota
	PlayerStateAuth
	PlayerStateSelectJob
	PlayerStateGame
)

const (
	EventPlayerInit      fsm.Event = "init"
	EventPlayerAuth                = "auth"
	EventPlayerSelectJob           = "selectJob"
	EventPlayerGame                = "game"
)

var (
	robotStateMachine *fsm.StateMachine
)

var (
	transitions = []*fsm.Trasition{
		//初始化->认证
		&fsm.Trasition{
			From:  PlayerStateInit,
			To:    PlayerStateAuth,
			Event: EventPlayerAuth,
		},
		//认证->进入选择职业
		&fsm.Trasition{
			From:  PlayerStateAuth,
			To:    PlayerStateSelectJob,
			Event: EventPlayerSelectJob,
		},
		//进入选择职业->进入场景
		&fsm.Trasition{
			From:  PlayerStateSelectJob,
			To:    PlayerStateGame,
			Event: EventPlayerGame,
		},
	}
)

var (
	stateMachine *fsm.StateMachine
)

func init() {
	stateMachine = fsm.NewStateMachine(transitions)
}
