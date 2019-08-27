package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questeventtypes "fgame/fgame/game/quest/event/types"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
)

//任务完成
func questFinish(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	questId := data.(int32)

	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}

	if !p.IsGuaJi() {
		return
	}
	s := p.GetScene()
	if s == nil {
		return
	}
	if questTemplate.GetQuestSubType() != questtypes.QuestSubTypeWorldBoss {
		return
	}
	if !s.MapTemplate().IsBoss() {
		return
	}
	p.BackLastScene()
	return nil
}

func init() {
	gameevent.AddEventListener(questeventtypes.EventTypeQuestFinish, event.EventListenerFunc(questFinish))
}
