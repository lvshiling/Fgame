package quest_guaji

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/quest/guaji/guaji"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	_ "fgame/fgame/game/quest/quest_guaji/system"

	log "github.com/Sirupsen/logrus"
)

func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeSystemX, guaji.QuestGuaJiFunc(systemX))
}

//系统升阶
func systemX(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {

	guaJiSuccess := false
	demandMap := questTemplate.GetQuestDemandMap(p.GetRole())
	for k, _ := range demandMap {
		reachXType := questtypes.SystemReachXType(k)
		h := guaji.GetQuestGuaJiSystemX(reachXType)
		if h == nil {
			log.WithFields(
				log.Fields{
					"playerId":   p.GetId(),
					"questType":  questTemplate.GetQuestSubType().String(),
					"reachXType": reachXType,
				}).Warn("quest_guaji:系统进阶任务处理器不存在")
			continue
		}
		flag := h.GuaJiSystemX(p, questTemplate)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId":   p.GetId(),
					"questType":  questTemplate.GetQuestSubType().String(),
					"reachXType": reachXType,
				}).Warn("quest_guaji:系统进阶任务挂机失败")
			continue
		}
		guaJiSuccess = true
	}
	return guaJiSuccess
}
