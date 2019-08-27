package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	majoreventtypes "fgame/fgame/game/major/event/types"
	majorlogic "fgame/fgame/game/major/logic"
	"fgame/fgame/game/major/pbutil"
	"fgame/fgame/game/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

//双修邀请决策
func majorInviteDeal(target event.EventTarget, data event.EventData) (err error) {
	eventData, ok := data.(*majoreventtypes.MajorInviteDealEventData)
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
		scMajorSpouseRefused := pbutil.BuildSCMajorSpouseRefused(spl.GetName(), int32(eventData.GetMajorType()), eventData.GetFubenId())
		pl.SendMsg(scMajorSpouseRefused)
	}
	return
}

//玩家进入场景
func onPlayerEnterScene(ctx context.Context, result interface{}, err error) (terr error) {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)
	eventData := result.(*majoreventtypes.MajorInviteDealEventData)
	spl := eventData.GetSpousePlayer()
	if spl == nil {
		return
	}

	s, _ := majorlogic.PlayerEnterMajor(tpl, spl.GetId(), eventData.GetMajorType(), eventData.GetFubenId())
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
	gameevent.AddEventListener(majoreventtypes.EventTypeMajorInviteDeal, event.EventListenerFunc(majorInviteDeal))
}
