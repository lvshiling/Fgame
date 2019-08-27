package quest_guaji

import (
	clicktypes "fgame/fgame/game/click/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/quest/guaji/guaji"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

//系统升级点击
func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeUpgradeSysOperation, guaji.QuestGuaJiFunc(clickUpgradeSysOperation))
}

//系统升级点击
func clickUpgradeSysOperation(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {

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
				"playerId":  p.GetId(),
				"questType": questTemplate.GetQuestSubType().String(),
			}).Warn("quest_guaji:任务不存在")
		return false
	}
	if q.QuestState != questtypes.QuestStateAccept {
		log.WithFields(
			log.Fields{
				"playerId":  p.GetId(),
				"questType": questTemplate.GetQuestSubType().String(),
			}).Warn("quest_guaji:任务没在进行中")
		return false
	}

	for k, v := range demandMap {
		clickSubTypeUpgradeSys := clicktypes.ClickSubTypeUpgradeSys(k)
		if !clickSubTypeUpgradeSys.Valid() {
			log.WithFields(
				log.Fields{
					"playerId":               p.GetId(),
					"clickSubTypeUpgradeSys": k,
				}).Warn("quest_guaji:升级类型无效")
			return false
		}

		num := q.QuestDataMap[k]
		if num >= v {
			continue
		}
		needNum := v - num
		for i := 0; i < int(needNum); i++ {
			quest.ClickHandle(p, clicktypes.ClickTypeUpgradeSys, clickSubTypeUpgradeSys)
		}
	}
	return true
}