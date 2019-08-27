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

//竞猜退还
func arenapvpGuessReturn(target event.EventTarget, data event.EventData) (err error) {
	attendList, ok := data.([]*arenapvp.ArenapvpGuessRecordObject)
	if !ok {
		return
	}

	//结算
	for _, attendObj := range attendList {
		pl := player.GetOnlinePlayerManager().GetPlayerById(attendObj.GetPlayerId())
		ctx := scene.WithPlayer(context.Background(), pl)
		arenapvpGuessMsg := message.NewScheduleMessage(onArenapvpGuessReturn, ctx, attendObj, nil)
		pl.Post(arenapvpGuessMsg)
	}

	return
}

//竞猜退还
func onArenapvpGuessReturn(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	data := result.(*arenapvp.ArenapvpGuessRecordObject)

	arenapvplogic.GuessReturn(pl, data)
	return nil
}

func init() {
	gameevent.AddEventListener(arenapvpeventtypes.EventTypeArenapvpGuessReturn, event.EventListenerFunc(arenapvpGuessReturn))
}
