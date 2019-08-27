package player

import (
	gameevent "fgame/fgame/game/event"
	onearenaeventtypes "fgame/fgame/game/onearena/event/types"
	onearenatypes "fgame/fgame/game/onearena/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"time"
)

const (
	oneArenaTaskTime = 1 * time.Minute
)

type OneArenaTask struct {
	p player.Player
}

func (wt *OneArenaTask) Run() {
	manager := wt.p.GetPlayerDataManager(types.PlayerOneArenaDataManagerType).(*PlayerOneArenaDataManager)
	oneArenaObj := manager.GetOneArena()
	if oneArenaObj == nil {
		return
	}
	if oneArenaObj.Level < onearenatypes.OneArenaLevelTypeHuangJi {
		return
	}
	gameevent.Emit(onearenaeventtypes.EventTypeOneArenaOccupyTime, wt.p, oneArenaObj)
}

func (wt *OneArenaTask) ElapseTime() time.Duration {
	return oneArenaTaskTime
}

func CreateOneArenaTask(p player.Player) *OneArenaTask {
	qTask := &OneArenaTask{
		p: p,
	}
	return qTask
}
