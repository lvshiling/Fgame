package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	playermarry "fgame/fgame/game/marry/player"
	marrytemplate "fgame/fgame/game/marry/template"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	propertytypes "fgame/fgame/game/property/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeDingQingTokenActivite, event.EventListenerFunc(coupleTokenForce))
}

//夫妻信物战斗力变化
func coupleTokenForce(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	marryManager := pl.GetPlayerDataManager(playertypes.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	if !marryManager.IsTrueMarry() {
		return
	}

	power := calculateCoupleTokenForce(pl)

	questlogic.SetQuestData(pl, questtypes.QuestSubTypeCoupleToeknForce, 0, int32(power))

	return
}

func calculateCoupleTokenForce(pl player.Player) int64 {
	marryManager := pl.GetPlayerDataManager(playertypes.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
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
