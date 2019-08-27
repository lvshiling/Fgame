package logic

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/dingshi/dingshi"
	dingshitemplate "fgame/fgame/game/dingshi/template"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"

	log "github.com/Sirupsen/logrus"
)

func CheckPlayerIfCanDingShiBossChallenge(pl player.Player, biologyId int32) (flag bool) {
	// if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeCangJingGe) {
	// 	return
	// }

	// huiYuanManager := pl.GetPlayerDataManager(types.PlayerHuiYuanDataManagerType).(*playerhuiyuan.PlayerHuiYuanManager)
	// isHuiYuan := huiYuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypePlus)
	// tempHuiYuan := huiYuanManager.IsHuiYuan(huiyuantypes.HuiYuanTypeInterim)
	// if !(isHuiYuan || tempHuiYuan) {
	// 	return
	// }

	// if !playerlogic.CheckCanEnterScene(pl) {
	// 	return
	// }

	// bossTemp := cangjinggetemplate.GetCangJingGeTemplateService().GetCangJingGeTemplateByBiologyId(biologyId)
	// if bossTemp == nil {
	// 	return
	// }

	// s := scene.GetSceneService().GetCangJingGeSceneByMapId(bossTemp.MapId)
	// if s == nil {
	// 	return
	// }
	// flag = true
	return
}

func HandleKillDingShiBoss(pl player.Player, bossBiologyId int32) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeDingShiBoss) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": bossBiologyId,
			}).Warn("dingshi:定时boss挑战请求，功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	bossTemp := dingshitemplate.GetDingShiTemplateService().GetDingShiBossTemplateByBiologyId(bossBiologyId)
	if bossTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"biologyId": bossBiologyId,
			}).Warn("dingshi:boss挑战请求，模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	boss := dingshi.GetDingShiService().GetDingShiBoss(bossBiologyId)
	if boss == nil {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"bossBiologyId": bossBiologyId,
			}).Warn("dingshi:boss不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	s := boss.GetScene()
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("dingshi:场景不存在")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
		return
	}

	//进入场景
	if !scenelogic.PlayerEnterScene(pl, s, s.MapTemplate().GetBornPos()) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("dingshi:场景进不去")
		return
	}

	scMsg := pbutil.BuildSCChallengeWorldBoss(boss.GetBornPosition(), int32(worldbosstypes.BossTypeDingShi))
	pl.SendMsg(scMsg)
	return
}
