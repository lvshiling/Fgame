package logic

import (
	guidereplica "fgame/fgame/game/guidereplica/guidereplica"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func PlayerEnterQuestGuideReplica(pl player.Player, questId int32) (flag bool) {
	if !playerlogic.CheckCanEnterScene(pl) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
			}).Warn("guidereplica:引导副本挑战请求，不能进入副本")
		return
	}

	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
			}).Warn("guidereplica:引导副本挑战请求，任务不存在")
		return false
	}
	guideReplicaTemplate := questTemplate.GetGuideReplicaTemplate()
	if guideReplicaTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
			}).Warn("guidereplica:引导副本挑战请求，引导副本不存在")
		return
	}
	//身上是否有要求的任务
	questManager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	questObj := questManager.GetQuestById(questId)
	if questObj == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
			}).Warn("guidereplica:引导副本挑战请求，任务不存在")
		return false
	}

	if questObj.QuestState != questtypes.QuestStateAccept {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"questId":  questId,
			}).Warn("guidereplica:引导副本挑战请求，任务不是接受状态")
		return false
	}

	sh := guidereplica.GetGuideReplicaSd(pl, guideReplicaTemplate, questId)
	if sh == nil {
		panic(fmt.Errorf("guidereplica:创建副本sh应该成功"))
	}
	s := scene.CreateFuBenScene(guideReplicaTemplate.MapId, sh)
	if s == nil {
		panic(fmt.Errorf("guidereplica:创建副本应该成功"))
	}
	scenelogic.PlayerEnterSingleFuBenScene(pl, s)

	flag = true
	return
}

func PlayerEnterGuideReplica(pl player.Player, temp *gametemplate.GuideReplicaTemplate, questId int32) (flag bool) {
	sh := guidereplica.GetGuideReplicaSd(pl, temp, questId)
	if sh == nil {
		panic(fmt.Errorf("guidereplica:创建副本sh应该成功"))
	}
	s := scene.CreateFuBenScene(temp.MapId, sh)
	if s == nil {
		panic(fmt.Errorf("guidereplica:创建副本应该成功"))
	}
	scenelogic.PlayerEnterSingleFuBenScene(pl, s)

	flag = true
	return
}
