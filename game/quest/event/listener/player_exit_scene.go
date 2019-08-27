package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	materiallogic "fgame/fgame/game/material/logic"
	playermaterial "fgame/fgame/game/material/player"
	materialtemplate "fgame/fgame/game/material/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	realmlogic "fgame/fgame/game/realm/logic"
	scenetypes "fgame/fgame/game/scene/types"
	xianfutypes "fgame/fgame/game/xianfu/types"
)

//玩家退出场景
func playerExitScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}

	err = enterRealm(pl)
	if err != nil {
		return
	}
	err = enterSpecialXianFu(pl)
	if err != nil {
		return
	}
	err = enter1v1(pl)
	if err != nil {
		return
	}
	err = attandActivityPlay(pl)
	if err != nil {
		return
	}

	err = challengeMaterialFuBen(pl)
	if err != nil {
		return
	}

	err = challengeSpecialMaterialFuBen(pl)
	if err != nil {
		return
	}
	return
}

//进入天劫塔X次
func enterRealm(pl player.Player) (err error) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	mapType := s.MapTemplate().GetMapType()
	switch mapType {
	case scenetypes.SceneTypeTianJieTa:
		sd := s.SceneDelegate()
		if sd == nil {
			return
		}
		sceneData := sd.(*realmlogic.TianJieTaSceneData)
		if sceneData.GetOwerId() == pl.GetId() {
			return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeEnterRealm, 0, 1)
		}
	}
	return
}

//进入指定X次秘境仙府
func enterSpecialXianFu(pl player.Player) (err error) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	mapType := s.MapTemplate().GetMapType()
	switch mapType {
	case scenetypes.SceneTypeExperience,
		scenetypes.SceneTypeYinLiang,
		scenetypes.SceneTypeMaterial:
		if mapType == scenetypes.SceneTypeExperience {
			questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeEnterSpecialXianFu, int32(xianfutypes.XianfuTypeExp), 1)
		} else if mapType == scenetypes.SceneTypeYinLiang {
			questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeEnterSpecialXianFu, int32(xianfutypes.XianfuTypeSilver), 1)
		} else {
			questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeEnterSpecialXianFu, int32(xianfutypes.XianfuTypeItem), 1)
		}

		if mapType == scenetypes.SceneTypeExperience ||
			mapType == scenetypes.SceneTypeYinLiang {
			questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeXianFuPersonal, 0, 1)
		}
	}
	return
}

//参加1V1竞技场X次
func enter1v1(pl player.Player) (err error) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	mapType := s.MapTemplate().GetMapType()
	switch mapType {
	case scenetypes.SceneTypeLingChiFighting:
		return questlogic.IncreaseQuestData(pl, questtypes.QuestSubType1V1, 0, 1)
	}
	return
}

//参与活动玩法
func attandActivityPlay(pl player.Player) (err error) {
	s := pl.GetScene()
	if s == nil {
		return
	}
	mapType := s.MapTemplate().GetMapType()
	activityType, flag := mapType.ToActivityType()
	if !flag {
		return
	}
	return questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeAttandActivityPlay, int32(activityType), 1)
}

//挑战材料副本x次
func challengeMaterialFuBen(pl player.Player) (err error) {
	s := pl.GetScene()
	if s == nil {
		return
	}

	mapType := s.MapTemplate().GetMapType()
	switch mapType {
	case scenetypes.SceneTypeMaterial:
		manager := pl.GetPlayerDataManager(types.PlayerMaterialDataManagerType).(*playermaterial.PlayerMaterialDataManager)
		allLeftTimes := manager.GetAllLeftTimes()
		if allLeftTimes <= 0 {
			return questlogic.FillQuestData(pl, questtypes.QuestSubTypechallengeMaterialFuBen, 0)
		} else {
			return questlogic.SetQuestData(pl, questtypes.QuestSubTypechallengeMaterialFuBen, 0, 1)
		}
	}
	return
}

//进入指定的材料副本
func challengeSpecialMaterialFuBen(pl player.Player) (err error) {
	s := pl.GetScene()
	if s == nil {
		return
	}

	mapType := s.MapTemplate().GetMapType()
	switch mapType {
	case scenetypes.SceneTypeMaterial:
		sd := s.SceneDelegate()
		if sd == nil {
			return
		}
		sceneData, ok := sd.(materiallogic.MaterialSceneData)
		if !ok {
			return
		}
		typ := sceneData.GetMaterialType()

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

		if leftNum <= 0 {
			return questlogic.FillQuestData(pl, questtypes.QuestSubTypechallengeSpecialMaterialFuBen, int32(typ))
		} else {
			return questlogic.SetQuestData(pl, questtypes.QuestSubTypechallengeSpecialMaterialFuBen, int32(typ), allTimes-leftNum)
		}
	}
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerExitScene, event.EventListenerFunc(playerExitScene))
}
