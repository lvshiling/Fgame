package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	skillcommon "fgame/fgame/game/skill/common"
	skilleventtypes "fgame/fgame/game/skill/event/types"
	"fgame/fgame/game/skill/pbutil"
)

//玩家技能移除下发
func playerSkillRemove(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	skillObj, ok := data.(skillcommon.SkillObject)
	if !ok {
		return
	}
	skillTypeId := skillObj.GetSkillId()
	level := skillObj.GetLevel()

	scSkillRemove := pbutil.BuildSCSkillRemove(skillTypeId, level)
	pl.SendMsg(scSkillRemove)
	return
}

func init() {
	gameevent.AddEventListener(skilleventtypes.EventTypeSkillRemove, event.EventListenerFunc(playerSkillRemove))
}
