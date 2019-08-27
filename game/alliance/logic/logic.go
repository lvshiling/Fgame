package logic

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	activitylogic "fgame/fgame/game/activity/logic"
	activitytypes "fgame/fgame/game/activity/types"
	"fgame/fgame/game/alliance/alliance"
	"fgame/fgame/game/alliance/pbutil"
	playeralliance "fgame/fgame/game/alliance/player"
	alliancescene "fgame/fgame/game/alliance/scene"
	alliancetemplate "fgame/fgame/game/alliance/template"
	alliancetypes "fgame/fgame/game/alliance/types"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fgame/fgame/game/robot/robot"
	"fgame/fgame/game/scene/scene"
	skilllogic "fgame/fgame/game/skill/logic"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/game/transportation/transpotation"
	yuxitemplate "fgame/fgame/game/yuxi/template"
)

//获胜
func AllianceWin(allianceId int64) {
	alliance.GetAllianceService().AllianceWin(allianceId)
	al := alliance.GetAllianceService().GetAlliance(allianceId)

	//生成新的主城雕像
	LoadWinnerModel(al.GetAllianceMengZhuId())
	//全服广播获胜霸主
	plList := player.GetOnlinePlayerManager().GetAllPlayers()
	for _, pl := range plList {
		ctx := scene.WithPlayer(context.Background(), pl)
		msg := message.NewScheduleMessage(onHegemonChanged, ctx, nil, nil)
		pl.Post(msg)
	}
	//连胜奖励
	var sceneRewardOccupyTemplate *gametemplate.WarAwardOccupyTemplate
	hegemon := alliance.GetAllianceService().GetAllianceHegemon()
	if allianceId == hegemon.GetAllianceId() {
		sceneRewardOccupyTemplate = alliancetemplate.GetAllianceTemplateService().GetAllianceSceneOccupyReward(hegemon.GetWinNum())
	}

	for _, mem := range al.GetMemberList() {
		//获胜奖励
		sceneRewardType := alliancetypes.AllianceSceneRewardTypeMember
		if mem.IsMengZhu() {
			sceneRewardType = alliancetypes.AllianceSceneRewardTypeMengZhu
		}
		sceneRewardTemplate := alliancetemplate.GetAllianceTemplateService().GetAllianceSceneReward(sceneRewardType)

		pl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetMemberId())
		if pl == nil {
			//添加离线邮件
			winTitle := lang.GetLangService().ReadLang(lang.AllianceSceneWinTitle)
			winContent := lang.GetLangService().ReadLang(lang.AllianceSceneWinContent)
			emaillogic.AddOfflineEmail(mem.GetMemberId(), winTitle, winContent, sceneRewardTemplate.GetItemMap())
			if sceneRewardOccupyTemplate != nil {
				continueWinTitle := lang.GetLangService().ReadLang(lang.AllianceSceneContinueWinTitle)
				continueWinContent := lang.GetLangService().ReadLang(lang.AllianceSceneContinueWinContent)
				emaillogic.AddOfflineEmail(mem.GetMemberId(), continueWinTitle, continueWinContent, sceneRewardOccupyTemplate.GetItemMap())
			}

		} else {
			ctx := scene.WithPlayer(context.Background(), pl)
			pl.Post(message.NewScheduleMessage(onPlayerAllianceWin, ctx, sceneRewardOccupyTemplate, nil))
		}
	}
}

//发送霸主信息
func onHegemonChanged(ctx context.Context, result interface{}, err error) error {
	spl := scene.PlayerInContext(ctx)
	pl := spl.(player.Player)
	SendAllianceHegemonInfo(pl)
	return nil
}

//回调
func onPlayerAllianceWin(ctx context.Context, result interface{}, err error) error {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)
	mem := alliance.GetAllianceService().GetAllianceMember(tpl.GetId())
	if mem == nil {
		return nil
	}
	//连胜奖励
	sceneRewardOccupyTemplate := result.(*gametemplate.WarAwardOccupyTemplate)

	//获胜奖励
	sceneRewardType := alliancetypes.AllianceSceneRewardTypeMember
	if mem.IsMengZhu() {
		sceneRewardType = alliancetypes.AllianceSceneRewardTypeMengZhu
	}
	sceneRewardTemplate := alliancetemplate.GetAllianceTemplateService().GetAllianceSceneReward(sceneRewardType)

	//添加邮件
	winTitle := lang.GetLangService().ReadLang(lang.AllianceSceneWinTitle)
	winContent := lang.GetLangService().ReadLang(lang.AllianceSceneWinContent)
	emaillogic.AddEmail(tpl, winTitle, winContent, sceneRewardTemplate.GetItemMap())
	if sceneRewardOccupyTemplate != nil {
		continueWinTitle := lang.GetLangService().ReadLang(lang.AllianceSceneContinueWinTitle)
		continueWinContent := lang.GetLangService().ReadLang(lang.AllianceSceneContinueWinContent)
		emaillogic.AddEmail(tpl, continueWinTitle, continueWinContent, sceneRewardOccupyTemplate.GetItemMap())
	}
	//TODO胜利次数
	manager := tpl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	manager.AddWinTime()
	return nil
}

