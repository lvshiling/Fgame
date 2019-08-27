package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/baby/baby"
	babylogic "fgame/fgame/game/baby/logic"
	"fgame/fgame/game/baby/pbutil"
	playerbaby "fgame/fgame/game/baby/player"
	babytemplate "fgame/fgame/game/baby/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
)

// 玩家加载后
func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)
	babyManager.CheckBabyBorn()
	// 加载宝宝技能
	skillList := babyManager.GetEffectTalentSkillList()
	babylogic.LoadBabySkill(pl, nil, skillList)

	//套装技能（装备数量，自动激活）
	allToyGroup := babyManager.GetAllToyGroupNum()
	for _, suitGroupMap := range allToyGroup {
		for groupId, num := range suitGroupMap {
			suitGroupTemplate := babytemplate.GetBabyTemplateService().GetBabyToyTemplateBySuitGroup(groupId)
			if suitGroupTemplate == nil {
				continue
			}
			suitSkillList := suitGroupTemplate.GetSuitEffectSkillId(num)
			for _, skill := range suitSkillList {
				skilllogic.TempSkillChange(pl, 0, skill)
			}
		}
	}

	//下发宝宝信息
	babyList := babyManager.GetBabyInfoList()
	pregnantInfo := babyManager.GetPregnantInfo()
	allToySlotMap := babyManager.GetAllToySlotMap()
	power := babyManager.GetPower()
	scMsg := pbutil.BuildSCBabyInfo(pregnantInfo, babyList, allToySlotMap)
	pl.SendMsg(scMsg)
	scBabyPowerNotice := pbutil.BuildSCBabyPowerNotice(power)
	pl.SendMsg(scBabyPowerNotice)

	//同步配偶宝宝数据
	spouseId := pl.GetSpouseId()
	coupleBabyList := baby.GetBabyService().GetCoupleBabyInfo(spouseId)
	if len(coupleBabyList) > 0 {
		babyManager.LoadAllCoupleBaby(coupleBabyList)
	}

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}
