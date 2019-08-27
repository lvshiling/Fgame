package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	xianfueventtypes "fgame/fgame/game/xianfu/event/types"
	xianfutypes "fgame/fgame/game/xianfu/types"
)

//仙府挑战
func xianFuSweep(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	eventData, ok := data.(*xianfueventtypes.XianFuChallengeEventData)
	if !ok {
		return
	}
	typ := eventData.GetType()
	num := eventData.GetNum()
	if num <= 0 {
		return
	}

	err = sweepSpecialXianFu(pl, typ, num)
	if err != nil {
		return
	}

	err = sweepXianFu(pl, num)
	if err != nil {
		return
	}

	err = sweepXianFuPersonal(pl, typ, num)
	if err != nil {
		return
	}

	return
}

//进入指定X次秘境仙府
func sweepSpecialXianFu(pl player.Player, typ xianfutypes.XianfuType, num int32) (err error) {
	return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeEnterSpecialXianFu, int32(typ), num)
}

//进入X次秘境仙府
func sweepXianFu(pl player.Player, num int32) (err error) {
	return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeEnterXianFu, 0, num)
}

//仙府个人副本(进入)
func sweepXianFuPersonal(pl player.Player, typ xianfutypes.XianfuType, num int32) (err error) {
	switch typ {
	case xianfutypes.XianfuTypeSilver,
		xianfutypes.XianfuTypeExp:
		{
			return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeXianFuPersonal, 0, num)
		}
	}
	return
}

func init() {
	gameevent.AddEventListener(xianfueventtypes.EventTypeXianFuSweep, event.EventListenerFunc(xianFuSweep))
}
