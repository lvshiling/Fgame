package listener

// import (
// 	"fgame/fgame/core/event"
// 	shenmoeventtypes "fgame/fgame/cross/shenmo/event/types"
// 	gameevent "fgame/fgame/game/event"
// 	shenmoscene "fgame/fgame/game/shenmo/scene"
// 	"fgame/fgame/game/shenmo/shenmo"
// )

// //玩家退出神魔战场
// func shenMoPlayerExitScene(target event.EventTarget, data event.EventData) (err error) {
// 	sd, ok := target.(shenmoscene.ShenMoSceneData)
// 	if !ok {
// 		return
// 	}
// 	num := sd.GetScenePlayerNum()
// 	shenmo.GetShenMoService().RemoveFirstLineUpPlayer(num)
// 	return
// }

// func init() {
// 	gameevent.AddEventListener(shenmoeventtypes.EventTypeShenMoPlayerExit, event.EventListenerFunc(shenMoPlayerExitScene))
// }
