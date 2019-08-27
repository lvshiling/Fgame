package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	"fgame/fgame/game/compensate/compensate"
	compensateeventtypes "fgame/fgame/game/compensate/event/types"
	compensatelogic "fgame/fgame/game/compensate/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
)

// 全服补偿变化
func serverCompensateChanged(target event.EventTarget, data event.EventData) (err error) {
	compensateObj, ok := target.(*compensate.CompensateObject)
	if !ok {
		return
	}

	// 在线玩家
	allPl := player.GetOnlinePlayerManager().GetAllPlayers()
	for _, pl := range allPl {
		plCreateTime := pl.GetCreateTime()

		if plCreateTime > compensateObj.GetRoleCreateTime() {
			continue
		}

		// 发邮件
		ctx := scene.WithPlayer(context.Background(), pl)
		msg := message.NewScheduleMessage(onSendCompensate, ctx, compensateObj, nil)
		pl.Post(msg)
	}

	return
}

func onSendCompensate(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl, ok := tpl.(player.Player)
	if !ok {
		return nil
	}
	compensateObj, ok := result.(*compensate.CompensateObject)
	if !ok {
		return nil
	}

	compensatelogic.SendServerCompensate(pl, compensateObj)
	return nil
}

func init() {
	gameevent.AddEventListener(compensateeventtypes.EventTypeServerCompensateChanged, event.EventListenerFunc(serverCompensateChanged))
}
