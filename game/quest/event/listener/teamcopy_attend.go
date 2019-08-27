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

//参与组队副本
func teamCopyAttend(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	purpose := data.(teamtypes.TeamPurposeType)

	// 参与组队副本
	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeAttendTeamFuBen, 0, 1)
	if err != nil {
		return
	}

	// 参与银两组队副本即完成
	if purpose == teamtypes.TeamPurposeTypeFuBenSilver {
		err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeZuDuiFuBen, int32(teamtypes.TeamPurposeTypeFuBenSilver), 1)
		if err != nil {
			return
		}
	}

	return
}

func init() {
	gameevent.AddEventListener(teamcopyeventtypes.EventTypeTeamCopyAttend, event.EventListenerFunc(teamCopyAttend))
}
