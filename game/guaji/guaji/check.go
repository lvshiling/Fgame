package guaji

import (
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	"fmt"
)

//挂机进入检查
type GuaJiEnterCheckHandler interface {
	GuaJiEnterCheck(p player.Player) bool
}
type GuaJiEnterCheckHandlerFunc func(p player.Player) bool

func (f GuaJiEnterCheckHandlerFunc) GuaJiEnterCheck(p player.Player) bool {
	return f(p)
}

var (
	guaJiEnterCheckHandlerMap = map[guajitypes.GuaJiType]GuaJiEnterCheckHandler{}
)

func RegisterGuaJiEnterCheckHandler(guaJiType guajitypes.GuaJiType, h GuaJiEnterCheckHandler) {
	_, ok := guaJiEnterCheckHandlerMap[guaJiType]
	if ok {
		panic(fmt.Errorf("重复注册挂机进入检查[%s]", guaJiType.String()))
	}
	guaJiEnterCheckHandlerMap[guaJiType] = h
}

func GetGuaJiEnterCheckHandler(guaJiType guajitypes.GuaJiType) GuaJiEnterCheckHandler {
	h, ok := guaJiEnterCheckHandlerMap[guaJiType]
	if !ok {
		return nil
	}
	return h
}

//挂机退出检查
type GuaJiExitCheckHandler interface {
	GuaJiExitCheck(p player.Player) bool
}
type GuaJiExitCheckHandlerFunc func(p player.Player) bool

func (f GuaJiExitCheckHandlerFunc) GuaJiExitCheck(p player.Player) bool {
	return f(p)
}
