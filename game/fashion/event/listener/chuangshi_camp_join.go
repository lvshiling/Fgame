package listener

// import (
// 	"fgame/fgame/core/event"
// 	chuangshieventtypes "fgame/fgame/game/chuangshi/event/types"
// 	chuangshitypes "fgame/fgame/game/chuangshi/types"
// 	gameevent "fgame/fgame/game/event"
// 	fashionlogic "fgame/fgame/game/fashion/logic"
// 	"fgame/fgame/game/player"
// )

// //加入阵营结果
// func chuangShiCampJoinCampResult(target event.EventTarget, data event.EventData) (err error) {
// 	playerId, ok := target.(int64)
// 	if !ok {
// 		return
// 	}
// 	campType, ok := target.(chuangshitypes.ChuangShiCampType)
// 	if !ok {
// 		return
// 	}

// 	//结果
// 	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
// 	if pl == nil {
// 		return
// 	}
// 	fashionlogic.ActivateCampFashion(pl, campType)

// 	return
// }

// func init() {
// 	gameevent.AddEventListener(chuangshieventtypes.EventTypeChuangShiCampJoinCampResult, event.EventListenerFunc(chuangShiCampJoinCampResult))
// }
