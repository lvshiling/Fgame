package quest_guaji

import (
	"fgame/fgame/game/gem/gem"
	gemlogic "fgame/fgame/game/gem/logic"
	playergem "fgame/fgame/game/gem/player"
	gemtypes "fgame/fgame/game/gem/types"
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

//赌石
func init() {
	guaji.RegisterQuestGuaJi(questtypes.QuestSubTypeGamble, guaji.QuestGuaJiFunc(gamble))
}

//赌石
func gamble(p player.Player, questTemplate *gametemplate.QuestTemplate) bool {

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
	v := demandMap[0]
	num := q.QuestDataMap[0]
	if num >= v {
		return true
	}
	needNum := v - num
	for i := 0; i < int(needNum); i++ {
		flag := doGamble(p, gemtypes.GemGambleTypePrimary)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": p.GetId(),
				}).Warn("quest_guaji:赌石失败")
			return false
		}

	}

	return true
}

func doGamble(pl player.Player, typ gemtypes.GemGambleType) bool {
	gemManager := pl.GetPlayerDataManager(playertypes.PlayerGemDataManagerType).(*playergem.PlayerGemDataManager)

	gamblingTemplate := gem.GetGemService().GetGambleTemplateByTyp(typ)
	if gamblingTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"gemGambleType": typ,
			}).Warn("quest_guaji:赌石模板不存在")
		return false
	}
	needYinLiang := int64(gamblingTemplate.NeedYinLiang)
	needGold := gamblingTemplate.NeedGold
	needYuanShi := gamblingTemplate.NeedYuanShi
	needItem := int32(0)
	needItemNum := int32(0)
	useItem := gamblingTemplate.GetUseItemTemplate()

	if useItem != nil {
		needItem = gamblingTemplate.NeedItem
		needItemNum = gamblingTemplate.NeedItemNum
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//判断银两
	if needYinLiang != 0 {
		flag := propertyManager.HasEnoughSilver(needYinLiang)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId":      pl.GetId(),
					"gemGambleType": typ,
				}).Warn("quest_guaji:赌石银两不足")
			return false
		}
	}

	//判断元宝
	if needGold != 0 {
		flag := propertyManager.HasEnoughGold(int64(needGold), false)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":      pl.GetId(),
				"gemGambleType": typ,
			}).Warn("quest_guaji:赌石元宝不足")
			return false
		}
	}

	//判断原石
	if needYuanShi != 0 {
		flag := gemManager.HasEnoughYuanShi(needYuanShi)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":      pl.GetId(),
				"gemGambleType": typ,
			}).Warn("quest_guaji:赌石原石不足，无法赌石")
			return false
		}
	}

	//物品判断
	if needItem != 0 {
		flag := inventoryManager.HasEnoughItem(needItem, needItemNum)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":      pl.GetId(),
				"gemGambleType": typ,
			}).Warn("quest_guaji:赌石道具不足，无法赌石")
			return false
		}
	}

	gemlogic.HandleGemGamble(pl, typ, false)
	return true
}
