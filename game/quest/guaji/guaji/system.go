package guaji

import (
	"fgame/fgame/game/player"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

//系统进阶挂机
type QuestGuaJiSystemXHandler interface {
	GuaJiSystemX(p player.Player, questTemplate *gametemplate.QuestTemplate) bool
}

type QuestGuaJiSystemXHandlerFunc func(p player.Player, questTemplate *gametemplate.QuestTemplate) bool

func (f QuestGuaJiSystemXHandlerFunc) GuaJiSystemX(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {
	return f(p, questTemplate)
}

var (
	guaJiSystemXHandlerMap = map[questtypes.SystemReachXType]QuestGuaJiSystemXHandler{}
)

func RegisterQuestGuaJiSystemXHandler(systemReachXType questtypes.SystemReachXType, h QuestGuaJiSystemXHandler) {
	_, ok := guaJiSystemXHandlerMap[systemReachXType]
	if ok {
		panic(fmt.Errorf("additionsys:重复注册系统进阶任务挂机[%d]", systemReachXType))
	}
	guaJiSystemXHandlerMap[systemReachXType] = h
}

func GetQuestGuaJiSystemX(systemReachXType questtypes.SystemReachXType) QuestGuaJiSystemXHandler {
	h, ok := guaJiSystemXHandlerMap[systemReachXType]
	if !ok {
		return nil
	}
	return h
}
