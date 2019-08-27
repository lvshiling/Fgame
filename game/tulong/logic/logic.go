package logic

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	chattypes "fgame/fgame/game/chat/types"
	crosslogic "fgame/fgame/game/cross/logic"
	crosstypes "fgame/fgame/game/cross/types"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	tulongtemplate "fgame/fgame/game/tulong/template"
	"fmt"
)

func PlayerEnterTuLongScene(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error) {
	tuLongConstTemplate := tulongtemplate.GetTuLongTemplateService().GetTuLongConstTemplate()
	if tuLongConstTemplate == nil {
		return
	}

	crosslogic.PlayerEnterCross(pl, crosstypes.CrossTypeTuLong)
	flag = true
	return
}

//跨服屠龙击杀boss 仙盟成员奖励
func GiveTuLongKillBossReward(al *alliance.Alliance, tuLongTemplate *gametemplate.TuLongTemplate) {
	if al == nil || tuLongTemplate == nil {
		return
	}
	biologyTemplate := tuLongTemplate.GetBiologyTemplate()
	if biologyTemplate == nil {
		return
	}
	offRewItemMap := tuLongTemplate.GetOffRewItemMap()
	emailTitle := lang.GetLangService().ReadLang(lang.TuLongKillBossTitle)
	allianceName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(al.GetAllianceName()))
	bossName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(biologyTemplate.Name))
	emailContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.TuLongKillBossContent), allianceName, bossName)
	if len(offRewItemMap) == 0 {
		return
	}
	data := newTempData(offRewItemMap, emailTitle, emailContent)
	memberList := al.GetMemberList()
	for _, member := range memberList {
		playerId := member.GetMemberId()
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		//写离线邮件
		if pl == nil {
			emaillogic.AddOfflineEmail(playerId, emailTitle, emailContent, offRewItemMap)
		} else {
			ctx := scene.WithPlayer(context.Background(), pl)
			playerRewardMsg := message.NewScheduleMessage(onTuLongKillBossReward, ctx, data, nil)
			pl.Post(playerRewardMsg)
		}
	}
}

type tempData struct {
	itemMap      map[int32]int32
	emailTitle   string
	emailContent string
}

func newTempData(itemMap map[int32]int32, emailTitle string, emailContent string) *tempData {
	d := &tempData{
		itemMap:      itemMap,
		emailTitle:   emailTitle,
		emailContent: emailContent,
	}
	return d
}

func onTuLongKillBossReward(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	data := result.(*tempData)
	emailTitle := data.emailTitle
	emailContent := data.emailContent
	offRewItemMap := data.itemMap
	emaillogic.AddEmail(pl, emailTitle, emailContent, offRewItemMap)
	return nil
}
