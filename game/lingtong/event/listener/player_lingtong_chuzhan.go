package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	lingtonglogic "fgame/fgame/game/lingtong/logic"
	"fgame/fgame/game/player"
	"fmt"
)

//灵童出战
func playerLingTongChuZhanChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	lingTong := pl.GetLingTong()
	//还没有灵童
	if lingTong == nil {
		flag := lingtonglogic.InitLingTong(pl)
		if !flag {
			panic(fmt.Errorf("初始化灵童应该成功"))
		}
	} else {
		flag := lingtonglogic.UpdateLingTong(pl)
		if !flag {
			panic(fmt.Errorf("更新灵童应该成功"))
		}

	}

	return
}

func init() {
	gameevent.AddEventListener(lingtongeventtypes.EventTypeLingTongChuZhanChanged, event.EventListenerFunc(playerLingTongChuZhanChanged))
}
