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

//仙府开始升级
func xianFuStartUpgrade(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	xianFuType, ok := data.(xianfutypes.XianfuType)
	if !ok {
		return
	}

	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeUpgradeSpecialXianFu, int32(xianFuType), 1)
	return
}

func init() {
	gameevent.AddEventListener(xianfueventtypes.EventTypeXianFuStartUpgrade, event.EventListenerFunc(xianFuStartUpgrade))
}
