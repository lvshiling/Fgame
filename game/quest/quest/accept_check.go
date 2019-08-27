package quest

import (
	"fgame/fgame/game/player"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

func handleDefault(pl player.Player, questTemplate *gametemplate.QuestTemplate) error {
	return nil
}

type CheckHandler interface {
	CheckHandle(pl player.Player, questTemplate *gametemplate.QuestTemplate) error
}

type CheckHandlerFunc func(pl player.Player, questTemplate *gametemplate.QuestTemplate) error

func (hf CheckHandlerFunc) CheckHandle(pl player.Player, questTemplate *gametemplate.QuestTemplate) error {
	return hf(pl, questTemplate)
}

func RegisterCheck(questSubType questtypes.QuestSubType, h CheckHandler) {
	_, exist := checkHandlerMap[questSubType]
	if exist {
		panic(fmt.Sprintf("repeat register questSubType %d", questSubType))
	}
	checkHandlerMap[questSubType] = h
}

func CheckHandle(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	h, exist := checkHandlerMap[questTemplate.GetQuestSubType()]
	if !exist {
		return defaultCheckHandler(pl, questTemplate)
	}
	return h.CheckHandle(pl, questTemplate)
}

var (
	checkHandlerMap     = make(map[questtypes.QuestSubType]CheckHandler)
	defaultCheckHandler = CheckHandlerFunc(handleDefault)
)
