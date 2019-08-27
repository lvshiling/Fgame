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

// //神王投票结果
// func chuangShiShenWangVote(target event.EventTarget, data event.EventData) (err error) {
// 	voteObj, ok := target.(*chuangshi.ChuangShiShenWangVoteObject)
// 	if !ok {
// 		return
// 	}

// 	voteList, ok := data.([]*chuangshidata.VoteInfo)
// 	if !ok {
// 		return
// 	}

// 	//结果
// 	pl := player.GetOnlinePlayerManager().GetPlayerById(voteObj.GetPlayerId())
// 	if pl == nil {
// 		return
// 	}

// 	ctx := scene.WithPlayer(context.Background(), pl)
// 	chuangshiVoteMsg := message.NewScheduleMessage(onShenWangVoteResult, ctx, nil, nil)
// 	pl.Post(chuangshiVoteMsg)

// 	scMsg := pbutil.BuildSCChuangShiShengWangTouPiaoList(voteList)
// 	pl.SendMsg(scMsg)
// 	return
// }

// func onShenWangVoteResult(ctx context.Context, result interface{}, err error) error {
// 	tpl := scene.PlayerInContext(ctx)
// 	pl := tpl.(player.Player)

// 	chuangshilogic.ShenWangVoteResult(pl)
// 	return nil
// }

// func init() {
// 	gameevent.AddEventListener(chuangshieventtypes.EventTypeChuangShiShenWangVote, event.EventListenerFunc(chuangShiShenWangVote))
// }
