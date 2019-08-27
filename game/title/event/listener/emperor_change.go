package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	emperoreventtypes "fgame/fgame/game/emperor/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	titlelogic "fgame/fgame/game/title/logic"
	"fgame/fgame/game/title/title"
	titletypes "fgame/fgame/game/title/types"
)

//帝王改变
func emperorChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	oldPlayerId, ok := data.(int64)
	if !ok {
		return
	}
	olpl := player.GetOnlinePlayerManager().GetPlayerById(oldPlayerId)
	titleId, _ := title.GetTitleService().GetTitleId(titletypes.TitleTypeKing, titletypes.TitleCommonSubTypeDefault)
	if pl != nil {
		ctx := scene.WithPlayer(context.Background(), pl)
		pl.Post(message.NewScheduleMessage(titlelogic.OnTempTitleChangedGet, ctx, titleId, nil))
	}
	if olpl != nil {
		ctx := scene.WithPlayer(context.Background(), olpl)
		olpl.Post(message.NewScheduleMessage(titlelogic.OnTempTitleChangedGetRemove, ctx, titleId, nil))
	}
	return
}

func init() {
	gameevent.AddEventListener(emperoreventtypes.EmperorEventTypeRobed, event.EventListenerFunc(emperorChanged))
}
