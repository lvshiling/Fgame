package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	baguaeventtypes "fgame/fgame/game/bagua/event/types"
	bagualogic "fgame/fgame/game/bagua/logic"
	"fgame/fgame/game/bagua/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

//夫妻助战决策
func pairInviteDeal(target event.EventTarget, data event.EventData) (err error) {
	eventData, ok := data.(*baguaeventtypes.BaGuaPairInviteDealEventData)
	if !ok {
		return
	}
	agree := eventData.GetAgree()
	playerId := eventData.GetPlayerId()
	spl := eventData.GetSpousePlayer()
	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if spl == nil || pl == nil {
		return
	}

	if agree {
		ctx := scene.WithPlayer(context.Background(), pl)
		playerEnterSceneMsg := message.NewScheduleMessage(onPlayerEnterScene, ctx, eventData, nil)
		pl.Post(playerEnterSceneMsg)
	} else {
		scBaGuaSpouseRefused := pbutil.BuildSCBaGuaSpouseRefused(spl.GetName())
		pl.SendMsg(scBaGuaSpouseRefused)
	}
	return
}

//玩家进入场景
func onPlayerEnterScene(ctx context.Context, result interface{}, err error) (terr error) {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)
	eventData := result.(*baguaeventtypes.BaGuaPairInviteDealEventData)
	level := eventData.GetLevel()
	spl := eventData.GetSpousePlayer()
	if spl == nil {
		return
	}

	s, _ := bagualogic.PlayerEnterBaGua(tpl, spl.GetId(), level)
	if s == nil {
		return
	}
	splCtx := scene.WithPlayer(context.Background(), spl)
	spouseEnterSceneMsg := message.NewScheduleMessage(onSpouseEnterScene, splCtx, s, nil)
	spl.Post(spouseEnterSceneMsg)
	return nil
}

//配偶进入场景
func onSpouseEnterScene(ctx context.Context, result interface{}, err error) (terr error) {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)

	s := result.(scene.Scene)
	scenelogic.PlayerEnterSingleFuBenScene(tpl, s)
	return
}

func init() {
	gameevent.AddEventListener(baguaeventtypes.EventTypeBaGuaPairInviteDeal, event.EventListenerFunc(pairInviteDeal))
}
