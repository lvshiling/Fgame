package listener

import (
	"fgame/fgame/core/event"
	equipbaokueventtypes "fgame/fgame/game/equipbaoku/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家装备宝库抽奖
func playerEquipBaoKuAttend(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	attendNum, ok := data.(int32)
	if !ok {
		return
	}

	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeEquipBaoKuAttend, 0, attendNum)
	return
}

func init() {
	gameevent.AddEventListener(equipbaokueventtypes.EventTypeEquipBaoKuAttend, event.EventListenerFunc(playerEquipBaoKuAttend))
}
