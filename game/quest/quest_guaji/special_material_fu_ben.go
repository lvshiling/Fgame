package quest_guaji

import (
	materiallogic "fgame/fgame/game/material/logic"
	materialtypes "fgame/fgame/game/material/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/guaji/guaji"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

//进入指定的材料副本
func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypechallengeSpecialMaterialFuBen, guaji.QuestGuaJiFunc(specialMaterialFuBen))
}

//进入指定的材料副本
func specialMaterialFuBen(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {

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
		materialType := materialtypes.MaterialType(k)
		if !materialType.Valid() {
			log.WithFields(
				log.Fields{
					"playerId": p.GetId(),
					"questId":  questTemplate.TemplateId(),
				}).Warn("quest_guaji:材料副本类型无效")
			return false
		}

		num := q.QuestDataMap[k]
		if num >= v {
			continue
		}
		//进入材料副本挑战
		materiallogic.HandlePlayerMaterialChallenge(p, materialType)
		break
	}
	return true
}
