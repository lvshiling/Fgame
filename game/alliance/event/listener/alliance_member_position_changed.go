package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	playeralliance "fgame/fgame/game/alliance/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
)

//成员职位变更
func memberPositionChanged(target event.EventTarget, data event.EventData) (err error) {
	memObj, ok := data.(*alliance.AllianceMemberObject)
	if !ok {
		return
	}

	memId := memObj.GetMemberId()
	pl := player.GetOnlinePlayerManager().GetPlayerById(memId)
	if pl == nil {
		return
	}

	ctx := scene.WithPlayer(context.Background(), pl)
	memberJoinMsg := message.NewScheduleMessage(onMemberPosChanged, ctx, memObj, nil)
	pl.Post(memberJoinMsg)

	return
}

//成员职位变更
func onMemberPosChanged(ctx context.Context, result interface{}, err error) error {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)
	mem := result.(*alliance.AllianceMemberObject)
	pos := mem.GetPosition()

	allianceManager := tpl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	allianceManager.SyncAlliancePos(pos)

	return nil
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceMemberPositionChanged, event.EventListenerFunc(memberPositionChanged))
}
