package logic

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/cangjingge/cangjingge"
	cangjinggetemplate "fgame/fgame/game/cangjingge/template"
	funcopentypes "fgame/fgame/game/funcopen/types"
	playerhuiyuan "fgame/fgame/game/huiyuan/player"
	huiyuantypes "fgame/fgame/game/huiyuan/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"

	log "github.com/Sirupsen/logrus"
)

func CheckPlayerIfCanCangJingGeBossChallenge(pl player.Player, biologyId int32) (flag bool) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeCangJingGe) {
		return
	}

	huiYuanManager := pl.GetPlayerDataManager(types.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	isHuiYuan := huiYuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypePlus)
	tempHuiYuan := huiYuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypeInterim)
	if !(isHuiYuan || tempHuiYuan) {
		return
	}

	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	bossTemp := cangjinggetemplate.GetCangJingGeTemplateService().GetCangJingGeTemplateByBiologyId(biologyId)
	if bossTemp == nil {
		return
	}

	s := scene.GetSceneService().GetBossSceneByMapId(bossTemp.MapId)
	if s == nil {
		return
	}
	flag = true
	return
}

func HandleKillCangJingGeBoss(pl player.Player, bossBiologyId int32) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeCangJingGe) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": bossBiologyId,
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

	bossTemp := cangjinggetemplate.GetCangJingGeTemplateService().GetCangJingGeTemplateByBiologyId(bossBiologyId)
	if bossTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": bossBiologyId,
			}).Warn("cangjingge:藏经阁boss挑战请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	boss := cangjingge.GetCangJingGeService().GetCangJingGeBoss(bossBiologyId)
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

	scMsg := pbutil.BuildSCChallengeWorldBoss(boss.GetBornPosition(), int32(worldbosstypes.BossTypeCangJingGe))
	pl.SendMsg(scMsg)
	return
}
