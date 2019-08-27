package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
	"fgame/fgame/game/tulongequip/pbutil"
	playertulongequip "fgame/fgame/game/tulongequip/player"
	tulongequiptemplate "fgame/fgame/game/tulongequip/template"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)

	//套装技能（装备数量，自动激活）
	loadEquipSuitGroupSkill(pl)
	//装备技能（手动激活）
	loadEquipSkill(pl)

	//推送所有槽位
	tulongequipManager := pl.GetPlayerDataManager(playertypes.PlayerTuLongEquipDataManagerType).(*playertulongequip.PlayerTuLongEquipDataManager)
	allSlotMap := tulongequipManager.GetAllEquipSlotMap()
	scMsg := pbutil.BuildSCTuLongEquipInfoNotice(allSlotMap)
	pl.SendMsg(scMsg)

	return
}

// 套装技能（装备数量，自动激活）
func loadEquipSuitGroupSkill(pl player.Player) {
	tulongequipManager := pl.GetPlayerDataManager(playertypes.PlayerTuLongEquipDataManagerType).(*playertulongequip.PlayerTuLongEquipDataManager)
	allTuLongGroup := tulongequipManager.GetAllTuLongEquipGroupNum()
	for _, suitGroupMap := range allTuLongGroup {
		for groupId, num := range suitGroupMap {
			suitGroupTemplate := tulongequiptemplate.GetTuLongEquipTemplateService().GetTuLongEquipTemplateBySuitGroup(groupId)
			if suitGroupTemplate == nil {
				continue
			}
			suitSkillList := suitGroupTemplate.GetSuitEffectSkillId(num)
			for _, skill := range suitSkillList {
				skilllogic.TempSkillChange(pl, 0, skill)
			}
		}
	}
}

//装备技能（手动激活）
func loadEquipSkill(pl player.Player) {
	tulongequipManager := pl.GetPlayerDataManager(playertypes.PlayerTuLongEquipDataManagerType).(*playertulongequip.PlayerTuLongEquipDataManager)
	allSkillMap := tulongequipManager.GetSuitSkillList()
	for suitType, skillObj := range allSkillMap {
		skillTemp := tulongequiptemplate.GetTuLongEquipTemplateService().GetTuLongEquipTemplateSkill(suitType, skillObj.GetLevel())
		if skillTemp == nil {
			continue
		}
		skilllogic.TempSkillChange(pl, 0, skillTemp.SkillId)
	}

	scMsg := pbutil.BuildSCTuLongEquipSkillNotice(allSkillMap)
	pl.SendMsg(scMsg)
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
