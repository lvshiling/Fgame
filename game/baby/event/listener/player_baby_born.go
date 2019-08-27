package listener

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	babyeventtypes "fgame/fgame/game/baby/event/types"
	"fgame/fgame/game/baby/pbutil"
	playerbaby "fgame/fgame/game/baby/player"
	babytemplate "fgame/fgame/game/baby/template"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
)

// 玩家宝宝出生
func playerBabyBorn(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	baby, ok := data.(*playerbaby.PlayerBabyObject)
	if !ok {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	babyConstantTemplate := babytemplate.GetBabyTemplateService().GetBabyConstantTemplate()
	babyItemData := droptemplate.CreateBaoBaoCardItemData(babyConstantTemplate.BaoBaoCard, 1, baby.GetDanBei(), baby.GetQuality(), baby.GetBabySex(), baby.GetSkillList())
	// 给配偶发宝宝邮件
	spl := player.GetOnlinePlayerManager().GetPlayerById(pl.GetSpouseId())
	if spl == nil {
		title := lang.GetLangService().ReadLang(lang.BabyBornTitle)
		content := lang.GetLangService().ReadLang(lang.BabyBornContent)
		emaillogic.AddOfflineEmailItemLevel(pl.GetSpouseId(), title, content, now, []*droptemplate.DropItemData{babyItemData})
	} else {
		ctx := scene.WithPlayer(context.Background(), spl)
		msg := message.NewScheduleMessage(spouseBabyBorn, ctx, babyItemData, nil)
		spl.Post(msg)
	}

	//出生推送
	//scMsg := pbutil.BuildSCBabyBornNotice(baby)
	//pl.SendMsg(scMsg)
	return
}

// 配偶宝宝出生
func spouseBabyBorn(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	itemData := result.(*droptemplate.DropItemData)

	now := global.GetGame().GetTimeService().Now()
	title := lang.GetLangService().ReadLang(lang.BabyBornTitle)
	content := lang.GetLangService().ReadLang(lang.BabyBornContent)
	emailObj := emaillogic.AddEmailItemLevel(pl, title, content, now, []*droptemplate.DropItemData{itemData})

	scMsg := pbutil.BuildSCBabyBornSpouseNotice(emailObj.GetEmailId())
	pl.SendMsg(scMsg)
	return nil
}

func init() {
	gameevent.AddEventListener(babyeventtypes.EventTypeBabyBorn, event.EventListenerFunc(playerBabyBorn))
}
