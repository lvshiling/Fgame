package battle

import (
	"fgame/fgame/core/template"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"time"
)

// TODO 控制同步时间
const (
	propertyTaskTime = time.Second
)

type propertyTask struct {
	p               scene.Player
	lastRecoverTime int64
}

func (t *propertyTask) Run() {
	now := global.GetGame().GetTimeService().Now()
	tempLevelTemplate := template.GetTemplateService().Get(int(t.p.GetLevel()), (*gametemplate.CharacterLevelTemplate)(nil))

	levelTemplate := tempLevelTemplate.(*gametemplate.CharacterLevelTemplate)
	if levelTemplate.TpReply <= 0 {
		t.lastRecoverTime = now
		return
	}
	elaspse := now - t.lastRecoverTime
	if elaspse > int64(levelTemplate.TpReplyTime) {
		recover := elaspse / int64(levelTemplate.TpReplyTime)
		t.p.AddTP(recover * int64(levelTemplate.TpReply))
		t.lastRecoverTime += recover * int64(levelTemplate.TpReplyTime)
	}
}

func (t *propertyTask) ElapseTime() time.Duration {
	return propertyTaskTime
}

func CreatePropertyTask(p scene.Player) *propertyTask {
	t := &propertyTask{
		p: p,
	}
	now := global.GetGame().GetTimeService().Now()
	t.lastRecoverTime = now
	return t
}
