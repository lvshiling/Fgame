package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	xianfueventtypes "fgame/fgame/game/xianfu/event/types"
	playerxianfu "fgame/fgame/game/xianfu/player"
	xianfutemplate "fgame/fgame/game/xianfu/template"
	xianfutypes "fgame/fgame/game/xianfu/types"
)

//仙府挑战
func xianFuChallenge(target event.EventTarget, data event.EventData) (err error) {

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

	manager := pl.GetPlayerDataManager(types.PlayerXianfuDtatManagerType).(*playerxianfu.PlayerXinafuDataManager)
	xianFuObj := manager.GetPlayerXianfuInfo(typ)
	if xianFuObj == nil {
		return
	}
	xianFuId := xianFuObj.GetXianfuId()
	useTimes := xianFuObj.GetUseTimes()
	freeTimes := xianfutemplate.GetXianfuTemplateService().GetFreePlayTimes(typ, xianFuId)
	leftNum := freeTimes - useTimes

	specialXianFu(pl, typ, leftNum, num)
	return
}

//进入指定秘境仙府x次数(仅限免费次数)
func specialXianFu(pl player.Player, typ xianfutypes.XianfuType, leftNum int32, num int32) (err error) {
	if leftNum <= 0 {
		return questlogic.FillQuestData(pl, questtypes.QuestSubTypeSpecialXianFu, int32(typ))
	} else {
		return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeSpecialXianFu, int32(typ), num)
	}
	return
}

func init() {
	gameevent.AddEventListener(xianfueventtypes.EventTypeXianFuChallenge, event.EventListenerFunc(xianFuChallenge))
}
