package quest_guaji

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/guaji/guaji"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"
	worldbosslogic "fgame/fgame/game/worldboss/logic"
	"fgame/fgame/game/worldboss/worldboss"

	log "github.com/Sirupsen/logrus"
)

//击杀世界boss
func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeWorldBoss, guaji.QuestGuaJiFunc(killWorldBoss))
}

//击杀世界boss
func killWorldBoss(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {

	demandMap := questTemplate.GetQuestDemandMap(p.GetRole())
	if len(demandMap) <= 0 {
		return true
	}
	//获取当前任务数据
	questManager := p.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	q := questManager.GetQuestById(int32(questTemplate.TemplateId()))
	needNum := demandMap[0]
	num := q.QuestDataMap[0]
	if num >= needNum {
		return true
	}
	//击杀世界boss
	flag := doKillWorldBoss(p)
	if !flag {
		return false
	}
	return true
}

//击杀世界boss
func doKillWorldBoss(pl player.Player) bool {
	worldBossList := worldboss.GetWorldBossService().GetGuaiJiWorldBossList(pl.GetForce())
	lenOfWorldBossList := len(worldBossList)
	if lenOfWorldBossList <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("quest_guaji:查找不到世界boss")
		return false
	}
	worldbosslogic.HandleKillWorldBoss(pl, int32(worldBossList[0].GetBiologyTemplate().TemplateId()))
	return true
}
