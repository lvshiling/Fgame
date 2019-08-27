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

//仙盟等级改变
func allianceLevelChanged(target event.EventTarget, data event.EventData) (err error) {
	al := target.(*alliance.Alliance)

	for _, member := range al.GetMemberList() {
		memberPl := player.GetOnlinePlayerManager().GetPlayerById(member.GetMemberId())
		if memberPl == nil {
			continue
		}
		ctx := scene.WithPlayer(context.Background(), memberPl)
		allianceLevelChangedMsg := message.NewScheduleMessage(onAllianceLevelChanged, ctx, al, nil)
		memberPl.Post(allianceLevelChangedMsg)
	}

	return
}

//仙盟等级变化回调
func onAllianceLevelChanged(ctx context.Context, result interface{}, err error) error {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)
	al := result.(*alliance.Alliance)

	allianceManager := tpl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	allianceManager.SyncAllianceLevel(al.GetAllianceLevel())
	return nil
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceLevelChanged, event.EventListenerFunc(allianceLevelChanged))
}
