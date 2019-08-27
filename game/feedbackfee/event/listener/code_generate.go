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
func codeGenerate(target event.EventTarget, data event.EventData) (err error) {
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
	onPlayerCodeGenerateMsg := message.NewScheduleMessage(onPlayerCodeGenerate, ctx, obj, nil)
	p.Post(onPlayerCodeGenerateMsg)
	return
}

func onPlayerCodeGenerate(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	feedbackExchangeObject := result.(*feedbackfee.FeedbackExchangeObject)

	feedbackfeelogic.PlayerCodeGenerate(pl, feedbackExchangeObject)
	return nil
}

func init() {
	gameevent.AddEventListener(feedbackfeeeventtypes.EventTypeCodeGenerate, event.EventListenerFunc(codeGenerate))
}
