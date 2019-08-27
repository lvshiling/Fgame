package boss_handler

import (
	"fgame/fgame/common/lang"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	unrealbosstemplate "fgame/fgame/game/unrealboss/template"
	"fgame/fgame/game/unrealboss/unrealboss"
	"fgame/fgame/game/worldboss/pbutil"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/game/worldboss/worldboss"

	log "github.com/Sirupsen/logrus"
)

func init() {
	worldboss.RegistKillBossHandler(worldbosstypes.BossTypeUnrealBoss, worldboss.KillBossHandlerFunc(killUnrealBoss))
}

func killUnrealBoss(pl player.Player, typ worldbosstypes.BossType, biologyId int32) {
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

	scMsg := pbutil.BuildSCChallengeWorldBoss(boss.GetBornPosition(), int32(worldbosstypes.BossTypeOutlandBoss))
	pl.SendMsg(scMsg)
}
