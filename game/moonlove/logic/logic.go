package logic

import (
	"context"
	"fgame/fgame/common/lang"
	commonlang "fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/common/message"
	coreutils "fgame/fgame/core/utils"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/moonlove/pbutil"
	playermoonlove "fgame/fgame/game/moonlove/player"
	moonlovetemplate "fgame/fgame/game/moonlove/template"
	moonlovetypes "fgame/fgame/game/moonlove/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func PlayerEnterMoonlove(pl player.Player, activityTemplate *gametemplate.ActivityTemplate) (flag bool, err error) {
	return PlayerEnterMoonloveArgs(pl, activityTemplate, "")
}

func PlayerEnterMoonloveArgs(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error) {
	s := getMoonLoveScene(activityTemplate)
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("moonlove:月下情缘场景不存在")
		playerlogic.SendSystemMessage(pl, lang.ActivityNotAtTime)
		return
	}
	bornPos := s.MapTemplate().GetBornPos()
	if !scenelogic.PlayerEnterScene(pl, s, bornPos) {
		return
	}
	flag = true
	return
}

//完成月下情缘活动
func onFinishMoonlove(pl player.Player, expCount int64) {
	//发送提示
	playerId := pl.GetId()
	scMoonloveSceneResult := pbutil.BuildMoonloveSceneResult(playerId, expCount)
	pl.SendMsg(scMoonloveSceneResult)
}

//推送场景信息
func onPushSceneInfo(pl player.Player, charmRank []*moonlovetypes.RankData, generousRank []*moonlovetypes.RankData, exp, startTime int64) {
	moonloveManager := pl.GetPlayerDataManager(types.PlayerMoonloveDataManagerType).(*playermoonlove.PlayerMoonloveDataManager)
	charmNum := moonloveManager.GetCharmNum()
	generousNum := moonloveManager.GetGenerousNum()

	scMoonloveSceneInfo := pbutil.BuildMoonloveSceneInfo(charmRank, generousRank, exp, startTime, charmNum, generousNum, pl.GetId())
	pl.SendMsg(scMoonloveSceneInfo)
}

//更新玩家月下情缘信息
func onScenePlayerEnter(pl player.Player, endTime int64) {
	moonloveManager := pl.GetPlayerDataManager(types.PlayerMoonloveDataManagerType).(*playermoonlove.PlayerMoonloveDataManager)
	moonloveManager.UpdateEnterTime(endTime)
}

//排行榜邮件回调
type rankCallBackData struct {
	title    string
	content  string
	rankTemp *gametemplate.MoonloveRankTemplate
}

//发送排行榜奖励
func onSendRankRewards(charmRank []*moonlovetypes.RankData, generousRank []*moonlovetypes.RankData) {
	for index, rankData := range charmRank {
		ranking := int32(index + 1)
		rankingText := fmt.Sprintf("第%d名", ranking)
		charmRankTmp := moonlovetemplate.GetMoonloveTemplateService().GetMoonloveCharmRankMap(ranking)
		mailName := commonlang.GetLangService().ReadLang(commonlang.EmailCharmName)
		mailContent := commonlang.GetLangService().ReadLang(commonlang.EmailCharmContent)
		mailContentF := fmt.Sprintf(mailContent, coreutils.FormatColor(moonlovetypes.ColorTypeEmailRanking, rankingText))

		pl := player.GetOnlinePlayerManager().GetPlayerById(rankData.PlayerId)
		if pl == nil {
			emaillogic.AddOfflineEmail(rankData.PlayerId, mailName, mailContentF, charmRankTmp.GetRewItemMap())
		} else {
			data := &rankCallBackData{}
			data.title = mailName
			data.content = mailContentF
			data.rankTemp = charmRankTmp
			ctx := scene.WithPlayer(context.Background(), pl)
			msg := message.NewScheduleMessage(playerAddRankRew, ctx, data, nil)
			pl.Post(msg)
		}
	}

	for index, rankData := range generousRank {
		ranking := int32(index + 1)
		rankingText := fmt.Sprintf("第%d名", ranking)
		generousRankTmp := moonlovetemplate.GetMoonloveTemplateService().GetMoonloveCharmRankMap(ranking)
		mailName := commonlang.GetLangService().ReadLang(commonlang.EmailGenerousName)
		mailContent := commonlang.GetLangService().ReadLang(commonlang.EmailGenerouseContent)
		mailContentF := fmt.Sprintf(mailContent, coreutils.FormatColor(moonlovetypes.ColorTypeEmailRanking, rankingText))

		pl := player.GetOnlinePlayerManager().GetPlayerById(rankData.PlayerId)
		if pl == nil {
			emaillogic.AddOfflineEmail(rankData.PlayerId, mailName, mailContentF, generousRankTmp.GetRewItemMap())
		} else {
			data := &rankCallBackData{}
			data.title = mailName
			data.content = mailContentF
			data.rankTemp = generousRankTmp
			ctx := scene.WithPlayer(context.Background(), pl)
			msg := message.NewScheduleMessage(playerAddRankRew, ctx, data, nil)
			pl.Post(msg)
		}
	}

	return
}

