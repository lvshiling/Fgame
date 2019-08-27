package listener

// import (
// 	"fgame/fgame/core/event"
// 	chuangshidata "fgame/fgame/game/chuangshi/data"
// 	chuangshieventtypes "fgame/fgame/game/chuangshi/event/types"
// 	gameevent "fgame/fgame/game/event"
// 	fashionlogic "fgame/fgame/game/fashion/logic"
// 	"fgame/fgame/game/player"
// )

// //成员阵营更换
// func chuangShiMemChangedCampResult(target event.EventTarget, data event.EventData) (err error) {
// 	playerId, ok := target.(int64)
// 	if !ok {
// 		return
// 	}

// 	newCamp, ok := data.(*chuangshidata.CampData)
// 	if !ok {
// 		return
// 	}

// 	//结果
// 	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
// 	if pl == nil {
// 		return
// 	}

// 	fashionlogic.ActivateCampFashion(pl, newCamp.CampType)
// 	return
// }

// func init() {
// 	gameevent.AddEventListener(chuangshieventtypes.EventTypeChuangShiMemChangedCampResult, event.EventListenerFunc(chuangShiMemChangedCampResult))
// }
