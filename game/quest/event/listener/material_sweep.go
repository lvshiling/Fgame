package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	materialeventtypes "fgame/fgame/game/material/event/types"
	playermaterial "fgame/fgame/game/material/player"
	materialtemplate "fgame/fgame/game/material/template"
	materialtypes "fgame/fgame/game/material/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//材料副本扫荡
func materialSweep(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	eventData, ok := data.(*materialeventtypes.MaterialFinishEventData)
	if !ok {
		return
	}
	typ := eventData.GetType()
	num := eventData.GetNum()
	if num <= 0 {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerMaterialDataManagerType).(*playermaterial.PlayerMaterialDataManager)
	materialObj := manager.GetPlayerMaterialInfo(typ)
	if materialObj == nil {
		return
	}
	useTimes := materialObj.GetUseTimes()
	materialTemplate := materialtemplate.GetMaterialTemplateService().GetMaterialTemplate(typ)
	if materialTemplate == nil {
		return
	}
	allTimes := materialTemplate.AllTimes
	leftNum := allTimes - useTimes

	allLeftTimes := manager.GetAllLeftTimes()
	err = sweepMaterialFuBen(pl, allLeftTimes, num)
	if err != nil {
		return
	}

	err = sweepSpecialMaterialFuBen(pl, typ, allTimes, leftNum)
	if err != nil {
		return
	}
	return
}

//挑战材料副本x次
func sweepMaterialFuBen(pl player.Player, leftNum int32, num int32) (err error) {
	if leftNum <= 0 {
		return questlogic.FillQuestData(pl, questtypes.QuestSubTypechallengeMaterialFuBen, 0)
	} else {
		return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypechallengeMaterialFuBen, 0, num)
	}
	return
}

//扫荡指定的材料副本
func sweepSpecialMaterialFuBen(pl player.Player, typ materialtypes.MaterialType, allTimes int32, leftNum int32) (err error) {
	if leftNum <= 0 {
		return questlogic.FillQuestData(pl, questtypes.QuestSubTypechallengeSpecialMaterialFuBen, int32(typ))
	} else {
		return questlogic.SetQuestData(pl, questtypes.QuestSubTypechallengeSpecialMaterialFuBen, int32(typ), allTimes-leftNum)
	}
	return
}

func init() {
	gameevent.AddEventListener(materialeventtypes.EventTypeMaterialFinish, event.EventListenerFunc(materialSweep))
}