//邮件奖励回调
func playerAddRankRew(ctx context.Context, result interface{}, err error) error {
	tpl := scene.PlayerInContext(ctx)
	pl := tpl.(player.Player)
	data := result.(*rankCallBackData)

	rewData := addRankRew(pl, data.rankTemp)
	scRankRew := pbutil.BuildMoonloveRankRewards(rewData)
	pl.SendMsg(scRankRew)
	propertylogic.SnapChangedProperty(pl)

	emaillogic.AddEmail(pl, data.title, data.content, data.rankTemp.GetRewItemMap())
	return nil
}

//场景定时奖励
func onSceneTickRew(pl player.Player, isDouble bool) {
	moonloveTemplateMap := moonlovetemplate.GetMoonloveTemplateService().GetMoonloveMap()
	for _, moonTemplate := range moonloveTemplateMap {
		if pl.GetLevel() < moonTemplate.MinLev || pl.GetLevel() > moonTemplate.MaxLev {
			continue
		}

		moonloveManager := pl.GetPlayerDataManager(types.PlayerMoonloveDataManagerType).(*playermoonlove.PlayerMoonloveDataManager)
		now := global.GetGame().GetTimeService().Now()

		//判断时间间隔
		enterTime := moonloveManager.GetEnterTime()
		diffFirst := now - enterTime

		if moonloveManager.GetPreRewTime() == 0 {
			if diffFirst > moonTemplate.FristTiem {
				addSceneTickRew(pl, isDouble, moonTemplate)
			}
		} else {
			preRewTime := moonloveManager.GetPreRewTime()
			diffRew := now - preRewTime
			isRew := diffRew > moonTemplate.RewTiem
			if !isRew {
				return
			}
			addSceneTickRew(pl, isDouble, moonTemplate)
		}

		moonloveManager.UpdateRewTime()
	}
}

