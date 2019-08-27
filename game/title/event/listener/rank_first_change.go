package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	rankeventtypes "fgame/fgame/game/rank/event/types"
	ranktypes "fgame/fgame/game/rank/types"
	"fgame/fgame/game/scene/scene"
	titlelogic "fgame/fgame/game/title/logic"
	"fgame/fgame/game/title/title"
	titletypes "fgame/fgame/game/title/types"
)

//排行榜第一改变
func rankFirstChanged(target event.EventTarget, data event.EventData) (err error) {
	eventData, ok := data.(*rankeventtypes.RankEventData)
	if !ok {
		return
	}
	rankClassType := eventData.GetRankClassType()
	if rankClassType != ranktypes.RankClassTypeLocal {
		return
	}

	rankType := eventData.GetRankType()
	if !rankType.Valid() {
		return
	}
	playerId := eventData.GetPlayerId()
	oldPlayerId := eventData.GetOldPlayerId()
	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	olpl := player.GetOnlinePlayerManager().GetPlayerById(oldPlayerId)
	titileRankSubType := titletypes.RankTypeToTitleRankSubType(rankType)

	titleId, _ := title.GetTitleService().GetTitleId(titletypes.TitleTypeRank, titileRankSubType)
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
	gameevent.AddEventListener(rankeventtypes.RankEventTypeFirst, event.EventListenerFunc(rankFirstChanged))
}
