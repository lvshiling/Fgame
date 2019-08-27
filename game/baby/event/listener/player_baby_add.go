package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/baby/baby"
	babyeventtypes "fgame/fgame/game/baby/event/types"
	babylogic "fgame/fgame/game/baby/logic"
	playerbaby "fgame/fgame/game/baby/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

// 玩家宝宝添加
func playerBabyAdd(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	addBaby, ok := data.(*playerbaby.PlayerBabyObject)
	if !ok {
		return
	}

	baby.GetBabyService().AddBaby(pl, addBaby.GetDBId(), addBaby.GetQuality(), addBaby.GetLearnLevel(), addBaby.GetDanBei(), addBaby.GetSkillList())
	babylogic.BabyPropertyChanged(pl)
	return
}

func init() {
	gameevent.AddEventListener(babyeventtypes.EventTypeBabyAdd, event.EventListenerFunc(playerBabyAdd))
}
