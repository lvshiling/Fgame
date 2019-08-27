package listener

// import (
// 	"context"
// 	"fgame/fgame/common/message"
// 	"fgame/fgame/core/event"
// 	"fgame/fgame/game/chuangshi/chuangshi"
// 	chuangshidata "fgame/fgame/game/chuangshi/data"
// 	chuangshieventtypes "fgame/fgame/game/chuangshi/event/types"
// 	chuangshilogic "fgame/fgame/game/chuangshi/logic"
// 	"fgame/fgame/game/chuangshi/pbutil"
// 	gameevent "fgame/fgame/game/event"
// 	"fgame/fgame/game/player"
// 	"fgame/fgame/game/scene/scene"
// )

// //神王报名结果
// func chuangShiShenWangSignUp(target event.EventTarget, data event.EventData) (err error) {
// 	signUpObj, ok := target.(*chuangshi.ChuangShiShenWangSignUpObject)
// 	if !ok {
// 		return
// 	}

// 	signList, ok := data.([]*chuangshidata.MemberInfo)
// 	if !ok {
// 		return
// 	}

// 	//结果
// 	pl := player.GetOnlinePlayerManager().GetPlayerById(signUpObj.GetPlayerId())
// 	if pl == nil {
// 		return
// 	}

// 	ctx := scene.WithPlayer(context.Background(), pl)
// 	chuangshiSignUpMsg := message.NewScheduleMessage(onShenWangSignUpResult, ctx, nil, nil)
// 	pl.Post(chuangshiSignUpMsg)

// 	scMsg := pbutil.BuildSCChuangShiShengWangBaoMingList(signList)
// 	pl.SendMsg(scMsg)
// 	return
// }

// func onShenWangSignUpResult(ctx context.Context, result interface{}, err error) error {
// 	tpl := scene.PlayerInContext(ctx)
// 	pl := tpl.(player.Player)

// 	chuangshilogic.ShenWangSignUpResult(pl)
// 	return nil
// }

// func init() {
// 	gameevent.AddEventListener(chuangshieventtypes.EventTypeChuangShiShenWangSignUp, event.EventListenerFunc(chuangShiShenWangSignUp))
// }
