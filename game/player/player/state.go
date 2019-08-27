package player

import (
	"fgame/fgame/common/codec"
	"fgame/fgame/common/message"
	"fgame/fgame/core/fsm"
	"fgame/fgame/game/player"
)

func (pl *Player) OnEnter(state fsm.State) {
	pl.SubjectBase.OnEnter(state)
}

func (pl *Player) OnExit(state fsm.State) {
	pl.SubjectBase.OnExit(state)
}

func (pl *Player) CanProcess(msg message.Message) bool {
	tmsg, ok := msg.(message.SessionMessage)
	if !ok {
		return false
	}
	ttmsg := tmsg.Message()
	if ttmsg == nil {
		return false
	}
	tttmsg, ok := ttmsg.(*codec.Message)
	if !ok {
		return false
	}
	msgType := tttmsg.MessageType
	switch pl.CurrentState() {
	case player.PlayerStateInit:
		//不处理消息
		break
	case player.PlayerStateAuth:
		//不处理消息
		break
	case player.PlayerStateLoadingRoleList:
		//不处理消息
		break
	case player.PlayerStateWaitingSelectRole:
		//只处理创建角色
		if codec.IsCreateRoleMsg(msgType) {
			return true
		}
		break
	case player.PlayerStateCreatingRole:
		break
	case player.PlayerStateSelectRole:
		break
	case player.PlayerStateLoading:
		break
	case player.PlayerStateLoaded:
		break
	case player.PlayerStateEnterScene:
		break
	case player.PlayerStateGaming:
		return true
	case player.PlayerStateLeaveScene:
		break
	case player.PlayerStateEnterCross:
		break
	case player.PlayerStateCrossing:
		return true
	case player.PlayerStateLeaveCross:
		break
	case player.PlayerStateLogouting:
		break
	case player.PlayerStateLogouted:
		break
	}

	return false
}
