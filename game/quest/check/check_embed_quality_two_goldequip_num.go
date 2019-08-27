package check

import (
	playergoldequip "fgame/fgame/game/goldequip/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeEmbedQualityTwoGoldEquipNum, quest.CheckHandlerFunc(handleEmbedQualityTwoGoldEquipNum))
}

//check 穿戴元神金装的数量
func handleEmbedQualityTwoGoldEquipNum(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理穿戴元神金装的数量")

	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	//模板校验过数据配一个
	for demandId, _ := range questDemandMap {
		manager := pl.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
		slotList := manager.GetGoldEquipBag().GetAll()
		totalNum := int32(0)
		for _, slot := range slotList {
			if slot.IsEmpty() {
				continue
			}
			itemTemp := item.GetItemService().GetItem(int(slot.GetItemId()))
			if itemTemp == nil {
				continue
			}
			if itemTemp.Quality < int32(itemtypes.ItemQualityTypeBlue) {
				continue
			}
			if itemTemp.NeedZhuanShu < int32(demandId) {
				continue
			}
			totalNum++
		}
		if totalNum <= 0 {
			return
		}

		questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
		flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), demandId, totalNum)
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
		break
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理穿戴元神金装的数量,完成")
	return nil
}
