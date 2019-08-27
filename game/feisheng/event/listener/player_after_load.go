package listener

// import (
// 	"fgame/fgame/core/event"
// 	gameevent "fgame/fgame/game/event"
// 	"fgame/fgame/game/player"
// 	playereventtypes "fgame/fgame/game/player/event/types"
// 	viplogic "fgame/fgame/game/vip/logic"
// )

// // 玩家加载后
// func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
// 	pl, ok := target.(player.Player)
// 	if !ok {
// 		return
// 	}

// 	viplogic.VipInfoNotice(pl)
// 	return
// }

// func init() {
// 	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
// }
