package listener

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	jieyieventtypes "fgame/fgame/game/jieyi/event/types"
	"fgame/fgame/game/jieyi/jieyi"
	jieyitemplate "fgame/fgame/game/jieyi/template"
	jieyitypes "fgame/fgame/game/jieyi/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
)

func init() {
	gameevent.AddEventListener(jieyieventtypes.JieYiEventTypeJieYiInviteFail, event.EventListenerFunc(jieYiInviteFail))
}

func jieYiInviteFail(target event.EventTarget, data event.EventData) (err error) {
	inviteObj, ok := target.(*jieyi.JieYiInviteObject)
	if !ok {
		return
	}
	playerId := inviteObj.GetInviteId()
	daoJu := inviteObj.GetJieYiDaoJu()
	sendEmailNotJieYiSuccess(playerId, daoJu)
	return
}

// 结义失败,发邮件
func sendEmailNotJieYiSuccess(inviteId int64, daoJu jieyitypes.JieYiDaoJuType) {
	invitePl := player.GetOnlinePlayerManager().GetPlayerById(inviteId)
	if invitePl == nil {
		sendOfflineEmail(inviteId, daoJu)
		return
	}
	ctx := scene.WithPlayer(context.Background(), invitePl)
	msg := message.NewScheduleMessage(sendOnlineEmail, ctx, daoJu, nil)
	invitePl.Post(msg)
}

// 结义失败,邀请者在线,发送邮件
func sendOnlineEmail(ctx context.Context, result interface{}, err error) error {
	p := scene.PlayerInContext(ctx)
	pl := p.(player.Player)
	daoJu := result.(jieyitypes.JieYiDaoJuType)

	daoJuTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiDaoJuTemplate(daoJu)
	if daoJuTemp == nil {
		return nil
	}

	itemMap := daoJuTemp.GetNeedItemMap()
	title := lang.GetLangService().ReadLang(lang.EmailJieYiDaoJuReturnTitle)
	content := lang.GetLangService().ReadLang(lang.EmailJieYiDaoJuReturnContent)
	emaillogic.AddEmail(pl, title, content, itemMap)
	return nil
}

// 结义失败,邀请者不在线,发送离线邮件
func sendOfflineEmail(inviteId int64, daoJu jieyitypes.JieYiDaoJuType) {
	daoJuTemp := jieyitemplate.GetJieYiTemplateService().GetJieYiDaoJuTemplate(daoJu)
	if daoJuTemp == nil {
		return
	}
	itemMap := daoJuTemp.GetNeedItemMap()
	title := lang.GetLangService().ReadLang(lang.EmailJieYiDaoJuReturnTitle)
	content := lang.GetLangService().ReadLang(lang.EmailJieYiDaoJuReturnContent)
	emaillogic.AddOfflineEmail(inviteId, title, content, itemMap)
}
