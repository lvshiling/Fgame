package listener

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	"fgame/fgame/game/alliance/pbutil"
	chattypes "fgame/fgame/game/chat/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	"fmt"
)

//仙盟合并
func allianceMerge(target event.EventTarget, data event.EventData) (err error) {
	al := target.(*alliance.Alliance)
	mergeMemList := data.([]*alliance.AllianceMemberObject)

	//成员变更推送
	title := lang.GetLangService().ReadLang(lang.AllianceInviteAllianceMergeTitle)
	alName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(al.GetAllianceName()))
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.AllianceInviteAllianceMergeContent), alName, alName)
	for _, mem := range mergeMemList {
		p := player.GetOnlinePlayerManager().GetPlayerById(mem.GetMemberId())
		if p == nil {
			emaillogic.AddOfflineEmail(mem.GetMemberId(), title, content, nil)
		} else {
			ctx := scene.WithPlayer(context.Background(), p)
			doushenChangedMsg := message.NewScheduleMessage(onAllianceMerge, ctx, alName, nil)
			p.Post(doushenChangedMsg)

			scMsg := pbutil.BuildSCAllianceInfo(al, mem)
			p.SendMsg(scMsg)
		}
	}

	return
}

//仙盟合并通知
func onAllianceMerge(ctx context.Context, result interface{}, err error) error {

	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	alName := result.(string)

	title := lang.GetLangService().ReadLang(lang.AllianceInviteAllianceMergeTitle)
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.AllianceInviteAllianceMergeContent), alName, alName)
	emaillogic.AddEmail(pl, title, content, nil)

	return nil
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceMerge, event.EventListenerFunc(allianceMerge))
}
