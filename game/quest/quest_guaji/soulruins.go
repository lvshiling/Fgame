package quest_guaji

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/guaji/guaji"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	soulruinslogic "fgame/fgame/game/soulruins/logic"
	"fgame/fgame/game/soulruins/soulruins"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

//通关指定帝魂副本X次
func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeSpecifiedSoulRuins, guaji.QuestGuaJiFunc(specifiedSoulRuins))
}

//通关指定帝魂副本X次
func specifiedSoulRuins(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {

	demandMap := questTemplate.GetQuestDemandMap(p.GetRole())
	if len(demandMap) <= 0 {
		return true
	}
	//获取当前任务数据
	questManager := p.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	q := questManager.GetQuestById(int32(questTemplate.TemplateId()))
	if q == nil {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Warn("quest_guaji:任务不存在")
		return false
	}
	if q.QuestState != questtypes.QuestStateAccept {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Warn("quest_guaji:任务没在进行中")
		return false
	}
	for k, v := range demandMap {
		soulruinsTemplate := soulruins.GetSoulRuinsService().GetSoulRuinsTemplateById(k)
		if soulruinsTemplate == nil {
			continue
		}
		num := q.QuestDataMap[k]
		if num >= v {
			continue
		}
		//进入帝魂挑战
		soulruinslogic.SoulRuinsChallenge(p, soulruinsTemplate.Chapter, soulruinsTemplate.GetType(), soulruinsTemplate.Level)
		break
	}
	return true
}
