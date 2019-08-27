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

//元神装备强化
func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeGoldEquipmentStrength, guaji.QuestGuaJiFunc(clickGoldEquipmentStrength))
}

const (
	goldEquipmentAll = 8
)

//点击元神装备强化
func clickGoldEquipmentStrength(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {

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
		if k == goldEquipmentAll {
			num := q.QuestDataMap[k]
			if num >= v {
				continue
			}
			needNum := v - num
			for i := 0; i < int(needNum); i++ {
				quest.ClickHandle(p, clicktypes.ClickTypeGoldEquipment, clicktypes.ClickSubTypeGoldEquipmentWeapon)
			}
			continue
		}
		clickSubTypeGoldEquipment := clicktypes.ClickSubTypeGoldEquipment(k)
		if !clickSubTypeGoldEquipment.Valid() {
			log.WithFields(
				log.Fields{
					"playerId":                  p.GetId(),
					"clickSubTypeGoldEquipment": k,
				}).Warn("quest_guaji:元神金装类型无效")
			return false
		}
		num := q.QuestDataMap[k]
		if num >= v {
			continue
		}
		needNum := v - num
		for i := 0; i < int(needNum); i++ {
			quest.ClickHandle(p, clicktypes.ClickTypeGoldEquipment, clickSubTypeGoldEquipment)
		}
	}

	return true
}
