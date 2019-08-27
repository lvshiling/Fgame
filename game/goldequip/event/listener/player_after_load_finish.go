package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	"fgame/fgame/game/goldequip/pbutil"
	playergoldequip "fgame/fgame/game/goldequip/player"
	goldequiptemplate "fgame/fgame/game/goldequip/template"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
)

//加载完成后
func playerAfterLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)

	goldequiplogic.RestitutionGemUnlock(pl)     //宝石槽检测
	goldequiplogic.CheckStrengthenLevelMove(pl) //强化等级迁移检测
	//装备套装技能
	goldequipManager := pl.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	goldEquipGroup := goldequipManager.GetGoldEquipGroupNum()
	for groupId, num := range goldEquipGroup {

		goldEquipTemplate := goldequiptemplate.GetGoldEquipTemplateService().GetGoldEquipTemplateByGroup(groupId)
		if goldEquipTemplate == nil {
			return
		}
		suitSkillList := goldEquipTemplate.GetGoldEquipGroupSuitSkill(num)
		for _, skill := range suitSkillList {
			err = skilllogic.TempSkillChange(pl, 0, skill)
			if err != nil {
				return
			}
		}
	}

	//推送所有物品
	equipBag := goldequipManager.GetGoldEquipBag()
	goldEquipSlotInfoList := pbutil.BuildGoldEquipSlotInfoList(equipBag.GetAll())
	pl.SendMsg(goldEquipSlotInfoList)

	//推送所有物品
	goldEquipSetting := goldequipManager.GetGoldEquipSetting()
	scGoldEquipAutoFenJie := pbutil.BuildSCGoldEquipAutoFenJie(goldEquipSetting.GetFenJieIsAuto(), int32(goldEquipSetting.GetFenJieQuality()), goldEquipSetting.GetFenJieZhuanShu())
	pl.SendMsg(scGoldEquipAutoFenJie)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoadFinish))
}
