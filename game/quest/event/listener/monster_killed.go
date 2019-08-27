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
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
)

//怪物被击杀
func monsterKilled(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	n, ok := data.(scene.NPC)
	if !ok {
		return
	}
	monsterId := int32(n.GetBiologyTemplate().TemplateId())
	manager := pl.GetPlayerDataManager(types.PlayerQuestDataManagerType).(*playerquest.PlayerQuestDataManager)
	questMap := manager.GetQuestMap(questtypes.QuestStateAccept)
	questList := make([]*playerquest.PlayerQuestObject, 0, 8)
	for _, qu := range questMap {
		questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(qu.QuestId)
		subType := questTemplate.GetQuestSubType()
		if subType != questtypes.QuestSubTypeKillMonster {
			continue
		}
		questDemandMap := questTemplate.GetQuestDemandMap(pl.GetRole())
		_, ok := questDemandMap[monsterId]
		if !ok {
			continue
		}
		flag := manager.IncreaseQuestData(qu.QuestId, monsterId, 1)
		if !flag {
			panic("quest:杀怪应该成功")
		}
		//TODO 检测任务是否完成
		_, err = questlogic.CheckQuestIfFinish(pl, qu.QuestId)
		if err != nil {
			return err
		}
		questList = append(questList, qu)
	}
	if len(questList) != 0 {
		scQuestUpdate := pbutil.BuildSCQuestListUpdate(questList)
		pl.SendMsg(scQuestUpdate)
	}

	to := template.GetTemplateService().Get(int(monsterId), (*gametemplate.BiologyTemplate)(nil))
	if to == nil {
		return
	}
	bt := to.(*gametemplate.BiologyTemplate)
	scriptType := bt.GetBiologyScriptType()
	setType := bt.GetBiologySetType()

	err = killSetType(pl, setType)
	if err != nil {
		return
	}

	err = killBiologyType(pl, scriptType)
	if err != nil {
		return
	}
	err = killWorldBoss(pl, scriptType)
	if err != nil {
		return
	}

	err = attendKillSpecialMonster(pl, scriptType)
	if err != nil {
		return
	}

	err = killBoss(pl, scriptType)
	if err != nil {
		return
	}

	return nil
}

//击杀指定的策划类型怪物
func killSetType(pl player.Player, setType scenetypes.BiologySetType) (err error) {
	return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeSetTypeKilled, int32(setType), 1)
}

//击杀某种类型怪物
func killBiologyType(pl player.Player, scriptType scenetypes.BiologyScriptType) (err error) {
	return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeBiologyTypeKilled, int32(scriptType), 1)
}

//击杀X只世界BOSS
func killWorldBoss(pl player.Player, scriptType scenetypes.BiologyScriptType) (err error) {
	switch scriptType {
	case scenetypes.BiologyScriptTypeWorldBoss,
		scenetypes.BiologyScriptTypeBossCallTicket,
		scenetypes.BiologyScriptTypeMyBoss,
		scenetypes.BiologyScriptTypeVIPMyBoss:
		{
			return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeWorldBoss, 0, 1)
		}
	}
	return
}

//击杀个人boss和付费boss
func attendKillSpecialMonster(pl player.Player, scriptType scenetypes.BiologyScriptType) (err error) {
	switch scriptType {
	case scenetypes.BiologyScriptTypeMyBoss,
		scenetypes.BiologyScriptTypeVIPMyBoss:
		{
			return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeAttandKillSpecialMonster, 0, 1)
		}
	}
	return
}

//击杀世界boss,跨服boss
func killBoss(pl player.Player, scriptType scenetypes.BiologyScriptType) (err error) {
	switch scriptType {
	case scenetypes.BiologyScriptTypeWorldBoss,
		scenetypes.BiologyScriptTypeCrossWorldBoss:
		{
			return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeBossKilled, 0, 1)
		}
	}
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeMonsterKilled, event.EventListenerFunc(monsterKilled))
}
