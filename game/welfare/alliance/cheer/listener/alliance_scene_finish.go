package listener

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	alliancescene "fgame/fgame/game/alliance/scene"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	welfarelogic "fgame/fgame/game/welfare/logic"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

//九霄城战结束
func allianceSceneFinish(target event.EventTarget, data event.EventData) (err error) {
	sceneData := target.(alliancescene.AllianceSceneData)
	allianceId := sceneData.GetCurrentDefendAllianceId()
	al := alliance.GetAllianceService().GetAlliance(allianceId)

	typ := welfaretypes.OpenActivityTypeAlliance
	subType := welfaretypes.OpenActivityAllianceSubTypeAlliance
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		//获胜记录
		welfare.GetWelfareService().AddAllianceWinRecord(groupId, allianceId)

		//在线玩家返还
		for _, mem := range al.GetMemberList() {
			memPl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetMemberId())
			if memPl == nil {
				continue
			}

			ctx := scene.WithPlayer(context.Background(), memPl)
			msg := message.NewScheduleMessage(allianceCheerWin, ctx, groupId, nil)
			memPl.Post(msg)
		}
	}

	return
}

func allianceCheerWin(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	groupId := result.(int32)

	welfarelogic.AllianceCheerEndMail(pl, groupId)
	return nil
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceSceneFinish, event.EventListenerFunc(allianceSceneFinish))
}
