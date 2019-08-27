package logic

import (
	"context"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/common/message"
	"fgame/fgame/game/foe/pbutil"
	playerfoe "fgame/fgame/game/foe/player"
	friendtemplate "fgame/fgame/game/friend/template"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/scene/scene"
	"fmt"
)

type FoeFeedbackData struct {
	feedbackPlId int64
	feedbackName string
	args         string
}

// 仇人反馈推送
func FoeFeedbackNotice(foePl player.Player, feedbackPlId int64, feedbackName, args string) {
	data := &FoeFeedbackData{
		feedbackPlId: feedbackPlId,
		feedbackName: feedbackName,
		args:         args,
	}
	ctx := scene.WithPlayer(context.Background(), foePl)
	msg := message.NewScheduleMessage(onFoeFeedback, ctx, data, nil)
	foePl.Post(msg)
}

func onFoeFeedback(ctx context.Context, result interface{}, err error) error {
	p := scene.PlayerInContext(ctx)
	data := result.(*FoeFeedbackData)
	pl := p.(player.Player)

	manager := pl.GetPlayerDataManager(playertypes.PlayerFoeDataManagerType).(*playerfoe.PlayerFoeDataManager)
	noticeConstantTemp := friendtemplate.GetFriendNoticeTemplateService().GetFriendNoticeConstanTemplate()
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	needSilver := int64(noticeConstantTemp.ChouRenSunShiSilver)
	curSilever := propertyManager.GetSilver()
	if curSilever < needSilver {
		needSilver = curSilever
	}
	isProtected := manager.IsOnProtected()
	if !isProtected {
		useReason := commonlog.SilverLogReasonFoeFeedbackCost
		flag := propertyManager.CostSilver(needSilver, useReason, useReason.String())
		if !flag {
			panic(fmt.Errorf("foe:被报复消耗银两应该成功"))
		}

		propertylogic.SnapChangedProperty(pl)
	}

	//报复结果
	feedbackPl := player.GetOnlinePlayerManager().GetPlayerById(data.feedbackPlId)
	if feedbackPl != nil {
		scMsg := pbutil.BuildSCFoeFeedback(isProtected, pl.GetName(), noticeConstantTemp.ChouRenSunShiSilver, data.args, pl.GetSex())
		feedbackPl.SendMsg(scMsg)
	}

	frScMsg := pbutil.BuildSCFoeFeedbackNotice(data.feedbackPlId, data.feedbackName, isProtected, needSilver)
	pl.SendMsg(frScMsg)
	return nil
}
