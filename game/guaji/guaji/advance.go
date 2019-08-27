package guaji

import (
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	"fmt"
)

//挂机升阶获取
type GuaJiAdvanceGetHandler interface {
	GetAdvance(p player.Player, t guajitypes.GuaJiAdvanceType) int32
}

type GuaJiAdvanceGetHandlerFunc func(player.Player, guajitypes.GuaJiAdvanceType) int32

func (f GuaJiAdvanceGetHandlerFunc) GetAdvance(p player.Player, t guajitypes.GuaJiAdvanceType) int32 {
	return f(p, t)
}

var (
	guaJiAdvanceGetHandlerMap = map[guajitypes.GuaJiAdvanceType]GuaJiAdvanceGetHandler{}
)

func RegisterGuaJiAdvanceGetHandler(guaJiAdvanceType guajitypes.GuaJiAdvanceType, h GuaJiAdvanceGetHandler) {
	_, ok := guaJiAdvanceGetHandlerMap[guaJiAdvanceType]
	if ok {
		panic(fmt.Errorf("重复注册进阶获取[%s]", guaJiAdvanceType.String()))
	}
	guaJiAdvanceGetHandlerMap[guaJiAdvanceType] = h
}

func GetGuaJiAdvanceGetHandler(guaJiAdvanceType guajitypes.GuaJiAdvanceType) GuaJiAdvanceGetHandler {
	h, ok := guaJiAdvanceGetHandlerMap[guaJiAdvanceType]
	if !ok {
		return nil
	}
	return h
}

//挂机升阶检查
type GuaJiAdvanceCheckHandler interface {
	AdvanceCheck(p player.Player, t guajitypes.GuaJiAdvanceType, advanceId int32, autobuy bool)
}

type GuaJiAdvanceCheckHandlerFunc func(p player.Player, t guajitypes.GuaJiAdvanceType, advanceId int32, autobuy bool)

func (f GuaJiAdvanceCheckHandlerFunc) AdvanceCheck(p player.Player, t guajitypes.GuaJiAdvanceType, advanceId int32, autobuy bool) {
	f(p, t, advanceId, autobuy)
}

var (
	guaJiAdvanceCheckHandlerMap = map[guajitypes.GuaJiAdvanceType]GuaJiAdvanceCheckHandler{}
)

func RegisterGuaJiAdvanceCheckHandler(guaJiAdvanceType guajitypes.GuaJiAdvanceType, h GuaJiAdvanceCheckHandler) {
	_, ok := guaJiAdvanceCheckHandlerMap[guaJiAdvanceType]
	if ok {
		panic(fmt.Errorf("重复注册进阶检查[%s]", guaJiAdvanceType.String()))
	}
	guaJiAdvanceCheckHandlerMap[guaJiAdvanceType] = h
}

func GetGuaJiAdvanceCheckHandler(guaJiAdvanceType guajitypes.GuaJiAdvanceType) GuaJiAdvanceCheckHandler {
	h, ok := guaJiAdvanceCheckHandlerMap[guaJiAdvanceType]
	if !ok {
		return nil
	}
	return h
}

func GetGuaJiAdvanceCheckHandlerMap() map[guajitypes.GuaJiAdvanceType]GuaJiAdvanceCheckHandler {
	return guaJiAdvanceCheckHandlerMap
}
