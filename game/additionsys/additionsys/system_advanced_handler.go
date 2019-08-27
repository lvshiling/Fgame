package additionsys

import (
	additionsystypes "fgame/fgame/game/additionsys/types"
	"fgame/fgame/game/player"
	"fmt"
)

const (
	defaultAdvancedNum = 0
)

type SystemAdvancedHandler interface {
	GetAdvancedNum(player.Player) (number int32)
}

type SystemAdvancedHandlerFunc func(player.Player) (number int32)

func (f SystemAdvancedHandlerFunc) GetAdvancedNum(pl player.Player) (number int32) {
	return f(pl)
}

var (
	systemAdvancedHandlerMap = make(map[additionsystypes.AdditionSysType]SystemAdvancedHandler)
)

func RegisterSystemAdvancedHandler(typ additionsystypes.AdditionSysType, h SystemAdvancedHandler) {
	_, ok := systemAdvancedHandlerMap[typ]
	if ok {
		panic(fmt.Errorf("system_advanced:repeat register %s", typ.String()))
	}
	systemAdvancedHandlerMap[typ] = h
}

func GetSystemAdvancedNum(pl player.Player, typ additionsystypes.AdditionSysType) int32 {
	h, ok := systemAdvancedHandlerMap[typ]
	if !ok {
		return defaultAdvancedNum
	}

	return h.GetAdvancedNum(pl)
}
