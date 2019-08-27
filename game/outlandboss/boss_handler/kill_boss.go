package boss_handler

import (
	"fgame/fgame/common/lang"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/outlandboss/outlandboss"
	playeroutlandboss "fgame/fgame/game/outlandboss/player"
	outlandbosstemplate "fgame/fgame/game/outlandboss/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/game/worldboss/worldboss"

	log "github.com/Sirupsen/logrus"
)

func init() {
	worldboss.RegistKillBossHandler(worldbosstypes.BossTypeOutlandBoss, worldboss.KillBossHandlerFunc(killOutBoss))
}

func killOutBoss(pl player.Player, typ worldbosstypes.BossType, biologyId int32) {
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

	scMsg := pbutil.BuildSCChallengeWorldBoss(boss.GetBornPosition(), int32(worldbosstypes.BossTypeOutlandBoss))
	pl.SendMsg(scMsg)
}
