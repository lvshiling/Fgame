package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	majoreventtypes "fgame/fgame/game/major/event/types"
	majortypes "fgame/fgame/game/major/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//双修完成
func playerMajorSuccess(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	eventData, ok := data.(*majoreventtypes.MajorSuccessEventData)
	if !ok {
		return
	}

	majorTemp := eventData.GetTemp()
	num := eventData.GetNum()

	majorType := majorTemp.GetMajorType()
	switch majorType {
	case majortypes.MajorTypeShuangXiu:
		{
			questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeFinishMajorFuBen, 0, num)
		}
	case majortypes.MajorTypeFuQi:
		{
			questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeFinishFuQiFuBen, 0, num)
		}
	default:
		break
	}
	return
}

func init() {
	gameevent.AddEventListener(majoreventtypes.EventTypePlayerMajorSuccess, event.EventListenerFunc(playerMajorSuccess))
}
