package guaji

import (
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	"fmt"
)

//挂机提升检查
type GuaJiCheckHandler interface {
	Check(p player.Player)
}

type GuaJiCheckHandlerFunc func(p player.Player)

func (f GuaJiCheckHandlerFunc) Check(p player.Player) {
	f(p)
}

var (
	guaJiCheckHandlerMap = map[guajitypes.GuaJiCheckType]GuaJiCheckHandler{}
)

func RegisterGuaJiCheckHandler(guaJiCheckType guajitypes.GuaJiCheckType, h GuaJiCheckHandler) {
	_, ok := guaJiCheckHandlerMap[guaJiCheckType]
	if ok {
		panic(fmt.Errorf("重复注册挂机提升检查[%s]", guaJiCheckType.String()))
	}
	guaJiCheckHandlerMap[guaJiCheckType] = h
}

func GetGuaJiCheckHandlerMap() map[guajitypes.GuaJiCheckType]GuaJiCheckHandler {
	return guaJiCheckHandlerMap
}

func GetGuaJiCheckHandler(guaJiType guajitypes.GuaJiCheckType) GuaJiCheckHandler {
	h, ok := guaJiCheckHandlerMap[guaJiType]
	if !ok {
		return nil
	}
	return h
}
