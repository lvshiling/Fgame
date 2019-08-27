package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/chat/chat"
	chateventtypes "fgame/fgame/game/chat/event/types"
	chattypes "fgame/fgame/game/chat/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

func channelFinish(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData := data.(*chat.ChatEventData)

	err = chatOnWorld(pl, eventData)
	if err != nil {
		return
	}

	err = chatOnBangPai(pl, eventData)
	if err != nil {
		return
	}
	return
}

//完成一次世界频道聊天
func chatOnWorld(pl player.Player, eventData *chat.ChatEventData) (err error) {
	channel := eventData.GetChannel()
	if channel != chattypes.ChannelTypeWorld {
		return
	}
	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeWorldChannel, 0, 1)
	return
}

//完成一次帮派频道聊天
func chatOnBangPai(pl player.Player, eventData *chat.ChatEventData) (err error) {
	channel := eventData.GetChannel()
	if channel != chattypes.ChannelTypeBangPai {
		return
	}

	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeAllianceChat, 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(chateventtypes.EventTypeChat, event.EventListenerFunc(channelFinish))
}
