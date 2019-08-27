package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/baby/baby"
	babyeventtypes "fgame/fgame/game/baby/event/types"
	babylogic "fgame/fgame/game/baby/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

// 玩家宝宝天赋变化
func playerBabyTalentChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	eventData, ok := data.(*babyeventtypes.PlayerBabyTalentChangedEventData)
	if !ok {
		return
	}

	//加载技能变化
	babylogic.LoadBabySkill(pl, eventData.GetOldEffectSkillList(), eventData.GetNewEffectSkillList())

	//同步到全局
	baby.GetBabyService().SyncBabyTalentList(pl, eventData.GetBabyId(), eventData.GetTalentList())
	return
}

func init() {
	gameevent.AddEventListener(babyeventtypes.EventTypeBabyTalentChanged, event.EventListenerFunc(playerBabyTalentChanged))
}
