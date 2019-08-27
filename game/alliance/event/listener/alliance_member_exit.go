package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	playeralliance "fgame/fgame/game/alliance/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fgame/fgame/game/scene/scene"
)

//成员退出
func memberExit(target event.EventTarget, data event.EventData) (err error) {
	al := target.(*alliance.Alliance)
	eventData := data.(*allianceeventtypes.AllianceMemberExitEventData)

	//成员变更推送
	scAllianceMemberChanged := pbutil.BuildSCAllianceMemberChanged(al.GetMemberList())
	for _, mem := range al.GetMemberList() {
		p := player.GetOnlinePlayerManager().GetPlayerById(mem.GetMemberId())
		if p == nil {
			continue
		}

		p.SendMsg(scAllianceMemberChanged)
	}

	pl := player.GetOnlinePlayerManager().GetPlayerById(eventData.GetMemberId())
	if pl == nil {
		return
	}
	ctx := scene.WithPlayer(context.Background(), pl)
	memberExitMsg := message.NewScheduleMessage(onMemberExit, ctx, eventData.IsClearPlayerData(), nil)
	pl.Post(memberExitMsg)

	return
}

//成员退出回调
func onMemberExit(ctx context.Context, result interface{}, err error) error {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)
	isClearPlayerData := result.(bool)

	allianceManager := tpl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	allianceManager.ExitAlliance(isClearPlayerData)

	tpl.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeLingyuAura.Mask())

	return nil
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceMemberExit, event.EventListenerFunc(memberExit))
}
