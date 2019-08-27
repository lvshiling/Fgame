package logic

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/game/compensate/compensate"
	playercompensate "fgame/fgame/game/compensate/player"
	compensatetypes "fgame/fgame/game/compensate/types"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
)

type compensateEmailData struct {
	title       string
	content     string
	sendTime    int64
	attachement []*droptemplate.DropItemData
}

func newCompensateEmailData(title, content string, sendTime int64, attachement []*droptemplate.DropItemData) *compensateEmailData {
	d := &compensateEmailData{
		title:       title,
		content:     content,
		sendTime:    sendTime,
		attachement: attachement,
	}

	return d
}

// 发补偿邮件
func SendPlayerCompensateEmail(playerId int64, title, content string, sendTime int64, attachement []*droptemplate.DropItemData) {
	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl == nil {
		emaillogic.AddOfflineEmailItemLevel(playerId, title, content, sendTime, attachement)
	} else {
		data := newCompensateEmailData(title, content, sendTime, attachement)
		ctx := scene.WithPlayer(context.Background(), pl)
		msg := message.NewScheduleMessage(onSendCompensate, ctx, data, nil)
		pl.Post(msg)
	}

}

func onSendCompensate(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	data := result.(*compensateEmailData)

	emaillogic.AddEmailItemLevel(pl, data.title, data.content, data.sendTime, data.attachement)
	return nil
}

// 添加全服补偿
func AddServerCompensate(needLevel int32, needCreateTime int64, title, content string, attachement []*droptemplate.DropItemData) {
	compensate.GetCompensateSrevice().AddCompensate(needLevel, needCreateTime, title, content, attachement)

}

//发全服补偿
func SendServerCompensate(pl player.Player, compensateObj *compensate.CompensateObject) {
	plCreateTime := pl.GetCreateTime()
	plLevel := pl.GetLevel()
	compensateId := compensateObj.GetCompensateId()
	compensateManager := pl.GetPlayerDataManager(playertypes.PlayerCompensateDataManagerType).(*playercompensate.PlayerCompensateDataManager)

	if plCreateTime > compensateObj.GetRoleCreateTime() {
		return
	}
	if plLevel < compensateObj.GetRoleLevel() {
		//添加补偿记录
		compensateManager.AddCompensate(compensateId, compensatetypes.CompensateRecordSateNotGet)
		return
	}

	if compensateManager.IsHadCompensate(compensateId) {
		return
	}

	// 发邮件
	title := compensateObj.GetTitle()
	content := compensateObj.GetConetent()
	attachement := compensateObj.GetAttachment()
	sendTime := compensateObj.GetCreateTime()
	emaillogic.AddEmailItemLevel(pl, title, content, sendTime, attachement)

	//添加补偿记录
	compensateManager.AddCompensate(compensateId, compensatetypes.CompensateRecordSateHadGet)
}
