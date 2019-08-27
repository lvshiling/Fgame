package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
)

//绝版首饰
func discountBuyZhuanShengGift(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*welfareeventtypes.PlayerAllianceCheerEventData)
	if !ok {
		return
	}
	groupId := eventData.GetGroupId()
	giftType := eventData.GetGiftType()
	if groupId != questtypes.QiYuEquipGiftGroupId && giftType != questtypes.QiYuEquipGiftIndex {
		return
	}

	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeBuyEquipGift, 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeDiscountBuyZhuanShengGift, event.EventListenerFunc(discountBuyZhuanShengGift))
}
