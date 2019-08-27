package check

import (
	playermarry "fgame/fgame/game/marry/player"
	marrytemplate "fgame/fgame/game/marry/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	propertytypes "fgame/fgame/game/property/types"
	playerquest "fgame/fgame/game/quest/player"
	"fgame/fgame/game/quest/quest"
	questtypes "fgame/fgame/game/quest/types"
	gametemplate "fgame/fgame/game/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	quest.RegisterCheck(questtypes.QuestSubTypeCoupleToeknForce, quest.CheckHandlerFunc(handleCoupleTokenForce))
}

func handleCoupleTokenForce(pl player.Player, questTemplate *gametemplate.QuestTemplate) (err error) {
	log.Debug("quest:处理夫妻信物战斗力")

	marryManager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	if !marryManager.IsTrueMarry() {
		return
	}

	power := calculateCoupleTokenForce(pl)

	questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
	//模板校验过数据配一个
	for demandId, _ := range questDemandMap {
		questManager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
		flag := questManager.SetQuestData(int32(questTemplate.TemplateId()), demandId, int32(power))
		if !flag {
			panic("quest:设置 SetQuestData 应该成功")
		}
		break
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("quest:处理夫妻信物战斗力,完成")
	return nil
}

func calculateCoupleTokenForce(pl player.Player) int64 {
	marryManager := pl.GetPlayerDataManager(types.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	hp := int64(0)
	attack := int64(0)
	defence := int64(0)

	playerSuitMap := marryManager.GetAllDingQingMap()
	if len(playerSuitMap) == 0 {
		return 0
	}
	spouseSuitMap := marryManager.GetSpouseSuit()

	for suitId, posMap := range playerSuitMap {
		//己算碎片
		for posId, _ := range posMap {
			item := marrytemplate.GetMarryTemplateService().GetMarryXinWuItem(suitId, posId)
			if item == nil {
				continue
			}
			addTimes := int32(1)
			_, exists := spouseSuitMap[suitId]
			if exists {
				_, exists = spouseSuitMap[suitId][posId]
				if exists {
					addTimes = int32(2)
				}
			}
			hp += int64(item.Hp * addTimes)
			attack += int64(item.Attack * addTimes)
			defence += int64(item.Defence * addTimes)
		}

		//开始计算套装
		suitLen := len(posMap)
		spouseLen := 0
		_, exists := spouseSuitMap[suitId]
		if exists {
			spouseLen = len(spouseSuitMap[suitId])
		}
		suitTemplate := marrytemplate.GetMarryTemplateService().GetMarryXinWuGroupTemplate(suitId)
		if suitTemplate == nil {
			continue
		}
		for i := 1; i <= suitLen; i++ {
			suitAddMap := suitTemplate.GetSuitAddMap()
			_, exists := suitAddMap[int32(i)]
			if exists {
				suitItem := suitAddMap[int32(i)]
				suitTime := int32(1)
				if i <= spouseLen { //伴侣也有
					suitTime = int32(2)
				}
				hp += int64(suitItem.Hp * suitTime)
				attack += int64(suitItem.Attack * suitTime)
				defence += int64(suitItem.Defence * suitTime)
			}
		}
	}

	battlePropertyMap := make(map[propertytypes.BattlePropertyType]int64)
	battlePropertyMap[propertytypes.BattlePropertyTypeMaxHP] = hp
	battlePropertyMap[propertytypes.BattlePropertyTypeAttack] = attack
	battlePropertyMap[propertytypes.BattlePropertyTypeDefend] = defence

	power := propertylogic.CulculateForce(battlePropertyMap)
	return power
}
