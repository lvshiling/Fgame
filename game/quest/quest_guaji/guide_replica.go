package quest_guaji

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/guaji/guaji"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

//任务引导
func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeGuideReplica, guaji.QuestGuaJiFunc(guideReplica))
}

//任务引导
func guideReplica(pl player.Player, questTemplate *gametemplate.QuestTemplate) bool {

	demandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	if len(demandMap) <= 0 {
		return true
	}
	questId := int32(questTemplate.TemplateId())
	//获取当前任务数据
	questManager := pl.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	q := questManager.GetQuestById(questId)
	if q == nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"questType": questTemplate.GetQuestSubType().String(),
			}).Warn("quest_guaji:任务不存在")
		return false
	}
	if q.QuestState != questtypes.QuestStateAccept {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"questId":   questId,
				"questType": questTemplate.GetQuestSubType().String(),
			}).Warn("quest_guaji:任务没在进行中")
		return false
	}
	for demandId, demandNum := range questTemplate.GetQuestDemandMap(pl.GetRole()) {
		questManager.SetQuestData(questId, demandId, demandNum)
	}

	return true
}
