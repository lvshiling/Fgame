package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	huiyuaneventtypes "fgame/fgame/game/huiyuan/event/types"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//购买会员
func playerHuiYuanBuy(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	huiyuanType := data.(huiyuantypes.HuiYuanType)
	if huiyuanType != huiyuantypes.HuiYuanTypePlus {
		return
	}

	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeBuyHuiYuan, 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(huiyuaneventtypes.EventTypeHuiYuanBuy, event.EventListenerFunc(playerHuiYuanBuy))
}
