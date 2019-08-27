package logic

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/team/team"
	"fgame/fgame/game/teamcopy/pbutil"
	"fmt"

	commonlog "fgame/fgame/common/log"
	crosslogic "fgame/fgame/game/cross/logic"
	crosstypes "fgame/fgame/game/cross/types"
	emaillogic "fgame/fgame/game/email/logic"
	gameevent "fgame/fgame/game/event"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	teamtypes "fgame/fgame/game/team/types"
	teamcopyeventtypes "fgame/fgame/game/teamcopy/event/types"
	teamcopytemplate "fgame/fgame/game/teamcopy/template"

	log "github.com/Sirupsen/logrus"
)

//组队副本队伍开始战斗
func OnTeamCopyStartBattle(pl player.Player, teamObject *team.TeamObject) {
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"teamId":   teamObject.GetTeamId(),
		}).Info("teamcopy:队伍开始战斗")

	//进入跨服
	crosslogic.PlayerEnterCross(pl, crosstypes.CrossTypeTeamCopy)
	scTeamCopyStartBattleBroadcast := pbutil.BuildSCTeamCopyStartBattleBroadcast(int32(pl.GetTeamPurpose()))
	//广播消息
	for _, mem := range teamObject.GetMemberList() {
		memPl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
		if memPl == nil {
			continue
		}
		memPl.SendMsg(scTeamCopyStartBattleBroadcast)
	}
}

//组队副本发送队伍信息
func TeamCopyStartBattleSend(pl player.Player) {
	teamData := team.GetTeamService().GetTeam(pl.GetTeamId())
	if teamData == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Info("teamcopy:队伍不存在")
		//退出竞技
		crosslogic.PlayerExitCross(pl)
		return
	}
	pList := make([]player.Player, 0, len(teamData.GetMemberList()))
	for _, mem := range teamData.GetMemberList() {
		p := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
		if p == nil {
			continue
		}
		pList = append(pList, p)
	}
	siTeamCopyStartBattle := pbutil.BuildSITeamCopyStartBattle(pList)
	pl.SendCrossMsg(siTeamCopyStartBattle)
}

//组队副本开始战斗失败
func OnTeamCopyStartBattleFailed(pl player.Player, teamObject *team.TeamObject) {
	log.WithFields(
		log.Fields{
			"teamId":      teamObject.GetTeamId(),
			"teamPurpose": teamObject.GetTeamPurpose(),
		}).Info("teamcopy:组队副本开始战斗失败")
	//开始战斗失败
	purpose := int32(pl.GetTeamPurpose())
	scResult := pbutil.BuildSCTeamCopyStartBattleResultBroadcast(purpose, false)
	for _, mem := range teamObject.GetMemberList() {
		memPl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
		if memPl == nil {
			continue
		}
		memPl.SendMsg(scResult)
	}
}

//组队副本开始战斗成功
func OnTeamCopyStartBattleSucess(pl player.Player, teamObject *team.TeamObject) {
	log.WithFields(
		log.Fields{
			"teamId":      teamObject.GetTeamId(),
			"teamPurpose": teamObject.GetTeamPurpose(),
		}).Info("teamcopy:组队副本开始战斗成功")

	purpose := int32(pl.GetTeamPurpose())
	scResult := pbutil.BuildSCTeamCopyStartBattleResultBroadcast(purpose, true)
	pl.SendMsg(scResult)
	//成功
	for _, mem := range teamObject.GetMemberList() {
		tpl := player.GetOnlinePlayerManager().GetPlayerById(mem.GetPlayerId())
		if tpl == nil {
			continue
		}
		if tpl == pl {
			continue
		}
		ctx := scene.WithPlayer(context.Background(), tpl)
		scheduleMsg := message.NewScheduleMessage(onTeamCopyBattle, ctx, nil, nil)
		tpl.Post(scheduleMsg)
	}

	ctx := scene.WithPlayer(context.Background(), pl)
	scheduleMsg := message.NewScheduleMessage(onTeamCaptainTeamCopyBattle, ctx, nil, nil)
	pl.Post(scheduleMsg)

}

func onTeamCopyBattle(ctx context.Context, result interface{}, err error) error {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)
	purpose := pl.GetTeamPurpose()
	scResult := pbutil.BuildSCTeamCopyStartBattleResultBroadcast(int32(purpose), true)
	pl.SendMsg(scResult)

	//进入跨服
	crosslogic.PlayerEnterCross(tpl, crosstypes.CrossTypeTeamCopy)

	//参加组队副本
	gameevent.Emit(teamcopyeventtypes.EventTypeTeamCopyAttend, tpl, purpose)
	return nil
}

func onTeamCaptainTeamCopyBattle(ctx context.Context, result interface{}, err error) error {
	p := scene.PlayerInContext(ctx)
	pl := p.(player.Player)

	purpose := pl.GetTeamPurpose()
	crosslogic.CrossPlayerDataLogin(pl)

	//参加组队副本
	gameevent.Emit(teamcopyeventtypes.EventTypeTeamCopyAttend, pl, purpose)
	return nil
}

//奖励
func GiveRewardTeamCopyPurpose(pl player.Player, purpose teamtypes.TeamPurposeType) (err error) {
	teamCopyTemaplate := teamcopytemplate.GetTeamCopyTemplateService().GetTeamCopyTempalte(purpose)
	if teamCopyTemaplate == nil {
		return
	}
	rewData := teamCopyTemaplate.GetRewData()
	rewItemMap := teamCopyTemaplate.GetRewItemMap()

	if len(rewItemMap) != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		flag := inventoryManager.HasEnoughSlots(rewItemMap)
		if !flag {
			//TODO 发送邮件
			emailTitle := lang.GetLangService().ReadLang(lang.TeamCopyTitle)
			emailContent := lang.GetLangService().ReadLang(lang.TeamCopyContent)
			emaillogic.AddEmail(pl, emailTitle, emailContent, rewItemMap)
		} else {
			inventoryLogReason := commonlog.InventoryLogReasonTeamCopyReward
			reasonText := fmt.Sprintf(commonlog.InventoryLogReasonTeamCopyReward.String(), purpose)
			flag = inventoryManager.BatchAdd(rewItemMap, inventoryLogReason, reasonText)
			if !flag {
				panic(fmt.Errorf("teamcopy: commit  BatchAdd should be ok"))
			}
			//同步物品
			inventorylogic.SnapInventoryChanged(pl)
		}
	}

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if rewData != nil {
		reasonGold := commonlog.GoldLogReasonTeamCopyReward
		reasonSilver := commonlog.SilverLogResonTeamCopyReward
		reasonLevel := commonlog.LevelLogReasonTeamCopyReward
		reasonGoldText := fmt.Sprintf(reasonGold.String(), purpose)
		reasonSliverText := fmt.Sprintf(reasonSilver.String(), purpose)
		reasonlevelText := fmt.Sprintf(reasonLevel.String(), purpose)
		flag := propertyManager.AddRewData(rewData, reasonGold, reasonGoldText, reasonSilver, reasonSliverText, reasonLevel, reasonlevelText)
		if !flag {
			panic(fmt.Errorf("teamcopy:  AddRewData  should be ok"))
		}
	}
	propertylogic.SnapChangedProperty(pl)
	return
}
