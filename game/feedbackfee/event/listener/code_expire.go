package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	feedbackfeeeventtypes "fgame/fgame/game/feedbackfee/event/types"
	"fgame/fgame/game/feedbackfee/feedbackfee"
	feedbackfeelogic "fgame/fgame/game/feedbackfee/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
)

// 码过期了
func codeExpire(target event.EventTarget, data event.EventData) (err error) {
	obj, ok := target.(*feedbackfee.FeedbackExchangeObject)
	if !ok {
		return
	}
	playerId := obj.GetPlayerId()
	p := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if p == nil {
		return
	}
	ctx := scene.WithPlayer(context.Background(), p)
	onPlayerCodeExpireMsg := message.NewScheduleMessage(onPlayerCodeExpire, ctx, obj, nil)
	p.Post(onPlayerCodeExpireMsg)
	return
}

func onPlayerCodeExpire(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	feedbackExchangeObject := result.(*feedbackfee.FeedbackExchangeObject)

	feedbackfeelogic.PlayerCodeExpire(pl, feedbackExchangeObject)
	return nil
}

func init() {
	gameevent.AddEventListener(feedbackfeeeventtypes.EventTypeCodeExpire, event.EventListenerFunc(codeExpire))
}
