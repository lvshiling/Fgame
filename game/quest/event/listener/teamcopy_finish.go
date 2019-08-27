package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	teamtypes "fgame/fgame/game/team/types"
	teamcopyeventtypes "fgame/fgame/game/teamcopy/event/types"
)

//通关组队副本完成
func teamCopyFinsh(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	purpose := int32(data.(teamtypes.TeamPurposeType))

	// 完成指定副本类型任务
	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypePassTeamCopy, purpose, 1)
	if err != nil {
		return
	}

	// 完成一次组队副本
	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeFinishTeamFuBen, 0, 1)
	if err != nil {
		return
	}

	return
}

func init() {
	gameevent.AddEventListener(teamcopyeventtypes.EventTypeTeamCopyFinishSucess, event.EventListenerFunc(teamCopyFinsh))
}
