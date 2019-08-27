package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	"fgame/fgame/game/arenapvp/arenapvp"
	arenapvpeventtypes "fgame/fgame/game/arenapvp/event/types"
	arenapvplogic "fgame/fgame/game/arenapvp/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
)

//竞猜结算
func arenapvpGuessResult(target event.EventTarget, data event.EventData) (err error) {
	attendList, ok := data.([]*arenapvp.ArenapvpGuessRecordObject)
	if !ok {
		return
	}

	//结算
	for _, attendObj := range attendList {
		pl := player.GetOnlinePlayerManager().GetPlayerById(attendObj.GetPlayerId())
		if pl == nil {
			continue
		}
		ctx := scene.WithPlayer(context.Background(), pl)
		arenapvpGuessMsg := message.NewScheduleMessage(onArenapvpGuessResult, ctx, attendObj, nil)
		pl.Post(arenapvpGuessMsg)
	}

	return
}

//竞猜结算
func onArenapvpGuessResult(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	data := result.(*arenapvp.ArenapvpGuessRecordObject)

	arenapvplogic.GuessResult(pl, data)
	return nil
}

func init() {
	gameevent.AddEventListener(arenapvpeventtypes.EventTypeArenapvpGuessResult, event.EventListenerFunc(arenapvpGuessResult))
}
