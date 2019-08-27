package quest

import (
	"fgame/fgame/game/player"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

func handleCommitFlowDefault(pl player.Player, questTemplate *gametemplate.QuestTemplate) error {
	return nil
}

type CommitFlowHandler interface {
	CommitFlowHandle(pl player.Player, questTemplate *gametemplate.QuestTemplate) error
}

type CommitFlowHandlerFunc func(pl player.Player, questTemplate *gametemplate.QuestTemplate) error

func (hf CommitFlowHandlerFunc) CommitFlowHandle(pl player.Player, questTemplate *gametemplate.QuestTemplate) error {
	return hf(pl, questTemplate)
}

func RegisterCommitFlow(questType questtypes.QuestType, h CommitFlowHandler) {
	_, exist := commitFlowHandlerMap[questType]
	if exist {
		panic(fmt.Sprintf("repeat register questType %d", questType))
	}
	commitFlowHandlerMap[questType] = h
}

func CommitFlowHandle(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	questType := questTemplate.GetQuestType()
	h, exist := commitFlowHandlerMap[questType]
	if !exist {
		return defaultCommitFlowHandler(pl, questTemplate)
	}
	return h.CommitFlowHandle(pl, questTemplate)
}

var (
	commitFlowHandlerMap     = make(map[questtypes.QuestType]CommitFlowHandler)
	defaultCommitFlowHandler = CommitFlowHandlerFunc(handleCommitFlowDefault)
)
