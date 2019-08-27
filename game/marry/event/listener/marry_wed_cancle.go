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
	"fgame/fgame/game/marry/marry"
	pbuitl "fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
)

//婚礼取消
func marryWedCancle(target event.EventTarget, data event.EventData) (err error) {
	marryWed, ok := target.(*marry.MarryWedObject)
	if !ok {
		return
	}

	playerId := marryWed.PlayerId
	spouseId := marryWed.SpouseId
	marryCancleDealwith(playerId, marryWed)
	marryCancleDealwith(spouseId, marryWed)
	return
}

func marryCancleDealwith(playerId int64, marryWed *marry.MarryWedObject) {
	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl != nil {
		ctx := scene.WithPlayer(context.Background(), pl)
		marryCancleMsg := message.NewScheduleMessage(onPlayerOffonlineMarryCancle, ctx, marryWed, nil)
		pl.Post(marryCancleMsg)
	} else {
		emailTitle := lang.GetLangService().ReadLang(lang.MarryWedCancleTitle)
		emailContent := lang.GetLangService().ReadLang(lang.MarryWedCancleContent)
		//发离线邮件
		if playerId == marryWed.PlayerId {
			grade := marryWed.Grade
			hunCheGrade := marryWed.HunCheGrade
			sugarGrade := marryWed.SugarGrade
			itemMap := marrylogic.MarryPreWedItemMap(grade, hunCheGrade, sugarGrade)
			emaillogic.AddOfflineEmail(playerId, emailTitle, emailContent, itemMap)
		} else {
			emaillogic.AddOfflineEmail(playerId, emailTitle, emailContent, nil)
		}
	}
}

func onPlayerOffonlineMarryCancle(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	marryWed := result.(*marry.MarryWedObject)
	period := marryWed.Period
	grade := marryWed.Grade
	hunCheGrade := marryWed.HunCheGrade
	sugarGrade := marryWed.SugarGrade
	proporalId := marryWed.PlayerId

	itemMap := marrylogic.MarryPreWedItemMap(grade, hunCheGrade, sugarGrade)
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	manager.CancleWedding()
	scMarryCancle := pbuitl.BuildSCMarryCancle(period)
	pl.SendMsg(scMarryCancle)
	scMarryInfoStatusChange := pbuitl.BuildSCMarryInfoStatusChange(int32(marrytypes.MarryStatusTypeProposal))
	pl.SendMsg(scMarryInfoStatusChange)

	//发邮件
	emailTitle := lang.GetLangService().ReadLang(lang.MarryWedCancleTitle)
	emailContent := lang.GetLangService().ReadLang(lang.MarryWedCancleContent)
	if proporalId == pl.GetId() {
		emaillogic.AddEmail(pl, emailTitle, emailContent, itemMap)
	} else {
		emaillogic.AddEmail(pl, emailTitle, emailContent, nil)
	}
	return nil
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryWedCancle, event.EventListenerFunc(marryWedCancle))
}
