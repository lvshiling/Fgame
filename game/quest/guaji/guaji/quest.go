package guaji

import (
	"fgame/fgame/game/player"

	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

type QuestGuaJi interface {
	DoQuest(p player.Player, questTemplate *gametemplate.QuestTemplate) bool
}

type QuestGuaJiFunc func(p player.Player, questTemplate *gametemplate.QuestTemplate) bool

func (f QuestGuaJiFunc) DoQuest(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {
	return f(p, questTemplate)
}

var (
	questGuaJiMap = map[questtypes.QuestSubType]QuestGuaJi{}
)

func GetQuestGuaJi(subType questtypes.QuestSubType) QuestGuaJi {
	guaJi, ok := questGuaJiMap[subType]
	if !ok {
		return nil
	}
	return guaJi
}

func RegisterQuestGuaJi(subType questtypes.QuestSubType, guaJi QuestGuaJi) {
	_, ok := questGuaJiMap[subType]
	if ok {
		panic(fmt.Errorf("重复注册%s任务挂机", subType.String()))
	}
	questGuaJiMap[subType] = guaJi
}
