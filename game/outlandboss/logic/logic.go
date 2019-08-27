package logic

import (
	"fgame/fgame/common/lang"
	droptemplate "fgame/fgame/game/drop/template"
	funcopentypes "fgame/fgame/game/funcopen/types"
	mybosspbutil "fgame/fgame/game/myboss/pbutil"
	mybosstemplate "fgame/fgame/game/myboss/template"
	"fgame/fgame/game/outlandboss/outlandboss"
	"fgame/fgame/game/outlandboss/pbutil"

	playeroutlandboss "fgame/fgame/game/outlandboss/player"
	outlandbosstemplate "fgame/fgame/game/outlandboss/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func CheckPlayerIfCanOutlandbossChallenge(pl player.Player, biologyId int32) (flag bool) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeOutlandBoss) {
		return
	}

	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerOutlandBossDataManagerType).(*playeroutlandboss.PlayerOutlandBossDataManager)
	manager.RefreshZhuoQi()
	if pl.IsZhuoQiLimit() {
		return
	}

	outlandbossTemp := outlandbosstemplate.GetOutlandBossTemplateService().GetOutlandBossTemplate(biologyId)
	if outlandbossTemp == nil {
		return
	}

	boss := outlandboss.GetOutlandBossService().GetOutlandBoss(biologyId)
	if boss == nil {
		return
	}

	s := boss.GetScene()
	if s == nil {
		return
	}
	flag = true
	return
}

func HandleOutlandbossChallenge(pl player.Player, biologyId int32) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeOutlandBoss) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": biologyId,
			}).Warn("outlandboss:外域boss挑战请求，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerOutlandBossDataManagerType).(*playeroutlandboss.PlayerOutlandBossDataManager)
	manager.RefreshZhuoQi()
	if pl.IsZhuoQiLimit() {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": biologyId,
			}).Warn("outlandboss:外域boss挑战请求，浊气值上限")
		playerlogic.SendSystemMessage(pl, lang.OutlandBossZhuoQiNumNotEnough)
		return
	}

	outlandbossTemp := outlandbosstemplate.GetOutlandBossTemplateService().GetOutlandBossTemplate(biologyId)
	if outlandbossTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": biologyId,
			}).Warn("outlandboss:外域boss挑战请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	boss := outlandboss.GetOutlandBossService().GetOutlandBoss(biologyId)
	if boss == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("outlandboss:boss不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	s := boss.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("outlandboss:场景不存在")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
		return
	}

	//进入场景
	if !scenelogic.PlayerEnterScene(pl, s, s.MapTemplate().GetBornPos()) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("outlandboss:场景进不去")
		return
	}

	scMsg := pbutil.BuildSCOutlandBossChallenge(boss.GetBornPosition())
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
