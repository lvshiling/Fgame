package listener

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	marrylogic "fgame/fgame/game/marry/logic"
	"fgame/fgame/game/marry/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
)

//玩家婚礼预定失败归还
func playerMarryPreWedGiveBack(target event.EventTarget, data event.EventData) (err error) {
	playerId, ok := target.(int64)
	if !ok {
		return
	}
	eventData, ok := data.(*marryeventtypes.MarryPreWedGiveBackEventData)
	if !ok {
		return
	}
	isRefuse := eventData.GetIsRefuse()
	emailTitle := lang.GetLangService().ReadLang(lang.MarryPreWedGiveBackTitle)
	emailRefuseContent := lang.GetLangService().ReadLang(lang.MarryPreWedGiveBackRefuseContent)
	emailRobbedContent := lang.GetLangService().ReadLang(lang.MarryPreWedGiveBackRobbedContent)
	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl != nil {
		if isRefuse {
			scMarryWedGradeRefuseToPeer := pbuitl.BuildSCMarryWedGradeRefuseToPeer()
			pl.SendMsg(scMarryWedGradeRefuseToPeer)
		}
		ctx := scene.WithPlayer(context.Background(), pl)
		playerPreWedGiveBackMsg := message.NewScheduleMessage(onMarryPreWedGiveBack, ctx, eventData, nil)
		pl.Post(playerPreWedGiveBackMsg)
	} else {
		grade := eventData.GetGrade()
		hunCheGrade := eventData.GetHunCheGrade()
		sugarGrade := eventData.GetSugarGrade()
		itemMap := marrylogic.MarryPreWedItemMap(grade, hunCheGrade, sugarGrade)
		//被拒邮件
		if isRefuse {
			emaillogic.AddOfflineEmail(playerId, emailTitle, emailRefuseContent, itemMap)
		} else {
			emaillogic.AddOfflineEmail(playerId, emailTitle, emailRobbedContent, itemMap)
		}
	}
	return
}

func onMarryPreWedGiveBack(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	eventData := result.(*marryeventtypes.MarryPreWedGiveBackEventData)

	isRefuse := eventData.GetIsRefuse()
	grade := eventData.GetGrade()
	hunCheGrade := eventData.GetHunCheGrade()
	sugarGrade := eventData.GetSugarGrade()
	itemMap := marrylogic.MarryPreWedItemMap(grade, hunCheGrade, sugarGrade)

	emailTitle := lang.GetLangService().ReadLang(lang.MarryPreWedGiveBackTitle)
	emailRefuseContent := lang.GetLangService().ReadLang(lang.MarryPreWedGiveBackRefuseContent)
	emailRobbedContent := lang.GetLangService().ReadLang(lang.MarryPreWedGiveBackRobbedContent)
	if isRefuse {
		// 被拒邮件
		emaillogic.AddEmail(pl, emailTitle, emailRefuseContent, itemMap)
	} else {
		// 被抢先预定邮件
		emaillogic.AddEmail(pl, emailTitle, emailRobbedContent, itemMap)
	}
	return nil
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryPreWedGiveBack, event.EventListenerFunc(playerMarryPreWedGiveBack))
}
