package logic

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/worldboss/pbutil"
	worldbosstemplate "fgame/fgame/game/worldboss/template"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fgame/fgame/game/worldboss/worldboss"

	log "github.com/Sirupsen/logrus"
)

func CheckPlayerIfCanKillBoss(pl player.Player, bossBiologyId int32) (flag bool) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	bossTemp := worldbosstemplate.GetWorldBossTemplateService().GetWorldBossTemplateByBiologyId(bossBiologyId)
	if bossTemp == nil {
		return false
	}

	s := scene.GetSceneService().GetBossSceneByMapId(bossTemp.MapId)
	if s == nil {
		return false
	}

	return true
}

func HandleKillWorldBoss(pl player.Player, bossBiologyId int32) (err error) {
	if !playerlogic.CheckCanEnterScene(pl) {
		return
	}

	bossTemp := worldbosstemplate.GetWorldBossTemplateService().GetWorldBossTemplateByBiologyId(bossBiologyId)
	if bossTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("worldBoss:boss不存在")

		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	s := scene.GetSceneService().GetBossSceneByMapId(bossTemp.MapId)
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("worldBoss:场景不存在")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
		return
	}

	if !scenelogic.PlayerEnterScene(pl, s, s.MapTemplate().GetBornPos()) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("worldBoss:场景进不去")
		return
	}

	bossId := int32(bossTemp.GetBiologyTemplate().TemplateId())
	boss := worldboss.GetWorldBossService().GetWorldBoss(bossId)
	scMsg := pbutil.BuildSCChallengeWorldBoss(boss.GetBornPosition(), int32(worldbosstypes.BossTypeOutlandBoss))
	pl.SendMsg(scMsg)

	return
}
