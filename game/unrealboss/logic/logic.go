package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/constant/constant"
	constantypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	funcopentypes "fgame/fgame/game/funcopen/types"
	mybosspbutil "fgame/fgame/game/myboss/pbutil"
	mybosstemplate "fgame/fgame/game/myboss/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/unrealboss/pbutil"
	playerunrealboss "fgame/fgame/game/unrealboss/player"
	unrealbosstemplate "fgame/fgame/game/unrealboss/template"
	"fgame/fgame/game/unrealboss/unrealboss"
	viplogic "fgame/fgame/game/vip/logic"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//幻境boss挑战逻辑
func CheckIfPlayerUnrealbossChallenge(pl player.Player, biologyId int32) (flag bool) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeBossHuanJing) {
		return
	}

	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	unrealbossTemp := unrealbosstemplate.GetUnrealBossTemplateService().GetUnrealBossTemplate(biologyId)
	if unrealbossTemp == nil {

		return
	}

	boss := unrealboss.GetUnrealBossService().GetUnrealBoss(biologyId)
	if boss == nil {
		return
	}

	if !pl.IsEnoughPilao(boss.GetBiologyTemplate().NeedPilao) {
		return
	}
	s := boss.GetScene()
	if s == nil {

		return
	}
	flag = true
	return
}

//幻境boss挑战逻辑
func HandleUnrealbossChallenge(pl player.Player, biologyId int32) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeBossHuanJing) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": biologyId,
			}).Warn("unrealboss:幻境boss挑战请求，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	unrealbossTemp := unrealbosstemplate.GetUnrealBossTemplateService().GetUnrealBossTemplate(biologyId)
	if unrealbossTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": biologyId,
			}).Warn("unrealboss:幻境boss挑战请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	boss := unrealboss.GetUnrealBossService().GetUnrealBoss(biologyId)
	if boss == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("unrealboss:boss不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	s := boss.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("unrealboss:场景不存在")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
		return
	}

	//进入场景
	if !scenelogic.PlayerEnterScene(pl, s, s.MapTemplate().GetBornPos()) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("unrealboss:场景进不去")
		return
	}

	scMsg := pbutil.BuildSCUnrealBossChallenge(boss.GetBornPosition())
	pl.SendMsg(scMsg)
	return
}

func PlayerEnterBoss(pl player.Player, biologyId int32) (flag bool) {
	mybossTemplate := mybosstemplate.GetMyBossTemplateService().GetMyBossTemplate(biologyId)
	if mybossTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": biologyId,
			}).Warn("myboss:处理跳转个人BOSS,boss不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	sh := createMyBossSceneData(pl.GetId(), mybossTemplate.BiologyId)
	s := scene.CreateFuBenScene(mybossTemplate.MapId, sh)
	if s == nil {
		panic(fmt.Errorf("myboss:创建副本应该成功"))
	}
	scenelogic.PlayerEnterSingleFuBenScene(pl, s)

	flag = true
	return
}

//挑战结束
func onMyBossFinish(p player.Player, itemList []*droptemplate.DropItemData, isSuccess bool) {
	scMsg := mybosspbutil.BuildSCMyBossChallengeResult(isSuccess, itemList)
	p.SendMsg(scMsg)
}

//下发场景信息
func onPushSceneInfo(p player.Player, startTime int64, bossId int32) {
	scMsg := mybosspbutil.BuildSCMyBossSceneInfo(startTime, bossId)
	p.SendMsg(scMsg)
}

const (
	limitLevel = 4
)

//幻境boss购买疲劳
func HandleUnrealbossBuy(pl player.Player, buyNum int32) (err error) {
	if pl.GetVip() < limitLevel {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"vip":      pl.GetVip(),
			}).Warn("unrealboss:幻境boss购买次数，VIP等级不足")
		playerlogic.SendSystemMessage(pl, lang.VipLevelToLow)
		return
	}

	// 购买次数
	maxBuyTimes := viplogic.GetUnrealBossMaxBuyTimes(pl)
	unrealManager := pl.GetPlayerDataManager(playertypes.PlayerUnrealBossDataManagerType).(*playerunrealboss.PlayerUnrealBossDataManager)
	curBuyTimes := unrealManager.GetPilaoBuyTimes()
	if curBuyTimes+buyNum > maxBuyTimes {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"buyNum":      buyNum,
				"curBuyTimes": curBuyTimes,
			}).Warn("unrealboss:幻境boss购买次数，已达最大购买次数")
		playerlogic.SendSystemMessage(pl, lang.UnrealBossBuyNumReachLimit)
		return
	}

	//元宝足够
	price := constant.GetConstantService().GetConstant(constantypes.ConstantTypePiLaoPrice)
	needGold := int64(price * buyNum)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if !propertyManager.HasEnoughGold(needGold, false) {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"buyNum":      buyNum,
				"curBuyTimes": curBuyTimes,
				"needGold":    needGold,
			}).Warn("unrealboss:幻境boss购买次数，元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	useReason := commonlog.GoldLogReasonUnrealBossBuyPilao
	flag := propertyManager.CostGold(needGold, false, useReason, useReason.String())
	if !flag {
		panic("unrealboss:购买疲劳，消耗元宝应该成功")
	}

	unrealManager.BuyPilaoNum(buyNum)
	propertylogic.SnapChangedProperty(pl)

	curPilao := unrealManager.GetCurPilaoNum()
	scMsg := pbutil.BuildSCUnrealBossBuyPilaoNum(curPilao)
	pl.SendMsg(scMsg)
	return
}

//幻境boss购买疲劳
func CheckIfPlayerUnrealbossBuy(pl player.Player, buyNum int32) (flag bool) {
	if pl.GetVip() < limitLevel {
		return false
	}

	// 购买次数
	maxBuyTimes := viplogic.GetUnrealBossMaxBuyTimes(pl)
	unrealManager := pl.GetPlayerDataManager(playertypes.PlayerUnrealBossDataManagerType).(*playerunrealboss.PlayerUnrealBossDataManager)
	curBuyTimes := unrealManager.GetPilaoBuyTimes()
	if curBuyTimes+buyNum > maxBuyTimes {
		return false
	}

	//元宝足够
	price := constant.GetConstantService().GetConstant(constantypes.ConstantTypePiLaoPrice)
	needGold := int64(price * buyNum)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if !propertyManager.HasEnoughGold(needGold, false) {
		return false
	}

	return true
}
