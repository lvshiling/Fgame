package listener

// import (
// 	"context"
// 	"fgame/fgame/common/message"
// 	"fgame/fgame/core/event"
// 	"fgame/fgame/game/chuangshi/chuangshi"
// 	chuangshieventtypes "fgame/fgame/game/chuangshi/event/types"
// 	chuangshilogic "fgame/fgame/game/chuangshi/logic"
// 	gameevent "fgame/fgame/game/event"
// 	"fgame/fgame/game/player"
// 	"fgame/fgame/game/scene/scene"
// )

// //城防建设结果
// func chuangShiChengFangJianShe(target event.EventTarget, data event.EventData) (err error) {
// 	jianSheObj, ok := target.(*chuangshi.ChuangShiChengFangJianSheObject)
// 	if !ok {
// 		return
// 	}

// 	//结果
// 	pl := player.GetOnlinePlayerManager().GetPlayerById(jianSheObj.GetPlayerId())
// 	if pl == nil {
// 		return
// 	}

// 	ctx := scene.WithPlayer(context.Background(), pl)
// 	chuangshiSignUpMsg := message.NewScheduleMessage(onChengFangJianSheResult, ctx, nil, nil)
// 	pl.Post(chuangshiSignUpMsg)
// 	return
// }

// func onChengFangJianSheResult(ctx context.Context, result interface{}, err error) error {
// 	tpl := scene.PlayerInContext(ctx)
// 	pl := tpl.(player.Player)

// 	chuangshilogic.ChengFangJianSheResult(pl)
// 	return nil
// }

// func init() {
// 	gameevent.AddEventListener(chuangshieventtypes.EventTypeChuangShiChengFangJianShe, event.EventListenerFunc(chuangShiChengFangJianShe))
// }