//推送腰牌变化
func SnapYaoPaiChanged(pl player.Player) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	yaoPai := manager.GetYaoPai()
	pl.SendMsg(pbutil.BuildSCAlliancePlayerYaoPaiChanged(yaoPai))
}

//斗神殿变化
func DoushenChanged(pl player.Player) {
	ctx := scene.WithPlayer(context.Background(), pl)
	doushenChangedMsg := message.NewScheduleMessage(onDoushenChanged, ctx, nil, nil)
	pl.Post(doushenChangedMsg)
}

//斗神殿回调
func onDoushenChanged(ctx context.Context, result interface{}, err error) error {

	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)

	pl.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeLingyuAura.Mask())

	return nil
}

// 虎符捐献
func DonateHufu(donor player.Player) {
	al := alliance.GetAllianceService().GetAlliance(donor.GetAllianceId())
	curHufu := al.GetAllianceObject().GetHuFu()
	scAllianceSceneHuFuChanged := pbutil.BuildSCAllianceSceneHuFuChanged(curHufu, donor.GetId(), donor.GetName())
	for _, member := range al.GetMemberList() {
		mpl := player.GetOnlinePlayerManager().GetPlayerById(member.GetMemberId())
		if mpl == nil {
			continue
		}

		ctx := scene.WithPlayer(context.Background(), mpl)
		hufuChangedMsg := message.NewScheduleMessage(onHuFuChanged, ctx, nil, nil)
		mpl.Post(hufuChangedMsg)
		mpl.SendMsg(scAllianceSceneHuFuChanged)
	}
}

//虎符改变回调
func onHuFuChanged(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)

	pl.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeHuFu.Mask())
	return nil
}

// 重新加载仙盟仙术
func ReloadAllianceSkill(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	skillMap := manager.GetPlayerAllianceSkillMap()
	allianceId := manager.GetAllianceId()
	for _, skillObj := range skillMap {
		temp := alliancetemplate.GetAllianceTemplateService().GetAllianceSkillTemplateByType(skillObj.GetLevel(), skillObj.GetSkillType())
		if temp == nil {
			continue
		}

		//卸下仙术
		if allianceId == 0 {
			err = skilllogic.TempSkillChange(pl, temp.SkillId, 0)
			continue
		}

		if !manager.IsOpenAllianceSkill(skillObj.GetLevel(), skillObj.GetSkillType()) {
			continue
		}

		err = skilllogic.TempSkillChange(pl, 0, temp.SkillId)
		if err != nil {
			return
		}
	}
	return
}

// 城门奖励
func DoorRewFinish(sd alliancescene.AllianceSceneData) {
	s := sd.GetScene()
	firstDefenceAllianceId := sd.GetFirstDefendAllianceId()
	curDoor := sd.GetCurrentDoor()
	for _, pl := range s.GetAllPlayers() {
		if pl.GetAllianceId() == firstDefenceAllianceId {
			continue
		}

		ctx := scene.WithPlayer(context.Background(), pl)
		msg := message.NewScheduleMessage(onDoorRew, ctx, curDoor, nil)
		pl.Post(msg)
	}
}

func onDoorRew(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	curDoor := result.(int32)

	allianceManager := pl.GetPlayerDataManager(playertypes.PlayerAllianceDataManagerType).(*playeralliance.PlayerAllianceDataManager)
	doorTempMap := alliancetemplate.GetAllianceTemplateService().GetDoorRewardTemplateMap()
	for _, doorTemp := range doorTempMap {
		door := int32(doorTemp.Id)
		if door > curDoor {
			continue
		}

		if !allianceManager.IfCanGetReward(door) {
			continue
		}

		if !allianceManager.IsEnoughWarPoint(doorTemp.NeedJiFen) {
			continue
		}

		doorRewTitle := lang.GetLangService().ReadLang(lang.AllianceSceneDoorRewTitle)
		doorRewContent := lang.GetLangService().ReadLang(lang.AllianceSceneDoorRewContent)
		emaillogic.AddEmail(pl, doorRewTitle, doorRewContent, doorTemp.GetEmailItemMap())
	}

	return nil
}

