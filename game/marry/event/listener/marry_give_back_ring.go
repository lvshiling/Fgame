package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	marrylogic "fgame/fgame/game/marry/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
)

//TODO:zrc 离线直接退还发离线邮件
//玩家婚戒归还
func playerRingGiveBack(target event.EventTarget, data event.EventData) (err error) {
	playerId, ok := target.(int64)
	if !ok {
		return
	}
	ringTypeData, ok := data.(*marryeventtypes.MarryGiveBackRingEventData)
	if !ok {
		return
	}

	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl == nil {
		return
	}
	ctx := scene.WithPlayer(context.Background(), pl)
	playerRingReplaceMsg := message.NewScheduleMessage(onPlayerRingGiveBack, ctx, ringTypeData, nil)
	pl.Post(playerRingReplaceMsg)
	return
}

func onPlayerRingGiveBack(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	marryInfo := result.(*marryeventtypes.MarryGiveBackRingEventData)

	marrylogic.PlayerMarryRingGiveBack(pl, marryInfo.GetRingType(), marryInfo.GetPeerName())
	return nil
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryRingGiveBack, event.EventListenerFunc(playerRingGiveBack))
}
