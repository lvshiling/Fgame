package listener

// import (
// 	"context"
// 	"fgame/fgame/common/message"
// 	"fgame/fgame/core/event"
// 	gameevent "fgame/fgame/game/event"
// 	"fgame/fgame/game/player"
// 	realmeventtypes "fgame/fgame/game/realm/event/types"
// 	"fgame/fgame/game/scene/scene"
// 	titlelogic "fgame/fgame/game/title/logic"
// 	"fgame/fgame/game/title/title"
// 	titletypes "fgame/fgame/game/title/types"
// )

// //天劫塔第一改变
// func realmFirstChanged(target event.EventTarget, data event.EventData) (err error) {
// 	newFirstId, ok := target.(int64)
// 	if !ok {
// 		return
// 	}
// 	oldFirstId, ok := data.(int64)
// 	if !ok {
// 		return
// 	}

// 	pl := player.GetOnlinePlayerManager().GetPlayerById(newFirstId)
// 	olpl := player.GetOnlinePlayerManager().GetPlayerById(oldFirstId)
// 	titleId, _ := title.GetTitleService().GetTitleId(titletypes.TitleTypeRank, titletypes.TitleRankSubTypeRealm)
// 	if pl != nil {
// 		ctx := scene.WithPlayer(context.Background(), pl)
// 		pl.Post(message.NewScheduleMessage(titlelogic.OnTempTitleChangedGet, ctx, titleId, nil))
// 	}
// 	if olpl != nil {
// 		ctx := scene.WithPlayer(context.Background(), olpl)
// 		olpl.Post(message.NewScheduleMessage(titlelogic.OnTempTitleChangedGetRemove, ctx, titleId, nil))
// 	}
// 	return
// }

// func init() {
// 	gameevent.AddEventListener(realmeventtypes.EventTypeRealmFirstChange, event.EventListenerFunc(realmFirstChanged))
// }
