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
	marrynpc "fgame/fgame/game/marry/npc/hunche"
	pbuitl "fgame/fgame/game/marry/pbutil"
	playermarry "fgame/fgame/game/marry/player"
	marryscene "fgame/fgame/game/marry/scene"
	marrytemplate "fgame/fgame/game/marry/template"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	playerproperty "fgame/fgame/game/property/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/team/team"
	"fmt"
)

//婚礼开始
func marryWedStart(target event.EventTarget, data event.EventData) (err error) {
	eventData, ok := data.(*marryeventtypes.MarryWedStartEventData)
	if !ok {
		return
	}
	grade := eventData.GetGrade()
	hunCheGrade := eventData.GetHunCheGrade()
	sugarGrade := eventData.GetSugarGrade()

	playerId := eventData.GetPlayerId()
	spouseId := eventData.GetSpouseId()

	pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
	spl := player.GetOnlinePlayerManager().GetPlayerById(spouseId)
	if pl == nil || spl == nil {
		return
	}

	//是否正在3v3匹配 做取消匹配
	teamId := pl.GetTeamId()
	spouseTeamId := spl.GetTeamId()
	if teamId != 0 {
		teamDataObj := team.GetTeamService().GetTeam(teamId)
		if teamDataObj != nil && teamDataObj.IsMatch() {
			team.GetTeamService().LeaveTeam(pl)
			marrylogic.MarryStartAutoLeave(pl)
		}
	}

	if spouseTeamId != 0 {
		teamDataObj := team.GetTeamService().GetTeam(spouseTeamId)
		if teamDataObj != nil && teamDataObj.IsMatch() {
			team.GetTeamService().LeaveTeam(spl)
			marrylogic.MarryStartAutoLeave(pl)
		}
	}

	playerList := make([]player.Player, 0, 2)
	playerList = append(playerList, pl, spl)

	marryScene := marry.GetMarryService().GetScene()

	//改变个人婚宴状态
	for _, cpl := range playerList {
		cplCtx := scene.WithPlayer(context.Background(), cpl)
		playerHoldWeddingMsg := message.NewScheduleMessage(onPlayerHoldWedding, cplCtx, true, nil)
		cpl.Post(playerHoldWeddingMsg)
	}

	//zrc:特殊处理 预定婚礼不算元宝消费,等到举行婚礼才算
	_, costGold, _ := marrytemplate.GetMarryTemplateService().GetMarryGradeCost(marrytypes.MarryBanquetSubTypeWed(grade), marrytypes.MarryBanquetSubTypeHunChe(hunCheGrade), marrytypes.MarryBanquetSubTypeSugar(sugarGrade))
	plCtx := scene.WithPlayer(context.Background(), pl)
	//记录消费
	costGoldMsg := message.NewScheduleMessage(onPlayerCostGold, plCtx, costGold, nil)
	pl.Post(costGoldMsg)

	luxuryWedding(playerList, marryScene, eventData)

	//婚礼游街开始
	playerAllianceId := int64(0)
	playerInfo, _ := player.GetPlayerService().GetPlayerInfo(playerId)
	if playerInfo != nil {
		playerAllianceId = playerInfo.AllianceId
	}
	spouseAllianceId := int64(0)
	spousePlayerInfo, _ := player.GetPlayerService().GetPlayerInfo(spouseId)
	if spousePlayerInfo != nil {
		spouseAllianceId = spousePlayerInfo.AllianceId
	}

	sd := marry.GetMarryService().GetMarrySceneData()
	scMarryWedPushStatus := pbuitl.BuildSCMarryWedPushStatus(sd, true)
	marryPairMsgRelated := marrytypes.CreateMarryPushWedRelated(sd.Id, sd.PlayerId, sd.SpouseId, scMarryWedPushStatus, playerAllianceId, spouseAllianceId)
	player.GetOnlinePlayerManager().BroadcastMsgRelated(marrylogic.OnPlayerMarryPushWedHunChe, marryPairMsgRelated)
	//添加双方的结婚纪念次数
	wedType := marrytypes.MarryBanquetSubTypeWed(grade)
	if !eventData.GetIsFirst() {
		addPlayerJiNian(pl, wedType)
		addPlayerJiNian(spl, wedType)
	}
	return
}

func onPlayerCostGold(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	costGold := result.(int32)

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateCycleCostInfo(int64(costGold))
	gameevent.Emit(propertyeventtypes.EventTypePlayerGoldCost, pl, int64(costGold))
	gameevent.Emit(propertyeventtypes.EventTypePlayerGoldCostIncludeBind, pl, int64(costGold))
	return nil
}

