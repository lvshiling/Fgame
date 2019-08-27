package system

import (
	anqilogic "fgame/fgame/game/anqi/logic"
	playeranqi "fgame/fgame/game/anqi/player"
	anqitemplate "fgame/fgame/game/anqi/template"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/quest/guaji/guaji"
	playerquest "fgame/fgame/game/quest/player"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	guaji.RegisterQuestGuaJiSystemXHandler(questtypes.SystemReachXTypeAnQi, guaji.QuestGuaJiSystemXHandlerFunc(anqiAdvanced))
}

//暗器进阶
func anqiAdvanced(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {
	questType := questTemplate.GetQuestSubType()
	log.WithFields(
		log.Fields{
			"playerId":   p.GetId(),
			"questType":  questType.String(),
			"reachXType": questtypes.SystemReachXTypeAnQi,
		}).Info("quest_guaji:正在做任务")
	demandMap := questTemplate.GetQuestDemandMap(p.GetRole())
	if len(demandMap) <= 0 {
		log.WithFields(
			log.Fields{
				"playerId":   p.GetId(),
				"questType":  questType.String(),
				"reachXType": questtypes.SystemReachXTypeAnQi,
			}).Warn("quest_guaji:任务需求是空")
		return false
	}
	//获取当前任务数据
	questManager := p.GetPlayerDataManager(playertypes.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	q := questManager.GetQuestById(int32(questTemplate.TemplateId()))
	if q == nil {
		log.WithFields(
			log.Fields{
				"playerId":   p.GetId(),
				"questType":  questType.String(),
				"reachXType": questtypes.SystemReachXTypeAnQi,
			}).Warn("quest_guaji:任务不存在")
		return false
	}
	if q.QuestState != questtypes.QuestStateAccept {
		log.WithFields(
			log.Fields{
				"playerId":   p.GetId(),
				"questType":  questType.String(),
				"reachXType": questtypes.SystemReachXTypeAnQi,
			}).Warn("quest_guaji:任务没在进行中")
		return false
	}

	//TODO 验证是否可以升级
	num := q.QuestDataMap[int32(questtypes.SystemReachXTypeAnQi)]
	needNum := demandMap[int32(questtypes.SystemReachXTypeAnQi)]
	if num >= needNum {
		return false
	}

	anqiManager := p.GetPlayerDataManager(playertypes.PlayerAnqiDataManagerType).(*playeranqi.PlayerAnqiDataManager)
	anqiInfo := anqiManager.GetAnqiInfo()
	nextAdvancedId := anqiInfo.AdvanceId + 1
	anqiTempalte := anqitemplate.GetAnqiTemplateService().GetAnqiNumber(int32(nextAdvancedId))
	if anqiTempalte == nil {
		log.WithFields(
			log.Fields{
				"playerId":   p.GetId(),
				"questType":  questType.String(),
				"reachXType": questtypes.SystemReachXTypeAnQi,
			}).Warn("quest_guaji:已经最高阶")
		return false
	}

	inventoryManager := p.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := p.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	//进阶需要消耗的元宝
	costGold := anqiTempalte.UseMoney
	if !propertyManager.HasEnoughGold(int64(costGold), false) {
		log.WithFields(
			log.Fields{
				"playerId":   p.GetId(),
				"questType":  questType.String(),
				"reachXType": questtypes.SystemReachXTypeAnQi,
			}).Warn("quest_guaji:元宝不足")
		return false
	}
	//进阶需要消耗的银两
	costSilver := int64(anqiTempalte.UseYinliang)
	if !propertyManager.HasEnoughSilver(int64(costSilver)) {
		log.WithFields(
			log.Fields{
				"playerId":   p.GetId(),
				"questType":  questType.String(),
				"reachXType": questtypes.SystemReachXTypeAnQi,
			}).Warn("quest_guaji:银两不足")
		return false
	}

	useItemTemplate := anqiTempalte.GetUseItemTemplate()
	if useItemTemplate != nil {
		if !inventoryManager.HasEnoughItem(anqiTempalte.UseItem, anqiTempalte.ItemCount) {
			log.WithFields(
				log.Fields{
					"playerId":   p.GetId(),
					"questType":  questType.String(),
					"reachXType": questtypes.SystemReachXTypeAnQi,
				}).Warn("quest_guaji:物品不足")
			return false
		}
	}
	anqilogic.HandleAnqiAdvanced(p, false)
	return true
}
