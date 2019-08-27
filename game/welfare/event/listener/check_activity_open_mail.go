package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/welfare"
)

//检查活动开启邮件
func checkActivityOpenMail(target event.EventTarget, data event.EventData) (err error) {
	group, ok := target.(int32)
	if !ok {
		return
	}

	if !welfarelogic.IsOnActivityTime(group) {
		return
	}

	plList := player.GetOnlinePlayerManager().GetAllPlayers()
	for _, pl := range plList {
		ctx := scene.WithPlayer(context.Background(), pl)
		msg := message.NewScheduleMessage(sendNoticeMail, ctx, group, nil)
		pl.Post(msg)
	}

	welfare.GetWelfareService().AddStartMailRecord(group)
	return
}

func sendNoticeMail(ctx context.Context, result interface{}, err error) error {
	p := scene.PlayerInContext(ctx)
	pl := p.(player.Player)

	group := result.(int32)
	welfarelogic.SendOpenNoticeMail(pl, group)
	return nil
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeCheckActivityOpenMail, event.EventListenerFunc(checkActivityOpenMail))
}
