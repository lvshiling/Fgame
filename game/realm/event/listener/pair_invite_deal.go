package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	realmeventtypes "fgame/fgame/game/realm/event/types"
	realmlogic "fgame/fgame/game/realm/logic"
	"fgame/fgame/game/realm/pbutil"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

//夫妻助战决策
func pairInviteDeal(target event.EventTarget, data event.EventData) (err error) {
	eventData, ok := data.(*realmeventtypes.RealmPairInviteDealEventData)
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
		scRealmSpouseRefused := pbutil.BuildSCRealmSpouseRefused(spl.GetName())
		pl.SendMsg(scRealmSpouseRefused)
	}
	return
}

//玩家进入场景
func onPlayerEnterScene(ctx context.Context, result interface{}, err error) (terr error) {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)
	eventData := result.(*realmeventtypes.RealmPairInviteDealEventData)
	level := eventData.GetLevel()
	spl := eventData.GetSpousePlayer()
	if spl == nil {
		return
	}

	s, _ := realmlogic.PlayerEnterTianJieTa(tpl, spl.GetId(), level)
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
	gameevent.AddEventListener(realmeventtypes.EventTypeRealmPairInviteDeal, event.EventListenerFunc(pairInviteDeal))
}
