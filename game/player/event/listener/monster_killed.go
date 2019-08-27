package listener

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/event"
	"fgame/fgame/core/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	propertylogic "fgame/fgame/game/property/logic"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
)

//怪物被击杀
func monsterKilled(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}
	n := data.(scene.NPC)
	monsterId := int32(n.GetBiologyTemplate().TemplateId())

	//TODO xzk:注册处理器
	//打宝塔特殊处理
	if !pl.IsOnDabao() && s.MapTemplate().IsTower() {
		if pl.IfNotDaBaoNotice() {
			playerlogic.SendSystemMessage(pl, lang.TowerNotOnDaBaoNotice)
		}
		return
	}

	//获取怪物经验
	to := template.GetTemplateService().Get(int(monsterId), (*gametemplate.BiologyTemplate)(nil))
	if to == nil {
		return
	}
	bt := to.(*gametemplate.BiologyTemplate)
	expBase := bt.ExpBase
	expPoint := bt.ExpPoint
	propertylogic.AddExpKillMonster(pl, monsterId, int64(expBase), int64(expPoint))

	return nil
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeMonsterKilled, event.EventListenerFunc(monsterKilled))
}
