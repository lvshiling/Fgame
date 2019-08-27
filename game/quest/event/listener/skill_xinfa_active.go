package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	skilleventtypes "fgame/fgame/game/skill/event/types"
	playerskill "fgame/fgame/game/skill/player"
)

// 玩家激活技能心法

func init() {
	gameevent.AddEventListener(skilleventtypes.EventTypeSkillTianFuAwaken, event.EventListenerFunc(playerSkillXinFaActive))
}

func playerSkillXinFaActive(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
	skillMap := manager.GetRoleSkillMap()
	skillXinFaNum := 0
	for _, obj := range skillMap {
		if len(obj.TianFuMap) != 0 {
			skillXinFaNum += len(obj.TianFuMap)
		}
	}

	questlogic.SetQuestData(pl, questtypes.QuestSubTypeSkillXinFa, 0, int32(skillXinFaNum))
	return
}
