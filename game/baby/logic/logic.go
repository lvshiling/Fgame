package logic

import (
	"fgame/fgame/game/baby/pbutil"
	playerbaby "fgame/fgame/game/baby/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	skilllogic "fgame/fgame/game/skill/logic"
)

//推送玩具改变
func SnapBabyToyChanged(pl player.Player) (err error) {
	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)
	slotChangedMap := babyManager.GetChangedToySlotAndResetMap()
	if len(slotChangedMap) <= 0 {
		return
	}
	scMsg := pbutil.BuildSCBabyToySlotChanged(slotChangedMap)
	pl.SendMsg(scMsg)
	return
}

//宝宝属性改变
func BabyPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeBaby.Mask())

	return
}

func LoadBabySkill(pl player.Player, oldEffectSkillList, newEffectSkillList map[int32]int32) {
	//先把所有技能卸下来,再装新的
	for _, skillId := range oldEffectSkillList {
		skilllogic.TempSkillChange(pl, skillId, 0)
	}

	for _, skillId := range newEffectSkillList {
		skilllogic.TempSkillChange(pl, 0, skillId)
	}
}
