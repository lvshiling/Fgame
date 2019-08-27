package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/tulong/pbutil"
	tulongscene "fgame/fgame/cross/tulong/scene"
	"fgame/fgame/cross/tulong/tulong"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	tulongtemplate "fgame/fgame/game/tulong/template"
)

//boss被杀
func battleObjectDead(target event.EventTarget, data event.EventData) (err error) {
	bo, ok := target.(scene.BattleObject)
	if !ok {
		return
	}
	npc, ok := bo.(scene.NPC)
	if !ok {
		return
	}
	attackId, ok := data.(int64)
	if !ok {
		return
	}

	pl := npc.GetScene().GetPlayer(attackId)
	if pl == nil {
		return
	}

	s := pl.GetScene()
	sd := s.SceneDelegate()
	if sd == nil {
		return
	}
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeCrossTuLong {
		return
	}
	scendData, ok := sd.(tulongscene.TuLongSceneData)
	if !ok {
		return
	}

	alliaceId := pl.GetAllianceId()
	if alliaceId == 0 {
		return
	}
	biologyTemplate := npc.GetBiologyTemplate()
	if biologyTemplate == nil {
		return
	}
	tuLongTemplate, flag := tulongtemplate.GetTuLongTemplateService().GetTuLongTemplateByBiologyId(int32(biologyTemplate.TemplateId()))
	if !flag {
		return
	}

	tulong.GetTuLongService().KillBoss(pl.GetServerId(), alliaceId, pl.GetAllianceName())
	scendData.KillBoss(pl)
	//发邮件使用
	biologyId := tuLongTemplate.BiologyId
	isTuLongKillBoss := pbutil.BuildISTuLongKillBoss(biologyId)
	pl.SendMsg(isTuLongKillBoss)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(battleObjectDead))
}
