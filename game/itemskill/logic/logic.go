package logic

import (
	playeritemskill "fgame/fgame/game/itemskill/player"
	itemskilltemplate "fgame/fgame/game/itemskill/template"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerrealm "fgame/fgame/game/realm/player"
	realmtemplate "fgame/fgame/game/realm/template"
)

//返回老玩家护体盾技能
func RestitutionHuTiDun(pl player.Player) {
	realmManager := pl.GetPlayerDataManager(playertypes.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
	level := realmManager.GetTianJieTaLevel()
	rSkillId := realmtemplate.GetRealmTemplateService().GetSkillId(level)
	if rSkillId == 0 {
		return
	}

	skTemplate := itemskilltemplate.GetItemSkillTemplateService().GetTemplateBySkillId(rSkillId)
	if skTemplate == nil {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerItemSkillDataManagerType).(*playeritemskill.PlayerItemSkillDataManager)
	manager.CreateItemSkillObjByRestitution(rSkillId)
	return
}
