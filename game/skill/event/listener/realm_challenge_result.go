package listener

// //天劫塔挑战结果
// func realmChallengeResult(target event.EventTarget, data event.EventData) (err error) {
// 	pl := target.(player.Player)
// 	if pl == nil {
// 		return
// 	}
// 	sucessful := data.(bool)
// 	if !sucessful {
// 		return
// 	}
// 	manager := pl.GetPlayerDataManager(types.PlayerRealmDataManagerType).(*playerrealm.PlayerRealmDataManager)
// 	level := manager.GetTianJieTaLevel()

// 	newSkillId := realmtemplate.GetRealmTemplateService().GetSkillId(level)
// 	oldSkillId := realmtemplate.GetRealmTemplateService().GetSkillId(level - 1)
// 	err = skilllogic.TempSkillChange(pl, oldSkillId, newSkillId)
// 	return
// }

// func init() {
// 	gameevent.AddEventListener(realmeventtypes.EventTypeRealmResult, event.EventListenerFunc(realmChallengeResult))
// }