func addSceneTickRew(pl player.Player, isDouble bool, moonTemplate *gametemplate.MoonloveTemplate) {
	//添加资源
	addTickRew(pl, moonTemplate, isDouble)

	//添加物品
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(moonTemplate.GetRewItemMap()) > 0 {

		flag := inventoryManager.HasEnoughSlots(moonTemplate.GetRewItemMap())
		if !flag {
			mailName := commonlang.GetLangService().ReadLang(commonlang.EmailMoonloveTickName)
			mailContent := commonlang.GetLangService().ReadLang(commonlang.EmailMoonloveTickConten)
			emaillogic.AddEmail(pl, mailName, mailContent, moonTemplate.GetRewItemMap())

		} else {
			flag := inventoryManager.BatchAdd(moonTemplate.GetRewItemMap(), commonlog.InventoryLogReasonMoonloveTickRew, commonlog.InventoryLogReasonMoonloveTickRew.String())
			if !flag {
				panic("moonlove:moonlove tick rewards add item should be ok")
			}
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	propertylogic.SnapChangedProperty(pl)

	return
}

//获取排行榜奖励
func addRankRew(pl player.Player, template *gametemplate.MoonloveRankTemplate) *propertytypes.RewData {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	reasonGold := commonlog.GoldLogReasonMoonloveRankRew
	reasonSilver := commonlog.SilverLogReasonMoonloveRankRew
	reasonLevel := commonlog.LevelLogReasonMoonloveRankRew
	saodangGoldReasonText := fmt.Sprintf(reasonGold.String(), template.Rank, template.Type)
	saodangSilverReasonText := fmt.Sprintf(reasonSilver.String(), template.Rank, template.Type)
	expReasonText := fmt.Sprintf(reasonLevel.String(), template.Rank, template.Type)

	rewSilver := int32(template.RewSilver)
	rewBindGold := template.RewBindGold
	rewGold := template.RewGold
	rewExp := int32(template.RewExp)
	rewExpPoint := int32(template.RewExpPoint)
	totalRewData := propertytypes.CreateRewData(rewExp, rewExpPoint, rewSilver, rewGold, rewBindGold)

	flag := propertyManager.AddRewData(totalRewData, reasonGold, saodangGoldReasonText, reasonSilver, saodangSilverReasonText, reasonLevel, expReasonText)
	if !flag {
		panic("moonlove:moonlove ranking rewards add RewData should be ok")
	}

	return totalRewData

}

//获取定时奖励
func addTickRew(pl player.Player, template *gametemplate.MoonloveTemplate, isDouble bool) {
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	rate := int32(1)
	if isDouble {
		rate = template.DoubleMan
	}

	reasonGold := commonlog.GoldLogReasonMoonloveTickRew
	reasonSilver := commonlog.SilverLogReasonMoonloveTickRew
	reasonLevel := commonlog.LevelLogReasonMoonloveTickRew
	reasonGoldText := fmt.Sprintf(reasonGold.String())
	reasonSilverText := fmt.Sprintf(reasonSilver.String())
	reasonLevelText := fmt.Sprintf(reasonLevel.String())

	rewSilver := template.RewSilver * int64(rate)
	rewBindGold := template.RewBindGold * rate
	rewGold := template.RewGold * rate
	rewExp := template.RewExp * int64(rate)
	rewExpPoint := template.RewExpPoint * rate

	flag := propertyManager.AddMoney(int64(rewBindGold), int64(rewGold), reasonGold, reasonGoldText, rewSilver, reasonSilver, reasonSilverText)
	if !flag {
		panic("moonlove:moonlove ranking rewards add RewData should be ok")
	}
	propertyManager.AddExpAndExpPoint(rewExp, int64(rewExpPoint), reasonLevel, reasonLevelText)

	return

}

//挑战奖励信息更新
func onPushSceneExpInfo(pl player.Player, exp int64) {
	scMoonloveExpCountNotice := pbutil.BuildSCMoonloveExpCountNotice(exp)
	pl.SendMsg(scMoonloveExpCountNotice)
}

func pushCharmRankChanged(s scene.Scene, sortRanks []*moonlovetypes.RankData, afterIndex int) {
	for index := afterIndex + 1; index < len(sortRanks); index++ {
		rankData := sortRanks[index]
		playerId := rankData.PlayerId
		number := rankData.Number
		ranking := int32(index + 1)
		if ranking > MAX_RANKING {
			ranking = 0
		}

		//推送玩家排名变化
		pl := s.GetPlayer(playerId)
		if pl == nil {
			continue
		}
		scMoonloveCharmChanged := pbutil.BuildSCMoonloveCharmChanged(playerId, number, ranking)
		pl.SendMsg(scMoonloveCharmChanged)
	}
}

func pushGenerousChanged(s scene.Scene, sortRanks []*moonlovetypes.RankData, afterIndex int) {
	for index := afterIndex + 1; index < len(sortRanks); index++ {
		rankData := sortRanks[index]
		playerId := rankData.PlayerId
		number := rankData.Number
		ranking := int32(index + 1)
		if ranking > MAX_RANKING {
			ranking = 0
		}
		//推送玩家排名变化
		pl := s.GetPlayer(playerId)
		if pl == nil {
			continue
		}
		scMoonloveGenerousChanged := pbutil.BuildSCMoonloveGenerousChanged(playerId, number, ranking)
		pl.SendMsg(scMoonloveGenerousChanged)
	}
}
