package listener

// import (
// 	"fgame/fgame/core/event"
// 	chuangshidata "fgame/fgame/game/chuangshi/data"
// 	chuangshieventtypes "fgame/fgame/game/chuangshi/event/types"
// 	"fgame/fgame/game/chuangshi/pbutil"
// 	gameevent "fgame/fgame/game/event"
// 	"fgame/fgame/game/player"
// )

// //神王变更推送
// func chuangShiShenWangChanged(target event.EventTarget, data event.EventData) (err error) {
// 	newCamp, ok := target.(*chuangshidata.CampData)
// 	if !ok {
// 		return
// 	}

// 	for _, mem := range newCamp.MemberList {
// 		pl := player.GetOnlinePlayerManager().GetPlayerById(mem.PlayerId)
// 		if pl == nil {
// 			continue
// 		}

// 		scMsg := pbutil.BuildSCChuangShiShenWangBroadcast(newCamp.KingMem)
// 		pl.SendMsg(scMsg)
// 	}

// 	// TODO xzk27 同步阵营信息 ChuangShiCamp

// 	return
// }

// func init() {
// 	gameevent.AddEventListener(chuangshieventtypes.EventTypeChuangShiShenWangChanged, event.EventListenerFunc(chuangShiShenWangVote))
// }
