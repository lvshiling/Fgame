package systemcompensate

import (
	"fgame/fgame/game/player"
	systemcompensatetypes "fgame/fgame/game/systemcompensate/types"
	"fmt"
)

const (
	defaultAdvancedNum = 0
)

type SystemCompensateHandler interface {
	GetAdvancedNum(player.Player) (number int32)
}

type SystemCompensateHandlerFunc func(player.Player) (number int32)

func (f SystemCompensateHandlerFunc) GetAdvancedNum(pl player.Player) (number int32) {
	return f(pl)
}

var (
	systemCompensateHandlerMap = make(map[systemcompensatetypes.SystemCompensateType]SystemCompensateHandler)
)

func RegisterSystemCompensateHandler(tag systemcompensatetypes.SystemCompensateType, h SystemCompensateHandler) {
	_, ok := systemCompensateHandlerMap[tag]
	if ok {
		panic(fmt.Errorf("systemskill:repeat register %s", tag))
	}
	systemCompensateHandlerMap[tag] = h
}

func GetSystemAdvancedNum(pl player.Player, tag systemcompensatetypes.SystemCompensateType) int32 {
	h, ok := systemCompensateHandlerMap[tag]
	if !ok {
		return defaultAdvancedNum
	}

	return h.GetAdvancedNum(pl)
}
