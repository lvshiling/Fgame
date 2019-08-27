package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/cangjingge/cangjingge"
	"fgame/fgame/game/cangjingge/pbutil"
	cangjinggetemplate "fgame/fgame/game/cangjingge/template"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_CANGJINGGE_BOSS_CHALLENGE_TYPE), dispatch.HandlerFunc(handlerCangJingGeBossChallenge))
}

//藏经阁boss挑战
func handlerCangJingGeBossChallenge(s session.Session, msg interface{}) (err error) {
	log.Debug("cangjingge:处理藏经阁boss挑战请求")

	pl := gamesession.SessionInContext(s.Context()).Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSCangjinggeBossChallenge)
	biologyId := csMsg.GetBiologyId()

	err = cangJingGeBossChallenge(tpl, biologyId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  tpl.GetId(),
				"biologyId": biologyId,
				"err":       err,
			}).Error("cangjingge:处理藏经阁boss挑战请求，错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  tpl.GetId(),
			"biologyId": biologyId,
		}).Debug("outlandboss：处理藏经阁boss挑战请求完成")

	return
}

//藏经阁boss挑战逻辑
func cangJingGeBossChallenge(pl player.Player, biologyId int32) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeCangJingGe) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": biologyId,
			}).Warn("cangjingge:藏经阁boss挑战请求，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	// huiYuanManager := pl.GetPlayerDataManager(types.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	// isHuiYuan := huiYuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypePlus)
	// tempHuiYuan := huiYuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypeInterim)
	// if !isHuiYuan && !tempHuiYuan {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId":  pl.GetId(),
	// 			"biologyId": biologyId,
	// 		}).Warn("cangjingge:藏经阁boss挑战请求，不是至尊会员")
	// 	playerlogic.SendSystemMessage(pl, lang.HuiYuanNotHuiYuan)
	// 	return
	// }

	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	bossTemp := cangjinggetemplate.GetCangJingGeTemplateService().GetCangJingGeTemplateByBiologyId(biologyId)
	if bossTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": biologyId,
			}).Warn("cangjingge:藏经阁boss挑战请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	boss := cangjingge.GetCangJingGeService().GetCangJingGeBoss(biologyId)
	if boss == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("cangjingge:boss不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	s := boss.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("cangjingge:场景不存在")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
		return
	}

	//进入场景
	if !scenelogic.PlayerEnterScene(pl, s, s.MapTemplate().GetBornPos()) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("cangjingge:场景进不去")
		return
	}

	scMsg := pbutil.BuildSCCangJingGeBossChallenge(boss.GetBornPosition())
	pl.SendMsg(scMsg)
	return
}