func onPlayerHoldWedding(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	isCruise := result.(bool)
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	manager.Holdwedding(isCruise)

	scMarryInfoStatusChange := pbuitl.BuildSCMarryInfoStatusChange(int32(marrytypes.MarryStatusTypeMarried))
	pl.SendMsg(scMarryInfoStatusChange)
	manager.AddMarryCount()
	marryCount := manager.GetMarryCount()
	sendMarryCountChange(pl, marryCount)
	return nil
}

func luxuryWedding(playerList []player.Player, marryScene scene.Scene, eventData *marryeventtypes.MarryWedStartEventData) {
	marryConstTemplate := marrytemplate.GetMarryTemplateService().GetMarryConstTempalte()
	moveScene := scene.GetSceneService().GetWorldSceneByMapId(marryConstTemplate.CarMapId)
	//刷婚车
	ctx := scene.WithScene(context.Background(), moveScene)
	moveScene.Post(message.NewScheduleMessage(onRefreshMarryCar, ctx, eventData, nil))
	//刷酒席
	marryctx := scene.WithScene(context.Background(), marryScene)
	marryScene.Post(message.NewScheduleMessage(onRefreshMarryBanquet, marryctx, eventData, nil))

	//zrc: 修改不同场景操作引起的奔溃
	//传送玩家
	for _, cpl := range playerList {
		// curScene := cpl.GetScene()
		// if curScene != moveScene {
		cplCtx := scene.WithPlayer(context.Background(), cpl)
		playerEnterSceneMsg := message.NewScheduleMessage(onPlayerEnterHunCheScene, cplCtx, moveScene, nil)
		cpl.Post(playerEnterSceneMsg)
		// 	continue
		// }
		// //玩家移动到婚车
		// marryMoveTemplate := marrytemplate.GetMarryTemplateService().GetMarryMoveFirstTeamplate()
		// movePos := marryMoveTemplate.GetPos()
		// if curScene == moveScene && cpl.GetPosition() != movePos {
		// 	scenelogic.FixPosition(cpl, movePos)
		// }
	}
}

//中低档婚宴
func noluxuryWedding(playerList []player.Player, marryScene scene.Scene, eventData *marryeventtypes.MarryWedStartEventData) (err error) {
	ctx := scene.WithScene(context.Background(), marryScene)
	marryScene.Post(message.NewScheduleMessage(onRefreshMarryBanquet, ctx, eventData, nil))

	scMarryWorish := pbuitl.BuildSCMarryWorish()
	for _, cpl := range playerList {
		if cpl == nil {
			continue
		}
		//传送玩家
		curScene := cpl.GetScene()
		if curScene != marryScene {
			cplCtx := scene.WithPlayer(context.Background(), cpl)
			playerEnterSceneMsg := message.NewScheduleMessage(onPlayerEnterMarryScene, cplCtx, marryScene, nil)
			cpl.Post(playerEnterSceneMsg)
		}
		//通知拜堂
		cpl.SendMsg(scMarryWorish)
	}
	return
}

//刷酒席
func onRefreshMarryBanquet(ctx context.Context, result interface{}, err error) error {
	marryScene := marry.GetMarryService().GetScene()
	eventData, ok := result.(*marryeventtypes.MarryWedStartEventData)
	if !ok {
		return nil
	}
	grade := eventData.GetGrade()
	banquetType := marrytypes.MarryBanquetSubTypeWed(grade)
	banquetTemplate := marrytemplate.GetMarryTemplateService().GetMarryBanquetTeamplate(marrytypes.MarryBanquetTypeWed, banquetType)
	biologyTemplate := banquetTemplate.GetBiologyTemplate()

	//更新场景数据
	marryDelegate := marryScene.SceneDelegate()
	marryDelegate.(marryscene.MarrySceneData).OnWeddingBegin(eventData)
	//刷新酒席
	for _, pos := range banquetTemplate.GetPosList() {
		n := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, 0, int64(0), int32(0), biologyTemplate, pos, 0, 0)
		if n != nil {
			//设置场景
			marryScene.AddSceneObject(n)
		}
	}
	return nil
}

//刷婚车
func onRefreshMarryCar(ctx context.Context, result interface{}, err error) error {
	moveScene := scene.SceneInContext(ctx)
	eventData := result.(*marryeventtypes.MarryWedStartEventData)
	periodId := eventData.GetPeriod()
	ownerId := eventData.GetOwerId()
	playerId := eventData.GetPlayerId()
	spouseId := eventData.GetSpouseId()
	hunCheGrade := eventData.GetHunCheGrade()
	sugarGrade := eventData.GetSugarGrade()
	//婚车移动
	n := marrynpc.CreateHunCheNPC(periodId, scenetypes.OwnerTypeMarry, ownerId, 0, playerId, spouseId, hunCheGrade, sugarGrade)
	marry.GetMarryService().WeddingCarNPC(n)
	//设置场景
	moveScene.AddSceneObject(n)
	return nil
}

