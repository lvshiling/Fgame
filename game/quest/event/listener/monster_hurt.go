package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/core/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	"fgame/fgame/game/quest/pbutil"
	playerquest "fgame/fgame/game/quest/player"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	gametemplate "fgame/fgame/game/template"
)

//怪物被伤害
func monsterHurted(target event.EventTarget, data event.EventData) (err error) {

	monsterId, ok := data.(int32)
	if !ok {
		return
	}
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	questMap := manager.GetQuestMap(questtypes.QuestStateAccept)
	questList := make([]*playerquest.PlayerQuestObject, 0, 8)
	for _, qu := range questMap {
		questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(qu.QuestId)
		subType := questTemplate.GetQuestSubType()
		if subType != questtypes.QuestSubTypeHurtMonster &&
			subType != questtypes.QuestSubTypeSomeMonster {
			continue
		}
		questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
		switch subType {
		case questtypes.QuestSubTypeHurtMonster:
			{
				_, ok := questDemandMap[monsterId]
				if !ok {
					continue
				}
				flag := manager.IncreaseQuestData(qu.QuestId, monsterId, 1)
				if !flag {
					panic("quest:伤害怪应该成功")
				}
				break
			}
		case questtypes.QuestSubTypeSomeMonster:
			{
				//获取怪物模板
				to := template.GetTemplateService().Get(int(monsterId), (*gametemplate.BiologyTemplate)(nil))
				if to == nil {
					return
				}
				bt := to.(*gametemplate.BiologyTemplate)
				scriptTypeInt := int32(bt.GetBiologyScriptType())
				_, ok := questDemandMap[scriptTypeInt]
				if !ok {
					continue
				}
				flag := manager.IncreaseQuestData(qu.QuestId, scriptTypeInt, 1)
				if !flag {
					panic("quest:伤害怪应该成功")
				}
				break
			}
		}

		//TODO 检测任务是否完成
		_, err = questlogic.CheckQuestIfFinish(pl, qu.QuestId)
		if err != nil {
			return err
		}
		questList = append(questList, qu)
		// scQuestUpdate := pbutil.BuildSCQuestUpdate(qu)
		// pl.SendMsg(scQuestUpdate)
	}
	if len(questList) != 0 {
		scQuestUpdate := pbutil.BuildSCQuestListUpdate(questList)
		pl.SendMsg(scQuestUpdate)
	}
	return nil
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeMonsterHurted, event.EventListenerFunc(monsterHurted))
}
