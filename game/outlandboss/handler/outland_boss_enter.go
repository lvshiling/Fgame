package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/outlandboss/outlandboss"
	"fgame/fgame/game/outlandboss/pbutil"
	playeroutlandboss "fgame/fgame/game/outlandboss/player"
	outlandbosstemplate "fgame/fgame/game/outlandboss/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OUTLAND_BOSS_CHALLENGE_TYPE), dispatch.HandlerFunc(handlerOutlandBossChallenge))
}

//外域boss挑战
func handlerOutlandBossChallenge(s session.Session, msg interface{}) (err error) {
	log.Debug("outlandboss:处理外域boss挑战请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOutlandBossChallenge)
	biologyId := csMsg.GetBiologyId()

	err = outlandbossChallenge(tpl, biologyId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  tpl.GetId(),
				"biologyId": biologyId,
				"err":       err,
			}).Error("outlandboss:处理外域boss挑战请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  tpl.GetId(),
			"biologyId": biologyId,
		}).Debug("outlandboss：处理外域boss挑战请求完成")

	return
}

//外域boss挑战逻辑
func outlandbossChallenge(pl player.Player, biologyId int32) (err error) {
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
