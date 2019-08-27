package logic

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/game/alliance/alliance"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	droptemplate "fgame/fgame/game/drop/template"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/yuxi/pbutil"
	playeryuxi "fgame/fgame/game/yuxi/player"
	yuxitemplate "fgame/fgame/game/yuxi/template"
	"fmt"
)

// 玉玺之战获胜
func WinYuXiWar(allianceId int64) {
	al := alliance.GetAllianceService().GetAlliance(allianceId)
	if al == nil {
		return
	}

	// 设置城战守方
	alliance.GetAllianceService().SetHegemonDefence(allianceId)
	//全服广播获胜仙盟信息
	plList := player.GetOnlinePlayerManager().GetAllPlayers()
	for _, pl := range plList {

		ctx := scene.WithPlayer(context.Background(), pl)
		msg := message.NewScheduleMessage(onSendYuXiInfo, ctx, nil, nil)
		pl.Post(msg)
	}

	// 公告
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.YuXiWinNoticeContent), al.GetAllianceName())
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)

	// 发获胜奖励
	constantTemp := yuxitemplate.GetYuXiTemplateService().GetYuXiConstTemplate()
	for _, mem := range al.GetMemberList() {
		pl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetMemberId())
		if pl == nil {
			//邮件
			createTime := global.GetGame().GetTimeService().Now()
			rewItemMap := constantTemp.GetWinItemMap()
			itemDataList := droptemplate.ConvertToItemDataListDefault(rewItemMap)
			emailTitle := lang.GetLangService().ReadLang(lang.YuXiWinTitle)
			emailContent := lang.GetLangService().ReadLang(lang.YuXiWinContent)
			emaillogic.AddOfflineEmailItemLevel(mem.GetMemberId(), emailTitle, emailContent, createTime, itemDataList)
		} else {
			ctx := scene.WithPlayer(context.Background(), pl)
			scheduleMsg := message.NewScheduleMessage(onWinYuXiWar, ctx, nil, nil)
			pl.Post(scheduleMsg)
		}
	}
}

// 玉玺之战邮件奖励
func onWinYuXiWar(ctx context.Context, result interface{}, err error) error {
	spl := scene.PlayerInContext(ctx)
	pl := spl.(player.Player)

	//写邮件
	createTime := global.GetGame().GetTimeService().Now()
	rewItemMap := yuxitemplate.GetYuXiTemplateService().GetYuXiConstTemplate().GetWinItemMap()
	itemDataList := droptemplate.ConvertToItemDataListDefault(rewItemMap)
	title := lang.GetLangService().ReadLang(lang.YuXiWinTitle)
	content := lang.GetLangService().ReadLang(lang.YuXiWinContent)
	emaillogic.AddEmailItemLevel(pl, title, content, createTime, itemDataList)
	return nil
}

func onSendYuXiInfo(ctx context.Context, result interface{}, err error) error {
	spl := scene.PlayerInContext(ctx)
	pl := spl.(player.Player)
	SendYuXiWarInfo(pl)
	return nil
}

func SendYuXiWarInfo(pl player.Player) {
	yuXiManager := pl.GetPlayerDataManager(playertypes.PlayerYuXiDataManagerType).(*playeryuxi.PlayerYuXiDataManager)
	yuXiManager.RefreshTimes()

	isReceive := yuXiManager.GetPlayerYuXiInfo().GetIsReceive()
	hegemon := alliance.GetAllianceService().GetAllianceHegemon()
	winAllianceId := hegemon.GetDefenceAllianceId()
	al := alliance.GetAllianceService().GetAlliance(winAllianceId)
	winAllianceName := ""
	winMengZhuName := ""
	if al != nil {
		winAllianceName = al.GetAllianceName()
		winMengZhuName = al.GetMengzhuName()
	}

	scMsg := pbutil.BuildSCYuXiGetInfo(isReceive, winAllianceId, winAllianceName, winMengZhuName)
	pl.SendMsg(scMsg)
}
