package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	majoreventtypes "fgame/fgame/game/major/event/types"
	majortemplate "fgame/fgame/game/major/template"
	majortypes "fgame/fgame/game/major/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//双修挑战
func majorChallenge(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	majorTemp, ok := data.(majortemplate.MajorTemplate)
	if !ok {
		return
	}

	majorType := majorTemp.GetMajorType()
	switch majorType {
	case majortypes.MajorTypeShuangXiu:
		{
			questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeAttendMajorFuBen, 0, 1)
		}
	case majortypes.MajorTypeFuQi:
		{
			questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeAttendFuQiFuBen, 0, 1)
		}
	default:
		break
	}

	return
}

func init() {
	gameevent.AddEventListener(majoreventtypes.EventTypePlayerEnterMajorScene, event.EventListenerFunc(majorChallenge))
}
