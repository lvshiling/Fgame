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

//装备强化
func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeClickEquipmentStrengthenButton, guaji.QuestGuaJiFunc(clickEquipmentStrengthenLevel))
}

//点击装备强化
func clickEquipmentStrengthenLevel(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {

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
	//TODO 验证是否可以升级
	num := q.QuestDataMap[0]
	needNum := demandMap[0]
	for i := 0; i < int(needNum-num); i++ {
		quest.ClickHandle(p, clicktypes.ClickTypeEquip, clicktypes.ClickSubTypeEquipStrength)
	}
	return true
}