//更新主城盟主雕像
func LoadWinnerModel(mengZhuId int64) {
	constantTemp := yuxitemplate.GetYuXiTemplateService().GetYuXiConstTemplate()
	modelS := scene.GetSceneService().GetSceneByMapId(constantTemp.ModelMapId)
	ctx := scene.WithScene(context.Background(), modelS)
	msg := message.NewScheduleMessage(onAddWinnerModel, ctx, mengZhuId, nil)
	modelS.Post(msg)
}

// 场景生成获胜雕像
func onAddWinnerModel(ctx context.Context, result interface{}, err error) error {
	s := scene.SceneInContext(ctx)
	mengZhuId := result.(int64)

	// 移除原本雕像
	curModelList := alliance.GetAllianceService().GetCurModelList()
	for _, model := range curModelList {
		s.RemoveSceneObject(model, true)
		alliance.GetAllianceService().RemoveModel(model)
	}

	mengZhuInfo, _ := player.GetPlayerService().GetPlayerInfo(mengZhuId)
	if mengZhuInfo == nil {
		return nil
	}

	constantTemp := yuxitemplate.GetYuXiTemplateService().GetYuXiConstTemplate()
	winnerModel := robot.GetRobotService().CreateModelRobot(mengZhuInfo.ServerId, mengZhuInfo.Sex, mengZhuInfo.Role, mengZhuInfo.FashionId, mengZhuInfo.AllWeaponInfo.Wear, mengZhuInfo.WingInfo.GetWingId(), false)
	winnerModel.SetEnterPos(constantTemp.GetWinnerModelPos())
	s.AddSceneObject(winnerModel)
	alliance.GetAllianceService().AddModel(winnerModel)

	// 盟主配偶
	spuoseId := mengZhuInfo.MarryInfo.SpouseId
	if spuoseId != 0 {
		mengZhuSpouseInfo, _ := player.GetPlayerService().GetPlayerInfo(spuoseId)
		if mengZhuSpouseInfo == nil {
			return nil
		}
		winnerCoupleModel := robot.GetRobotService().CreateModelRobot(mengZhuSpouseInfo.ServerId, mengZhuSpouseInfo.Sex, mengZhuSpouseInfo.Role, mengZhuSpouseInfo.FashionId, mengZhuSpouseInfo.AllWeaponInfo.Wear, mengZhuSpouseInfo.WingInfo.GetWingId(), false)
		winnerCoupleModel.SetEnterPos(constantTemp.GetWinnerModelCouplePos())
		s.AddSceneObject(winnerCoupleModel)
		alliance.GetAllianceService().AddModel(winnerCoupleModel)
	}

	return nil
}

//推送城战霸主信息
func SendAllianceHegemonInfo(pl player.Player) {
	hegemon := alliance.GetAllianceService().GetAllianceHegemon()
	al := alliance.GetAllianceService().GetAlliance(hegemon.GetAllianceId())
	if al == nil {
		scAllianceHegemonInfo := pbutil.BuildSCAllianceHegemonInfo(0, "", 0, "", 0, 0, 0)
		pl.SendMsg(scAllianceHegemonInfo)
		return
	}

	winNum := hegemon.GetWinNum()
	allianceId := al.GetAllianceId()
	allianceName := al.GetAllianceObject().GetName()
	totalForce := al.GetAllianceObject().GetTotalForce()
	mengZhuId := al.GetAllianceObject().GetMengzhuId()
	mengZhuName := al.GetMengzhuName()
	mengzhuSex := al.GetMengzhuSex()

	scAllianceHegemonInfo := pbutil.BuildSCAllianceHegemonInfo(allianceId, allianceName, mengZhuId, mengZhuName, mengzhuSex, totalForce, winNum)
	pl.SendMsg(scAllianceHegemonInfo)
}

// 是否处于仙盟集体活动
func IfOnAllianceActivity(allianceId int64) (flag bool) {
	al := alliance.GetAllianceService().GetAlliance(allianceId)
	if al == nil {
		return true
	}

	// 仙盟屠龙、
	if activitylogic.IfActivityTime(activitytypes.ActivityTypeCoressTuLong) {
		flag = true
		return
	}

	// 玉玺争夺战、
	if activitylogic.IfActivityTime(activitytypes.ActivityTypeYuXi) {
		flag = true
		return
	}

	// 九霄城战、
	if activitylogic.IfActivityTime(activitytypes.ActivityTypeAlliance) {
		flag = true
		return
	}

	// 仙盟镖车、
	biaocheObj := transpotation.GetTransportService().GetTransportation(al.GetAllianceMengZhuId())
	if biaocheObj != nil {
		flag = true
		return
	}
	// 仙盟BOSS
	s := alliance.GetAllianceService().AllianceBossScene(allianceId)
	if s != nil {
		flag = true
		return
	}

	// 仙盟圣坛
	if activitylogic.IfActivityTime(activitytypes.ActivityTypeAllianceShengTan) {
		flag = true
		return
	}

	return
}
