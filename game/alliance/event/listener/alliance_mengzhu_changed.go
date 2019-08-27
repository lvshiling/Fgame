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
	"fgame/fgame/game/scene/scene"
)

//仙盟盟主改变
func allianceMengzhuChanged(target event.EventTarget, data event.EventData) (err error) {
	al := target.(*alliance.Alliance)

	for _, member := range al.GetMemberList() {
		memberPl := player.GetOnlinePlayerManager().GetPlayerById(member.GetMemberId())
		if memberPl == nil {
			continue
		}
		ctx := scene.WithPlayer(context.Background(), memberPl)
		allianceMengzhuChangedMsg := message.NewScheduleMessage(onAllianceMengzhuChanged, ctx, al, nil)
		memberPl.Post(allianceMengzhuChangedMsg)
	}

	return
}

//仙盟盟主变化回调
func onAllianceMengzhuChanged(ctx context.Context, result interface{}, err error) error {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)
	al := result.(*alliance.Alliance)
	allianceManager := tpl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	allianceManager.SyncAllianceMengzhu(al.GetAllianceMengZhuId())

	//盟主信息推送
	scMsg := pbutil.BuildSCAllianceMengZhuInfoNotice(al)
	pl.SendMsg(scMsg)

	return nil
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceMengzhuChanged, event.EventListenerFunc(allianceMengzhuChanged))
}
