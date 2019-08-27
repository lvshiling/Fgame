package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questeventtypes "fgame/fgame/game/quest/event/types"
	questtemplate "fgame/fgame/game/quest/template"
	questtypes "fgame/fgame/game/quest/types"
	scenetypes "fgame/fgame/game/scene/types"
)

//任务完成
func questCommit(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	questId := data.(int32)

	questTemplate := questtemplate.GetQuestTemplateService().GetQuestTemplate(questId)
	if questTemplate == nil {
		return
	}
	//不是日常任务
	if questTemplate.GetQuestType() != questtypes.QuestTypeDaily {
		return
	}

	if !p.IsGuaJi() {
		return
	}

	s := p.GetScene()
	if s == nil {
		return
	}

	//日常任务挂机
	t, flag := p.GetLastGuaJiType()
	if !flag {
		return
	}

	//主线挂机
	if t == scenetypes.GuaJiTypeQuest {
		p.ExitGuaJi()
	}

	return nil
}

func init() {
	gameevent.AddEventListener(questeventtypes.EventTypeQuestCommit, event.EventListenerFunc(questCommit))
}
