package listener

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/event"
	coreutils "fgame/fgame/core/utils"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	marrylogic "fgame/fgame/game/marry/logic"
	"fgame/fgame/game/marry/marry"
	pbuitl "fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marrytypes "fgame/fgame/game/marry/types"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fmt"
)

// //玩家预定婚礼
// func playerMarryWedding(target event.EventTarget, data event.EventData) (err error) {
// 	wedCardObj, ok := target.(*marry.MarryWedCardObject)
// 	if !ok {
// 		return
// 	}
// 	eventData, ok := data.(*marryeventtypes.MarryWedEventData)
// 	if !ok {
// 		return
// 	}
// 	period := eventData.GetPeriod()
// 	pl := eventData.GetPlayer()
// 	spouseId := eventData.GetSpouseId()
// 	spouseName := eventData.GetSpouseName()
// 	playerName := pl.GetName()

// 	excludeIdList := make([]int64, 0, 2)
// 	excludeIdList = append(excludeIdList, pl.GetId(), spouseId)

// 	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
// 	manager.DueToWedding()

// 	//推送婚烟信息
// 	scMarryInfoChangeWedding := pbuitl.BuildSCMarryInfoChangeWedding(int32(marrytypes.MarryStatusTypeEngagement), period)
// 	pl.SendMsg(scMarryInfoChangeWedding)

// 	spouseAllianceId := int64(0)
// 	spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
// 	if spl != nil {
// 		spouseAllianceId = spl.GetAllianceId()
// 		splCtx := scene.WithPlayer(context.Background(), spl)
// 		playerDueToWeddingMsg := message.NewScheduleMessage(onDueToWedding, splCtx, scMarryInfoChangeWedding, nil)
// 		spl.Post(playerDueToWeddingMsg)
// 	} else {
// 		playerInfo, _ := player.GetPlayerService().GetPlayerInfo(spouseId)
// 		spouseAllianceId = playerInfo.AllianceId
// 	}

// 	//系统频道
// 	name := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(playerName))
// 	peerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(spouseName))
// 	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryDueWedding), name, peerName, wedCardObj.HoldTime)
// 	noticelogic.NoticeNumBroadcast([]byte(content), marrytypes.WeddingInterval, 3)
// 	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))

// 	//好友和仙盟发送喜帖
// 	scMarryPushWedCard := pbuitl.BuildSCMarryPushWedCard(wedCardObj)
// 	marryPairMsgRelated := marrytypes.CreateMarryPairMsgRelated(pl.GetId(), spouseId, pl.GetAllianceId(), spouseAllianceId, marrytypes.MarryMsgRelateTypeWedCard, scMarryPushWedCard, nil)
// 	player.GetOnlinePlayerManager().BroadcastMsgRelated(marrylogic.OnPlayerMarryCeremonyMsg, marryPairMsgRelated)
// 	return
// }

// func onDueToWedding(ctx context.Context, result interface{}, err error) (terr error) {
// 	tpl := scene.PlayerInContext(ctx)
// 	pl := tpl.(player.Player)

// 	scMarryInfoChangeWedding := result.(*uipb.SCMarryInfoChange)
// 	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
// 	manager.DueToWedding()
// 	pl.SendMsg(scMarryInfoChangeWedding)
// 	return nil
// }

const (
	weddingNoticeTimes = 3
)

//玩家预定婚礼
func playerMarryWedding(target event.EventTarget, data event.EventData) (err error) {
	wedCardObj, ok := target.(*marry.MarryWedCardObject)
	if !ok {
		return
	}
	eventData, ok := data.(*marryeventtypes.MarryWedEventData)
	if !ok {
		return
	}
	spl := eventData.GetPlayer()
	playerId := eventData.GetPlayerId()
	playerName := eventData.GetPlayerName()
	holdTime := eventData.GetHoldTime()

	//推送婚烟信息
	onMarryWedSucess(spl, eventData)

	playerAllianceId := int64(0)
	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	if pl != nil {
		playerAllianceId = pl.GetAllianceId()
		plCtx := scene.WithPlayer(context.Background(), pl)
		playerDueToWeddingMsg := message.NewScheduleMessage(onDueToWedding, plCtx, eventData, nil)
		pl.Post(playerDueToWeddingMsg)
	} else {
		playerInfo, _ := player.GetPlayerService().GetPlayerInfo(playerId)
		playerAllianceId = playerInfo.AllianceId

		//发邮件
		emailTitle := lang.GetLangService().ReadLang(lang.MarryWedSuccessTitle)
		emailContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryWedSuccessContent), holdTime)
		emaillogic.AddOfflineEmail(playerId, emailTitle, emailContent, nil)
	}

	//系统频道
	name := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(playerName))
	peerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(spl.GetName()))
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryDueWedding), name, peerName, wedCardObj.HoldTime)
	noticelogic.NoticeNumBroadcast([]byte(content), marrytypes.WeddingInterval, weddingNoticeTimes)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))

	//好友和仙盟发送喜帖
	scMarryPushWedCard := pbuitl.BuildSCMarryPushWedCard(wedCardObj)
	marryPairMsgRelated := marrytypes.CreateMarryPairMsgRelated(playerId, spl.GetId(), playerAllianceId, spl.GetAllianceId(), marrytypes.MarryMsgRelateTypeWedCard, scMarryPushWedCard, nil)
	player.GetOnlinePlayerManager().BroadcastMsgRelated(marrylogic.OnPlayerMarryCeremonyMsg, marryPairMsgRelated)
	return
}

func onDueToWedding(ctx context.Context, result interface{}, err error) (terr error) {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	eventData := result.(*marryeventtypes.MarryWedEventData)
	onMarryWedSucess(pl, eventData)
	return nil
}

func onMarryWedSucess(pl player.Player, eventData *marryeventtypes.MarryWedEventData) {
	period := eventData.GetPeriod()
	holdTime := eventData.GetHoldTime()

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	manager.DueToWedding()
	marryInfo := manager.GetMarryInfo()
	marryStatus := marryInfo.Status
	scMarryInfoChangeWedding := pbuitl.BuildSCMarryInfoChangeWedding(int32(marryStatus), period)
	scMarrryWedSucess := pbuitl.BuildSCMarryWedSucess(period)

	// scMarryInfoChangeWedding := pbuitl.BuildSCMarryInfoChangeWedding(int32(marrytypes.MarryStatusTypeEngagement), period)
	pl.SendMsg(scMarryInfoChangeWedding)
	// scMarrryWedSucess := pbuitl.BuildSCMarryWedSucess(period)
	pl.SendMsg(scMarrryWedSucess)

	//写邮件
	emailTitle := lang.GetLangService().ReadLang(lang.MarryWedSuccessTitle)
	emailContent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryWedSuccessContent), holdTime)
	emaillogic.AddEmail(pl, emailTitle, emailContent, nil)
	return
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryWed, event.EventListenerFunc(playerMarryWedding))
}
