package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/core/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	propertylogic "fgame/fgame/game/property/logic"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
)

//普通怪被击杀 同仙盟分享经验
func monsterKilled(target event.EventTarget, data event.EventData) (err error) {
	//场景过滤
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	n := data.(scene.NPC)
	monsterId := int32(n.GetBiologyTemplate().TemplateId())

	s := pl.GetScene()
	if s == nil {
		return
	}
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeFourGodWar {
		return
	}

	//获取怪物模板
	to := template.GetTemplateService().Get(int(monsterId), (*gametemplate.BiologyTemplate)(nil))
	if to == nil {
		return
	}
	bt := to.(*gametemplate.BiologyTemplate)

	//不是普通怪
	if bt.GetBiologyScriptType() != scenetypes.BiologyScriptTypeMonster {
		return
	}

	//仙盟id
	curAllianceId := pl.GetAllianceId()
	if curAllianceId == 0 {
		return
	}

	sceneAlliancePlayers := make([]player.Player, 0, 8)
	allPlayers := s.GetAllPlayers()
	for curPlayerId, apl := range allPlayers {
		if curAllianceId != apl.GetAllianceId() {
			continue
		}
		if curPlayerId == pl.GetId() {
			continue
		}
		spl, ok := apl.(player.Player)
		if !ok {
			continue
		}
		sceneAlliancePlayers = append(sceneAlliancePlayers, spl)
	}

	//同仙盟的共享杀普通怪经验
	for _, spl := range sceneAlliancePlayers {
		propertylogic.AddExpKillMonster(spl, monsterId, int64(bt.ExpBase), int64(bt.ExpPoint))
	}
	return nil
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeMonsterKilled, event.EventListenerFunc(monsterKilled))
}
