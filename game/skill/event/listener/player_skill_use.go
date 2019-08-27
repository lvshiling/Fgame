package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	skilleventtypes "fgame/fgame/game/skill/event/types"
)

//玩家技能使用
func playerSkillUse(target event.EventTarget, data event.EventData) (err error) {
	// pl, ok := target.(player.Player)
	// if !ok {
	// 	return
	// }
	// skillId, ok := data.(int32)
	// if !ok {
	// 	return
	// }
	// manager := pl.GetPlayerDataManager(types.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
	// manager.SkillCdTime(skillId)
	return
}

func init() {
	gameevent.AddEventListener(skilleventtypes.EventTypeSkillUse, event.EventListenerFunc(playerSkillUse))
}