//玩家传送到婚车
func onPlayerEnterHunCheScene(ctx context.Context, result interface{}, err error) (terr error) {
	marryMoveTemplate := marrytemplate.GetMarryTemplateService().GetMarryMoveFirstTeamplate()
	movePos := marryMoveTemplate.GetPos()
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)
	moveScene := result.(scene.Scene)
	if pl.GetScene() == moveScene {
		scenelogic.FixPosition(pl, movePos)
	} else {
		scenelogic.PlayerEnterScene(tpl, moveScene, movePos)
	}
	return nil
}

//玩家传送到婚宴场地
func onPlayerEnterMarryScene(ctx context.Context, result interface{}, err error) (terr error) {
	marryTemplate := marrytemplate.GetMarryTemplateService().GetMarryConstTempalte()
	marryMapTemplate := marryTemplate.GetMarryMapTemplate()
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	marryScene := result.(scene.Scene)

	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	manager.EndCruise()
	scenelogic.PlayerEnterScene(pl, marryScene, marryMapTemplate.GetBornPos())
	return nil
}

//婚宴开始时候增加婚礼纪念的类型次数
func addPlayerJiNian(pl player.Player, wedType marrytypes.MarryBanquetSubTypeWed) error {
	splCtx := scene.WithPlayer(context.Background(), pl)
	msg := message.NewScheduleMessage(onAddMarryJiNian, splCtx, wedType, nil)
	pl.Post(msg)
	return nil
}

//婚宴增加结婚纪念次数
func onAddMarryJiNian(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	wedType := result.(marrytypes.MarryBanquetSubTypeWed)
	manager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	// inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	jiNianTemplate := marrytemplate.GetMarryTemplateService().GetMarryJiNianTemplate(wedType)
	if manager.IsCanAddJiNian(wedType, jiNianTemplate.NeedNum) { //如果可以发送纪念
		// reasonText := fmt.Sprintf(commonlog.InventoryLogReasonMarryJiNian.String(), int32(wedType))
		// flag := inventoryManager.BatchAdd(jiNianTemplate.GetItemMap(), commonlog.InventoryLogReasonMarryJiNian, reasonText)
		// if !flag {
		// 	panic(fmt.Errorf("addMarryJiNian should be ok!"))
		// }
		banquetTypeTemplate := marrytemplate.GetMarryTemplateService().GetMarryBanquetTeamplate(marrytypes.MarryBanquetTypeWed, wedType)
		banquetName := ""
		if banquetTypeTemplate != nil {
			banquetName = banquetTypeTemplate.Name
		}
		chengHaoItem := jiNianTemplate.GetItemMap()
		title := lang.GetLangService().ReadLang(lang.MarryJiNianChengHaoTitle)
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryJiNianChengHaoContent), jiNianTemplate.NeedNum, banquetName)
		emaillogic.AddEmail(pl, title, content, chengHaoItem)
	}
	manager.UpdateMarryJiNianCount(wedType, jiNianTemplate.NeedNum) //更新次数
	if manager.IfCanGetJiNianShiZhuang() {                          //可以领取,发送邮件
		// reasonText := commonlog.InventoryLogReasonMarryJiNianShiZhuang.String()
		jiNianSj := marrytemplate.GetMarryTemplateService().GetMarryJiNianSjItems()
		// flag := inventoryManager.BatchAdd(jiNianSj, commonlog.InventoryLogReasonMarryJiNianShiZhuang, reasonText)
		// if !flag {
		// 	panic(fmt.Errorf("addMarryJiNian ShiZhuang should be ok!"))
		// }
		title := lang.GetLangService().ReadLang(lang.MarryJiNianSendShiZhuangTitle)
		content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.MarryJiNianSendShiZhuangContent))
		emaillogic.AddEmail(pl, title, content, jiNianSj)
		manager.GetJiNianShiZhuang()
	}
	return nil
}

func sendMarryCountChange(pl player.Player, marryCount int32) {
	marryCoutMsg := pbuitl.BuildSCMarryMarryCountChange(pl.GetId(), marryCount)
	pl.SendMsg(marryCoutMsg)
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryWedStart, event.EventListenerFunc(marryWedStart))
}
